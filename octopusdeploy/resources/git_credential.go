package resources

// gitCredential is the embedded struct used for all GIT-based credentials.
type gitCredential struct {
	Type GitCredentialType `validate:"omitempty,oneof=Anonymous UsernamePassword"`
}
