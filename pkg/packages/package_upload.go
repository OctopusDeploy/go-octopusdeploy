package packages

import (
	"io"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// PackageUploadCommand has no json tags because it is never sent to the server; we use a multipart/binary form post
// to send packages, not JSON
type PackageUploadCommand struct {
	FileName   string    // the name of the file (no directory information)
	FileReader io.Reader // provides the contents

	OverwriteMode OverwriteMode `uri:"overwriteMode,omitempty"` // sent as a querystring parameter
	Replace       string        `uri:"replace,omitempty"`       // sent as a querystring parameter
	SpaceID       string        `uri:"spaceId"`                 // sent in the URL route
}

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

func NewPackageUploadCommand(spaceID string) *PackageUploadCommand {
	return &PackageUploadCommand{
		SpaceID: spaceID,
	}
}
