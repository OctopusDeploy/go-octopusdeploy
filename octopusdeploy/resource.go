package octopusdeploy

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Resource struct {
	ID string `json:"Id,omitempty"`

	IResource
}

type SpaceScopedResource struct {
	SpaceID string `json:"SpaceId,omitempty"`
}

type AuditedResource struct {
	modifiedBy string     `json:"LastModifiedBy,omitempty"`
	modifiedOn *time.Time `json:"LastModifiedOn,omitempty"`

	IAuditedResource
}

func NewResource() *Resource {
	return &Resource{}
}

// GetID returns the ID value of the Resource.
func (r *Resource) GetID() string {
	return r.ID
}

// GetModifiedBy returns the name of the account that modified the value of
// this Resource.
func (r *AuditedResource) GetModifiedBy() string {
	return r.modifiedBy
}

// GetModifiedOn returns the time when the value of this Resource was changed.
func (r *AuditedResource) GetModifiedOn() *time.Time {
	return r.modifiedOn
}

// SetID sets the ID value of the Resource.
func (r *Resource) SetID(id string) {
	r.ID = id
}

// SetModifiedBy set the name of the account that modified the value of
// this Resource.
func (r *AuditedResource) SetModifiedBy(modifiedBy string) {
	r.modifiedBy = modifiedBy
}

// SetModifiedOn set the time when the value of this Resource was changed.
func (r *AuditedResource) SetModifiedOn(modifiedOn *time.Time) {
	r.modifiedOn = modifiedOn
}

// Validate checks the state of the Resource and returns an error if invalid.
func (r *Resource) Validate() error {
	return validator.New().Struct(r)
}

var _ IResource = &Resource{}
