package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// AuthenticationService handles communication with Authentication-related
// methods of the Octopus API.
type AuthenticationService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

// NewAuthenticationService returns an AuthenticationService with a
// preconfigured client.
func NewAuthenticationService(sling *sling.Sling) *AuthenticationService {
	if sling == nil {
		return nil
	}

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

func (s *AuthenticationService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("AuthenticationService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("AuthenticationService: the internal path is not set")
	}

	return nil
}

var _ ServiceInterface = &AuthenticationService{}
