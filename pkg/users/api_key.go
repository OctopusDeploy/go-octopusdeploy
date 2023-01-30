package users

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// APIKey represents an API key.
type APIKey struct {
	APIKey  string     `json:"ApiKey,omitempty"`
	Created *time.Time `json:"Created,omitempty"`
	Purpose string     `json:"Purpose,omitempty"`
	UserID  string     `json:"UserId,omitempty"`
	Expires *time.Time `json:"Expires,omitempty"`

	resources.Resource
}

// NewAPIKey initializes an API key with a purpose.
func NewAPIKey(purpose string, userID string) *APIKey {
	return &APIKey{
		Purpose:  purpose,
		UserID:   userID,
		Resource: *resources.NewResource(),
	}
}

var _ resources.IResource = &APIKey{}
