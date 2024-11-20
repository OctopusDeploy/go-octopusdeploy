package tenants

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type Tenant struct {
	ClonedFromTenantID  string              `json:"ClonedFromTenantId"`
	Description         string              `json:"Description"`
	Name                string              `json:"Name" validate:"required"`
	ProjectEnvironments map[string][]string `json:"ProjectEnvironments,omitempty"`
	SpaceID             string              `json:"SpaceId"`
	TenantTags          []string            `json:"TenantTags,omitempty"`
	IsDisabled          bool                `json:"IsDisabled"`

	resources.Resource
}

// NewTenant initializes a Tenant with a name.
func NewTenant(name string) *Tenant {
	return &Tenant{
		Name:                name,
		ProjectEnvironments: map[string][]string{},
		Resource:            *resources.NewResource(),
	}
}

// Validate checks the state of the tenant and returns an error if invalid.
func (t Tenant) Validate() error {
	return validator.New().Struct(t)
}

func (t *Tenant) GetName() string {
	return t.Name
}
