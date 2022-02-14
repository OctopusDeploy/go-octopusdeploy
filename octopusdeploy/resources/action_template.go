package resources

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

// ActionTemplate represents an action template in Octopus Deploy.
type ActionTemplate struct {
	ActionType                string                    `json:"ActionType" validate:"required,notblank"`
	CommunityActionTemplateID string                    `json:"CommunityActionTemplateId,omitempty"`
	Description               string                    `json:"Description,omitempty"`
	Name                      string                    `json:"Name" validate:"required"`
	Packages                  []PackageReference        `json:"Packages,omitempty"`
	Parameters                []ActionTemplateParameter `json:"Parameters,omitempty"`
	Properties                map[string]PropertyValue  `json:"Properties,omitempty"`
	SpaceID                   string                    `json:"SpaceId,omitempty"`
	Version                   int32                     `json:"Version,omitempty"`

	Resource
}

// NewActionTemplate creates and initializes an action template.
func NewActionTemplate(name string, actionType string) *ActionTemplate {
	return &ActionTemplate{
		ActionType: actionType,
		Name:       name,
		Packages:   []PackageReference{},
		Parameters: []ActionTemplateParameter{},
		Properties: map[string]PropertyValue{},
		Resource:   *NewResource(),
	}
}

// Validate checks the state of this ActionTemplate and returns an error if
// invalid.
func (a *ActionTemplate) Validate() error {
	v := validator.New()
	if err := v.RegisterValidation("notblank", validators.NotBlank); err != nil {
		return err
	}

	return v.Struct(a)
}