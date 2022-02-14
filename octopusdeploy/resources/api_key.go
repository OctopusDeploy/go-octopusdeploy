package resources

import (
	"time"
)

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
		Resource: *NewResource(),
	}
}

var _ IResource = &APIKey{}
