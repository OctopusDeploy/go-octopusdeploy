package credentials

// GitHubApp defines a GitHub App connection Git credential.
type GitHubApp struct {
	ID string `json:"Id"`

	gitCredential
}

// NewGitHubApp creates and initializes a GitHub App Git credential.
func NewGitHubApp(id string) *GitHubApp {
	return &GitHubApp{
		ID: id,
		gitCredential: gitCredential{
			CredentialType: GitCredentialTypeGitHubApp,
		},
	}
}

// Type returns the type for this Git credential.
func (g *GitHubApp) Type() Type {
	return g.CredentialType
}

var _ GitCredential = &GitHubApp{}
