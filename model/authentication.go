package model

// Authentication represents enabled authentication providers.
type Authentication struct {
	AnyAuthenticationProvidersSupportPasswordManagement bool                             `json:"AnyAuthenticationProvidersSupportPasswordManagement"`
	AuthenticationProviders                             []*AuthenticationProviderElement `json:"AuthenticationProviders"`
	AutoLoginEnabled                                    bool                             `json:"AutoLoginEnabled"`

	Resource
}

func NewAuthentication() *Authentication {
	return &Authentication{
		Resource: *newResource(),
	}
}
