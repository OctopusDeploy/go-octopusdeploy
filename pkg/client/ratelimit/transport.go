// Package ratelimit contains an http.RoundTripper that obeys the rate limits of Octopus Server.
//
// See https://octopus.com/docs/administration/managing-infrastructure/rate-limiting
package ratelimit

import (
	"context"
	"io"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// PolicyHeader is the advisory header that tells the burst limit and the requests per minute.
	// For example: "l=200;rpm=600".
	PolicyHeader = "Octopus-RateLimit-Policy"

	// StateHeader is the advisory header that tells the remaining tokens and the seconds until all tokens are available again.
	// For example: "r=7;t=18".
	StateHeader = "Octopus-RateLimit"

	retryAfterHeader = "Retry-After"
	initialBackoff   = time.Second
	maxBackoffShift  = 10
)

// Options are the settings for a Transport.
type Options struct {
	// MaxRetries is the maximum number of times that the Transport sends a request again after a 429 response.
	MaxRetries int

	// MaxWait is the maximum time that the Transport waits before it sends a request.
	// If the server tells the Transport to wait longer before a retry, the Transport does not retry. It returns the 429 response.
	MaxWait time.Duration

	// UseAdvisoryHeaders makes the Transport wait before it sends a request when the advisory headers
	// show that no tokens remain. This prevents most 429 responses.
	UseAdvisoryHeaders bool
}

// DefaultOptions returns the recommended settings for a Transport.
func DefaultOptions() Options {
	return Options{
		MaxRetries:         3,
		MaxWait:            60 * time.Second,
		UseAdvisoryHeaders: true,
	}
}

// Transport is an http.RoundTripper that obeys the rate limits of Octopus Server.
//
// When the server returns 429 (Too Many Requests), the Transport waits and sends the request again.
// The server does not process a request that it rejects, thus a retry is safe for all HTTP methods.
// The Transport retries a request only if it can read the request body again (the body is empty, or
// http.Request.GetBody is set).
//
// When UseAdvisoryHeaders is true, the Transport also reads the advisory headers on each response. If no tokens remain,
// the Transport waits for one token before it sends the next request.
//
// Use NewTransport to create a Transport. A Transport is safe for concurrent use by multiple goroutines.
// Do not copy a Transport after its first use.
type Transport struct {
	base    http.RoundTripper
	options Options

	mutex           sync.Mutex
	nextRequestTime time.Time
}

var _ http.RoundTripper = (*Transport)(nil)

// NewTransport returns a Transport that sends requests with base. If base is nil, the Transport uses http.DefaultTransport.
func NewTransport(base http.RoundTripper, options Options) *Transport {
	if base == nil {
		base = http.DefaultTransport
	}

	return &Transport{
		base:    base,
		options: options,
	}
}

// RoundTrip implements http.RoundTripper.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	for attempt := 0; ; attempt++ {
		if t.options.UseAdvisoryHeaders {
			if err := wait(req.Context(), t.throttleDelay()); err != nil {
				// The base transport did not get the request, so the body must be closed here.
				if attempt == 0 && req.Body != nil {
					_ = req.Body.Close()
				}
				return nil, err
			}
		}

		attemptRequest, err := requestForAttempt(req, attempt)
		if err != nil {
			return nil, err
		}

		resp, err := t.base.RoundTrip(attemptRequest)
		if err != nil {
			return nil, err
		}

		if t.options.UseAdvisoryHeaders {
			t.updateThrottle(resp)
		}

		if resp.StatusCode != http.StatusTooManyRequests || attempt >= t.options.MaxRetries || !canRetry(req) {
			return resp, nil
		}

		delay, found := retryAfter(resp)
		if !found {
			delay = min(backoff(attempt), t.options.MaxWait)
		}
		if delay > t.options.MaxWait {
			return resp, nil
		}

		closeResponse(resp)
		if err := wait(req.Context(), delay); err != nil {
			return nil, err
		}
	}
}

// throttleDelay returns the time to wait before the next request, as the advisory headers tell.
func (t *Transport) throttleDelay() time.Duration {
	t.mutex.Lock()
	nextRequestTime := t.nextRequestTime
	t.mutex.Unlock()

	return min(time.Until(nextRequestTime), t.options.MaxWait)
}

// updateThrottle reads the advisory headers on resp. If no tokens remain, it sets the time of the next request
// to the time when one token is available again.
func (t *Transport) updateThrottle(resp *http.Response) {
	state := parseHeader(resp.Header.Get(StateHeader))
	remaining, ok := state["r"]
	if !ok {
		return
	}

	nextRequestTime := time.Time{}
	if remaining <= 0 {
		policy := parseHeader(resp.Header.Get(PolicyHeader))
		if requestsPerMinute := policy["rpm"]; requestsPerMinute > 0 {
			nextRequestTime = time.Now().Add(time.Minute / time.Duration(requestsPerMinute))
		}
	}

	t.mutex.Lock()
	t.nextRequestTime = nextRequestTime
	t.mutex.Unlock()
}

// canRetry returns true if the body of req can be read again.
func canRetry(req *http.Request) bool {
	return req.Body == nil || req.Body == http.NoBody || req.GetBody != nil
}

// requestForAttempt returns req for the first attempt. For a retry, it returns a copy of req with a new body.
func requestForAttempt(req *http.Request, attempt int) (*http.Request, error) {
	if attempt == 0 {
		return req, nil
	}

	retryRequest := req.Clone(req.Context())
	if req.Body != nil && req.Body != http.NoBody {
		body, err := req.GetBody()
		if err != nil {
			return nil, err
		}
		retryRequest.Body = body
	}

	return retryRequest, nil
}

// retryAfter returns the time in the Retry-After header of resp. It returns false if the header is missing or not valid.
func retryAfter(resp *http.Response) (time.Duration, bool) {
	seconds, err := strconv.Atoi(strings.TrimSpace(resp.Header.Get(retryAfterHeader)))
	if err != nil || seconds < 0 {
		return 0, false
	}

	return time.Duration(seconds) * time.Second, true
}

// backoff returns an exponential backoff with jitter for the attempt.
func backoff(attempt int) time.Duration {
	delay := initialBackoff << min(attempt, maxBackoffShift)
	return delay/2 + rand.N(delay/2)
}

// parseHeader parses a header value such as "l=200;rpm=600" into a map of names to integer values.
// It ignores the parts that it cannot parse.
func parseHeader(value string) map[string]int {
	values := map[string]int{}
	for part := range strings.SplitSeq(value, ";") {
		name, number, found := strings.Cut(part, "=")
		if !found {
			continue
		}

		parsedNumber, err := strconv.Atoi(strings.TrimSpace(number))
		if err != nil {
			continue
		}

		values[strings.TrimSpace(name)] = parsedNumber
	}

	return values
}

// wait stops for the time in delay, or until ctx is cancelled.
func wait(ctx context.Context, delay time.Duration) error {
	if delay <= 0 {
		return nil
	}

	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// closeResponse reads all of the response body and closes it, so that the connection can be used again.
func closeResponse(resp *http.Response) {
	_, _ = io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()
}
