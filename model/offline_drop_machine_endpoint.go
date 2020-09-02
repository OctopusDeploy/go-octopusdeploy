package model

type OfflineDropMachineEndpoint struct {
	Destination                          OfflineDropDestination `json:"OfflineDropDestination"`
	SensitiveVariablesEncryptionPassword SensitiveValue         `json:"SensitiveVariablesEncryptionPassword"`
	ApplicationsDirectory                string                 `json:"ApplicationsDirectory,omitempty"`
	WorkingDirectory                     string                 `json:"OctopusWorkingDirectory,omitempty"`
}
