package ratelimit

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// fakeServer answers requests with canned responses, in order, and records what it was sent.
type fakeServer struct {
	responses []*http.Response
	requests  []*http.Request
	bodies    []string
}

func (s *fakeServer) RoundTrip(req *http.Request) (*http.Response, error) {
	s.requests = append(s.requests, req)

	body := ""
	if req.Body != nil {
		b, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		_ = req.Body.Close()
		body = string(b)
	}
	s.bodies = append(s.bodies, body)

	if len(s.responses) == 0 {
		return nil, errors.New("fakeServer has run out of responses")
	}
	response := s.responses[0]
	s.responses = s.responses[1:]
	return response, nil
}

// trackedBody lets a test see whether a response body was closed.
type trackedBody struct {
	io.Reader
	closed bool
}

func (b *trackedBody) Close() error {
	b.closed = true
	return nil
}

func response(statusCode int, headers ...string) *http.Response {
	r := &http.Response{
		StatusCode: statusCode,
		Header:     http.Header{},
		Body:       &trackedBody{Reader: strings.NewReader("{}")},
	}
	for i := 0; i+1 < len(headers); i += 2 {
		r.Header.Set(headers[i], headers[i+1])
	}
	return r
}

func tooManyRequests(retryAfter string) *http.Response {
	return response(http.StatusTooManyRequests, retryAfterHeaderName, retryAfter)
}

// newTestTransport returns a transport over server that records how long it was asked to wait, rather than waiting.
func newTestTransport(server *fakeServer, options Options) (*transport, *[]time.Duration) {
	sleeps := &[]time.Duration{}
	t := NewTransport(server, options).(*transport)
	t.random = func() float64 { return 0 }
	t.sleep = func(ctx context.Context, d time.Duration) error {
		*sleeps = append(*sleeps, d)
		return ctx.Err()
	}
	if t.pacer != nil {
		// A clock that never moves, so the pacing delays are exact.
		t.pacer.now = func() time.Duration { return time.Hour }
	}
	return t, sleeps
}

func get(t *testing.T) *http.Request {
	req, err := http.NewRequest(http.MethodGet, "http://octopus.example/api/projects", nil)
	require.NoError(t, err)
	return req
}

