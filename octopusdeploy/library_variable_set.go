package octopusdeploy

import (
	"github.com/go-playground/validator/v10"
)

type LibraryVariableSets struct {
	Items []*LibraryVariableSet `json:"Items"`
	PagedResults
}

type LibraryVariableSet struct {
	ContentType   string                    `json:"ContentType" validate:"required,oneof=ScriptModule Variables"`
	Description   string                    `json:"Description,omitempty"`
	Name          string                    `json:"Name" validate:"required"`
	SpaceID       string                    `json:"SpaceId,omitempty"`
	Templates     []ActionTemplateParameter `json:"Templates,omitempty"`
	VariableSetID string                    `json:"VariableSetId,omitempty"`

	resource
}

func NewLibraryVariableSet(name string) *LibraryVariableSet {
	return &LibraryVariableSet{
		ContentType: "Variables",
		Name:        name,
		Templates:   []ActionTemplateParameter{},
		resource:    *newResource(),
	}
}

// ValidateLibraryVariableSetValues checks the values of a library variable set
// to see if they are suitable for sending to Octopus Deploy. Used when adding
// or updating library variable sets.
func ValidateLibraryVariableSetValues(LibraryVariableSet *LibraryVariableSet) error {
	validate := validator.New()
	err := validate.Struct(LibraryVariableSet)
	return err
}

// Validate checks the state of the library variable set and returns an error
// if invalid.
func (l LibraryVariableSet) Validate() error {
	return validator.New().Struct(l)
}
