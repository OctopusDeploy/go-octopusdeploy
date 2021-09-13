package octopusdeploy

type scopedDeploymentAction struct {
	Channel    string   `json:"ChannelId,omitempty"`
	Tenants    []string `json:"TenantIds,omitempty"`
	TenantTags []string `json:"TenantTags,omitempty"`

	triggerAction
}

func newScopedDeploymentAction(actionType ActionType) *scopedDeploymentAction {
	return &scopedDeploymentAction{
		Tenants:       []string{},
		TenantTags:    []string{},
		triggerAction: *newTriggerAction(actionType),
	}
}

func (a *scopedDeploymentAction) GetActionType() ActionType {
	return a.Type
}

func (a *scopedDeploymentAction) SetActionType(actionType ActionType) {
	a.Type = actionType
}

var _ ITriggerAction = &scopedDeploymentAction{}
