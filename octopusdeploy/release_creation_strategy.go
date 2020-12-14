package octopusdeploy

type ReleaseCreationStrategy struct {
	ChannelID                    string                   `json:"ChannelId,omitempty"`
	ReleaseCreationPackage       *DeploymentActionPackage `json:"ReleaseCreationPackage,omitempty"`
	ReleaseCreationPackageStepID string                   `json:"ReleaseCreationPackageStepId,omitempty"`
}
