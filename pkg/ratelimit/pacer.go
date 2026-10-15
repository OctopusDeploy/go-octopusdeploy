package ratelimit

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// PolicyHeaderName advertises the policy in force for us, e.g. `l=200;rpm=600`: l is the burst limit and rpm
	// how fast it refills. Static per policy, so it rarely changes.
	PolicyHeaderName = "Octopus-RateLimit-Policy"

	// RateLimitHeaderName advertises where we currently stand against that policy, e.g. `r=7;t=18`: r is what we
	// have left and t how long until the whole allowance is back. Sent on every response.
	RateLimitHeaderName = "Octopus-RateLimit"
)

// reservedRequests is how much of the advertised allowance we hold back rather than spend. The server's remaining
// count is whole requests and its bucket refills on a timer rather than continuously, so its arithmetic and ours
// disagree slightly; stopping a request short of the line means that disagreement costs us a small wait instead of
// a rejection.
const reservedRequests = 1

// pacer paces outgoing requests against the rate limit the server advertises, using GCRA (the generic cell rate
// algorithm) - the same shape of algorithm the server's token bucket uses, run in reverse. It is a port of the .NET
// client's RateLimitPacer and the web portal's rateLimitPacer.ts, so all three behave the same way.
//
// The state is a single "theoretical arrival time" (TAT): the point at which the requests we have already sent will
// have been paid for at the policy's steady rate. A request may go as soon as the TAT is within burstTolerance of
// now, and each request pushes the TAT out by one emissionInterval. That gives us the same "burst then settle to the
// steady rate" shape as the server's bucket, without needing a timer to refill anything.
//
// Every response resyncs the TAT to what the server says is left, so we self-correct: if something else is spending
// the same user's allowance, or the server's replenishment doesn't line up exactly with our arithmetic, the next
// response pulls us back into line. The resync only ever moves the TAT later, never earlier, so a response that
// overtook a newer one can't talk us into speeding up.
//
// A pacer with no policy is inert. Nothing paces until the server tells us what the limit is, so an older server or
// an unlimited policy behaves exactly as it did before.
type pacer struct {
	mu sync.Mutex

	// now returns elapsed time on a monotonic clock rather than wall clock time: we only ever compare our own
	// timestamps, and this can't jump. Replaceable in tests.
	now func() time.Duration

	havePolicy        bool
	burstLimit        float64
	requestsPerMinute float64
	// emissionInterval is how long a single request costs at the policy's steady rate.
	emissionInterval time.Duration
	// burstTolerance is how far ahead of the steady rate we are allowed to run.
	burstTolerance         time.Duration
	theoreticalArrivalTime time.Duration
}

func newPacer() *pacer {
	start := time.Now()
	return &pacer{now: func() time.Duration { return time.Since(start) }}
}

// reserveSlot claims the next slot in the schedule and returns how long the caller should wait before sending - zero
// when we have allowance to spare, which is the normal case.
//
// The slot is claimed whether or not the caller ends up waiting, so concurrent callers get consecutive slots and
// space themselves out, rather than each being told it can go now.
func (p *pacer) reserveSlot() time.Duration {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Until the server has told us what the limit is, there's nothing to pace against.
	if !p.havePolicy {
		return 0
	}

	now := p.now()
	earliestSendTime := p.theoreticalArrivalTime - p.burstTolerance

	p.theoreticalArrivalTime = max(p.theoreticalArrivalTime, now) + p.emissionInterval

	if earliestSendTime > now {
		return earliestSendTime - now
	}
	return 0
}

