package resources

type ProjectVariableSetUsage struct {
	IsCurrentlyBeingUsedInProject bool                         `json:"IsCurrentlyBeingUsedInProject"`
	ProjectID                     string                       `json:"ProjectId,omitempty"`
	ProjectName                   string                       `json:"ProjectName,omitempty"`
	ProjectSlug                   string                       `json:"ProjectSlug,omitempty"`
	Releases                      []*ReleaseUsageEntry         `json:"Releases"`
	RunbookSnapshots              []*RunbookSnapshotUsageEntry `json:"RunbookSnapshots"`
}
