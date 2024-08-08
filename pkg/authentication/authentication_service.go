package authentication

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

// AuthenticationService handles communication with Authentication-related methods of the Octopus API.
type AuthenticationService struct {
	loginInitiatedPath string

	services.Service
}

// NewAuthenticationService returns an AuthenticationService with a preconfigured client.
func NewAuthenticationService(sling *sling.Sling, uriTemplate string, loginInitiatedPath string) *AuthenticationService {
	return &AuthenticationService{
		loginInitiatedPath: loginInitiatedPath,
		Service:            services.NewService(constants.ServiceAuthenticationService, sling, uriTemplate),
	}
}

// Deprecated: Use authentication.Get
func (s *AuthenticationService) Get() (*Authentication, error) {
	path, err := services.GetPath(s)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(Authentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*Authentication), nil
}

// ----- new -----
const template = "/api/authentication"

func Get(client newclient.Client) (*Authentication, error) {
	return newclient.Get[Authentication](client.HttpSession(), template)
}
