package packages

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/buildinformation"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type PackageUploadCommand struct {
	Contents      []byte        `json:"Contents" uri:"-"`
	FileName      string        `json:"FileName" uri:"-"`
	OverwriteMode OverwriteMode `json:"-" uri:"overwriteMode,omitempty"`
	Replace       string        `json:"-" uri:"replace,omitempty"`
	SpaceID       string        `json:"-" uri:"spaceId"`
}

type PackageUploadResponse struct {
	PackageSizeBytes               int
	Hash                           string
	NuGetPackageId                 string
	PackageId                      string
	NuGetFeedId                    string
	FeedId                         string
	Title                          string
	Summary                        string
	Version                        string
	Description                    string
	Published                      *time.Time
	ReleaseNotes                   string
	FileExtension                  string
	PackageVersionBuildInformation buildinformation.PackageVersionBuildInformation

	resources.Resource
}

func NewPackageUploadCommand(spaceID string) *PackageUploadCommand {
	return &PackageUploadCommand{
		SpaceID: spaceID,
	}
}
