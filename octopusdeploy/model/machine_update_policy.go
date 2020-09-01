package model

type MachineUpdatePolicy struct {
	CalamariUpdateBehavior  string `json:"CalamariUpdateBehavior,omitempty"`
	TentacleUpdateAccountID string `json:"TentacleUpdateAccountId,omitempty"`
	TentacleUpdateBehavior  string `json:"TentacleUpdateBehavior,omitempty"`
}
