package actions

type CreateReleaseAction struct {
	ChannelID string `json:"ChannelId,omitempty"`

	triggerAction
}

func NewCreateReleaseAction(channel string) *CreateReleaseAction {
	return &CreateReleaseAction{
		ChannelID:     channel,
		triggerAction: *newTriggerAction(CreateRelease),
	}
}

func (a *CreateReleaseAction) GetActionType() ActionType {
	return a.Type
}

func (a *CreateReleaseAction) SetActionType(actionType ActionType) {
	a.Type = actionType
}

var _ ITriggerAction = &CreateReleaseAction{}
