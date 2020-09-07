package model

import (
	"fmt"
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

func (a *Artifact) GetID() string {
	return a.ID
}

func (a *Artifact) Validate() error {
	validate := validator.New()
	err := validate.Struct(a)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return err
	}

	return nil
}

var _ ResourceInterface = &Artifact{}
