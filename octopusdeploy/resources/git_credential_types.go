package resources

type GitCredentialType string

const (
	GitCredentialTypeAnonymous        = GitCredentialType("Anonymous")
	GitCredentialTypeUsernamePassword = GitCredentialType("UsernamePassword")
)
