package model

import (
	"github.com/go-playground/validator/v10"
)

// NewAccountUsage initializes an AccountUsage.
func NewAccountUsage() (*AccountUsage, error) {
	return &AccountUsage{}, nil
}

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

func (a *AccountUsage) GetID() string {
	return a.ID
}

func (a *AccountUsage) Validate() error {
	validate := validator.New()
	err := validate.Struct(a)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	return nil
}

var _ ResourceInterface = &AccountUsage{}
