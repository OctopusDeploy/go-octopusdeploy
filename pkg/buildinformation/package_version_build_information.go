package buildinformation

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"

type PackageVersionBuildInformation struct {
	PackageID               string                  `json:"PackageId"`
	Version                 string                  `json:"Version"`
	OctopusBuildInformation OctopusBuildInformation `json:"OctopusBuildInformation"`

	resources.Resource
}

func NewPackageVersionBuildInformation() *PackageVersionBuildInformation {
	return &PackageVersionBuildInformation{}
}
