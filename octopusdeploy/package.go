package octopusdeploy

import "time"

type Package struct {
	Description      string            `json:"Description,omitempty"`
	FeedID           string            `json:"FeedId,omitempty"`
	FileExtension    string            `json:"FileExtension,omitempty"`
	NuGetFeedID      string            `json:"NuGetFeedId,omitempty"`
	NuGetPackageID   string            `json:"NuGetPackageId,omitempty"`
	PackageID        string            `json:"PackageId,omitempty"`
	BuildInformation *BuildInformation `json:"PackageVersionBuildInformation,omitempty"`
	Published        time.Time         `json:"ReleaseNotes,omitempty"`
	Summary          string            `json:"Summary,omitempty"`
	Title            string            `json:"Title,omitempty"`
	Version          string            `json:"Version,omitempty"`

	resource
}

type Packages struct {
	Items []*Package `json:"Items"`
	PagedResults
}

func NewPackage() *Package {
	return &Package{
		resource: *newResource(),
	}
}
