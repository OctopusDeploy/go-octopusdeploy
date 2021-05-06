package octopusdeploy

type TenantVariable struct {
	LibraryVariables map[string]LibraryVariable `json:"LibraryVariables,omitempty"`
	ProjectVariables map[string]ProjectVariable `json:"ProjectVariables,omitempty"`
	SpaceID          string                     `json:"SpaceId,omitempty"`
	TenantID         string                     `json:"TenantId,omitempty"`
	TenantName       string                     `json:"TenantName,omitempty"`

	resource
}

func NewTenantVariable(tenantID string) *TenantVariable {
	return &TenantVariable{
		TenantID: tenantID,
		resource: *newResource(),
	}
}
