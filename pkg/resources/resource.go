package resources

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// IResource defines the interface for resources.
type IResource interface {
	GetID() string
	GetModifiedBy() string
	GetModifiedOn() *time.Time
	GetLinks() map[string]string
	SetID(string)
	SetLinks(map[string]string)
	SetModifiedBy(string)
	SetModifiedOn(*time.Time)
	Validate() error
}

type Resource struct {
	ID         string            `json:"Id,omitempty"`
	ModifiedBy string            `json:"LastModifiedBy,omitempty"`
	ModifiedOn *time.Time        `json:"LastModifiedOn,omitempty"`
	Links      map[string]string `json:"Links,omitempty"`
}

func NewResource() *Resource {
	return &Resource{
		Links: map[string]string{},
	}
}

// GetID returns the ID value of the resource.
func (r *Resource) GetID() string {
	return r.ID
}

// GetModifiedBy returns the name of the account that modified the value of
// this resource.
func (r *Resource) GetModifiedBy() string {
	return r.ModifiedBy
}

// GetModifiedOn returns the time when the value of this resource was changed.
func (r *Resource) GetModifiedOn() *time.Time {
	return r.ModifiedOn
}

// GetLinks returns the associated links with the value of this resource.
func (r *Resource) GetLinks() map[string]string {
	return r.Links
}

// SetID sets the ID value of the resource.
func (r *Resource) SetID(id string) {
	r.ID = id
}

// SetLinks sets the associated links with the value of this resource.
func (r *Resource) SetLinks(links map[string]string) {
	r.Links = links
}

// SetModifiedBy set the name of the account that modified the value of
// this resource.
func (r *Resource) SetModifiedBy(modifiedBy string) {
	r.ModifiedBy = modifiedBy
}

// SetModifiedOn set the time when the value of this resource was changed.
func (r *Resource) SetModifiedOn(modifiedOn *time.Time) {
	r.ModifiedOn = modifiedOn
}

// Validate checks the state of the resource and returns an error if invalid.
func (r *Resource) Validate() error {
	return validator.New().Struct(r)
}

var _ IResource = &Resource{}
