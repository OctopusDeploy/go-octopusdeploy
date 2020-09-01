package model

type ReleaseCreationStrategy struct {
	ChannelID              string                   `json:"ChannelId,omitempty"`
	ReleaseCreationPackage *DeploymentActionPackage `json:"ReleaseCreationPackage,omitempty"`
}
