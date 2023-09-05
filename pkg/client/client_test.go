package client

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := &http.Client{}
	octopusURL := os.Getenv("OCTOPUS_HOST")
	apiKey := os.Getenv("OCTOPUS_API_KEY")
	spaceID := os.Getenv("OCTOPUS_SPACE")

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	testCases := []struct {
		name    string
		isValid bool
		client  *http.Client
		url     *url.URL
		apiKey  string
		spaceID string
	}{
		{"NilURL", false, client, nil, apiKey, spaceID},
		{"EmptyAPIKey", false, client, apiURL, "", ""},
		{"EmptyAPIKeyWithSpace", false, client, apiURL, " ", spaceID},
		{"InvalidAPIKey", false, client, apiURL, "API-***************************", spaceID},
		{"ValidAPIKey", true, client, apiURL, apiKey, spaceID},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewClientForTool(tc.client, tc.url, tc.apiKey, tc.spaceID, "test")

			if !tc.isValid {
				require.Error(t, err)
				require.Nil(t, client)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, client)
			assert.NotNil(t, client.Accounts)
		})
	}
}

func TestGetHeadersWithApiKeySetsCorrectHeader(t *testing.T) {
	apiKey := "API-API1234"
	headers := getHeaders(ApiCredentials{
		ApiKey: &apiKey,
	}, "test")

	require.Equal(t, headers[constants.ClientAPIKeyHTTPHeader], apiKey)
}

func TestGetHeadersWithAccessTokenSetsCorrectHeader(t *testing.T) {
	accessToken := "token"
	headers := getHeaders(ApiCredentials{
		AccessToken: &accessToken,
	}, "test")

	require.Equal(t, headers["Authorization"], fmt.Sprintf("Bearer %s", accessToken))
}

func TestGetHeadersSetsCorrectUserAgent(t *testing.T) {
	expectedUserAgent := api.GetUserAgentString("test")
	accessToken := "token"
	headers := getHeaders(ApiCredentials{
		AccessToken: &accessToken,
	}, "test")

	require.Equal(t, headers["User-Agent"], expectedUserAgent)
}

func TestNewClientWithAccessToken(t *testing.T) {
	client := &http.Client{}
	octopusURL := os.Getenv("OCTOPUS_HOST")
	accessToken := os.Getenv("OCTOPUS_ACCESS_TOKEN")
	spaceID := os.Getenv("OCTOPUS_SPACE")

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	testCases := []struct {
		name        string
		isValid     bool
		client      *http.Client
		url         *url.URL
		accessToken string
		spaceID     string
	}{
		{"NilURL", false, client, nil, accessToken, spaceID},
		{"EmptyAccessToken", false, client, apiURL, "", ""},
		{"ValidAccessToken", true, client, apiURL, accessToken, spaceID},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := NewClientWithCredentials(tc.client, tc.url, ApiCredentials{AccessToken: &tc.accessToken}, tc.spaceID, "test")

			if !tc.isValid {
				require.Error(t, err)
				require.Nil(t, client)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, client)
			assert.NotNil(t, client.Accounts)
		})
	}
}

func TestGetUserAgentString(t *testing.T) {
	userAgentString := api.GetUserAgentString("test")
	assert.NotNil(t, userAgentString)
}

func TestGetWithEmptyParameters(t *testing.T) {
	resource, err := api.ApiGet(nil, nil, "")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestGetWithEmptySling(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	resource, err := api.ApiGet(nil, input, "fake-path")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestGetWithEmptyPath(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	resource, err := api.ApiGet(sling.New(), input, "")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = api.ApiGet(sling.New(), input, " ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestAddWithEmptyParameters(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := services.ApiAdd(nil, input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestAddWithEmptySling(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := services.ApiAdd(nil, input, response, "fake-path")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestAddWithEmptyPath(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := services.ApiAdd(sling.New(), input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = services.ApiAdd(sling.New(), input, response, " ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestPostWithEmptyParameters(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := services.ApiPost(nil, input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestPostWithEmptySling(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := services.ApiPost(nil, input, response, "fake-path")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestPostWithEmptyPath(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := services.ApiPost(sling.New(), input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = services.ApiPost(sling.New(), input, response, " ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestUpdateWithEmptyParameters(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := services.ApiUpdate(nil, input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestUpdateWithEmptySling(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := services.ApiUpdate(nil, input, response, "fake-path")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestUpdateWithEmptyPath(t *testing.T) {
	input := &inputTestValueStruct{test: "fake-value"}
	response := &inputTestResponseStruct{test: "fake-value"}

	resource, err := services.ApiUpdate(sling.New(), input, response, "")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = services.ApiUpdate(sling.New(), input, response, " ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestDeleteWithEmptyParameters(t *testing.T) {
	err := services.ApiDelete(nil, "")
	assert.Error(t, err)
}

func TestDeleteWithEmptySling(t *testing.T) {
	err := services.ApiDelete(nil, "fake-path")
	assert.Error(t, err)
}

func TestDeleteWithEmptyPath(t *testing.T) {
	err := services.ApiDelete(nil, "")
	assert.Error(t, err)

	err = services.ApiDelete(nil, " ")
	assert.Error(t, err)
}

type inputTestValueStruct struct {
	test string
}

func (i *inputTestValueStruct) GetID() string {
	return "fake-ID"
}

func (i *inputTestValueStruct) Validate() error {
	return nil
}

type inputTestResponseStruct struct {
	test string
}

func (i *inputTestResponseStruct) GetID() string {
	return "fake-ID"
}

func (i *inputTestResponseStruct) Validate() error {
	return nil
}
