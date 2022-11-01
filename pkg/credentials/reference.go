package credentials

// Reference defines a reference Git credential.
type Reference struct {
	ID string `json:"Id"`

	gitCredential
}

// NewReference creates and initializes a reference Git credential.
func NewReference(id string) *Reference {
	return &Reference{
		ID: id,
		gitCredential: gitCredential{
			CredentialType: GitCredentialTypeReference,
		},
	}
}

// Type returns the type for this Git credential.
func (u *Reference) Type() Type {
	return u.CredentialType
}

var _ GitCredential = &Reference{}
