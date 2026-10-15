package ratelimit

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const testURL = "https://octopus.example/api"

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// fakeServer returns the responses in order and records each request that it gets.
// After the last response, it returns the last response again.
type fakeServer struct {
	mutex     sync.Mutex
	responses []func() *http.Response
	requests  []*http.Request
	bodies    []string
}

func (s *fakeServer) RoundTrip(req *http.Request) (*http.Response, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	body := ""
	if req.Body != nil {
		content, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		body = string(content)
	}

	s.requests = append(s.requests, req)
	s.bodies = append(s.bodies, body)

	index := min(len(s.requests)-1, len(s.responses)-1)
	return s.responses[index](), nil
}

func (s *fakeServer) requestCount() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return len(s.requests)
}

func newResponse(statusCode int, headers map[string]string) func() *http.Response {
	return func() *http.Response {
		header := http.Header{}
		for name, value := range headers {
			header.Set(name, value)
		}

		return &http.Response{
			StatusCode: statusCode,
			Header:     header,
			Body:       io.NopCloser(strings.NewReader("")),
		}
	}
}

func testOptions() Options {
	return Options{
		MaxRetries:         3,
		MaxWait:            time.Second,
		UseAdvisoryHeaders: true,
	}
}

func newGetRequest(t *testing.T) *http.Request {
	req, err := http.NewRequest(http.MethodGet, testURL, nil)
	require.NoError(t, err)
	return req
}

func TestDefaultOptions(t *testing.T) {
	options := DefaultOptions()

	require.Equal(t, 3, options.MaxRetries)
	require.Equal(t, 60*time.Second, options.MaxWait)
	require.True(t, options.UseAdvisoryHeaders)
}

func TestNewTransportWithNilBaseUsesDefaultTransport(t *testing.T) {
	transport := NewTransport(nil, testOptions())

	require.Equal(t, http.DefaultTransport, transport.base)
}

func TestRoundTripReturnsSuccessfulResponse(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{newResponse(http.StatusOK, nil)}}
	transport := NewTransport(server, testOptions())

	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, 1, server.requestCount())
}

func TestRoundTripDoesNotRetryOtherErrorResponses(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{newResponse(http.StatusServiceUnavailable, nil)}}
	transport := NewTransport(server, testOptions())

	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusServiceUnavailable, resp.StatusCode)
	require.Equal(t, 1, server.requestCount())
}

func TestRoundTripRetriesAfterTooManyRequests(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusTooManyRequests, map[string]string{"Retry-After": "0"}),
		newResponse(http.StatusOK, nil),
	}}
	transport := NewTransport(server, testOptions())

	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, 2, server.requestCount())
}

func TestRoundTripStopsAfterMaxRetries(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusTooManyRequests, map[string]string{"Retry-After": "0"}),
	}}
	transport := NewTransport(server, testOptions())

	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	require.Equal(t, 4, server.requestCount())
}

func TestRoundTripDoesNotRetryWhenMaxRetriesIsZero(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusTooManyRequests, map[string]string{"Retry-After": "0"}),
	}}
	options := testOptions()
	options.MaxRetries = 0
	transport := NewTransport(server, options)

	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	require.Equal(t, 1, server.requestCount())
}

func TestRoundTripDoesNotRetryWhenRetryAfterIsLongerThanMaxWait(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusTooManyRequests, map[string]string{"Retry-After": "120"}),
		newResponse(http.StatusOK, nil),
	}}
	transport := NewTransport(server, testOptions())

	start := time.Now()
	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	require.Equal(t, 1, server.requestCount())
	require.Less(t, time.Since(start), time.Second)
}

func TestRoundTripWaitsForRetryAfter(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusTooManyRequests, map[string]string{"Retry-After": "1"}),
		newResponse(http.StatusOK, nil),
	}}
	options := testOptions()
	options.MaxWait = 2 * time.Second
	transport := NewTransport(server, options)

	start := time.Now()
	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.GreaterOrEqual(t, time.Since(start), time.Second)
}

func TestRoundTripUsesBackoffCappedAtMaxWaitWhenRetryAfterIsMissing(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusTooManyRequests, nil),
		newResponse(http.StatusOK, nil),
	}}
	options := testOptions()
	options.MaxWait = 50 * time.Millisecond
	transport := NewTransport(server, options)

	start := time.Now()
	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, 2, server.requestCount())
	require.GreaterOrEqual(t, time.Since(start), 50*time.Millisecond)
	require.Less(t, time.Since(start), 500*time.Millisecond)
}

