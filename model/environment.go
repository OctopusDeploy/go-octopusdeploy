package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Environments struct {
	Items []Environment `json:"Items"`
	PagedResults
}

type Environment struct {
	Name                       string `json:"Name"`
	Description                string `json:"Description"`
	SortOrder                  int    `json:"SortOrder"`
	UseGuidedFailure           bool   `json:"UseGuidedFailure"`
	AllowDynamicInfrastructure bool   `json:"AllowDynamicInfrastructure"`

	Resource
}

// GetID returns the ID value of the Environment.
func (resource Environment) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Environment.
func (resource Environment) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Environment was changed.
func (resource Environment) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Environment.
func (resource Environment) GetLinks() map[string]string {
	return resource.Links
}

func (resource Environment) SetID(id string) {
	resource.ID = id
}

func (resource Environment) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource Environment) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the Environment and returns an error if invalid.
func (resource Environment) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

func NewEnvironment(name, description string, useguidedfailure bool) *Environment {
	return &Environment{
		Name:             name,
		Description:      description,
		UseGuidedFailure: useguidedfailure,
	}
}

var _ ResourceInterface = &Environment{}
