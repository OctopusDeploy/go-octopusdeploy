package runbookprocess

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type RunbookProcess struct {
	LastSnapshotID string                        `json:"LastSnapshotId,omitempty"`
	ProjectID      string                        `json:"ProjectId,omitempty"`
	RunbookID      string                        `json:"RunbookId,omitempty"`
	SpaceID        string                        `json:"SpaceId,omitempty"`
	Steps          []*deployments.DeploymentStep `json:"Steps"`
	Version        *int32                        `json:"Version"`

	resources.Resource
}

func NewRunbookProcess() *RunbookProcess {
	return &RunbookProcess{
		Resource: *resources.NewResource(),
	}
}
