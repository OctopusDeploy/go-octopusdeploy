package client

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
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
				Body:       io.NopCloser(strings.NewReader(constants.ApiReplyRoot)),
			}, nil
		}

		return &http.Response{
			StatusCode: statusCode,
			Body:       io.NopCloser(strings.NewReader(responseBody)),
		}, nil
	})

	host := os.Getenv(constants.EnvironmentVariableOctopusHost)
	apiKey := os.Getenv(constants.EnvironmentVariableOctopusApiKey)

	if len(host) == 0 {
		host = os.Getenv(constants.ClientURLEnvironmentVariable)
	}

	if len(apiKey) == 0 {
		apiKey = os.Getenv(constants.ClientAPIKeyEnvironmentVariable)
	}

	url, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	octopusClient, err := NewClient(&httpClient, url, apiKey, "")
	if err != nil {
		return nil, err
	}

	return octopusClient, nil
}
