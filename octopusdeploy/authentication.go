package octopusdeploy

// Authentication represents enabled authentication providers.
type Authentication struct {
	AnyAuthenticationProvidersSupportPasswordManagement bool                             `json:"AnyAuthenticationProvidersSupportPasswordManagement"`
	AuthenticationProviders                             []*AuthenticationProviderElement `json:"AuthenticationProviders"`
	AutoLoginEnabled                                    bool                             `json:"AutoLoginEnabled"`

	resource
}

func NewAuthentication() *Authentication {
	return &Authentication{
		resource: *newResource(),
	}
}
