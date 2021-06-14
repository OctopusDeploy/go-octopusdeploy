package octopusdeploy

type TenantVariables struct {
	LibraryVariables map[string]LibraryVariable `json:"LibraryVariables,omitempty"`
	ProjectVariables map[string]ProjectVariable `json:"ProjectVariables,omitempty"`
	SpaceID          string                     `json:"SpaceId,omitempty"`
	TenantID         string                     `json:"TenantId,omitempty"`
	TenantName       string                     `json:"TenantName,omitempty"`

	resource
}

func NewTenantVariables(tenantID string) *TenantVariables {
	return &TenantVariables{
		TenantID: tenantID,
		resource: *newResource(),
	}
}
