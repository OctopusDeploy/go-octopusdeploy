package octopusdeploy

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

type octopusServerEndpoint struct {
	BaseURLWithAPI *url.URL
	ApiKey         string
}

func NewOctopusServerEndpoint(apiURL *url.URL, apiKey string) (*octopusServerEndpoint, error) {
	octopusApiURLFromEnvVar := os.Getenv(clientURLEnvironmentVariable)
	var urlError error
	if apiURL == nil && !isEmpty(octopusApiURLFromEnvVar) {
		apiURL, urlError = url.Parse(octopusApiURLFromEnvVar)
		if urlError != nil {
			return nil, urlError
		}
	}

	octopusAPIKeyFromEnvVar := os.Getenv(clientAPIKeyEnvironmentVariable)
	if isEmpty(apiKey) && !isEmpty(octopusAPIKeyFromEnvVar) {
		apiKey = octopusAPIKeyFromEnvVar
	}

	if apiURL == nil {
		return nil, createInvalidParameterError(clientNewOctopusServerEndpoint, ParameterOctopusURL)
	}

	if isEmpty(apiKey) {
		return nil, createInvalidParameterError(clientNewOctopusServerEndpoint, ParameterAPIKey)
	}

	if !isAPIKey(apiKey) {
		return nil, createInvalidParameterError(clientNewOctopusServerEndpoint, ParameterAPIKey)
	}

	baseURLWithAPIString := strings.TrimRight(apiURL.String(), "/")
	baseURLWithAPIString = fmt.Sprintf("%s/api", baseURLWithAPIString)
	baseURLWithAPI, urlErr := url.Parse(baseURLWithAPIString)
	if urlErr != nil {
		return nil, urlErr
	}

	endpoint := octopusServerEndpoint{
		BaseURLWithAPI: baseURLWithAPI,
		ApiKey:         apiKey,
	}

	return &endpoint, nil
}
