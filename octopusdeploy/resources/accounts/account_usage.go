package accounts

import "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"

// AccountUsage contains the projects and deployments which are using an
// account.
type AccountUsage struct {
	DeploymentProcesses []*octopusdeploy.StepUsage                    `json:"DeploymentProcesses,omitempty"`
	LibraryVariableSets []*octopusdeploy.LibraryVariableSetUsageEntry `json:"LibraryVariableSets,omitempty"`
	ProjectVariableSets []*octopusdeploy.ProjectVariableSetUsage      `json:"ProjectVariableSets,omitempty"`
	Releases            []*octopusdeploy.ReleaseUsage                 `json:"Releases,omitempty"`
	RunbookProcesses    []*octopusdeploy.RunbookStepUsage             `json:"RunbookProcesses,omitempty"`
	RunbookSnapshots    []*octopusdeploy.RunbookSnapshotUsage         `json:"RunbookSnapshots,omitempty"`
	Targets             []*octopusdeploy.TargetUsageEntry             `json:"Targets,omitempty"`

	octopusdeploy.resource
}

// NewAccountUsage initializes an AccountUsage.
func NewAccountUsage() *AccountUsage {
	accountUsage := &AccountUsage{
		resource: *octopusdeploy.newResource(),
	}

	return accountUsage
}
