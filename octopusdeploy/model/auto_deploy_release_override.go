package model

type AutoDeployReleaseOverride struct {
	EnvironmentID string `json:"EnvironmentId,omitempty"`
	ReleaseID     string `json:"ReleaseId,omitempty"`
	TenantID      string `json:"TenantId,omitempty"`
}