func TestRoundTripSendsTheSameBodyOnRetry(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusTooManyRequests, map[string]string{"Retry-After": "0"}),
		newResponse(http.StatusOK, nil),
	}}
	transport := NewTransport(server, testOptions())

	req, err := http.NewRequest(http.MethodPost, testURL, bytes.NewReader([]byte(`{"Name":"test"}`)))
	require.NoError(t, err)

	resp, err := transport.RoundTrip(req)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, []string{`{"Name":"test"}`, `{"Name":"test"}`}, server.bodies)
	require.Equal(t, http.MethodPost, server.requests[1].Method)
}

func TestRoundTripDoesNotRetryWhenBodyCannotBeReadAgain(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusTooManyRequests, map[string]string{"Retry-After": "0"}),
		newResponse(http.StatusOK, nil),
	}}
	transport := NewTransport(server, testOptions())

	req, err := http.NewRequest(http.MethodPost, testURL, io.NopCloser(strings.NewReader("package content")))
	require.NoError(t, err)
	require.Nil(t, req.GetBody)

	resp, err := transport.RoundTrip(req)

	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	require.Equal(t, 1, server.requestCount())
}

func TestRoundTripReturnsErrorWhenContextIsCancelledDuringRetryWait(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusTooManyRequests, map[string]string{"Retry-After": "30"}),
		newResponse(http.StatusOK, nil),
	}}
	options := testOptions()
	options.MaxWait = time.Minute
	transport := NewTransport(server, options)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, testURL, nil)
	require.NoError(t, err)

	start := time.Now()
	resp, err := transport.RoundTrip(req)

	require.ErrorIs(t, err, context.DeadlineExceeded)
	require.Nil(t, resp)
	require.Equal(t, 1, server.requestCount())
	require.Less(t, time.Since(start), time.Second)
}

func TestRoundTripReturnsErrorFromBaseTransport(t *testing.T) {
	baseErr := io.ErrUnexpectedEOF
	transport := NewTransport(roundTripFunc(func(req *http.Request) (*http.Response, error) {
		return nil, baseErr
	}), testOptions())

	resp, err := transport.RoundTrip(newGetRequest(t))

	require.ErrorIs(t, err, baseErr)
	require.Nil(t, resp)
}

func TestRoundTripWaitsWhenNoTokensRemain(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusOK, map[string]string{PolicyHeader: "l=200;rpm=600", StateHeader: "r=0;t=18"}),
		newResponse(http.StatusOK, nil),
	}}
	transport := NewTransport(server, testOptions())

	_, err := transport.RoundTrip(newGetRequest(t))
	require.NoError(t, err)

	start := time.Now()
	_, err = transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	// 600 requests per minute gives one token each 100 milliseconds.
	require.GreaterOrEqual(t, time.Since(start), 90*time.Millisecond)
}

func TestRoundTripDoesNotWaitWhenTokensRemain(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusOK, map[string]string{PolicyHeader: "l=200;rpm=1", StateHeader: "r=7;t=18"}),
	}}
	transport := NewTransport(server, testOptions())

	_, err := transport.RoundTrip(newGetRequest(t))
	require.NoError(t, err)

	start := time.Now()
	_, err = transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Less(t, time.Since(start), 50*time.Millisecond)
}

func TestUpdateThrottleClearsWaitWhenTokensAreAvailableAgain(t *testing.T) {
	transport := NewTransport(nil, testOptions())

	transport.updateThrottle(newResponse(http.StatusOK, map[string]string{PolicyHeader: "l=200;rpm=1", StateHeader: "r=0;t=60"})())
	require.Greater(t, transport.throttleDelay(), time.Duration(0))

	transport.updateThrottle(newResponse(http.StatusOK, map[string]string{PolicyHeader: "l=200;rpm=1", StateHeader: "r=5;t=55"})())
	require.LessOrEqual(t, transport.throttleDelay(), time.Duration(0))
}

func TestUpdateThrottleIgnoresResponsesWithoutAdvisoryHeaders(t *testing.T) {
	transport := NewTransport(nil, testOptions())

	transport.updateThrottle(newResponse(http.StatusOK, map[string]string{PolicyHeader: "l=200;rpm=1", StateHeader: "r=0;t=60"})())
	transport.updateThrottle(newResponse(http.StatusOK, nil)())

	require.Greater(t, transport.throttleDelay(), time.Duration(0))
}

