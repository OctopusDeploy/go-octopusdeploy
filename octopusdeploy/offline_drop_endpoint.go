package octopusdeploy

type OfflineDropEndpoint struct {
	ApplicationsDirectory                string                  `json:"ApplicationsDirectory,omitempty"`
	Destination                          *OfflineDropDestination `json:"Destination"`
	WorkingDirectory                     string                  `json:"OctopusWorkingDirectory,omitempty"`
	SensitiveVariablesEncryptionPassword *SensitiveValue         `json:"SensitiveVariablesEncryptionPassword"`

	endpoint
}

func NewOfflineDropEndpoint() *OfflineDropEndpoint {
	return &OfflineDropEndpoint{
		endpoint: *newEndpoint("OfflineDrop"),
	}
}
