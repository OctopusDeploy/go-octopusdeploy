package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
)

type Resource struct {
	ID string `json:"Id,omitempty"`

	IResource
}

func NewResource() *Resource {
	return &Resource{}
}

// GetID returns the ID value of the Resource.
func (r *Resource) GetID() string {
	return r.ID
}

// SetID sets the ID value of the Resource.
func (r *Resource) SetID(id string) {
	r.ID = id
}

// Validate checks the state of the Resource and returns an error if invalid.
func (r *Resource) Validate() error {
	return validator.New().Struct(r)
}

var _ IResource = &Resource{}
