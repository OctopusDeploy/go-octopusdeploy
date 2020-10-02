package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// APIKeys defines a collection of API keys with built-in support for paged
// results.
type APIKeys struct {
	Items []APIKey `json:"Items"`
	PagedResults
}

// NewAPIKey initializes an API key with a purpose.
func NewAPIKey(purpose string, userID string) (*APIKey, error) {
	return &APIKey{
		Purpose: purpose,
		UserID:  userID,
	}, nil
}

type APIKey struct {
	APIKey  string     `json:"ApiKey,omitempty"`
	Created *time.Time `json:"Created,omitempty"`
	Purpose string     `json:"Purpose,omitempty"`
	UserID  string     `json:"UserId,omitempty"`

	Resource
}

// GetID returns the ID value of the APIKey.
func (resource APIKey) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this APIKey.
func (resource APIKey) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this APIKey was changed.
func (resource APIKey) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this APIKey.
func (resource APIKey) GetLinks() map[string]string {
	return resource.Links
}

// SetID
func (resource APIKey) SetID(id string) {
	resource.ID = id
}

// SetLastModifiedBy
func (resource APIKey) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

// SetLastModifiedOn
func (resource APIKey) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the APIKey and returns an error if invalid.
func (resource APIKey) Validate() error {
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

var _ ResourceInterface = &APIKey{}
