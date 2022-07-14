package channels

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/packages"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type ChannelRule struct {
	ActionPackages []packages.DeploymentActionPackage `json:"ActionPackages,omitempty"`
	ID             string                             `json:"Id,omitempty"`
	Tag            string                             `json:"Tag,omitempty"`

	//Use the NuGet or Maven versioning syntax (depending on the feed type)
	//to specify the range of versions to include
	VersionRange string `json:"VersionRange,omitempty"`

	resources.Resource
}
