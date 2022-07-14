package authentication

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"

// Authentication represents enabled authentication providers.
type Authentication struct {
	AnyAuthenticationProvidersSupportPasswordManagement bool                             `json:"AnyAuthenticationProvidersSupportPasswordManagement"`
	AuthenticationProviders                             []*AuthenticationProviderElement `json:"AuthenticationProviders"`
	AutoLoginEnabled                                    bool                             `json:"AutoLoginEnabled"`

	resources.Resource
}

func NewAuthentication() *Authentication {
	return &Authentication{
		Resource: *resources.NewResource(),
	}
}
