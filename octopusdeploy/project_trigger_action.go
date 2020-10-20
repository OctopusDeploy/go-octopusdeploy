package octopusdeploy

type ProjectTriggerAction struct {
	ActionType                                 string `json:"ActionType"`
	DestinationEnvironmentID                   string `json:"DestinationEnvironmentId"`
	ShouldRedeployWhenMachineHasBeenDeployedTo bool   `json:"ShouldRedeployWhenMachineHasBeenDeployedTo"`
	SourceEnvironmentID                        string `json:"SourceEnvironmentId"`
}
