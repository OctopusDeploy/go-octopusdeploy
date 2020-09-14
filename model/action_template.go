package model

import (
	"github.com/go-playground/validator/v10"
)

// ActionTemplates defines a collection of action templates with built-in
// support for paged results.
type ActionTemplates struct {
	Items []ActionTemplate `json:"Items"`
	PagedResults
}

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

// NewActionTemplate initializes an account with a name and action type.
func NewActionTemplate(name string, actionType string) (*ActionTemplate, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("NewActionTemplate", "name")
	}

	return &ActionTemplate{
		Name:       name,
		ActionType: actionType,
	}, nil
}

func (a *ActionTemplate) GetID() string {
	return a.ID
}

func (a *ActionTemplate) Validate() error {
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

var _ ResourceInterface = &ActionTemplate{}
