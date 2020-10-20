package octopusdeploy

import "github.com/go-playground/validator/v10"

type DeploymentStep struct {
	Actions            []DeploymentAction               `json:"Actions,omitempty"`
	Condition          DeploymentStepCondition          `json:"Condition,omitempty" validate:"required,oneof=Success Failure Always Variable"` // variable option adds a Property "Octopus.Action.ConditionVariableExpression"
	Name               string                           `json:"Name"`
	PackageRequirement DeploymentStepPackageRequirement `json:"PackageRequirement,omitempty"` // may need its own model / enum
	Properties         map[string]string                `json:"Properties"`                   // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	StartTrigger       DeploymentStepStartTrigger       `json:"StartTrigger,omitempty" validate:"required,oneof=StartAfterPrevious StartWithPrevious"`

	Resource
}

// NewDeploymentStep initializes a DeploymentStep with a name.
func NewDeploymentStep(name string) *DeploymentStep {
	return &DeploymentStep{
		Actions:    []DeploymentAction{},
		Name:       name,
		Properties: map[string]string{},
		Resource:   *newResource(),
	}
}

// Validate checks the state of the deployment step and returns an error if
// invalid.
func (d DeploymentStep) Validate() error {
	return validator.New().Struct(d)
}
