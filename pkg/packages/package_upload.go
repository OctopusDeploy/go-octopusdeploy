package packages

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

type PackageUploadResponse struct {
	PackageSizeBytes int
	Hash             string
	NuGetPackageId   string
	PackageId        string
	NuGetFeedId      string
	FeedId           string
	Title            string
	Summary          string
	Version          string
	Description      string
	Published        *time.Time
	ReleaseNotes     string
	FileExtension    string
	// PackageVersionBuildInformation buildinformation.PackageVersionBuildInformation

	resources.Resource
}
