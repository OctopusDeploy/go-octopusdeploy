package octopusdeploy

import (
	"github.com/dghubble/sling"
)

// authenticationService handles communication with Authentication-related methods of the Octopus API.
type authenticationService struct {
	loginInitiatedPath string

	service
}

// newAuthenticationService returns an authenticationService with a preconfigured client.
func newAuthenticationService(sling *sling.Sling, uriTemplate string, loginInitiatedPath string) *authenticationService {
	return &authenticationService{
		loginInitiatedPath: loginInitiatedPath,
		service:            newService(serviceAuthenticationService, sling, uriTemplate, new(Authentication)),
	}
}

func (s authenticationService) Get() (*Authentication, error) {
	path, err := getPath(s)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(Authentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Authentication), nil
}
