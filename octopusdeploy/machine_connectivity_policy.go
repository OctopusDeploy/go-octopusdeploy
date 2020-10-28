package octopusdeploy

type MachineConnectivityPolicy struct {
	MachineConnectivityBehavior string `json:"MachineConnectivityBehavior" validate:"oneof=ExpectedToBeOnline MayBeOfflineAndCanBeSkipped"`
}

func NewMachineConnectivityPolicy() *MachineConnectivityPolicy {
	return &MachineConnectivityPolicy{
		MachineConnectivityBehavior: "ExpectedToBeOnline",
	}
}
