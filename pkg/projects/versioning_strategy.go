package projects

import "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"

type VersioningStrategy struct {
	DonorPackage       *packages.DeploymentActionPackage `json:"DonorPackage,omitempty"`
	DonorPackageStepID *string                           `json:"DonorPackageStepId,omitempty"`
	Template           string                            `json:"Template,omitempty"`
}
