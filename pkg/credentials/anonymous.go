package credentials

type Anonymous struct {
	gitCredential
}

func NewAnonymous() *Anonymous {
	return &Anonymous{
		gitCredential: gitCredential{
			Type: GitCredentialTypeAnonymous,
		},
	}
}

// GitCredentialType returns the type for this Git credential.
func (a *Anonymous) GetType() Type {
	return a.Type
}

var _ IGitCredential = &Anonymous{}
