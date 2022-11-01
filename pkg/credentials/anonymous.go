package credentials

type Anonymous struct {
	gitCredential
}

func NewAnonymous() *Anonymous {
	return &Anonymous{
		gitCredential: gitCredential{
			CredentialType: GitCredentialTypeAnonymous,
		},
	}
}

// Type returns the type for this Git credential.
func (a *Anonymous) Type() Type {
	return a.CredentialType
}

var _ GitCredential = &Anonymous{}
