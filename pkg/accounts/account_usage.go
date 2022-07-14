package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/deployments"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/releases"
	resources "github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/runbooks"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/variables"
)

// AccountUsage contains the projects and deployments which are using an
// account.
type AccountUsage struct {
	DeploymentProcesses []*deployments.StepUsage                  `json:"DeploymentProcesses,omitempty"`
	LibraryVariableSets []*variables.LibraryVariableSetUsageEntry `json:"LibraryVariableSets,omitempty"`
	ProjectVariableSets []*variables.ProjectVariableSetUsage      `json:"ProjectVariableSets,omitempty"`
	Releases            []*releases.ReleaseUsage                  `json:"Releases,omitempty"`
	RunbookProcesses    []*runbooks.RunbookStepUsage              `json:"RunbookProcesses,omitempty"`
	RunbookSnapshots    []*runbooks.RunbookSnapshotUsage          `json:"RunbookSnapshots,omitempty"`
	Targets             []*TargetUsageEntry                       `json:"Targets,omitempty"`

	resources.Resource
}

// NewAccountUsage initializes an AccountUsage.
func NewAccountUsage() *AccountUsage {
	accountUsage := &AccountUsage{
		Resource: *resources.NewResource(),
	}

	return accountUsage
}
