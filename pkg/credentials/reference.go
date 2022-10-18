package credentials

// Reference defines a reference Git credential.
type Reference struct {
	Id string `json:"Id"`

	gitCredential
}

// NewReference creates and initializes a reference Git credential.
func NewReference(id string) *Reference {
	return &Reference{
		Id: id,
		gitCredential: gitCredential{
			Type: GitCredentialTypeReference,
		},
	}
}

// GitCredentialType returns the type for this Git credential.
func (u *Reference) GetType() Type {
	return u.Type
}

var _ IGitCredential = &Reference{}
