package machinepolicies

type MachineScriptPolicy struct {
	RunType    string  `json:"RunType" validate:"required,oneof=InheritFromDefault Inline OnlyConnectivity"`
	ScriptBody *string `json:"ScriptBody"`
}

func NewMachineScriptPolicy() *MachineScriptPolicy {
	return &MachineScriptPolicy{
		RunType: "InheritFromDefault",
	}
}
