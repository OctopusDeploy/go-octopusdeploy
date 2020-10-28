package octopusdeploy

import "time"

type PackageVersion struct {
	FeedID       string    `json:"FeedId,omitempty"`
	PackageID    string    `json:"PackageId,omitempty"`
	Published    time.Time `json:"Published,omitempty"`
	ReleaseNotes string    `json:"ReleaseNotes,omitempty"`
	SizeBytes    int64     `json:"SizeBytes,omitempty"`
	Title        string    `json:"Title,omitempty"`
	Version      string    `json:"Version,omitempty"`

	resource
}

func NewPackageVersion() *PackageVersion {
	return &PackageVersion{
		resource: *newResource(),
	}
}
