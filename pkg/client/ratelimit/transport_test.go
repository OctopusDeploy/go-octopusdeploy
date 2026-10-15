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

// tooManyRequests returns a 429 response with the Retry-After value. If retryAfter is empty, the response has no Retry-After header.
func tooManyRequests(retryAfter string) func() *http.Response {
	if retryAfter == "" {
		return newResponse(http.StatusTooManyRequests, nil)
	}

	return newResponse(http.StatusTooManyRequests, map[string]string{retryAfterHeader: retryAfter})
}

// testOptions are the settings for the transport tests. The delays are short, so that the tests are fast.
func testOptions() Options {
	return Options{
		RetryCount:          3,
		DefaultDelay:        50 * time.Millisecond,
		MaxDelay:            time.Second,
		UseRateLimitHeaders: true,
	}
}

func newGetRequest(t *testing.T) *http.Request {
	req, err := http.NewRequest(http.MethodGet, testURL, nil)
	require.NoError(t, err)
	return req
}

func TestNewTransportWithNilBaseUsesDefaultTransport(t *testing.T) {
	transport := NewTransport(nil, testOptions())

	require.Equal(t, http.DefaultTransport, transport.Base())
}

func TestBaseReturnsTheBaseTransport(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{newResponse(http.StatusOK, nil)}}

	transport := NewTransport(server, testOptions())

	require.Same(t, server, transport.Base())
}

func TestNewTransportUsesTheDefaultDelaysWhenTheDelaysAreNotPositive(t *testing.T) {
	testCases := []struct {
		name  string
		delay time.Duration
	}{
		{"Zero", 0},
		{"Negative", -time.Second},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			transport := NewTransport(nil, Options{RetryCount: 3, DefaultDelay: testCase.delay, MaxDelay: testCase.delay})

			require.Equal(t, DefaultOptions().DefaultDelay, transport.options.DefaultDelay)
			require.Equal(t, DefaultOptions().MaxDelay, transport.options.MaxDelay)
		})
	}
}

func TestNewTransportKeepsPositiveDelays(t *testing.T) {
	transport := NewTransport(nil, Options{RetryCount: 3, DefaultDelay: time.Second, MaxDelay: 5 * time.Second})

	require.Equal(t, time.Second, transport.options.DefaultDelay)
	require.Equal(t, 5*time.Second, transport.options.MaxDelay)
}

func TestNewTransportPacesOnlyWhenUseRateLimitHeadersIsTrue(t *testing.T) {
	options := testOptions()

	options.UseRateLimitHeaders = true
	require.NotNil(t, NewTransport(nil, options).pacer)

	options.UseRateLimitHeaders = false
	require.Nil(t, NewTransport(nil, options).pacer)
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

func TestRoundTripRetriesUntilTheServerAcceptsTheRequest(t *testing.T) {
	// The server rejects the first two requests, then accepts the request.
	server := &fakeServer{responses: []func() *http.Response{
		tooManyRequests("0"),
		tooManyRequests("0"),
		newResponse(http.StatusOK, nil),
	}}
	transport := NewTransport(server, testOptions())

	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, 3, server.requestCount())
}

func TestRoundTripWaitsForRetryAfter(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		tooManyRequests("1"),
		newResponse(http.StatusOK, nil),
	}}
	options := testOptions()
	options.MaxDelay = 2 * time.Second
	transport := NewTransport(server, options)

	start := time.Now()
	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, 2, server.requestCount())
	require.GreaterOrEqual(t, time.Since(start), time.Second)
}

func TestRoundTripWaitsTheDefaultDelayWhenThereIsNoRetryAfterHeader(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		tooManyRequests(""),
		newResponse(http.StatusOK, nil),
	}}
	transport := NewTransport(server, testOptions())

	start := time.Now()
	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, 2, server.requestCount())
	require.GreaterOrEqual(t, time.Since(start), testOptions().DefaultDelay)
}

func TestRoundTripGivesUpAfterTheConfiguredNumberOfRetries(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{tooManyRequests("0")}}
	transport := NewTransport(server, testOptions())

	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	// The first request, and RetryCount (3) retries.
	require.Equal(t, 4, server.requestCount())
}

func TestRoundTripDoesNotRetryWhenRetryCountIsZeroOrLess(t *testing.T) {
	testCases := []struct {
		name       string
		retryCount int
	}{
		{"Zero", 0},
		{"Negative", -1},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := &fakeServer{responses: []func() *http.Response{
				tooManyRequests("0"),
				newResponse(http.StatusOK, nil),
			}}
			options := testOptions()
			options.RetryCount = testCase.retryCount
			transport := NewTransport(server, options)

			resp, err := transport.RoundTrip(newGetRequest(t))

			require.NoError(t, err)
			require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
			require.Equal(t, 1, server.requestCount())
		})
	}
}

