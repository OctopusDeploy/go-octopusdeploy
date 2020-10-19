package model

// AuthenticationProviderElement represents an authentication provider element.
type AuthenticationProviderElement struct {
	CSSLinks          []string          `json:"CSSLinks"`
	FormsLoginEnabled bool              `json:"FormsLoginEnabled"`
	IdentityType      string            `json:"IdentityType,omitempty"`
	JavascriptLinks   []string          `json:"JavascriptLinks"`
	Links             map[string]string `json:"Links,omitempty"`
	Name              string            `json:"Name,omitempty"`
}
