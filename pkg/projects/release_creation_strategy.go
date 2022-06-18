package projects

import "github.com/OctopusDeploy/go-octopusdeploy/pkg/channels"

type ReleaseCreationStrategy struct {
	ChannelID                    string                            `json:"ChannelId,omitempty"`
	ReleaseCreationPackage       *channels.DeploymentActionPackage `json:"ReleaseCreationPackage,omitempty"`
	ReleaseCreationPackageStepID string                            `json:"ReleaseCreationPackageStepId,omitempty"`
}
