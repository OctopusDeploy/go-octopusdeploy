package packages

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/buildinformation"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
)

type Package struct {
	Description      string                             `json:"Description,omitempty"`
	FeedID           string                             `json:"FeedId,omitempty"`
	FileExtension    string                             `json:"FileExtension,omitempty"`
	NuGetFeedID      string                             `json:"NuGetFeedId,omitempty"`
	NuGetPackageID   string                             `json:"NuGetPackageId,omitempty"`
	PackageID        string                             `json:"PackageId,omitempty"`
	BuildInformation *buildinformation.BuildInformation `json:"PackageVersionBuildInformation,omitempty"`
	Published        time.Time                          `json:"ReleaseNotes,omitempty"`
	Summary          string                             `json:"Summary,omitempty"`
	Title            string                             `json:"Title,omitempty"`
	Version          string                             `json:"Version,omitempty"`

	resources.Resource
}

type Packages struct {
	Items []*Package `json:"Items"`
	resources.PagedResults
}

func NewPackage() *Package {
	return &Package{
		Resource: *resources.NewResource(),
	}
}
