package users

import (
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
)

// CreateAPIKey represents an API key when creating a new API key (POST)
type CreateAPIKey struct {
	APIKey  string     `json:"ApiKey,omitempty"`
	Created *time.Time `json:"Created,omitempty"`
	Purpose string     `json:"Purpose,omitempty"`
	UserID  string     `json:"UserId,omitempty"`
	Expires *time.Time `json:"Expires,omitempty"`

	resources.Resource
}

// APIKey represents an API key returned from a GET request
type APIKey struct {
	APIKey  *APIKeyStruct `json:"ApiKey,omitempty"`
	Created *time.Time    `json:"Created,omitempty"`
	Purpose string        `json:"Purpose,omitempty"`
	UserID  string        `json:"UserId,omitempty"`
	Expires *time.Time    `json:"Expires,omitempty"`

	resources.Resource
}

type APIKeyStruct struct {
	HasValue bool   `json:"HasValue"`
	NewValue string `json:"NewValue,omitempty"`
	Hint     string `json:"Hint"`
}

// NewAPIKey initializes an API key with a purpose.
func NewAPIKey(purpose string, userID string) *CreateAPIKey {
	return &CreateAPIKey{
		Purpose:  purpose,
		UserID:   userID,
		Resource: *resources.NewResource(),
	}
}

var _ resources.IResource = &APIKey{}
