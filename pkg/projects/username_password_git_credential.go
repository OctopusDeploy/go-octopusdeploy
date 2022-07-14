package projects

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"

// UsernamePasswordGitCredential defines a username-password Git credential.
type UsernamePasswordGitCredential struct {
	Password *core.SensitiveValue `json:"Password"`
	Username string               `json:"Username"`

	gitCredential
}

// NewUsernamePasswordGitCredential creates and initializes an username-password Git credential.
func NewUsernamePasswordGitCredential(username string, password *core.SensitiveValue) *UsernamePasswordGitCredential {
	return &UsernamePasswordGitCredential{
		gitCredential: gitCredential{
			Type: GitCredentialTypeUsernamePassword,
		},
		Password: password,
		Username: username,
	}
}

// GitCredentialType returns the type for this Git credential.
func (u *UsernamePasswordGitCredential) GetType() GitCredentialType {
	return u.Type
}

var _ IGitCredential = &UsernamePasswordGitCredential{}
