package octopusdeploy

// AutoDeployReleaseOverride represents an auto-deploy release override.
type AutoDeployReleaseOverride struct {
	EnvironmentID string `json:"EnvironmentId,omitempty"`
	ReleaseID     string `json:"ReleaseId,omitempty"`
	TenantID      string `json:"TenantId,omitempty"`
}
