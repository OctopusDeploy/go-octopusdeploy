package serverstatus

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type ServerStatus struct {
	IsDatabaseEncrypted                     bool   `json:"IsDatabaseEncrypted,omitempty"`
	IsMajorMinorUpgrade                     bool   `json:"IsMajorMinorUpgrade,omitempty"`
	IsInMaintenanceMode                     bool   `json:"IsInMaintenanceMode,omitempty"`
	IsUpgradeAvailable                      bool   `json:"IsUpgradeAvailable,omitempty"`
	MaintenanceExpires                      string `json:"MaintenanceExpires,omitempty"`
	MaximumAvailableVersion                 string `json:"MaximumAvailableVersion,omitempty"`
	MaximumAvailableVersionCoveredByLicense string `json:"MaximumAvailableVersionCoveredByLicense,omitempty"`

	resources.Resource
}
