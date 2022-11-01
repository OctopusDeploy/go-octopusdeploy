package credentials

type Anonymous struct {
	gitCredential
}

func NewAnonymous() *Anonymous {
	return &Anonymous{
		gitCredential: gitCredential{
			credentialType: GitCredentialTypeAnonymous,
		},
	}
}

// Type returns the type for this Git credential.
func (a *Anonymous) Type() Type {
	return a.credentialType
}

var _ GitCredential = &Anonymous{}
