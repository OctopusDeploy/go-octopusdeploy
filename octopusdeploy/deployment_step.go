package octopusdeploy

import "github.com/go-playground/validator/v10"

type DeploymentStep struct {
	Actions            []DeploymentAction               `json:"Actions,omitempty"`
	Condition          DeploymentStepConditionType      `json:"Condition,omitempty"` // variable option adds a Property "Octopus.Action.ConditionVariableExpression"
	Name               string                           `json:"Name"`
	PackageRequirement DeploymentStepPackageRequirement `json:"PackageRequirement,omitempty"`
	Properties         map[string]PropertyValue         `json:"Properties,omitempty"`
	StartTrigger       DeploymentStepStartTrigger       `json:"StartTrigger,omitempty" validate:"required,oneof=StartAfterPrevious StartWithPrevious"`

	resource
}

// NewDeploymentStep initializes a DeploymentStep with a name.
func NewDeploymentStep(name string) *DeploymentStep {
	return &DeploymentStep{
		Actions:    []DeploymentAction{},
		Name:       name,
		Properties: map[string]PropertyValue{},
		resource:   *newResource(),
	}
}

// Validate checks the state of the deployment step and returns an error if
// invalid.
func (d DeploymentStep) Validate() error {
	return validator.New().Struct(d)
}
