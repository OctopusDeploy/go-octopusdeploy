package octopusdeploy

type DeploymentStepConditionType string

const (
	DeploymentStepConditionTypeSuccess  DeploymentStepConditionType = "Success"
	DeploymentStepConditionTypeFailure  DeploymentStepConditionType = "Failure"
	DeploymentStepConditionTypeAlways   DeploymentStepConditionType = "Always"
	DeploymentStepConditionTypeVariable DeploymentStepConditionType = "Variable"
)
