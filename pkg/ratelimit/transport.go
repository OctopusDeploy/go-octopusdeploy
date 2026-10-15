package ratelimit

import (
	"context"
	"io"
	"math/rand/v2"
	"net/http"
	"time"
)

// NewTransport wraps base (or http.DefaultTransport, if base is nil) so that requests going through it are retried
// when the server rejects them with HTTP 429, and optionally paced against the server's advisory rate limit headers.
//
// The Octopus client constructors in pkg/client already do this for you; only use this directly if you're making
// requests with your own http.Client. Don't wrap a transport twice, or retries will multiply.
//
// Every request made through the returned transport shares one view of the rate limit allowance, because they are
// all spending the same one. Requests whose body can't be replayed (a body with no http.Request.GetBody, such as a
// streamed package upload) are still paced, but are sent only once, and an HTTP 429 is passed back to the caller.
func NewTransport(base http.RoundTripper, options Options) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}

	t := &transport{
		base:    base,
		options: options,
		sleep:   sleepContext,
		random:  rand.Float64,
	}
	if options.UseRateLimitHeaders {
		t.pacer = newPacer()
	}
	return t
}

type transport struct {
	base    http.RoundTripper
	options Options

	// pacer is nil when we aren't pacing.
	pacer *pacer

	// Replaceable in tests.
	sleep  func(ctx context.Context, d time.Duration) error
	random func() float64
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	retryCount := t.options.retryCount()
	if !canReplay(req) {
		retryCount = 0
	}

	attemptRequest := req
	for attempt := 0; ; attempt++ {
		// Wait our turn if we're going faster than the server's rate limit allows. Retries are paced too: a request
		// the limiter has already rejected is the last one that should go straight back at it.
		if err := t.waitForPacing(attemptRequest); err != nil {
			// A RoundTripper must close the request body, even on error. On a retry, the base transport already
			// closed the previous attempt's body and we haven't made this attempt's one yet.
			if attempt == 0 && req.Body != nil {
				_ = req.Body.Close()
			}
			return nil, err
		}

		response, err := t.base.RoundTrip(attemptRequest)
		if err != nil {
			return nil, err
		}

		if t.pacer != nil {
			t.pacer.observeResponse(response.Header.Get(PolicyHeaderName), response.Header.Get(RateLimitHeaderName))
		}

		if response.StatusCode != http.StatusTooManyRequests || attempt >= retryCount {
			return response, nil
		}

		delay, ok := retryDelay(response.Header.Get(retryAfterHeaderName), t.options, attempt, t.random())
		if !ok {
			return response, nil
		}

		nextRequest, err := cloneForRetry(req)
		if err != nil {
			// We can't rebuild the body, so let the caller see the 429 rather than an error of our own making.
			return response, nil
		}

		discard(response)

		if delay > 0 {
			if err := t.sleep(req.Context(), delay); err != nil {
				if nextRequest.Body != nil {
					_ = nextRequest.Body.Close()
				}
				return nil, err
			}
		}

		attemptRequest = nextRequest
	}
}

func (t *transport) waitForPacing(req *http.Request) error {
	if t.pacer == nil {
		return nil
	}

	delay := t.pacer.reserveSlot()
	if delay <= 0 {
		return nil
	}
	return t.sleep(req.Context(), delay)
}

// canReplay reports whether we can send the request again: it either has no body, or it has a GetBody to make a fresh
// copy of it. http.NewRequest sets GetBody for in-memory bodies; anything else is potentially a large or unrewindable
// stream.
func canReplay(req *http.Request) bool {
	return req.Body == nil || req.Body == http.NoBody || req.GetBody != nil
}

func cloneForRetry(req *http.Request) (*http.Request, error) {
	clone := req.Clone(req.Context())
	if req.Body != nil && req.Body != http.NoBody {
		body, err := req.GetBody()
		if err != nil {
			return nil, err
		}
		clone.Body = body
	}
	return clone, nil
}

// discard drains and closes a response we're not passing on, so its connection can be reused.
func discard(response *http.Response) {
	_, _ = io.Copy(io.Discard, response.Body)
	_ = response.Body.Close()
}

// sleepContext waits for d, or until ctx is done, in which case it returns ctx's error so the caller can give up
// rather than sending a request nobody is waiting for any more.
func sleepContext(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
