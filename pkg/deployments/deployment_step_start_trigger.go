package deployments

type DeploymentStepStartTrigger string

const (
	DeploymentStepStartTriggerStartAfterPrevious DeploymentStepStartTrigger = "StartAfterPrevious"
	DeploymentStepStartTriggerStartWithPrevious  DeploymentStepStartTrigger = "StartWithPrevious"
)
