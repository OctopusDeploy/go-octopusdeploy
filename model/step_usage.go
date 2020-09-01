package model

type StepUsage struct {
	ProjectID   string            `json:"ProjectId,omitempty"`
	ProjectName string            `json:"ProjectName,omitempty"`
	ProjectSlug string            `json:"ProjectSlug,omitempty"`
	Steps       []*StepUsageEntry `json:"Steps"`
}
