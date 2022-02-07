package octopusdeploy

type TenantVariables struct {
	LibraryVariables map[string]LibraryVariable `json:"LibraryVariables,omitempty"`
	ProjectVariables map[string]ProjectVariable `json:"ProjectVariables,omitempty"`
	SpaceID          string                     `json:"SpaceId,omitempty"`
	TenantID         string                     `json:"TenantId,omitempty"`
	TenantName       string                     `json:"TenantName,omitempty"`

	Resource
}

func NewTenantVariables(tenantID string) *TenantVariables {
	return &TenantVariables{
		TenantID: tenantID,
		Resource: *newResource(),
	}
}
