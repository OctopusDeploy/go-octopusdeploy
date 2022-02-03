package octopusdeploy

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Resources struct {
	Items []*resource `json:"Items"`
	PagedResults
}

type resource struct {
	ID         string            `json:"Id,omitempty"`
	ModifiedBy string            `json:"LastModifiedBy,omitempty"`
	ModifiedOn *time.Time        `json:"LastModifiedOn,omitempty"`
}

func newResource() *resource {
	return &resource{
	}
}

// GetID returns the ID value of the resource.
func (r *resource) GetID() string {
	return r.ID
}

// GetModifiedBy returns the name of the account that modified the value of
// this resource.
func (r *resource) GetModifiedBy() string {
	return r.ModifiedBy
}

// GetModifiedOn returns the time when the value of this resource was changed.
func (r *resource) GetModifiedOn() *time.Time {
	return r.ModifiedOn
}

// SetID sets the ID value of the resource.
func (r *resource) SetID(id string) {
	r.ID = id
}

// SetModifiedBy set the name of the account that modified the value of
// this resource.
func (r *resource) SetModifiedBy(modifiedBy string) {
	r.ModifiedBy = modifiedBy
}

// SetModifiedOn set the time when the value of this resource was changed.
func (r *resource) SetModifiedOn(modifiedOn *time.Time) {
	r.ModifiedOn = modifiedOn
}

// Validate checks the state of the resource and returns an error if invalid.
func (r *resource) Validate() error {
	return validator.New().Struct(r)
}

var _ IResource = &resource{}
