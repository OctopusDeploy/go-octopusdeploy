package client

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

type roundTripFunc func(r *http.Request) (*http.Response, error)

func (s roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return s(r)
}

// GetFakeOctopusClient -
func GetFakeOctopusClient(t *testing.T, apiPath string, statusCode int, responseBody string) (*Client, error) {
	httpClient := http.Client{}
	httpClient.Transport = roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.URL.Path == "/api" {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(apiReplyRoot)),
			}, nil
		}

		return &http.Response{
			StatusCode: statusCode,
			Body:       ioutil.NopCloser(strings.NewReader(responseBody)),
		}, nil
	})

	url, err := url.Parse(os.Getenv(clientURLEnvironmentVariable))
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv(clientAPIKeyEnvironmentVariable)

	octopusClient, err := NewClient(&httpClient, url, apiKey, emptyString)
	if err != nil {
		return nil, err
	}

	return octopusClient, nil
}
