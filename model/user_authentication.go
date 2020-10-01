package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type UserAuthentication struct {
	AuthenticationProviders             []*AuthenticationProviderElement `json:"AuthenticationProviders,omitempty"`
	CanCurrentUserEditIdentitiesForUser *bool                            `json:"AutoLoginEnabled,omitempty"`

	Resource
}

// GetID returns the ID value of the user authentication.
func (resource UserAuthentication) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of
// this user authentication.
func (resource UserAuthentication) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this user
// authentication was changed.
func (resource UserAuthentication) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this user
// authentication.
func (resource UserAuthentication) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the user authentication and returns an error if
// invalid.
func (resource UserAuthentication) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &UserAuthentication{}
