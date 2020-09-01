package model

import "time"

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
