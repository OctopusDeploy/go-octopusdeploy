package octopusdeploy

type Identity struct {
	Claims               map[string]IdentityClaim `json:"Claims,omitempty"`
	IdentityProviderName string                   `json:"IdentityProviderName,omitempty"`
}
