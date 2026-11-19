package ratelimit

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// jitterFraction is the maximum random jitter that is added to a retry delay, as a fraction of the delay.
// The limiter tells all of the requests that it rejects in a burst to wait the same time. Without jitter, they
// all come back at the same time, and the requests that it rejects again stay together for all of their retries.
const jitterFraction = 0.25

// retryAfterPattern matches a Retry-After value that is a whole number of seconds.
var retryAfterPattern = regexp.MustCompile(`^\s*-?\d+\s*$`)

// retryDelay returns the time to wait before a retry, from the Retry-After header value that the server sent and
// the number of retries that were already made. It returns false if the client must not retry, because the server
// tells it to wait longer than options.MaxDelay.
//
// attempt is 0 for the first retry. The delay doubles with each attempt. Retry-After tells the time until the server
// can accept one more request, not the time until it can accept all of the requests that it rejected. When the
// server rejects a burst of requests, a retry of all of them after the same delay lets only some of them through.
// The backoff gives the limiter time to refill for all of the requests.
//
// random is a value in the range [0, 1) that sets the jitter.
//
// The retry delay is the same as the getRateLimitRetryDelay function of the TypeScript client.
func retryDelay(retryAfterHeaderValue string, options Options, attempt int, random float64) (time.Duration, bool) {
	baseDelay, found := parseRetryAfterHeader(retryAfterHeaderValue)
	if !found {
		baseDelay = options.DefaultDelay
	}

	// A Retry-After of 0 or less tells the client that it can retry immediately.
	baseDelay = max(baseDelay, 0)

	// The server tells the client to wait longer than it is permitted to wait, thus do not retry.
	if baseDelay > options.MaxDelay {
		return 0, false
	}

	delay := math.Min(float64(baseDelay)*math.Pow(2, float64(attempt)), float64(options.MaxDelay))

	// The jitter only makes the delay longer, never shorter: the server told the earliest time to retry.
	jitter := math.Min(delay*jitterFraction, float64(options.MaxDelay)-delay)
	return time.Duration(delay + random*jitter), true
}

// parseRetryAfterHeader returns the delay in a Retry-After header value. It returns false if the value is empty or
// is not a whole number of seconds. Then the caller uses the default delay.
//
// Retry-After can also be an HTTP date, but parseRetryAfterHeader does not read dates, the same as the C# and
// TypeScript clients. A delay from a date uses the clock of the client, which can be wrong by minutes. If the clock
// is wrong in one direction, the client retries without a delay. If it is wrong in the other direction, the delay
// is longer than MaxDelay and the client does not retry. Octopus Server always sends a number of seconds.
func parseRetryAfterHeader(headerValue string) (time.Duration, bool) {
	if !retryAfterPattern.MatchString(headerValue) {
		return 0, false
	}

	// The pattern only matches whole numbers, thus ParseInt can only fail when the value is out of range.
	// Then ParseInt returns the largest or smallest int64, which the checks below handle.
	seconds, _ := strconv.ParseInt(strings.TrimSpace(headerValue), 10, 64)

	// Prevent an overflow. A delay that is too long is longer than MaxDelay, and a negative delay is 0.
	maxSeconds := int64(math.MaxInt64 / time.Second)
	if seconds > maxSeconds {
		return time.Duration(math.MaxInt64), true
	}
	if seconds < -maxSeconds {
		return 0, true
	}

	return time.Duration(seconds) * time.Second, true
}
