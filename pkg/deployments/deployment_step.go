package deployments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type DeploymentStep struct {
	Actions            []*DeploymentAction              `json:"Actions,omitempty"`
	Condition          DeploymentStepConditionType      `json:"Condition,omitempty"` // variable option adds a Property "Octopus.Action.ConditionVariableExpression"
	Name               string                           `json:"Name"`
	PackageRequirement DeploymentStepPackageRequirement `json:"PackageRequirement,omitempty"`
	Properties         map[string]core.PropertyValue    `json:"Properties,omitempty"`
	StartTrigger       DeploymentStepStartTrigger       `json:"StartTrigger,omitempty" validate:"required,oneof=StartAfterPrevious StartWithPrevious"`
	TargetRoles        []string                         `json:"-"`

	resources.Resource
}

// NewDeploymentStep initializes a DeploymentStep with a name.
func NewDeploymentStep(name string) *DeploymentStep {
	return &DeploymentStep{
		Actions:            []*DeploymentAction{},
		Condition:          "Success",
		Name:               name,
		PackageRequirement: "LetOctopusDecide",
		Properties:         map[string]core.PropertyValue{},
		StartTrigger:       "StartAfterPrevious",
		TargetRoles:        []string{},
		Resource:           *resources.NewResource(),
	}
}

// Validate checks the state of the deployment step and returns an error if
// invalid.
func (d DeploymentStep) Validate() error {
	return validator.New().Struct(d)
}
