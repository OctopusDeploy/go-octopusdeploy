package actions

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"

// ITriggerAction defines the interface for trigger actions.
type ITriggerAction interface {
	GetActionType() ActionType
	SetActionType(actionType ActionType)
}

type triggerAction struct {
	Type ActionType `json:"ActionType"`

	resources.Resource
}

func newTriggerAction(actionType ActionType) *triggerAction {
	return &triggerAction{
		Type:     actionType,
		Resource: *resources.NewResource(),
	}
}
