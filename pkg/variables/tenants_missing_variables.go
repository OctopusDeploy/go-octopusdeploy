package variables

type TenantsMissingVariables struct {
	Links            map[string]string `json:"Links,omitempty"`
	MissingVariables []MissingVariable `json:"MissingVariables,omitempty"`
	TenantID         string            `json:"TenantId,omitempty"`
}
