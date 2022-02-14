package resources

type DeploymentProcesses struct {
	Items []*DeploymentProcess `json:"Items"`
	PagedResults
}

type DeploymentProcess struct {
	Branch         string           `json:"-"`
	LastSnapshotID string           `json:"LastSnapshotId,omitempty"`
	ProjectID      string           `json:"ProjectId,omitempty"`
	SpaceID        string           `json:"SpaceId,omitempty"`
	Steps          []DeploymentStep `json:"Steps,omitempty"`
	Version        int32            `json:"Version"`

	Resource
}

// NewDeploymentProcess initializes a deployment process.
func NewDeploymentProcess(projectID string) *DeploymentProcess {
	return &DeploymentProcess{
		ProjectID: projectID,
		Resource:  *NewResource(),
	}
}

const (
	PackageAcquisitionLocationServer          = "Server"
	PackageAcquisitionLocationExecutionTarget = "ExecutionTarget"
	PackageAcquisitionLocationNotAcquired     = "NotAcquired"
)
