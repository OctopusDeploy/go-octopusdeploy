package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
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
		service:            newService(serviceAuthenticationService, sling, uriTemplate, new(model.Authentication)),
	}
}

func (s authenticationService) Get() (*model.Authentication, error) {
	path, err := getPath(s)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Authentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Authentication), nil
}
