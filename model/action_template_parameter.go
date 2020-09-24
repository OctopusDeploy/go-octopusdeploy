package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ActionTemplateParameter struct {
	DefaultValue    *PropertyValueResource `json:"DefaultValue,omitempty"`
	DisplaySettings map[string]string      `json:"DisplaySettings,omitempty"`
	HelpText        string                 `json:"HelpText,omitempty"`
	Label           string                 `json:"Label,omitempty"`
	Name            string                 `json:"Name,omitempty"`

	Resource
}

// GetID returns the ID value of the ActionTemplateParameter struct instance.
func (resource ActionTemplateParameter) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this ActionTemplateParameter.
func (resource ActionTemplateParameter) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this ActionTemplateParameter was changed.
func (resource ActionTemplateParameter) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this ActionTemplateParameter.
func (resource ActionTemplateParameter) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the ActionTemplateParameter and returns an error if invalid.
func (resource ActionTemplateParameter) Validate() error {
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

var _ ResourceInterface = &ActionTemplateParameter{}
