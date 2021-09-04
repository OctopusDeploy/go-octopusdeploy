package octopusdeploy

type RunbookSnapshotTemplate struct {
	NextNameIncrement string                    `json:"NextNameIncrement,omitempty"`
	Packages          []*ReleaseTemplatePackage `json:"Packages"`
	RunbookID         string                    `json:"RunbookId,omitempty"`
	RunbookProcessID  string                    `json:"RunbookProcessId,omitempty"`

	resource
}
