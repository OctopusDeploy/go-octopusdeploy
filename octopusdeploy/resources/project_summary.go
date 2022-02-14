package resources

type ProjectSummary struct {
	HasDeploymentProcess bool `json:"HasDeploymentProcess,omitempty"`
	HasRunbooks          bool `json:"HasRunbooks,omitempty"`
}
