package resources

type ReleaseUsage struct {
	ProjectID   string               `json:"ProjectId,omitempty"`
	ProjectName string               `json:"ProjectName,omitempty"`
	Releases    []*ReleaseUsageEntry `json:"Releases"`
}
