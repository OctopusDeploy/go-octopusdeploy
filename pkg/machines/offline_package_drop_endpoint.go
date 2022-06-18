package machines

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/core"

type OfflinePackageDropEndpoint struct {
	ApplicationsDirectory                string                         `json:"ApplicationsDirectory,omitempty"`
	Destination                          *OfflinePackageDropDestination `json:"Destination"`
	SensitiveVariablesEncryptionPassword *core.SensitiveValue                `json:"SensitiveVariablesEncryptionPassword"`
	WorkingDirectory                     string                         `json:"OctopusWorkingDirectory,omitempty"`

	endpoint
}

func NewOfflinePackageDropEndpoint() *OfflinePackageDropEndpoint {
	return &OfflinePackageDropEndpoint{
		Destination: NewOfflinePackageDropDestination(),
		endpoint:    *newEndpoint("OfflineDrop"),
	}
}
