// Package ratelimit contains an http.RoundTripper that obeys the rate limits of Octopus Server.
//
// See https://octopus.com/docs/administration/managing-infrastructure/rate-limiting
package ratelimit

import (
	"context"
	"io"
	"math/rand/v2"
	"net/http"
	"time"
)

const (
	// PolicyHeader is the advisory header that tells the burst limit and the requests per minute.
	// For example: "l=200;rpm=600".
	PolicyHeader = "Octopus-RateLimit-Policy"

	// RateLimitHeader is the advisory header that tells the remaining requests and the seconds until all of the
	// requests are available again. For example: "r=7;t=18".
	RateLimitHeader = "Octopus-RateLimit"

	retryAfterHeader = "Retry-After"
)

// Options are the settings for a Transport.
type Options struct {
	// RetryCount is the maximum number of times that the Transport sends a request again after a 429 response.
	// If RetryCount is 0 or less, the Transport does not retry, and the caller gets the 429 response.
	RetryCount int

	// DefaultDelay is the time to wait before a retry when the 429 response has no Retry-After header.
	// If DefaultDelay is 0 or less, the Transport uses the default of 2 seconds from DefaultOptions.
	DefaultDelay time.Duration

	// MaxDelay is the maximum time to wait before a retry. If the server tells the Transport to wait longer,
	// the Transport does not retry, and the caller gets the 429 response.
	// If MaxDelay is 0 or less, the Transport uses the default of 30 seconds from DefaultOptions.
	MaxDelay time.Duration

	// UseRateLimitHeaders makes the Transport pace its requests with the Octopus-RateLimit-Policy and
	// Octopus-RateLimit headers that the server sends. Then the Transport stays in the rate limit and
	// prevents most 429 responses.
	UseRateLimitHeaders bool
}

// DefaultOptions returns the recommended settings for a Transport.
func DefaultOptions() Options {
	return Options{
		RetryCount:          5,
		DefaultDelay:        2 * time.Second,
		MaxDelay:            30 * time.Second,
		UseRateLimitHeaders: true,
	}
}

// Transport is an http.RoundTripper that obeys the rate limits of Octopus Server.
//
// When the server returns 429 (Too Many Requests), the Transport waits and sends the request again.
// The server does not process a request that it rejects, thus a retry is safe for all HTTP methods.
// The Transport retries a request only if it can read the request body again (the body is empty, or
// http.Request.GetBody is set). The delay doubles with each retry, and has random jitter.
//
// When UseRateLimitHeaders is true, the Transport also paces each request, retries included, with the advisory
// headers that the server sends. The pacer uses GCRA (the Generic Cell Rate Algorithm), the same as the C# and
// TypeScript clients.
//
// Use NewTransport to create a Transport. A Transport is safe for concurrent use by multiple goroutines.
// Do not copy a Transport after its first use.
type Transport struct {
	base    http.RoundTripper
	options Options

	// pacer is nil when UseRateLimitHeaders is false.
	pacer *pacer
}

var _ http.RoundTripper = (*Transport)(nil)

// NewTransport returns a Transport that sends requests with base. If base is nil, the Transport uses http.DefaultTransport.
// If options.DefaultDelay or options.MaxDelay is 0 or less, the Transport uses the value from DefaultOptions.
func NewTransport(base http.RoundTripper, options Options) *Transport {
	if base == nil {
		base = http.DefaultTransport
	}

	if options.DefaultDelay <= 0 {
		options.DefaultDelay = DefaultOptions().DefaultDelay
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = DefaultOptions().MaxDelay
	}

	transport := &Transport{
		base:    base,
		options: options,
	}

	if options.UseRateLimitHeaders {
		transport.pacer = newPacer()
	}

	return transport
}

// RoundTrip implements http.RoundTripper.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	for attempt := 0; ; attempt++ {
		// The pacer also paces retries: a request that the limiter rejected must not go back to it immediately.
		if t.pacer != nil {
			if err := wait(req.Context(), t.pacer.reserveSlot()); err != nil {
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

		if t.pacer != nil {
			t.pacer.observeResponse(resp.Header.Get(PolicyHeader), resp.Header.Get(RateLimitHeader))
		}

		if resp.StatusCode != http.StatusTooManyRequests || attempt >= t.options.RetryCount || !canRetry(req) {
			return resp, nil
		}

		delay, ok := retryDelay(resp.Header.Get(retryAfterHeader), t.options, attempt, rand.Float64())
		if !ok {
			return resp, nil
		}

		closeResponse(resp)
		if err := wait(req.Context(), delay); err != nil {
			return nil, err
		}
	}
}

// Base returns the http.RoundTripper that the Transport uses to send requests.
func (t *Transport) Base() http.RoundTripper {
	return t.base
}

// CloseIdleConnections closes the idle connections of the base transport, if the base transport supports it.
// http.Client.CloseIdleConnections calls this method.
func (t *Transport) CloseIdleConnections() {
	if closer, ok := t.base.(interface{ CloseIdleConnections() }); ok {
		closer.CloseIdleConnections()
	}
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

// wait stops for the time in delay, or until ctx is cancelled. It returns the error of ctx if ctx is already
// cancelled, also when delay is 0, thus the Transport does not send a request that nobody waits for.
func wait(ctx context.Context, delay time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}

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
