package credentials

type Type string

const (
	GitCredentialTypeAnonymous        = Type("Anonymous")
	GitCredentialTypeReference        = Type("Reference")
	GitCredentialTypeUsernamePassword = Type("UsernamePassword")
)
