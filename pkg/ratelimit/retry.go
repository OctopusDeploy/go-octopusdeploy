package ratelimit

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// retryAfterHeaderName is the standard header the server sends with an HTTP 429, saying how many seconds to wait.
const retryAfterHeaderName = "Retry-After"

// jitterFraction is how much random jitter we add on top of a wait, as a fraction of that wait. Everything the limiter
// rejects in a burst is told to wait the same number of seconds, so without this they all come back at the same
// instant, and the ones that are rejected again go on moving in lockstep for the rest of their retries.
const jitterFraction = 0.25

// retryDelay works out how long we should wait before retrying, based on the Retry-After value the server sent us and
// how many times we have already retried. Returns false if we shouldn't retry because the server asked us to wait
// longer than we're prepared to.
//
// attempt is 0 for the first retry and counts up from there. The wait doubles with each attempt: the server's
// Retry-After tells us how long until it has capacity for one more request, not how long until it has capacity for
// everything it just rejected, so when a whole batch of requests is rate limited at once, retrying them all at that
// same interval only lets a handful through per round. Backing off gives the limiter time to refill for the whole
// queue instead of just the front of it.
//
// random is a value in [0, 1) that decides how much of the jitter allowance we use.
func retryDelay(retryAfterHeaderValue string, options Options, attempt int, random float64) (time.Duration, bool) {
	baseDelay, ok := parseRetryAfterHeader(retryAfterHeaderValue)
	if !ok {
		baseDelay = options.retryDefaultDelay()
	}

	maxDelay := options.retryMaxDelay()

	// A Retry-After of zero means we can go again straight away.
	if baseDelay < 0 {
		baseDelay = 0
	}

	// The server asked for longer than we are willing to wait, so give up now rather than backing off into it.
	if baseDelay > maxDelay {
		return 0, false
	}

	delay := baseDelay
	for i := 0; i < attempt && delay < maxDelay; i++ {
		delay *= 2
	}
	delay = min(delay, maxDelay)

	// Jitter always lands after the delay, never before it: the server told us the earliest it wants us back.
	jitter := min(time.Duration(float64(delay)*jitterFraction), maxDelay-delay)
	return delay + time.Duration(random*float64(jitter)), true
}

var retryAfterPattern = regexp.MustCompile(`^-?\d+$`)

// parseRetryAfterHeader parses a raw Retry-After header value into a delay, or returns false when the header is
// missing or we can't make sense of it, in which case the caller falls back to the default delay.
//
// Retry-After may also be an HTTP date, but we deliberately don't read it as one, the same as the .NET client and the
// web portal. Converting a date into a delay leans on the local clock, which we don't control and which can be minutes
// out: skew one way would have us retry with no delay at all, exactly the hammering the rate limiter exists to stop,
// and skew the other way would push us past the max delay and silently stop us retrying. The Octopus Server always
// sends a number of seconds, so this only comes up if something in between rewrites it.
func parseRetryAfterHeader(headerValue string) (time.Duration, bool) {
	headerValue = strings.TrimSpace(headerValue)
	if !retryAfterPattern.MatchString(headerValue) {
		return 0, false
	}

	// On overflow ParseInt saturates and reports ErrRange; the saturated value is what we want, as either way it's far
	// outside anything we would wait for.
	seconds, err := strconv.ParseInt(headerValue, 10, 64)
	if err != nil && !errors.Is(err, strconv.ErrRange) {
		return 0, false
	}

	// Clamp before converting, so a huge Retry-After can't overflow time.Duration and wrap around to a short wait.
	const maxSeconds = int64(math.MaxInt64 / time.Second)
	seconds = min(max(seconds, -maxSeconds), maxSeconds)
	return time.Duration(seconds) * time.Second, true
}
