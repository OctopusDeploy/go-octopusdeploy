package model

type RunbookStepUsage struct {
	ProcessID   string            `json:"ProcessId,omitempty"`
	ProjectID   string            `json:"ProjectId,omitempty"`
	ProjectName string            `json:"ProjectName,omitempty"`
	ProjectSlug string            `json:"ProjectSlug,omitempty"`
	RunbookID   string            `json:"RunbookId,omitempty"`
	RunbookName string            `json:"RunbookName,omitempty"`
	Steps       []*StepUsageEntry `json:"Steps"`
}
