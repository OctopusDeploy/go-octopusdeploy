package octopusdeploy

type triggerAction struct {
	Type ActionType `json:"ActionType"`

	resource
}

func newTriggerAction(actionType ActionType) *triggerAction {
	return &triggerAction{
		Type:     actionType,
		resource: *newResource(),
	}
}
