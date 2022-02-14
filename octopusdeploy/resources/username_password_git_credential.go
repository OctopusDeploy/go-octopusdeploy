package resources

// UsernamePasswordGitCredential defines a username-password Git credential.
type UsernamePasswordGitCredential struct {
	Password *SensitiveValue `json:"Password"`
	Username string          `json:"Username"`

	gitCredential
}

// NewUsernamePasswordGitCredential creates and initializes an username-password Git credential.
func NewUsernamePasswordGitCredential(username string, password *SensitiveValue) *UsernamePasswordGitCredential {
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
