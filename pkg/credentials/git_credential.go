package credentials

// GitCredential defines the interface for Git credentials.
type GitCredential interface {
	Type() Type
}

// gitCredential is the embedded struct used for Git credentials.
type gitCredential struct {
	credentialType Type `validate:"omitempty,oneof=Anonymous Reference UsernamePassword"`
}
