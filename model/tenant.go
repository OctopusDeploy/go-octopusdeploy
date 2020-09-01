package model

type Tenants struct {
	Items []Tenant `json:"Items"`
	PagedResults
}

type Tenant struct {
	Name                string              `json:"Name"`
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

func ValidateTenantValues(tenant *Tenant) error {
	return ValidateMultipleProperties([]error{
		ValidateRequiredPropertyValue("Name", tenant.Name),
	})
}
