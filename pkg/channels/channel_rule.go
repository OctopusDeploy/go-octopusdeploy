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

	// VersioningStrategy controls how packages are ordered to determine the latest version.
	// Set to "MostRecentlyPublished" to use publish date ordering instead of SemVer comparison.
	// When unset or "SemVer", the existing behaviour applies.
	VersioningStrategy string `json:"VersioningStrategy,omitempty"`

	// VersionTagRegex is a regex matched against the full version/tag string.
	// Used with VersioningStrategy "MostRecentlyPublished" as an alternative to
	// VersionRange and Tag, supporting non-SemVer versioning schemes.
	VersionTagRegex string `json:"VersionTagRegex,omitempty"`

	resources.Resource
}
