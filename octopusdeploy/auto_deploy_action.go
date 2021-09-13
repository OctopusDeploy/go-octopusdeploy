package octopusdeploy

type AutoDeployAction struct {
	ShouldRedeploy bool `json:"ShouldRedeployWhenMachineHasBeenDeployedTo"`

	triggerAction
}

func NewAutoDeployAction(shouldRedeploy bool) *AutoDeployAction {
	return &AutoDeployAction{
		ShouldRedeploy: shouldRedeploy,
		triggerAction:  *newTriggerAction(AutoDeploy),
	}
}

func (a *AutoDeployAction) GetActionType() ActionType {
	return a.Type
}

func (a *AutoDeployAction) SetActionType(actionType ActionType) {
	a.Type = actionType
}

var _ ITriggerAction = &AutoDeployAction{}
