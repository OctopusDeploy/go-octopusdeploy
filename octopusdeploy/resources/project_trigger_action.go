package resources

type ProjectTriggerAction struct {
	ActionType                                 string `json:"ActionType"`
	DestinationEnvironmentID                   string `json:"DestinationEnvironmentId"`
	ShouldRedeployWhenMachineHasBeenDeployedTo bool   `json:"ShouldRedeployWhenMachineHasBeenDeployedTo"`
	SourceEnvironmentID                        string `json:"SourceEnvironmentId"`
}

func (a *ProjectTriggerAction) GetActionType() ActionType {
	actionType, _ := ActionTypeString(a.ActionType)
	return actionType
}

func (a *ProjectTriggerAction) SetActionType(actionType ActionType) {
	a.ActionType = actionType.String()
}

var _ ITriggerAction = &ProjectTriggerAction{}
