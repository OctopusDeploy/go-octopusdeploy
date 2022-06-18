package packages

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/resources"
)

type PackageVersion struct {
	FeedID       string    `json:"FeedId,omitempty"`
	PackageID    string    `json:"PackageId,omitempty"`
	Published    time.Time `json:"Published,omitempty"`
	ReleaseNotes string    `json:"ReleaseNotes,omitempty"`
	SizeBytes    int64     `json:"SizeBytes,omitempty"`
	Title        string    `json:"Title,omitempty"`
	Version      string    `json:"Version,omitempty"`

	resources.Resource
}

func NewPackageVersion() *PackageVersion {
	return &PackageVersion{
		Resource: *resources.NewResource(),
	}
}
