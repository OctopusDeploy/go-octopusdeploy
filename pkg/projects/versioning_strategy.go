package projects

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/channels"

type VersioningStrategy struct {
	DonorPackage       *channels.DeploymentActionPackage `json:"DonorPackage,omitempty"`
	DonorPackageStepID *string                           `json:"DonorPackageStepId,omitempty"`
	Template           string                            `json:"Template,omitempty"`
}
