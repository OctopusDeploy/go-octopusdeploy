package tenants

type TenantsMissingVariablesQuery struct {
	EnvironmentID  []string `uri:"environmentId,omitempty" url:"environmentId,omitempty"`
	IncludeDetails bool     `uri:"includeDetails,omitempty" url:"includeDetails,omitempty"`
	ProjectID      string   `uri:"projectId,omitempty" url:"projectId,omitempty"`
	TenantID       string   `uri:"tenantId,omitempty" url:"tenantId,omitempty"`
}
