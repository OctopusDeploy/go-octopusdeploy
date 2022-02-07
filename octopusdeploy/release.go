package octopusdeploy

import (
	"time"
)

type Release struct {
	Assembled                          time.Time                                `json:"Assembled,omitempty"`
	BuildInformation                   []*ReleasePackageVersionBuildInformation `json:"BuildInformation,omitempty"`
	ChannelID                          string                                   `json:"ChannelId,omitempty"`
	IgnoreChannelRules                 bool                                     `json:"IgnoreChannelRules,omitempty"`
	LibraryVariableSetSnapshotIDs      []string                                 `json:"LibraryVariableSetSnapshotIds,omitempty"`
	ProjectDeploymentProcessSnapshotID string                                   `json:"ProjectDeploymentProcessSnapshotId,omitempty"`
	ProjectID                          string                                   `json:"ProjectId,omitempty"`
	ProjectVariableSetSnapshotID       string                                   `json:"ProjectVariableSetSnapshotId,omitempty"`
	ReleaseNotes                       string                                   `json:"ReleaseNotes,omitempty"`
	SelectedPackages                   []*SelectedPackage                       `json:"SelectedPackages,omitempty"`
	SpaceID                            string                                   `json:"SpaceId,omitempty"`
	Version                            string                                   `json:"Version"`

	Resource
}

func NewRelease(channelID string, projectID string, version string) *Release {
	return &Release{
		ChannelID: channelID,
		ProjectID: projectID,
		Version:   version,
		Resource:  *newResource(),
	}
}
