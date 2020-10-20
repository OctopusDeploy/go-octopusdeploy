package octopusdeploy

type MachineUpdatePolicy struct {
	CalamariUpdateBehavior  string `json:"CalamariUpdateBehavior" validate:"required,oneof=UpdateAlways UpdateOnDeployment UpdateOnNewMachine"`
	TentacleUpdateAccountID string `json:"TentacleUpdateAccountId,omitempty"`
	TentacleUpdateBehavior  string `json:"TentacleUpdateBehavior" validate:"required,oneof=NeverUpdate Update"`
}

func NewMachineUpdatePolicy() *MachineUpdatePolicy {
	return &MachineUpdatePolicy{
		CalamariUpdateBehavior: "UpdateOnDeployment",
		TentacleUpdateBehavior: "NeverUpdate",
	}
}
