package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/model"
	"github.com/dghubble/sling"
)

// AuthenticationService handles communication with Authentication-related
// methods of the Octopus API.
type AuthenticationService struct {
	sling *sling.Sling
	path  string
}

// NewAuthenticationService returns an AuthenticationService with a
// preconfigured client.
func NewAuthenticationService(sling *sling.Sling) *AuthenticationService {
	return &AuthenticationService{
		sling: sling,
		path:  "authentication",
	}
}

func (s *AuthenticationService) Get() (*model.Authentication, error) {
	resp, err := apiGet(s.sling, new(model.Authentication), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Authentication), nil
}
