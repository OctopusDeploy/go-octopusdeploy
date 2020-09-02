package model

type CloudServiceMachineEndpoint struct {
	CloudServiceName        string `json:CloudServiceName`
	Slot                    string `json:Slot`
	StorageAccountName      string `json:StorageAccountName`
	SwapIfPossible          bool   `json:SwapIfPossible`
	UseCurrentInstanceCount bool   `json:UseCurrentInstanceCount`
}
