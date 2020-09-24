package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Authentication represents enabled authentication providers.
type Authentication struct {
	AnyAuthenticationProvidersSupportPasswordManagement bool                             `json:"AnyAuthenticationProvidersSupportPasswordManagement"`
	AuthenticationProviders                             []*AuthenticationProviderElement `json:"AuthenticationProviders"`
	AutoLoginEnabled                                    bool                             `json:"AutoLoginEnabled"`

	Resource
}

// GetID returns the ID value of the Authentication.
func (resource Authentication) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Authentication.
func (resource Authentication) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Authentication was changed.
func (resource Authentication) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Authentication.
func (resource Authentication) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the Authentication and returns an error if invalid.
func (resource Authentication) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	return nil
}

var _ ResourceInterface = &Authentication{}
