package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type DeploymentStep struct {
	Actions            []DeploymentAction               `json:"Actions,omitempty"`
	Condition          DeploymentStepCondition          `json:"Condition,omitempty" validate:"oneof=Success Failure Always Variable"` // variable option adds a Property "Octopus.Action.ConditionVariableExpression"
	Name               string                           `json:"Name"`
	PackageRequirement DeploymentStepPackageRequirement `json:"PackageRequirement,omitempty"` // may need its own model / enum
	Properties         map[string]string                `json:"Properties"`                   // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	StartTrigger       DeploymentStepStartTrigger       `json:"StartTrigger,omitempty" validate:"oneof=StartAfterPrevious StartWithPrevious"`

	Resource
}

// NewDeploymentStep initializes a DeploymentStep with a name.
func NewDeploymentStep(name string) (*DeploymentStep, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewDeploymentStep", "name")
	}

	return &DeploymentStep{
		Name: name,
	}, nil
}

// GetID returns the ID value of the DeploymentStep.
func (resource DeploymentStep) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this DeploymentStep.
func (resource DeploymentStep) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this DeploymentStep was changed.
func (resource DeploymentStep) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this DeploymentStep.
func (resource DeploymentStep) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the DeploymentStep and returns an error if invalid.
func (resource DeploymentStep) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &DeploymentStep{}
