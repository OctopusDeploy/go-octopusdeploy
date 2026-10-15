package ratelimit

import (
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewHttpClientWithNilBaseUsesDefaultTransport(t *testing.T) {
	client := NewHttpClient(nil, testOptions())

	require.NotNil(t, client)
	require.IsType(t, &Transport{}, client.Transport)
	require.Equal(t, http.DefaultTransport, client.Transport.(*Transport).base)
}

func TestNewHttpClientWrapsTheTransportOfBase(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{newResponse(http.StatusOK, nil)}}
	base := &http.Client{Transport: server}

	client := NewHttpClient(base, testOptions())

	require.IsType(t, &Transport{}, client.Transport)
	require.Equal(t, server, client.Transport.(*Transport).base)
}

func TestNewHttpClientDoesNotChangeBase(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{newResponse(http.StatusOK, nil)}}
	base := &http.Client{Transport: server}

	client := NewHttpClient(base, testOptions())

	require.NotSame(t, base, client)
	require.Equal(t, server, base.Transport)
}

func TestNewHttpClientKeepsTheSettingsOfBase(t *testing.T) {
	jar, err := cookiejar.New(nil)
	require.NoError(t, err)
	checkRedirect := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	base := &http.Client{
		Timeout:       42 * time.Second,
		Jar:           jar,
		CheckRedirect: checkRedirect,
	}

	client := NewHttpClient(base, testOptions())

	require.Equal(t, base.Timeout, client.Timeout)
	require.Equal(t, base.Jar, client.Jar)
	require.NotNil(t, client.CheckRedirect)
	require.ErrorIs(t, client.CheckRedirect(nil, nil), http.ErrUseLastResponse)
}

func TestNewHttpClientCloseIdleConnectionsReachesTheBaseTransport(t *testing.T) {
	base := &idleConnectionCloser{}

	client := NewHttpClient(&http.Client{Transport: base}, testOptions())
	client.CloseIdleConnections()

	require.Equal(t, 1, base.closeCount)
}

func TestNewHttpClientRetriesTooManyRequests(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{
		newResponse(http.StatusTooManyRequests, map[string]string{"Retry-After": "0"}),
		newResponse(http.StatusOK, nil),
	}}

	client := NewHttpClient(&http.Client{Transport: server}, testOptions())
	resp, err := client.Get(testURL)

	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, 2, server.requestCount())
}
