package resources

import "time"

type AuditedResource struct {
	modifiedBy string     `json:"LastModifiedBy,omitempty"`
	modifiedOn *time.Time `json:"LastModifiedOn,omitempty"`

	IAuditedResource
}

// GetModifiedBy returns the name of the accountV1 that modified the value of
// this Resource.
func (r *AuditedResource) GetModifiedBy() string {
	return r.modifiedBy
}

// SetModifiedBy set the name of the accountV1 that modified the value of
// this Resource.
func (r *AuditedResource) SetModifiedBy(modifiedBy string) {
	r.modifiedBy = modifiedBy
}

// GetModifiedOn returns the time when the value of this Resource was changed.
func (r *AuditedResource) GetModifiedOn() *time.Time {
	return r.modifiedOn
}

// SetModifiedOn set the time when the value of this Resource was changed.
func (r *AuditedResource) SetModifiedOn(modifiedOn *time.Time) {
	r.modifiedOn = modifiedOn
}
