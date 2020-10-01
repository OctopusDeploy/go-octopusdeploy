package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// ActionTemplates defines a collection of ActionTemplate items with built-in support for paged results.
type ActionTemplates struct {
	Items []ActionTemplate `json:"Items"`
	PagedResults
}

// ActionTemplate represents an action template in Octopus.
type ActionTemplate struct {
	ActionType                string                     `json:"ActionType" validate:"required"`
	CommunityActionTemplateID string                     `json:"CommunityActionTemplateId,omitempty"`
	Description               string                     `json:"Description,omitempty"`
	Name                      string                     `json:"Name" validate:"required"`
	Packages                  []*PackageReference        `json:"Packages"`
	Parameters                []*ActionTemplateParameter `json:"Parameters"`
	Properties                map[string]PropertyValue   `json:"Properties,omitempty"`
	SpaceID                   string                     `json:"SpaceId,omitempty"`
	Version                   int32                      `json:"Version,omitempty"`

	Resource
}

// NewActionTemplate creates and initializes an action template.
func NewActionTemplate(name string, actionType string) (*ActionTemplate, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewActionTemplate", "name")
	}

	return &ActionTemplate{
		Name:       name,
		ActionType: actionType,
	}, nil
}

// GetID returns the ID value of the ActionTemplate.
func (resource ActionTemplate) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this ActionTemplate.
func (resource ActionTemplate) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this ActionTemplate was changed.
func (resource ActionTemplate) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this ActionTemplate.
func (resource ActionTemplate) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the ActionTemplate and returns an error if invalid.
func (resource ActionTemplate) Validate() error {
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

var _ ResourceInterface = &ActionTemplate{}
