package machinepolicies

type MachineUpdatePolicy struct {
	CalamariUpdateBehavior        string `json:"CalamariUpdateBehavior" validate:"required,oneof=UpdateAlways UpdateOnDeployment UpdateOnNewMachine"`
	TentacleUpdateAccountID       string `json:"TentacleUpdateAccountId,omitempty"`
	TentacleUpdateBehavior        string `json:"TentacleUpdateBehavior" validate:"required,oneof=NeverUpdate Update"`
	KubernetesAgentUpdateBehavior string `json:"KubernetesAgentUpdateBehavior" validate:"required,oneof=NeverUpdate Update Block"`
}

func NewMachineUpdatePolicy() *MachineUpdatePolicy {
	return &MachineUpdatePolicy{
		CalamariUpdateBehavior:        "UpdateOnDeployment",
		TentacleUpdateBehavior:        "NeverUpdate",
		KubernetesAgentUpdateBehavior: "Update",
	}
}
