package releases

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
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
	SelectedPackages                   []*packages.SelectedPackage              `json:"SelectedPackages,omitempty"`
	SpaceID                            string                                   `json:"SpaceId,omitempty"`
	Version                            string                                   `json:"Version"`

	resources.Resource
}

func NewRelease(channelID string, projectID string, version string) *Release {
	return &Release{
		ChannelID: channelID,
		ProjectID: projectID,
		Version:   version,
		Resource:  *resources.NewResource(),
	}
}
