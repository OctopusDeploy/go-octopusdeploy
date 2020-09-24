package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Tenants struct {
	Items []Tenant `json:"Items"`
	PagedResults
}

type Tenant struct {
	Name                string              `json:"Name" validate:"required"`
	TenantTags          []string            `json:"TenantTags,omitempty"`
	ProjectEnvironments map[string][]string `json:"ProjectEnvironments,omitempty"`
	SpaceID             string              `json:"SpaceId"`
	ClonedFromTenantID  string              `json:"ClonedFromTenantId"`
	Description         string              `json:"Description"`

	Resource
}

// NewTenant initializes a Tenant with a name and a description.
func NewTenant(name string, description string) *Tenant {
	return &Tenant{
		Name:        name,
		Description: description,
	}
}

// GetID returns the ID value of the Tenant.
func (resource Tenant) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Tenant.
func (resource Tenant) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Tenant was changed.
func (resource Tenant) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Tenant.
func (resource Tenant) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the Tenant and returns an error if invalid.
func (resource Tenant) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &Tenant{}
