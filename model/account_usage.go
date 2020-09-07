package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

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

func (a *AccountUsage) GetID() string {
	return a.ID
}

func (a *AccountUsage) Validate() error {
	validate := validator.New()
	err := validate.Struct(a)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return err
	}

	return nil
}

var _ ResourceInterface = &AccountUsage{}
