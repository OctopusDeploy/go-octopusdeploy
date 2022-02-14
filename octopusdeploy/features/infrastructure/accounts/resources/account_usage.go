package resources

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
)

// AccountUsage contains the projects and deployments which are using an
// account.
type AccountUsage struct {
	DeploymentProcesses []*resources.StepUsage                    `json:"DeploymentProcesses,omitempty"`
	LibraryVariableSets []*resources.LibraryVariableSetUsageEntry `json:"LibraryVariableSets,omitempty"`
	ProjectVariableSets []*resources.ProjectVariableSetUsage      `json:"ProjectVariableSets,omitempty"`
	Releases            []*resources.ReleaseUsage                 `json:"Releases,omitempty"`
	RunbookProcesses    []*resources.RunbookStepUsage             `json:"RunbookProcesses,omitempty"`
	RunbookSnapshots    []*resources.RunbookSnapshotUsage         `json:"RunbookSnapshots,omitempty"`
	Targets             []*resources.TargetUsageEntry             `json:"Targets,omitempty"`

	resources.Resource
}

// NewAccountUsage initializes an AccountUsage.
func NewAccountUsage() *AccountUsage {
	accountUsage := &AccountUsage{
		Resource: *resources.NewResource(),
	}

	return accountUsage
}
