package credentials

type Type string

const (
	GitCredentialTypeAnonymous        = Type("Anonymous")
	GitCredentialTypeGitHubApp        = Type("GitHub")
	GitCredentialTypeReference        = Type("Reference")
	GitCredentialTypeUsernamePassword = Type("UsernamePassword")
)
