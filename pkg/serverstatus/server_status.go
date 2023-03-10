package serverstatus

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type ServerStatus struct {
	IsDatabaseEncrypted                     bool   `json:"IsDatabaseEncrypted"`
	IsMajorMinorUpgrade                     bool   `json:"IsMajorMinorUpgrade"`
	IsInMaintenanceMode                     bool   `json:"IsInMaintenanceMode"`
	IsUpgradeAvailable                      bool   `json:"IsUpgradeAvailable"`
	MaintenanceExpires                      string `json:"MaintenanceExpires,omitempty"`
	MaximumAvailableVersion                 string `json:"MaximumAvailableVersion,omitempty"`
	MaximumAvailableVersionCoveredByLicense string `json:"MaximumAvailableVersionCoveredByLicense,omitempty"`

	resources.Resource
}
