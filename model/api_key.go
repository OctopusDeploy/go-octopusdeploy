package model

import (
	"fmt"
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

func (apiKey *APIKey) Validate() error {
	validate := validator.New()
	err := validate.Struct(apiKey)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
		return err
	}

	return nil
}

var _ ResourceInterface = &APIKey{}
