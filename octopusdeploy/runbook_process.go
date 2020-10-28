package octopusdeploy

type RunbookProcesses struct {
	Items []*RunbookProcess `json:"Items"`
	PagedResults
}

type RunbookProcess struct {
	LastSnapshotID string            `json:"LastSnapshotId,omitempty"`
	ProjectID      string            `json:"ProjectId,omitempty"`
	RunbookID      string            `json:"RunbookId,omitempty"`
	SpaceID        string            `json:"SpaceId,omitempty"`
	Steps          []*DeploymentStep `json:"Steps"`
	Version        *int32            `json:"Version"`

	resource
}

func NewRunbookProcess() *RunbookProcess {
	return &RunbookProcess{
		resource: *newResource(),
	}
}
