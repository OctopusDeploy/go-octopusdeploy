package resources

type DeployLatestReleaseAction struct {
	DestinationEnvironment string   `json:"DestinationEnvironmentId"`
	ShouldRedeploy         bool     `json:"ShouldRedeployWhenReleaseIsCurrent"`
	SourceEnvironments     []string `json:"SourceEnvironmentIds"`
	Variables              string   `json:"Variables"`

	scopedDeploymentAction
}

func NewDeployLatestReleaseAction(destinationEnvironment string, shouldRedeploy bool, sourceEnvironments []string, variables string) *DeployLatestReleaseAction {
	return &DeployLatestReleaseAction{
		DestinationEnvironment: destinationEnvironment,
		ShouldRedeploy:         shouldRedeploy,
		SourceEnvironments:     sourceEnvironments,
		Variables:              variables,
		scopedDeploymentAction: *newScopedDeploymentAction(DeployLatestRelease),
	}
}

func (a *DeployLatestReleaseAction) GetActionType() ActionType {
	return a.Type
}

func (a *DeployLatestReleaseAction) SetActionType(actionType ActionType) {
	a.Type = actionType
}

var _ ITriggerAction = &DeployLatestReleaseAction{}
