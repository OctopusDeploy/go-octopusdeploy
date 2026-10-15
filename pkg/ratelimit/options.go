// Package ratelimit keeps the client's requests inside the rate limit the Octopus Server enforces.
//
// It does two things, both of which mirror the .NET client and the web portal:
//   - Retrying: requests the server rejects with HTTP 429 (Too Many Requests) are retried after the delay in the
//     response's Retry-After header, backing off with each attempt. On by default.
//   - Pacing: the client spaces its requests out using the advisory Octopus-RateLimit-Policy and Octopus-RateLimit
//     headers, so that it stays inside the limit rather than running at it and being rejected. Off by default; see
//     Options.UseRateLimitHeaders.
package ratelimit

import "time"

const (
	defaultRetryCount        = 5
	defaultRetryDefaultDelay = 2 * time.Second
	defaultRetryMaxDelay     = 30 * time.Second
)

// Options configures how the client deals with the server's rate limiting. The zero value gives the defaults: 429s
// are retried, and requests are not paced.
type Options struct {
	// RetryCount is the maximum number of times a request will be retried after the server rejects it with HTTP 429
	// (Too Many Requests) because it hit a rate limit. Zero means the default of 5. Set it to a negative number to
	// turn off retrying, in which case the HTTP 429 surfaces to the caller as an error.
	RetryCount int

	// RetryDefaultDelay is how long to wait before retrying an HTTP 429 response that has no Retry-After header,
	// before backing off. Zero means the default of 2 seconds.
	RetryDefaultDelay time.Duration

	// RetryMaxDelay is the longest we are willing to wait between retries. If the server's Retry-After asks us to wait
	// longer than this, we stop retrying and the HTTP 429 surfaces to the caller as an error. Zero means the default of
	// 30 seconds.
	RetryMaxDelay time.Duration

	// UseRateLimitHeaders makes the client pace itself against the advisory Octopus-RateLimit-Policy and
	// Octopus-RateLimit headers the server sends, spacing requests out so that it stays inside the limit instead of
	// running at it and being rejected with HTTP 429. Off by default, because it makes the client wait on its own
	// initiative, which a caller that would rather see the HTTP 429 doesn't want.
	UseRateLimitHeaders bool
}

func (o Options) retryCount() int {
	if o.RetryCount < 0 {
		return 0
	}
	if o.RetryCount == 0 {
		return defaultRetryCount
	}
	return o.RetryCount
}

func (o Options) retryDefaultDelay() time.Duration {
	if o.RetryDefaultDelay <= 0 {
		return defaultRetryDefaultDelay
	}
	return o.RetryDefaultDelay
}

func (o Options) retryMaxDelay() time.Duration {
	if o.RetryMaxDelay <= 0 {
		return defaultRetryMaxDelay
	}
	return o.RetryMaxDelay
}