func TestRoundTripDoesNotRetryWhenRetryAfterIsLongerThanMaxDelay(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		tooManyRequests("120"),
		newResponse(http.StatusOK, nil),
	}}
	transport := NewTransport(server, testOptions())

	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusTooManyRequests, resp.StatusCode)
	require.Equal(t, 1, server.requestCount())
}

func TestRoundTripSendsTheSameBodyOnRetry(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		tooManyRequests("0"),
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
		tooManyRequests("0"),
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
		tooManyRequests("30"),
		newResponse(http.StatusOK, nil),
	}}
	options := testOptions()
	options.MaxDelay = time.Minute
	transport := NewTransport(server, options)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, testURL, nil)
	require.NoError(t, err)

	resp, err := transport.RoundTrip(req)

	require.ErrorIs(t, err, context.DeadlineExceeded)
	require.Nil(t, resp)
	require.Equal(t, 1, server.requestCount())
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

func TestRoundTripGivesTheRateLimitHeadersToThePacer(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
	}{
		{"SuccessfulResponse", http.StatusOK},
		// The server writes these headers before the endpoint runs, thus error responses also have them.
		{"ErrorResponse", http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			server := &fakeServer{responses: []func() *http.Response{
				newResponse(testCase.statusCode, map[string]string{PolicyHeader: testPolicy, RateLimitHeader: "r=0;t=10"}),
			}}
			transport := NewTransport(server, testOptions())
			testPacer, _ := newTestPacer()
			transport.pacer = testPacer

			_, err := transport.RoundTrip(newGetRequest(t))

			require.NoError(t, err)
			require.Equal(t, 2*testEmissionInterval, testPacer.reserveSlot())
		})
	}
}

func TestRoundTripWaitsForThePacer(t *testing.T) {
	// 600 requests per minute gives one request each 100 milliseconds. No requests remain, thus the
	// next request waits for two requests: the request that it sends, and the request that the pacer keeps back.
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusOK, map[string]string{PolicyHeader: "l=10;rpm=600", RateLimitHeader: "r=0;t=1"}),
	}}
	transport := NewTransport(server, testOptions())

	_, err := transport.RoundTrip(newGetRequest(t))
	require.NoError(t, err)

	start := time.Now()
	_, err = transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.GreaterOrEqual(t, time.Since(start), 150*time.Millisecond)
}

func TestRoundTripPacesRetries(t *testing.T) {
	// The 429 response tells the client to retry immediately, but no requests remain, thus the pacer makes the
	// retry wait. A request that the limiter rejected must not go back to it immediately.
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusTooManyRequests, map[string]string{retryAfterHeader: "0", PolicyHeader: "l=10;rpm=600", RateLimitHeader: "r=0;t=1"}),
		newResponse(http.StatusOK, nil),
	}}
	transport := NewTransport(server, testOptions())

	start := time.Now()
	resp, err := transport.RoundTrip(newGetRequest(t))

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, 2, server.requestCount())
	require.GreaterOrEqual(t, time.Since(start), 150*time.Millisecond)
}

func TestRoundTripReturnsErrorWhenContextIsCancelledDuringPacerWait(t *testing.T) {
	// 1 request per minute and no requests remain, thus the next request must wait for two minutes.
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusOK, map[string]string{PolicyHeader: "l=10;rpm=1", RateLimitHeader: "r=0;t=600"}),
	}}
	transport := NewTransport(server, testOptions())

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
	server := &fakeServer{responses: []func() *http.Response{newResponse(http.StatusOK, nil)}}
	transport := NewTransport(server, testOptions())

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	body := &closeRecorder{Reader: strings.NewReader(`{"Name":"test"}`)}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, testURL, body)
	require.NoError(t, err)

	_, err = transport.RoundTrip(req)

	require.ErrorIs(t, err, context.Canceled)
	require.True(t, body.closed)
	require.Equal(t, 0, server.requestCount())
}

// idleConnectionCloser is a base transport that records the calls to CloseIdleConnections.
type idleConnectionCloser struct {
	closeCount int
}

func (c *idleConnectionCloser) RoundTrip(req *http.Request) (*http.Response, error) {
	return newResponse(http.StatusOK, nil)(), nil
}

func (c *idleConnectionCloser) CloseIdleConnections() {
	c.closeCount++
}

func TestCloseIdleConnectionsCallsTheBaseTransport(t *testing.T) {
	base := &idleConnectionCloser{}
	transport := NewTransport(base, testOptions())

	transport.CloseIdleConnections()

	require.Equal(t, 1, base.closeCount)
}

func TestCloseIdleConnectionsDoesNothingWhenTheBaseTransportDoesNotSupportIt(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{newResponse(http.StatusOK, nil)}}
	transport := NewTransport(server, testOptions())

	require.NotPanics(t, transport.CloseIdleConnections)
}
