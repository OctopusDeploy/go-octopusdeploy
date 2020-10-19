package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// AccountResource is the embedded struct used for all accounts.
type AccountResource struct {
	Description            string   `json:"Description,omitempty"`
	EnvironmentIDs         []string `json:"EnvironmentIds,omitempty"`
	Name                   string   `json:"Name" validate:"required,notblank"`
	SpaceID                string   `json:"SpaceId,omitempty" validate:"omitempty,notblank"`
	TenantedDeploymentMode string   `json:"TenantedDeploymentParticipation" validate:"required,oneof=Untenanted TenantedOrUntenanted Tenanted"`
	TenantIDs              []string `json:"TenantIds,omitempty"`
	TenantTags             []string `json:"TenantTags,omitempty"`

	Resource
}

// newAccountResource creates and initializes an account resource.
func newAccountResource(name string) *AccountResource {
	return &AccountResource{
		EnvironmentIDs:         []string{},
		Name:                   name,
		TenantedDeploymentMode: "Untenanted",
		TenantIDs:              []string{},
		TenantTags:             []string{},
		Resource:               *newResource(),
	}
}

// GetName returns the name of the account resource.
func (a *AccountResource) GetName() string {
	return a.Name
}

// SetName sets the name of the account resource.
func (a *AccountResource) SetName(name string) {
	a.Name = name
}

// Validate checks the state of the account resource and returns an error if
// invalid.
func (a *AccountResource) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(a)
}

var _ IHasName = &AccountResource{}
