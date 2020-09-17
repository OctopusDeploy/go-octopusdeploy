package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// AuthenticationService handles communication with Authentication-related
// methods of the Octopus API.
type AuthenticationService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

// NewAuthenticationService returns an AuthenticationService with a
// preconfigured client.
func NewAuthenticationService(sling *sling.Sling, uriTemplate string) *AuthenticationService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &AuthenticationService{
		name:  "AuthenticationService",
		path:  path,
		sling: sling,
	}
}

func (s *AuthenticationService) Get() (*model.Authentication, error) {
	resp, err := apiGet(s.sling, new(model.Authentication), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Authentication), nil
}

func (s *AuthenticationService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &AuthenticationService{}
