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
func NewAPIKey(purpose string) (*APIKey, error) {
	return &APIKey{
		Purpose: &purpose,
	}, nil
}

type APIKey struct {
	APIKey  *string    `json:"ApiKey,omitempty"`
	Created *time.Time `json:"Created,omitempty"`
	Purpose *string    `json:"Purpose,omitempty"`
	UserID  *string    `json:"UserId,omitempty"`
	Resource
}

func (a *APIKey) GetID() string {
	return a.ID
}

func (a *APIKey) Validate() error {
	validate := validator.New()
	err := validate.Struct(a)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	return nil
}

var _ ResourceInterface = &APIKey{}
