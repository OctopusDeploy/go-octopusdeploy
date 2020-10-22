package octopusdeploy

type TriggerAction struct {
	ActionType string `json:"ActionType" validate:"required,oneof=AutoDeploy DeployLatestRelease DeployNewRelease RunRunbook"`

	resource
}

func NewTriggerAction() *TriggerAction {
	return &TriggerAction{
		resource: *newResource(),
	}
}
