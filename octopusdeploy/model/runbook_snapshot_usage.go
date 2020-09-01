package model

type RunbookSnapshotUsage struct {
	ProjectID   string                       `json:"ProjectId,omitempty"`
	ProjectName string                       `json:"ProjectName,omitempty"`
	RunbookID   string                       `json:"RunbookId,omitempty"`
	RunbookName string                       `json:"RunbookName,omitempty"`
	Snapshots   []*RunbookSnapshotUsageEntry `json:"Snapshots"`
}
