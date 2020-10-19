package model

import (
	"time"
)

// APIKeys defines a collection of API keys with built-in support for paged
// results.
type APIKeys struct {
	Items []*APIKey `json:"Items"`
	PagedResults
}

// APIKey represents an API key.
type APIKey struct {
	APIKey  string     `json:"ApiKey,omitempty"`
	Created *time.Time `json:"Created,omitempty"`
	Purpose string     `json:"Purpose,omitempty"`
	UserID  string     `json:"UserId,omitempty"`

	Resource
}

// NewAPIKey initializes an API key with a purpose.
func NewAPIKey(purpose string, userID string) *APIKey {
	return &APIKey{
		Purpose:  purpose,
		UserID:   userID,
		Resource: *newResource(),
	}
}

var _ IResource = &APIKey{}
