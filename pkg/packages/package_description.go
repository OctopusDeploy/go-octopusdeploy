package packages

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type PackageDescription struct {
	Description   string `json:"Description,omitempty"`
	LatestVersion string `json:"LatestVersion,omitempty"`
	Name          string `json:"Name,omitempty"`

	resources.Resource
}
