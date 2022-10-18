package credentials

// IGitCredential defines the interface for Git credentials.
type IGitCredential interface {
	GetType() Type
}

// gitCredential is the embedded struct used for Git credentials.
type gitCredential struct {
	Type Type `validate:"omitempty,oneof=Anonymous Reference UsernamePassword"`
}
