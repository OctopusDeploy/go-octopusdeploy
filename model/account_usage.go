package model

// AccountUsage contains the projects and deployments which are using an
// account.
type AccountUsage struct {
	DeploymentProcesses []*StepUsage                    `json:"DeploymentProcesses,omitempty"`
	LibraryVariableSets []*LibraryVariableSetUsageEntry `json:"LibraryVariableSets,omitempty"`
	ProjectVariableSets []*ProjectVariableSetUsage      `json:"ProjectVariableSets,omitempty"`
	Releases            []*ReleaseUsage                 `json:"Releases,omitempty"`
	RunbookProcesses    []*RunbookStepUsage             `json:"RunbookProcesses,omitempty"`
	RunbookSnapshots    []*RunbookSnapshotUsage         `json:"RunbookSnapshots,omitempty"`
	Targets             []*TargetUsageEntry             `json:"Targets,omitempty"`

	Resource
}

// NewAccountUsage initializes an AccountUsage.
func NewAccountUsage() *AccountUsage {
	accountUsage := &AccountUsage{
		Resource: *newResource(),
	}

	return accountUsage
}
