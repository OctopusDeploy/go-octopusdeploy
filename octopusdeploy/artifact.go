package octopusdeploy

import (
	"time"
)

// Artifacts defines a collection of artifacts with built-in support for paged
// results from the API.
type Artifacts struct {
	Items []*Artifact `json:"Items"`
	PagedResults
}

// Artifact represents an artifact.
type Artifact struct {
	Created          *time.Time `json:"Created,omitempty"`
	Filename         string     `json:"Filename" validate:"required"`
	LogCorrelationID string     `json:"LogCorrelationId,omitempty"`
	ServerTaskID     string     `json:"ServerTaskId,omitempty"`
	Source           string     `json:"Source,omitempty"`
	SpaceID          string     `json:"SpaceId,omitempty"`

	resource
}

// NewArtifact creates and initializes an artifact.
func NewArtifact(filename string) *Artifact {
	return &Artifact{
		Filename: filename,
		resource: *newResource(),
	}
}

var _ IResource = &Artifact{}
