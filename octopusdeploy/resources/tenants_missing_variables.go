package resources

type TenantsMissingVariables struct {
	Links            map[string]string `json:"Links,omitempty"`
	MissingVariables []MissingVariable `json:"MissingVariables,omitempty"`
	TenantID         string            `json:"TenantId,omitempty"`
}

type MissingVariable struct {
	EnvironmentID        string            `json:"EnvironmentId,omitempty"`
	LibraryVariableSetID string            `json:"LibraryVariableSetId,omitempty"`
	Links                map[string]string `json:"Links,omitempty"`
	ProjectID            string            `json:"ProjectId,omitempty"`
	VariableTemplateID   string            `json:"VariableTemplateId,omitempty"`
	VariableTemplateName string            `json:"VariableTemplateName,omitempty"`
}
