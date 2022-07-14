package projects

// IGitCredential defines the interface for Git-associated credentials.
type IGitCredential interface {
	GetType() GitCredentialType
}

// gitCredential is the embedded struct used for all GIT-based credentials.
type gitCredential struct {
	Type GitCredentialType `validate:"omitempty,oneof=Anonymous UsernamePassword"`
}
