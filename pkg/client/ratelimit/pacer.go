package ratelimit

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// reservedRequests is the number of requests of the allowance that the pacer does not spend. The server tells
// the remaining requests as a whole number and its bucket refills on a timer, thus the arithmetic of the server
// and the pacer can be different by a fraction of a request. When the pacer stops one request before the limit,
// that difference causes a small wait and not a 429 response.
const reservedRequests = 1

// numberPattern matches the numbers that the advisory headers can contain. strconv.ParseFloat also accepts
// values such as "0x10", "Inf" and "NaN", which the server does not send.
var numberPattern = regexp.MustCompile(`^[-+]?(\d+(\.\d*)?|\.\d+)([eE][-+]?\d+)?$`)

// pacer keeps requests in the rate limit that the server advertises in the PolicyHeader and RateLimitHeader
// headers. Thus the client spreads its requests out and does not send them until the server rejects them.
//
// The pacer uses GCRA (the Generic Cell Rate Algorithm), which has the same behaviour as the token bucket of the
// server. The pacer keeps one value, the theoretical arrival time (TAT). The TAT is the time when the requests
// that the pacer already let go are paid for. A request can go when the TAT is in the burst tolerance from now,
// and each request moves the TAT one emission interval later. Thus the client can send a burst of requests and
// then continues at the steady rate of the policy.
//
// Each response sets the TAT again from the remaining requests that the server tells. Thus the pacer corrects
// itself when something else spends the same allowance. The pacer only moves the TAT later, never earlier,
// thus an old response that arrives after a newer one cannot make the client go faster.
//
// The pacer is the same as the RateLimitPacer of the C# and TypeScript clients.
//
// A pacer is safe for concurrent use by multiple goroutines.
type pacer struct {
	// now returns a monotonic time. Tests replace it.
	now func() time.Duration

	mutex                  sync.Mutex
	havePolicy             bool
	burstLimit             int
	requestsPerMinute      float64
	emissionInterval       time.Duration
	burstTolerance         time.Duration
	theoreticalArrivalTime time.Duration
}

// newPacer returns a pacer that uses the monotonic clock of the time package.
func newPacer() *pacer {
	start := time.Now()
	return newPacerWithClock(func() time.Duration {
		return time.Since(start)
	})
}

// newPacerWithClock returns a pacer that uses the clock now.
func newPacerWithClock(now func() time.Duration) *pacer {
	return &pacer{now: now}
}

// reserveSlot takes the next slot in the schedule and returns the time to wait before the request is sent.
// The slot is taken also when the caller does not wait. Thus requests that start at the same time get
// consecutive slots, and they do not all go at the same time.
func (p *pacer) reserveSlot() time.Duration {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// Until the server tells the policy, there is no limit to pace against.
	if !p.havePolicy {
		return 0
	}

	now := p.now()
	earliestSendTime := p.theoreticalArrivalTime - p.burstTolerance

	p.theoreticalArrivalTime = max(p.theoreticalArrivalTime, now) + p.emissionInterval

	return max(earliestSendTime-now, 0)
}

// observeResponse reads the advisory headers of a response and sets the schedule again from them.
// Both values can be empty: an older server, or an endpoint that the server does not rate limit.
// Then the pacer keeps the policy that it has, and gives back the slot of the request (see refundSlot).
//
// Call observeResponse for each response, also for error responses. The server writes these headers before
// the endpoint runs, thus a 429 or a 500 response from a rate limited endpoint also tells the remaining requests.
func (p *pacer) observeResponse(policyHeaderValue string, rateLimitHeaderValue string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	limit, haveLimit := getParameter(policyHeaderValue, "l")
	rate, haveRate := getParameter(policyHeaderValue, "rpm")

	// Use only a policy that is valid, and only when it changes. The policy header is the same on almost all
	// responses, and if the pacer applies it again, it loses the schedule that it keeps.
	if haveLimit && haveRate && limit >= 1 && rate > 0 &&
		(!p.havePolicy || limit != float64(p.burstLimit) || rate != p.requestsPerMinute) {
		p.burstLimit = int(limit)
		p.requestsPerMinute = rate
		p.emissionInterval = time.Duration(float64(time.Minute) / rate)

		// One request for the request that is sent, and one request that the pacer keeps back.
		// The remaining requests of the bucket are the burst.
		allowedBurst := p.burstLimit - 1 - reservedRequests
		p.burstTolerance = 0
		if allowedBurst > 0 {
			p.burstTolerance = p.emissionInterval * time.Duration(allowedBurst)
		}

		// The old schedule uses the units of the old policy, thus it is not correct now.
		// The remaining requests in this response set the schedule again below.
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

	remaining = min(max(remaining, 0), float64(p.burstLimit))

	// The allowance that the server tells is spent, as a time in the schedule. The pacer only moves the
	// schedule later, never earlier: responses can arrive out of order, and an older response that arrives
	// last tells a bucket that is fuller than it is.
	spent := time.Duration(float64(p.emissionInterval) * (float64(p.burstLimit) - remaining))
	p.theoreticalArrivalTime = max(p.theoreticalArrivalTime, p.now()+spent)
}

// refundSlot gives back the slot that a request took, because the response tells that the request did not cost
// the limiter anything. A response without a remaining count is from an endpoint that the server does not rate
// limit, or the request did not get to the limiter.
//
// Without a refund, the pacer spends the schedule on requests that the server did not count. Then a client that
// sends requests to limited and unlimited endpoints paces all of its requests against a limit that only applies
// to some of them.
//
// A refund that is too large is safe: the next response that has a remaining count moves the schedule back.
// A schedule that is too small has no correction, because only time moves it earlier.
//
// The caller must hold the mutex.
func (p *pacer) refundSlot() {
	// Only give back time that the schedule holds. A schedule that is already behind the clock does not pace
	// anything, and if the pacer moves it further into the past, it gives credit for a burst that is not earned.
	now := p.now()
	if p.theoreticalArrivalTime > now {
		p.theoreticalArrivalTime = max(p.theoreticalArrivalTime-p.emissionInterval, now)
	}
}

// getParameter reads one name=value parameter from a header such as "l=200;rpm=600". The name is not case-sensitive.
// It returns false if the header does not have the parameter, or if the value is not a number.
// It ignores parameters that it does not know, thus the server can add parameters without a change to the client.
func getParameter(headerValue string, name string) (float64, bool) {
	for parameter := range strings.SplitSeq(headerValue, ";") {
		parameterName, rawValue, found := strings.Cut(parameter, "=")
		if !found {
			continue
		}

		if !strings.EqualFold(strings.TrimSpace(parameterName), name) {
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
