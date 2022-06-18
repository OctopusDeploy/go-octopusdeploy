package client

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
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
				Body:       ioutil.NopCloser(strings.NewReader(constants.ApiReplyRoot)),
			}, nil
		}

		return &http.Response{
			StatusCode: statusCode,
			Body:       ioutil.NopCloser(strings.NewReader(responseBody)),
		}, nil
	})

	url, err := url.Parse(os.Getenv(constants.ClientURLEnvironmentVariable))
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv(constants.ClientAPIKeyEnvironmentVariable)

	octopusClient, err := NewClient(&httpClient, url, apiKey, "")
	if err != nil {
		return nil, err
	}

	return octopusClient, nil
}
