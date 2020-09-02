package model

type CloudServiceMachineEndpoint struct {
	CloudServiceName        string `json:"CloudServiceName,omitempty"`
	Slot                    string `json:"Slot,omitempty"`
	StorageAccountName      string `json:"StorageAccountName,omitempty"`
	SwapIfPossible          bool   `json:"SwapIfPossible"`
	UseCurrentInstanceCount bool   `json:"UseCurrentInstanceCount"`
}
