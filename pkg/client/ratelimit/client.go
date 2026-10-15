package ratelimit

import "net/http"

// NewClient returns a copy of base that obeys the rate limits of Octopus Server.
// The copy sends requests through a Transport that wraps the transport of base.
// If base is nil, NewClient uses a new http.Client. NewClient does not change base.
func NewClient(base *http.Client, options Options) *http.Client {
	if base == nil {
		base = &http.Client{}
	}

	client := *base
	client.Transport = NewTransport(base.Transport, options)
	return &client
}
