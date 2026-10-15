package client

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/ratelimit"
	"github.com/stretchr/testify/require"
)

const testRateLimitAPIKey = "API-ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// newRateLimitedHttpClient returns an http.Client that returns 429 for the first request to the root resource.
// After that, it returns the root resource for /api and an empty list for all other paths.
func newRateLimitedHttpClient() (*http.Client, func() int) {
	var mutex sync.Mutex
	rootRequestCount := 0

	httpClient := &http.Client{
		Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
			if request.URL.Path == "/api" {
				mutex.Lock()
				rootRequestCount++
				count := rootRequestCount
				mutex.Unlock()

				if count == 1 {
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header:     http.Header{"Retry-After": []string{"0"}},
						Body:       io.NopCloser(strings.NewReader(`{"ErrorMessage":"Too many requests"}`)),
					}, nil
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(constants.ApiReplyRoot)),
				}, nil
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"Items":[]}`)),
			}, nil
		}),
	}

	return httpClient, func() int {
		mutex.Lock()
		defer mutex.Unlock()
		return rootRequestCount
	}
}

func newTestRateLimitOptions() *ClientOptions {
	return &ClientOptions{
		RateLimit: &ratelimit.Options{
			MaxRetries: 3,
			MaxWait:    time.Second,
		},
	}
}

func TestNewClientWithOptionsRetriesTooManyRequests(t *testing.T) {
	httpClient, rootRequestCount := newRateLimitedHttpClient()
	apiURL, err := url.Parse("https://octopus.example")
	require.NoError(t, err)
	apiKey, err := NewApiKey(testRateLimitAPIKey)
	require.NoError(t, err)

	octopusClient, err := NewClientWithOptions(httpClient, apiURL, apiKey, "", "test", newTestRateLimitOptions())

	require.NoError(t, err)
	require.NotNil(t, octopusClient)
	require.Equal(t, 2, rootRequestCount())
}

func TestNewClientWithOptionsDoesNotChangeTheHttpClient(t *testing.T) {
	httpClient, _ := newRateLimitedHttpClient()
	apiURL, err := url.Parse("https://octopus.example")
	require.NoError(t, err)
	apiKey, err := NewApiKey(testRateLimitAPIKey)
	require.NoError(t, err)

	_, err = NewClientWithOptions(httpClient, apiURL, apiKey, "", "test", newTestRateLimitOptions())

	require.NoError(t, err)
	require.IsType(t, roundTripFunc(nil), httpClient.Transport)
}

func TestNewClientWithOptionsDoesNotRetryWhenRateLimitIsNil(t *testing.T) {
	httpClient, rootRequestCount := newRateLimitedHttpClient()
	apiURL, err := url.Parse("https://octopus.example")
	require.NoError(t, err)
	apiKey, err := NewApiKey(testRateLimitAPIKey)
	require.NoError(t, err)

	octopusClient, err := NewClientWithOptions(httpClient, apiURL, apiKey, "", "test", &ClientOptions{})

	require.Error(t, err)
	require.Nil(t, octopusClient)
	require.Equal(t, 1, rootRequestCount())
}

func TestNewClientWithOptionsDoesNotRetryWhenOptionsAreNil(t *testing.T) {
	httpClient, rootRequestCount := newRateLimitedHttpClient()
	apiURL, err := url.Parse("https://octopus.example")
	require.NoError(t, err)
	apiKey, err := NewApiKey(testRateLimitAPIKey)
	require.NoError(t, err)

	octopusClient, err := NewClientWithOptions(httpClient, apiURL, apiKey, "", "test", nil)

	require.Error(t, err)
	require.Nil(t, octopusClient)
	require.Equal(t, 1, rootRequestCount())
}
