package credentials

// GitHub defines a reference GitHub connection.
type GitHub struct {
	ID string `json:"Id"`

	gitCredential
}

// NewReference creates and initializes a reference Git credential.
func NewGitHub(id string) *Reference {
	return &Reference{
		ID: id,
		gitCredential: gitCredential{
			CredentialType: GitCredentialTypeGitHub,
		},
	}
}

// Type returns the type for this Git credential.
func (u *GitHub) Type() Type {
	return u.CredentialType
}

var _ GitCredential = &GitHub{}