func TestRoundTripCapsThrottleWaitAtMaxWait(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusOK, map[string]string{PolicyHeader: "l=200;rpm=1", StateHeader: "r=0;t=60"}),
		newResponse(http.StatusOK, nil),
	}}
	options := testOptions()
	options.MaxWait = 50 * time.Millisecond
	transport := NewTransport(server, options)

	_, err := transport.RoundTrip(newGetRequest(t))
	require.NoError(t, err)

	start := time.Now()
	_, err = transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.GreaterOrEqual(t, time.Since(start), 50*time.Millisecond)
	require.Less(t, time.Since(start), 500*time.Millisecond)
}

func TestRoundTripDoesNotThrottleWhenAdvisoryHeadersAreNotUsed(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusOK, map[string]string{PolicyHeader: "l=200;rpm=1", StateHeader: "r=0;t=60"}),
	}}
	options := testOptions()
	options.UseAdvisoryHeaders = false
	transport := NewTransport(server, options)

	_, err := transport.RoundTrip(newGetRequest(t))
	require.NoError(t, err)

	start := time.Now()
	_, err = transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Less(t, time.Since(start), 50*time.Millisecond)
}

func TestRoundTripReturnsErrorWhenContextIsCancelledDuringThrottleWait(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusOK, map[string]string{PolicyHeader: "l=200;rpm=1", StateHeader: "r=0;t=60"}),
	}}
	options := testOptions()
	options.MaxWait = time.Minute
	transport := NewTransport(server, options)

	_, err := transport.RoundTrip(newGetRequest(t))
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, testURL, nil)
	require.NoError(t, err)

	resp, err := transport.RoundTrip(req)

	require.ErrorIs(t, err, context.DeadlineExceeded)
	require.Nil(t, resp)
	require.Equal(t, 1, server.requestCount())
}

type closeRecorder struct {
	io.Reader
	closed bool
}

func (c *closeRecorder) Close() error {
	c.closed = true
	return nil
}

func TestRoundTripClosesBodyWhenContextIsCancelledBeforeFirstAttempt(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusOK, map[string]string{PolicyHeader: "l=200;rpm=1", StateHeader: "r=0;t=60"}),
	}}
	options := testOptions()
	options.MaxWait = time.Minute
	transport := NewTransport(server, options)

	_, err := transport.RoundTrip(newGetRequest(t))
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	body := &closeRecorder{Reader: strings.NewReader(`{"Name":"test"}`)}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, testURL, body)
	require.NoError(t, err)

	_, err = transport.RoundTrip(req)

	require.ErrorIs(t, err, context.Canceled)
	require.True(t, body.closed)
	require.Equal(t, 1, server.requestCount())
}

func TestParseHeader(t *testing.T) {
	testCases := []struct {
		name     string
		value    string
		expected map[string]int
	}{
		{"Policy", "l=200;rpm=600", map[string]int{"l": 200, "rpm": 600}},
		{"State", "r=7;t=18", map[string]int{"r": 7, "t": 18}},
		{"Spaces", " r = 7 ; t = 18 ", map[string]int{"r": 7, "t": 18}},
		{"Empty", "", map[string]int{}},
		{"MissingValue", "r;t=18", map[string]int{"t": 18}},
		{"NotANumber", "r=abc;t=18", map[string]int{"t": 18}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			require.Equal(t, testCase.expected, parseHeader(testCase.value))
		})
	}
}

func TestRetryAfter(t *testing.T) {
	testCases := []struct {
		name          string
		value         string
		expectedDelay time.Duration
		expectedFound bool
	}{
		{"Seconds", "5", 5 * time.Second, true},
		{"Zero", "0", 0, true},
		{"Missing", "", 0, false},
		{"Negative", "-1", 0, false},
		{"HttpDate", "Wed, 21 Oct 2015 07:28:00 GMT", 0, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resp := newResponse(http.StatusTooManyRequests, map[string]string{"Retry-After": testCase.value})()

			delay, found := retryAfter(resp)

			require.Equal(t, testCase.expectedDelay, delay)
			require.Equal(t, testCase.expectedFound, found)
		})
	}
}

func TestBackoff(t *testing.T) {
	for attempt := range 4 {
		maxDelay := initialBackoff << attempt

		delay := backoff(attempt)

		require.GreaterOrEqual(t, delay, maxDelay/2)
		require.Less(t, delay, maxDelay)
	}
}

func TestBackoffDoesNotOverflowForLargeAttempts(t *testing.T) {
	maxDelay := initialBackoff << maxBackoffShift

	delay := backoff(1000)

	require.GreaterOrEqual(t, delay, maxDelay/2)
	require.Less(t, delay, maxDelay)
}
