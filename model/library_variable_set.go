package model

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/go-playground/validator/v10"
)

type LibraryVariableSets struct {
	Items []LibraryVariableSet `json:"Items"`
	PagedResults
}

type LibraryVariableSet struct {
	ContentType   enum.VariableSetContentType `json:"ContentType"`
	Description   string                      `json:"Description,omitempty"`
	Name          string                      `json:"Name" validate:"required"`
	SpaceID       string                      `json:"SpaceId,omitempty"`
	Templates     []*ActionTemplateParameter  `json:"Templates,omitempty"`
	VariableSetID string                      `json:"VariableSetId,omitempty"`
	Resource
}

func NewLibraryVariableSet(name string) *LibraryVariableSet {
	return &LibraryVariableSet{
		Name:        name,
		ContentType: enum.Variables,
	}
}

// ValidateLibraryVariableSetValues checks the values of a LibraryVariableSet object to see if they are suitable for
// sending to Octopus Deploy. Used when adding or updating libraryVariableSets.
func ValidateLibraryVariableSetValues(LibraryVariableSet *LibraryVariableSet) error {
	validate := validator.New()
	err := validate.Struct(LibraryVariableSet)
	return err
}

func (l *LibraryVariableSet) GetID() string {
	return l.ID
}

// Validate returns a collection of validation errors against the library
// variable set's internal values.
func (l *LibraryVariableSet) Validate() error {
	validate := validator.New()
	err := validate.Struct(l)

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

var _ ResourceInterface = &LibraryVariableSet{}
