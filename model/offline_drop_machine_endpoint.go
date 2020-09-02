package model

type OfflineDropMachineEndpoint struct {
	Destination                          OfflineDropDestination `json:"OfflineDropDestination"`
	SensitiveVariablesEncryptionPassword SensitiveValue         `json:"SensitiveVariablesEncryptionPassword"`
	ApplicationsDirectory                string                 `json:"ApplicationsDirectory" validate:"required"`
	WorkingDirectory                     string                 `json:"OctopusWorkingDirectory" validate:"required"`
}
