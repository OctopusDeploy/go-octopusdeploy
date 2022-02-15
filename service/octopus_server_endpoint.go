package service

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"net/url"
	"os"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

type octopusServerEndpoint struct {
	BaseURLWithAPI *url.URL
	ApiKey         string
}

func NewOctopusServerEndpoint(apiURL *url.URL, apiKey string) (*octopusServerEndpoint, error) {
	octopusApiURLFromEnvVar := os.Getenv(octopusdeploy.clientURLEnvironmentVariable)
	var urlError error
	if apiURL == nil && !internal.IsEmpty(octopusApiURLFromEnvVar) {
		apiURL, urlError = url.Parse(octopusApiURLFromEnvVar)
		if urlError != nil {
			return nil, urlError
		}
	}

	octopusAPIKeyFromEnvVar := os.Getenv(clientAPIKeyEnvironmentVariable)
	if IsEmpty(apiKey) && !IsEmpty(octopusAPIKeyFromEnvVar) {
		apiKey = octopusAPIKeyFromEnvVar
	}

	if apiURL == nil {
		return nil, internal.CreateInvalidParameterError(clientNewOctopusServerEndpoint, octopusdeploy.ParameterOctopusURL)
	}

	if IsEmpty(apiKey) {
		return nil, CreateInvalidParameterError(clientNewOctopusServerEndpoint, ParameterAPIKey)
	}

	if !isAPIKey(apiKey) {
		return nil, CreateInvalidParameterError(clientNewOctopusServerEndpoint, ParameterAPIKey)
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
