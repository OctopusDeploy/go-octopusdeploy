package projects

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/packages"

type ReleaseCreationStrategy struct {
	ChannelID                    string                            `json:"ChannelId,omitempty"`
	ReleaseCreationPackage       *packages.DeploymentActionPackage `json:"ReleaseCreationPackage,omitempty"`
	ReleaseCreationPackageStepID string                            `json:"ReleaseCreationPackageStepId,omitempty"`
}
