package testutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
)

func GoBegin[TResult any](action func() TResult) chan TResult {
	c := make(chan TResult)
	go func() {
		c <- action()
	}()
	return c
}

type Pair[T1 any, T2 any] struct {
	Item1 T1
	Item2 T2
}

func GoBegin2[TResult1 any, TResult2 any](action func() (TResult1, TResult2)) chan Pair[TResult1, TResult2] {
	c := make(chan Pair[TResult1, TResult2])
	go func() {
		r1, r2 := action()
		c <- Pair[TResult1, TResult2]{Item1: r1, Item2: r2}
	}()
	return c
}

func ReceivePair[T1 any, T2 any](receiver chan Pair[T1, T2]) (T1, T2) {
	pair := <-receiver
	return pair.Item1, pair.Item2
}

type responseOrError struct {
	response *http.Response
	error    error
}

type MockHttpServer struct {
	// when the client issues a request, we receive it here
	Request chan *http.Request
	// when we want to respond back to the client, we send it here
	Response chan responseOrError

	// so test code can detect unanswered requests or responses at the end.
	// Not strictly necessary as unanswered req/resp results in a channel deadlock
	// and go panics and kills the process, so we find out about it, but this is a bit
	// less confusing to troubleshoot
	pendingMsgCount int32
}

// MockHttpServer conforms to sling.Doer so we can use it for unit testing network calls mocking the server.
// TODO: Revisit all this in v3 of the Client Library; we'd like to get rid of sling and just use
// HttpClient which should alter a lot of this

func (m *MockHttpServer) Do(req *http.Request) (*http.Response, error) {
	// we're the client here, so we send a request down the request channel
	atomic.AddInt32(&m.pendingMsgCount, 1)
	m.Request <- req
	atomic.AddInt32(&m.pendingMsgCount, -1)
	// then we wait for a response via the response channel

	atomic.AddInt32(&m.pendingMsgCount, 1)
	x := <-m.Response
	atomic.AddInt32(&m.pendingMsgCount, -1)
	return x.response, x.error
}

func NewMockHttpServer() *MockHttpServer {
	return &MockHttpServer{
		Request:  make(chan *http.Request),
		Response: make(chan responseOrError),
	}
}

// jsonDecoder decodes http response JSON into a JSON-tagged struct value.
type jsonDecoder struct{}

// Decode decodes the Response Body into the value pointed to by v.
// Caller must provide a non-nil v and close the resp.Body.
func (d jsonDecoder) Decode(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}

func (m *MockHttpServer) Sling() *sling.Sling {
	result := &sling.Sling{}
	result.ResponseDecoder(jsonDecoder{})
	result.Doer(m)
	return result
}

func (m *MockHttpServer) GetPendingMessageCount() int {
	return int(m.pendingMsgCount)
}

func (m *MockHttpServer) ReceiveRequest() *http.Request {
	atomic.AddInt32(&m.pendingMsgCount, 1)
	request := <-m.Request
	atomic.AddInt32(&m.pendingMsgCount, -1)
	return request
}

func (m *MockHttpServer) Respond(response *http.Response, err error) {
	atomic.AddInt32(&m.pendingMsgCount, 1)
	m.Response <- responseOrError{response: response, error: err}
	atomic.AddInt32(&m.pendingMsgCount, -1)
}

// now we build some higher level methods on top of ReceiveRequest

func (m *MockHttpServer) ExpectRequest(t *testing.T, method string, pathAndQuery string) *RequestWrapper {
	r := m.ReceiveRequest()

	rPathAndQuery := r.URL.Path
	if r.URL.RawQuery != "" {
		rPathAndQuery = fmt.Sprintf("%s?%s", rPathAndQuery, r.URL.RawQuery)
	}
	assert.Equal(t, method, r.Method)
	assert.Equal(t, pathAndQuery, rPathAndQuery)

	return &RequestWrapper{r, m}
}

type RequestWrapper struct {
	// in case you need it
	Request *http.Request
	Server  *MockHttpServer
}

func (r *RequestWrapper) RespondWithText(body string) {
	r.Server.Respond(&http.Response{
		StatusCode:    http.StatusOK,
		Body:          ioutil.NopCloser(strings.NewReader(body)),
		ContentLength: int64(len(body)),
	}, nil)
}

func (r *RequestWrapper) RespondWith(responseObject any) {
	if responseObject == nil {
		panic("TODO: implement responses with no body")
	}

	body, _ := json.Marshal(responseObject)

	// Regarding response errors:
	// Note that we would use an error here for a low level thing like a network error.
	// An HTTP error like a 404 or 500 would be considered a valid response with an
	// appropriate status code
	r.Server.Respond(&http.Response{
		StatusCode:    http.StatusOK,
		Body:          ioutil.NopCloser(bytes.NewReader(body)),
		ContentLength: int64(len(body)),
	}, nil)
}