// observeResponse takes in what the server told us about the limit on a response, and resynchronises our schedule
// with it. Either header value may be empty (an older server, or an endpoint the server doesn't rate limit); we keep
// the policy we already knew, and give back the slot the request claimed - see refundSlot.
//
// Call this for every response, including error responses: the server writes these headers before the endpoint
// runs, so a 429 or a 500 from a rate limited endpoint still tells us where we stand.
func (p *pacer) observeResponse(policyHeaderValue string, rateLimitHeaderValue string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	limit, haveLimit := getParameter(policyHeaderValue, "l")
	rate, haveRate := getParameter(policyHeaderValue, "rpm")

	// Only take a policy that makes sense, and only do this work when it's actually new: the policy header is static
	// for as long as we stay on the same policy, so this is the same values on nearly every response, and reapplying
	// them would throw away the schedule we're keeping.
	if haveLimit && haveRate && limit >= 1 && rate > 0 && (!p.havePolicy || limit != p.burstLimit || rate != p.requestsPerMinute) {
		p.burstLimit = math.Floor(limit)
		p.requestsPerMinute = rate
		p.emissionInterval = time.Duration(float64(time.Minute) / rate)

		// One for the request being sent, one held back; the rest of the bucket is ours to burst through.
		allowedBurst := p.burstLimit - 1 - reservedRequests
		if allowedBurst > 0 {
			p.burstTolerance = time.Duration(float64(p.emissionInterval) * allowedBurst)
		} else {
			p.burstTolerance = 0
		}

		// The schedule we were keeping is denominated in the old policy's units, so it means nothing now. The
		// remaining count below, read from this same response, puts us back on track immediately.
		p.theoreticalArrivalTime = p.now()

		p.havePolicy = true
	}

	if !p.havePolicy {
		return
	}

	remaining, haveRemaining := getParameter(rateLimitHeaderValue, "r")
	if !haveRemaining {
		p.refundSlot()
		return
	}

	remaining = min(max(remaining, 0), p.burstLimit)

	// What the server says we've spent, expressed as a point in time the way we keep it. We only ever move the
	// schedule later, never earlier: responses can arrive out of order, and an older one arriving last reports a
	// fuller bucket than the truth, which would otherwise talk us into speeding up.
	spent := time.Duration(float64(p.emissionInterval) * (p.burstLimit - remaining))
	p.theoreticalArrivalTime = max(p.theoreticalArrivalTime, p.now()+spent)
}

// refundSlot gives back the slot a request claimed, because the response told us the request never cost the limiter
// anything: no remaining count means either the endpoint isn't rate limited (the server opts a handful of them out,
// and puts others on limiters that send no advisory headers) or the request never reached the limiter at all.
//
// Without this, reserveSlot spends schedule on requests the server didn't charge us for, and a caller that mixes
// limited and unlimited endpoints paces its whole request volume against a limit that only governs part of it.
//
// Refunding too much is safe in a way spending too much isn't: the resync in observeResponse only ever raises the
// schedule towards what the server says is left, so the next response that does carry a count puts the floor
// straight back under us. Spending too much has no such correction, because nothing lowers the schedule except time
// passing.
//
// The caller must hold p.mu.
func (p *pacer) refundSlot() {
	// Only ever give back time we're actually holding. A schedule that has already fallen behind the clock isn't
	// pacing anything, and pulling it further into the past would bank credit for a future burst.
	now := p.now()
	if p.theoreticalArrivalTime > now {
		p.theoreticalArrivalTime = max(p.theoreticalArrivalTime-p.emissionInterval, now)
	}
}

// numberPattern is the plain decimal numbers we accept as a parameter value. strconv.ParseFloat on its own is more
// forgiving than we want here: it reads "Inf", "NaN" and hex floats such as "0x1p4".
var numberPattern = regexp.MustCompile(`^[-+]?(\d+(\.\d*)?|\.\d+)([eE][-+]?\d+)?$`)

// getParameter reads one parameter out of a `name=value;name=value` header, e.g. `l` from `l=200;rpm=600`. Returns
// false if the header is empty or has no such parameter, and tolerates parameters we don't know about so the server
// can add more without breaking us.
func getParameter(headerValue string, name string) (float64, bool) {
	if headerValue == "" {
		return 0, false
	}

	for _, parameter := range strings.Split(headerValue, ";") {
		key, rawValue, found := strings.Cut(parameter, "=")
		if !found {
			continue
		}

		if !strings.EqualFold(strings.TrimSpace(key), name) {
			continue
		}

		rawValue = strings.TrimSpace(rawValue)
		if !numberPattern.MatchString(rawValue) {
			return 0, false
		}

		value, err := strconv.ParseFloat(rawValue, 64)
		if err != nil {
			return 0, false
		}
		return value, true
	}

	return 0, false
}
