package credentials

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
)

// UsernamePassword defines a username-password Git credential.
type UsernamePassword struct {
	Password *core.SensitiveValue `json:"Password"`
	Username string               `json:"Username"`

	gitCredential
}

// NewUsernamePassword creates and initializes an username-password Git credential.
func NewUsernamePassword(username string, password *core.SensitiveValue) *UsernamePassword {
	return &UsernamePassword{
		Password: password,
		Username: username,
		gitCredential: gitCredential{
			Type: GitCredentialTypeUsernamePassword,
		},
	}
}

// GitCredentialType returns the type for this Git credential.
func (u *UsernamePassword) GetType() Type {
	return u.Type
}

var _ IGitCredential = &UsernamePassword{}
