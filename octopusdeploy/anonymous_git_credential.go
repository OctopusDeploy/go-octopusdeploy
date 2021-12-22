package octopusdeploy

type AnonymousGitCredential struct {
	gitCredential
}

func NewAnonymousGitCredential() *AnonymousGitCredential {
	return &AnonymousGitCredential{
		gitCredential: gitCredential{
			Type: GitCredentialTypeAnonymous,
		},
	}
}

// GitCredentialType returns the type for this Git credential.
func (a *AnonymousGitCredential) GetType() GitCredentialType {
	return a.Type
}

var _ IGitCredential = &AnonymousGitCredential{}
