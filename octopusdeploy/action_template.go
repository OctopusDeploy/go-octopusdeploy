package octopusdeploy

import "github.com/go-playground/validator/v10"

// ActionTemplates defines a collection of action templates with built-in
// support for paged results.
type ActionTemplates struct {
	Items []*ActionTemplate `json:"Items"`
	PagedResults
}

// ActionTemplate represents an action template in Octopus.
type ActionTemplate struct {
	ActionType                string                     `json:"ActionType" validate:"required"`
	CommunityActionTemplateID string                     `json:"CommunityActionTemplateId,omitempty"`
	Description               string                     `json:"Description,omitempty"`
	Name                      string                     `json:"Name" validate:"required"`
	Packages                  []*PackageReference        `json:"Packages,omitempty"`
	Parameters                []*ActionTemplateParameter `json:"Parameters,omitempty"`
	Properties                map[string]PropertyValue   `json:"Properties,omitempty"`
	SpaceID                   string                     `json:"SpaceId,omitempty"`
	Version                   int32                      `json:"Version,omitempty"`

	resource
}

// NewActionTemplate creates and initializes an action template.
func NewActionTemplate(name string, actionType string) *ActionTemplate {
	return &ActionTemplate{
		ActionType: actionType,
		Name:       name,
		resource:   *newResource(),
	}
}

// Validate checks the state of the ActionTemplate and returns an error if
// invalid.
func (a *ActionTemplate) Validate() error {
	return validator.New().Struct(a)
}