func TestTransport_RetriesARateLimitedRequestAfterTheRetryAfterDelay(t *testing.T) {
	server := &fakeServer{responses: []*http.Response{tooManyRequests("3"), response(http.StatusOK)}}
	transport, sleeps := newTestTransport(server, Options{})

	resp, err := transport.RoundTrip(get(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Len(t, server.requests, 2)
	require.Equal(t, []time.Duration{3 * time.Second}, *sleeps)
}

func TestTransport_BacksOffWithEachRetry(t *testing.T) {
	server := &fakeServer{responses: []*http.Response{tooManyRequests("1"), tooManyRequests("1"), tooManyRequests("1"), response(http.StatusOK)}}
	transport, sleeps := newTestTransport(server, Options{})

	resp, err := transport.RoundTrip(get(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, []time.Duration{time.Second, 2 * time.Second, 4 * time.Second}, *sleeps)
}

func TestTransport_ClosesTheRateLimitedResponsesItDoesNotPassOn(t *testing.T) {
	rejected := tooManyRequests("1")
	server := &fakeServer{responses: []*http.Response{rejected, response(http.StatusOK)}}
	transport, _ := newTestTransport(server, Options{})

	_, err := transport.RoundTrip(get(t))

	require.NoError(t, err)
	require.True(t, rejected.Body.(*trackedBody).closed)
}

func TestTransport_PassesTheRateLimitedResponseOnOnceItRunsOutOfRetries(t *testing.T) {
	server := &fakeServer{responses: []*http.Response{tooManyRequests("0"), tooManyRequests("0"), tooManyRequests("0")}}
	transport, _ := newTestTransport(server, Options{RetryCount: 2})

	resp, err := transport.RoundTrip(get(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	require.Len(t, server.requests, 3)
	require.False(t, resp.Body.(*trackedBody).closed, "the caller gets the response to read")
}

func TestTransport_DoesNotRetryWhenRetryingIsTurnedOff(t *testing.T) {
	server := &fakeServer{responses: []*http.Response{tooManyRequests("0")}}
	transport, _ := newTestTransport(server, Options{RetryCount: -1})

	resp, err := transport.RoundTrip(get(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	require.Len(t, server.requests, 1)
}

func TestTransport_DoesNotRetryWhenTheServerAsksUsToWaitLongerThanWeAreWillingTo(t *testing.T) {
	server := &fakeServer{responses: []*http.Response{tooManyRequests("60")}}
	transport, sleeps := newTestTransport(server, Options{})

	resp, err := transport.RoundTrip(get(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	require.Len(t, server.requests, 1)
	require.Empty(t, *sleeps)
}

func TestTransport_DoesNotRetryOtherErrors(t *testing.T) {
	server := &fakeServer{responses: []*http.Response{response(http.StatusServiceUnavailable, retryAfterHeaderName, "1")}}
	transport, _ := newTestTransport(server, Options{})

	resp, err := transport.RoundTrip(get(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusServiceUnavailable, resp.StatusCode)
	require.Len(t, server.requests, 1)
}

func TestTransport_ReplaysTheRequestBodyOnARetry(t *testing.T) {
	server := &fakeServer{responses: []*http.Response{tooManyRequests("0"), response(http.StatusCreated)}}
	transport, _ := newTestTransport(server, Options{})

	req, err := http.NewRequest(http.MethodPost, "http://octopus.example/api/projects", bytes.NewBufferString(`{"Name":"x"}`))
	require.NoError(t, err)
	req.Header.Set("X-Custom", "kept")

	resp, err := transport.RoundTrip(req)

	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	require.Equal(t, []string{`{"Name":"x"}`, `{"Name":"x"}`}, server.bodies)
	require.Equal(t, http.MethodPost, server.requests[1].Method)
	require.Equal(t, "kept", server.requests[1].Header.Get("X-Custom"))
}

func TestTransport_DoesNotRetryARequestWhoseBodyCannotBeReplayed(t *testing.T) {
	// A streamed body, like a package upload, can't be rewound, so the request goes once and the caller sees the 429.
	server := &fakeServer{responses: []*http.Response{tooManyRequests("0")}}
	transport, _ := newTestTransport(server, Options{})

	req, err := http.NewRequest(http.MethodPost, "http://octopus.example/api/packages/raw", io.NopCloser(strings.NewReader("package")))
	require.NoError(t, err)
	require.Nil(t, req.GetBody)

	resp, err := transport.RoundTrip(req)

	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	require.Len(t, server.requests, 1)
}

func TestTransport_GivesUpWhenTheRequestIsCancelledWhileWaitingToRetry(t *testing.T) {
	server := &fakeServer{responses: []*http.Response{tooManyRequests("10")}}
	transport, _ := newTestTransport(server, Options{})
	transport.sleep = sleepContext

	ctx, cancel := context.WithCancel(context.Background())
	req := get(t).WithContext(ctx)
	go cancel()

	_, err := transport.RoundTrip(req)

	require.ErrorIs(t, err, context.Canceled)
	require.Len(t, server.requests, 1)
}

func TestTransport_DoesNotPaceByDefault(t *testing.T) {
	server := &fakeServer{responses: []*http.Response{
		response(http.StatusOK, PolicyHeaderName, testPolicy, RateLimitHeaderName, "r=0;t=10"),
		response(http.StatusOK),
	}}
	transport, sleeps := newTestTransport(server, Options{})

	_, err := transport.RoundTrip(get(t))
	require.NoError(t, err)
	_, err = transport.RoundTrip(get(t))
	require.NoError(t, err)

	require.Nil(t, transport.pacer)
	require.Empty(t, *sleeps)
}

func TestTransport_PacesRequestsAgainstTheAdvisoryHeadersWhenAskedTo(t *testing.T) {
	server := &fakeServer{responses: []*http.Response{
		response(http.StatusOK, PolicyHeaderName, testPolicy, RateLimitHeaderName, "r=0;t=10"),
		response(http.StatusOK),
	}}
	transport, sleeps := newTestTransport(server, Options{UseRateLimitHeaders: true})

	// The first request goes straight out; nothing has told us about a limit yet.
	_, err := transport.RoundTrip(get(t))
	require.NoError(t, err)
	require.Empty(t, *sleeps)

	// The server said the allowance is spent, so the next request waits for it to come back.
	_, err = transport.RoundTrip(get(t))
	require.NoError(t, err)
	require.Equal(t, []time.Duration{twoEmissionIntervals}, *sleeps)
}

func TestTransport_PacesRetriesTooAndLearnsFromRateLimitedResponses(t *testing.T) {
	// The 429 carries the advisory headers as well, so the retry is paced as well as waiting out Retry-After.
	server := &fakeServer{responses: []*http.Response{
		response(http.StatusTooManyRequests, retryAfterHeaderName, "1", PolicyHeaderName, testPolicy, RateLimitHeaderName, "r=0;t=10"),
		response(http.StatusOK),
	}}
	transport, sleeps := newTestTransport(server, Options{UseRateLimitHeaders: true})

	resp, err := transport.RoundTrip(get(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, []time.Duration{time.Second, twoEmissionIntervals}, *sleeps)
}

func TestTransport_ClosesTheRequestBodyWhenCancelledBeforeSending(t *testing.T) {
	server := &fakeServer{responses: []*http.Response{
		response(http.StatusOK, PolicyHeaderName, testPolicy, RateLimitHeaderName, "r=0;t=10"),
	}}
	transport, _ := newTestTransport(server, Options{UseRateLimitHeaders: true})
	_, err := transport.RoundTrip(get(t))
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	body := &trackedBody{Reader: strings.NewReader("x")}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://octopus.example/api/projects", body)
	require.NoError(t, err)

	_, err = transport.RoundTrip(req)

	require.ErrorIs(t, err, context.Canceled)
	require.True(t, body.closed)
	require.Len(t, server.requests, 1, "the cancelled request was never sent")
}
