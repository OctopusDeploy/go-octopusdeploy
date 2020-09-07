package model

type UserAuthentication struct {
	AuthenticationProviders             []*AuthenticationProviderElement `json:"AuthenticationProviders,omitempty"`
	CanCurrentUserEditIdentitiesForUser *bool                            `json:"AutoLoginEnabled,omitempty"`
	Resource
}
