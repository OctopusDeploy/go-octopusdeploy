package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Resources struct {
	Items []*Resource `json:"Items"`
	PagedResults
}

type Resource struct {
	ID             string            `json:"Id,omitempty"`
	LastModifiedBy string            `json:"LastModifiedBy,omitempty"`
	LastModifiedOn *time.Time        `json:"LastModifiedOn,omitempty"`
	Links          map[string]string `json:"Links,omitempty"`
}

func newResource() *Resource {
	return &Resource{
		Links: map[string]string{},
	}
}

// GetID returns the ID value of the resource.
func (r *Resource) GetID() string {
	return r.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of
// this resource.
func (r *Resource) GetLastModifiedBy() string {
	return r.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this resource was
// changed.
func (r *Resource) GetLastModifiedOn() *time.Time {
	return r.LastModifiedOn
}

// GetLinks returns the associated links with the value of this resource.
func (r *Resource) GetLinks() map[string]string {
	return r.Links
}

// SetID sets the ID value of this resource.
func (r *Resource) SetID(id string) {
	r.ID = id
}

// SetLastModifiedBy sets the name of the account that modified the value of
// this resource.
func (r *Resource) SetLastModifiedBy(name string) {
	r.LastModifiedBy = name
}

// SetLastModifiedOn sets the time when the value of this resource was changed.
func (r *Resource) SetLastModifiedOn(time *time.Time) {
	r.LastModifiedOn = time
}

// Validate checks the state of the resource and returns an error if invalid.
func (r *Resource) Validate() error {
	return validator.New().Struct(r)
}

var _ IResource = &Resource{}
