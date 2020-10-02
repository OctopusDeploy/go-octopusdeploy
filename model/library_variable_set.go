package model

import (
	"time"

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

// GetID returns the ID value of the LibraryVariableSet.
func (resource LibraryVariableSet) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this LibraryVariableSet.
func (resource LibraryVariableSet) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this LibraryVariableSet was changed.
func (resource LibraryVariableSet) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this LibraryVariableSet.
func (resource LibraryVariableSet) GetLinks() map[string]string {
	return resource.Links
}

func (resource LibraryVariableSet) SetID(id string) {
	resource.ID = id
}

func (resource LibraryVariableSet) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource LibraryVariableSet) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the LibraryVariableSet and returns an error if invalid.
func (resource LibraryVariableSet) Validate() error {
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

var _ ResourceInterface = &LibraryVariableSet{}
