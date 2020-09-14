package model

import "github.com/go-playground/validator/v10"

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

func NewTenant(name, description string) *Tenant {
	return &Tenant{
		Name:        name,
		Description: description,
	}
}

func (t *Tenant) GetID() string {
	return t.ID
}

func (t *Tenant) Validate() error {
	validate := validator.New()
	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &Tenant{}
