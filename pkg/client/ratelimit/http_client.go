package ratelimit

import "net/http"

// NewHttpClient returns a copy of base that obeys the rate limits of Octopus Server.
// The copy sends requests through a Transport that wraps the transport of base.
// If base is nil, NewHttpClient uses a new http.Client. NewHttpClient does not change base.
//
// The Timeout of base includes the time that the Transport waits before a retry or a paced request.
// If Timeout is less than MaxDelay, a request can fail with a timeout while the Transport waits.
func NewHttpClient(base *http.Client, options Options) *http.Client {
	if base == nil {
		base = &http.Client{}
	}

	client := *base
	client.Transport = NewTransport(base.Transport, options)
	return &client
}
