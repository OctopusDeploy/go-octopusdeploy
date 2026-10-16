package ratelimit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHttpClientWithNilBaseUsesDefaultTransport(t *testing.T) {
	client := NewHttpClient(nil, testOptions())

	require.NotNil(t, client)
	require.IsType(t, &Transport{}, client.Transport)
	require.Equal(t, http.DefaultTransport, client.Transport.(*Transport).Base())
}

func TestNewHttpClientWrapsTheTransportOfBase(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{newResponse(http.StatusOK, nil)}}
	base := &http.Client{Transport: server}

	client := NewHttpClient(base, testOptions())

	require.IsType(t, &Transport{}, client.Transport)
	require.Equal(t, server, client.Transport.(*Transport).Base())
}

func TestNewHttpClientDoesNotChangeBase(t *testing.T) {
	server := &fakeServer{responses: []func() *http.Response{newResponse(http.StatusOK, nil)}}
	base := &http.Client{Transport: server}

	client := NewHttpClient(base, testOptions())

	require.NotSame(t, base, client)
	require.Equal(t, server, base.Transport)
}
