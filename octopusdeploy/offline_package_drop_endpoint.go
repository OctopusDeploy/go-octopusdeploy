package octopusdeploy

type OfflinePackageDropEndpoint struct {
	ApplicationsDirectory                string                        `json:"ApplicationsDirectory,omitempty"`
	Destination                          OfflinePackageDropDestination `json:"Destination"`
	SensitiveVariablesEncryptionPassword *SensitiveValue               `json:"SensitiveVariablesEncryptionPassword"`
	WorkingDirectory                     string                        `json:"OctopusWorkingDirectory,omitempty"`

	endpoint
}

func NewOfflinePackageDropEndpoint() *OfflinePackageDropEndpoint {
	return &OfflinePackageDropEndpoint{
		Destination: *NewOfflinePackageDropDestination(),
		endpoint:    *newEndpoint("OfflineDrop"),
	}
}
