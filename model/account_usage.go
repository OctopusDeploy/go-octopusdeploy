package model

type AccountUsage struct {
	DeploymentProcesses []*StepUsage                    `json:"DeploymentProcesses"`
	LibraryVariableSets []*LibraryVariableSetUsageEntry `json:"LibraryVariableSets"`
	ProjectVariableSets []*ProjectVariableSetUsage      `json:"ProjectVariableSets"`
	Releases            []*ReleaseUsage                 `json:"Releases"`
	RunbookProcesses    []*RunbookStepUsage             `json:"RunbookProcesses"`
	RunbookSnapshots    []*RunbookSnapshotUsage         `json:"RunbookSnapshots"`
	Targets             []*TargetUsageEntry             `json:"Targets"`
	Resource
}
