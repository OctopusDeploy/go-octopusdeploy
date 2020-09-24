package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

// authenticationService handles communication with Authentication-related methods of the Octopus API.
type authenticationService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

// newAuthenticationService returns an authenticationService with a preconfigured client.
func newAuthenticationService(sling *sling.Sling, uriTemplate string) *authenticationService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &authenticationService{
		name:        serviceAuthenticationService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s authenticationService) getClient() *sling.Sling {
	return s.sling
}

func (s authenticationService) getName() string {
	return s.name
}

func (s authenticationService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
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

var _ ServiceInterface = &authenticationService{}
