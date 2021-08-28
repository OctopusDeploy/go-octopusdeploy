package octopusdeploy

type TriggerAction struct {
	ActionType                                 string   `json:"ActionType" validate:"required,oneof=AutoDeploy DeployLatestRelease DeployNewRelease RunRunbook"`
	ChannelID                                  string   `json:"ChannelId,omitempty"`
	DestinationEnvironmentID                   string   `json:"DestinationEnvironmentId,omitempty"`
	SourceEnvironmentIDs                       []string `json:"SourceEnvironmentIds,omitempty"`
	RedeployCurrent                            bool     `json:"ShouldRedeployWhenReleaseIsCurrent,omitempty"`
	ShouldRedeployWhenMachineHasBeenDeployedTo bool     `json:"ShouldRedeployWhenMachineHasBeenDeployedTo,omitempty"`
	RunbookID                                  string   `json:"RunbookId,omitempty"`
	TenantIDs                                  []string `json:"TenantIds,omitempty"`
	TenantTags                                 []string `json:"TenantTags,omitempty"`
	Variables                                  string   `json:"Variables,omitempty"`

	resource
}

func NewTriggerAction() *TriggerAction {
	return &TriggerAction{
		resource: *newResource(),
	}
}
