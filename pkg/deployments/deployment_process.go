package deployments

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type DeploymentProcess struct {
	Branch         string            `json:"-"`
	LastSnapshotID string            `json:"LastSnapshotId,omitempty"`
	ProjectID      string            `json:"ProjectId,omitempty"`
	SpaceID        string            `json:"SpaceId,omitempty"`
	Steps          []*DeploymentStep `json:"Steps,omitempty"`
	Version        int32             `json:"Version"`

	resources.Resource
}

// NewDeploymentProcess initializes a deployment process.
func NewDeploymentProcess(projectID string) *DeploymentProcess {
	return &DeploymentProcess{
		ProjectID: projectID,
		Resource:  *resources.NewResource(),
	}
}

const (
	PackageAcquisitionLocationServer          = "Server"
	PackageAcquisitionLocationExecutionTarget = "ExecutionTarget"
	PackageAcquisitionLocationNotAcquired     = "NotAcquired"
)
