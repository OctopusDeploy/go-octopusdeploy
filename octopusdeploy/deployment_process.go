package octopusdeploy

type DeploymentProcesses struct {
	Items []*DeploymentProcess `json:"Items"`
	PagedResults
}

type DeploymentProcess struct {
	LastSnapshotID string           `json:"LastSnapshotId,omitempty"`
	ProjectID      string           `json:"ProjectId,omitempty"`
	SpaceID        string           `json:"SpaceId,omitempty"`
	Steps          []DeploymentStep `json:"Steps,omitempty"`
	Version        int32            `json:"Version"`

	resource
}

// NewDeploymentProcess initializes a deployment process.
func NewDeploymentProcess(projectID string) *DeploymentProcess {
	return &DeploymentProcess{
		ProjectID: projectID,
		resource:  *newResource(),
	}
}

const (
	PackageAcquisitionLocationServer          = "Server"
	PackageAcquisitionLocationExecutionTarget = "ExecutionTarget"
	PackageAcquisitionLocationNotAcquired     = "NotAcquired"
)
