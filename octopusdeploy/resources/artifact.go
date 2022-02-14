package resources

import (
	"time"
)

// Artifact represents an artifact.
type Artifact struct {
	Created          *time.Time `json:"Created,omitempty"`
	Filename         string     `json:"Filename" validate:"required"`
	LogCorrelationID string     `json:"LogCorrelationId,omitempty"`
	ServerTaskID     string     `json:"ServerTaskId,omitempty"`
	Source           string     `json:"Source,omitempty"`
	SpaceID          string     `json:"SpaceId,omitempty"`

	Resource
}

// NewArtifact creates and initializes an artifact.
func NewArtifact(filename string) *Artifact {
	return &Artifact{
		Filename: filename,
		Resource: *NewResource(),
	}
}

var _ IResource = &Artifact{}
