package runbooks

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
)

type RunbookProcesses struct {
	Items []*RunbookProcess `json:"Items"`
	resources.PagedResults
}

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
