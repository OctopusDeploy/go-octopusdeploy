package model

type DeploymentStep struct {
	ID                 string                           `json:"Id,omitempty"`
	Name               string                           `json:"Name"`
	PackageRequirement DeploymentStepPackageRequirement `json:"PackageRequirement,omitempty"`                                         // may need its own model / enum
	Properties         map[string]string                `json:"Properties"`                                                           // TODO: refactor to use the PropertyValueResource for handling sensitive values - https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
	Condition          DeploymentStepCondition          `json:"Condition,omitempty" validate:"oneof=Success Failure Always Variable"` // variable option adds a Property "Octopus.Action.ConditionVariableExpression"
	StartTrigger       DeploymentStepStartTrigger       `json:"StartTrigger,omitempty" validate:"oneof=StartAfterPrevious StartWithPrevious"`
	Actions            []DeploymentAction               `json:"Actions,omitempty"`
}
