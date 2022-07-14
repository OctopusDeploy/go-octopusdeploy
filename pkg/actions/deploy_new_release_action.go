package actions

type DeployNewReleaseAction struct {
	Environment             string                   `json:"EnvironmentId,omitempty"`
	Variables               string                   `json:"Variable,omitempty"`
	VersionControlReference *VersionControlReference `json:"VersionControlReference,omitempty"`

	scopedDeploymentAction
}

func NewDeployNewReleaseAction(environment string, variables string, versionControlReference *VersionControlReference) *DeployNewReleaseAction {
	return &DeployNewReleaseAction{
		Environment:             environment,
		Variables:               variables,
		VersionControlReference: versionControlReference,
		scopedDeploymentAction:  *newScopedDeploymentAction(DeployNewRelease),
	}
}

func (a *DeployNewReleaseAction) GetActionType() ActionType {
	return a.Type
}

func (a *DeployNewReleaseAction) SetActionType(actionType ActionType) {
	a.Type = actionType
}

var _ ITriggerAction = &DeployNewReleaseAction{}
