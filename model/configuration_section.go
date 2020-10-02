package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ConfigurationSections struct {
	Items []ConfigurationSection `json:"Items"`
	PagedResults
}

type ConfigurationSection struct {
	Description string `json:"Description,omitempty"`
	Name        string `json:"Name,omitempty"`

	Resource
}

// GetID returns the ID value of the ConfigurationSection.
func (resource ConfigurationSection) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this ConfigurationSection.
func (resource ConfigurationSection) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this ConfigurationSection was changed.
func (resource ConfigurationSection) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this ConfigurationSection.
func (resource ConfigurationSection) GetLinks() map[string]string {
	return resource.Links
}

// SetID
func (resource ConfigurationSection) SetID(id string) {
	resource.ID = id
}

// SetLastModifiedBy
func (resource ConfigurationSection) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

// SetLastModifiedOn
func (resource ConfigurationSection) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the ConfigurationSection and returns an error if invalid.
func (resource ConfigurationSection) Validate() error {
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

var _ ResourceInterface = &ConfigurationSection{}
