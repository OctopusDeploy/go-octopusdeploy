package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Artifacts defines a collection of Artifact types with built-in support for
// paged results from the API.
type Artifacts struct {
	Items []Artifact `json:"Items"`
	PagedResults
}

type Artifact struct {
	Created          time.Time `json:"Created,omitempty"`
	Filename         *string   `json:"Filename"`
	LogCorrelationID string    `json:"LogCorrelationId,omitempty"`
	ServerTaskID     string    `json:"ServerTaskId,omitempty"`
	Source           string    `json:"Source,omitempty"`
	SpaceID          string    `json:"SpaceId,omitempty"`

	Resource
}

// GetID returns the ID value of the Artifact.
func (resource Artifact) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Artifact.
func (resource Artifact) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Artifact was changed.
func (resource Artifact) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Artifact.
func (resource Artifact) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the Artifact and returns an error if invalid.
func (resource Artifact) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	return nil
}

var _ ResourceInterface = &Artifact{}
