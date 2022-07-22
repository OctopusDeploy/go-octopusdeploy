package environments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type Environments struct {
	Items []*Environment `json:"Items"`
	resources.PagedResults
}

type Environment struct {
	AllowDynamicInfrastructure bool   `json:"AllowDynamicInfrastructure"`
	Description                string `json:"Description,omitempty"`
	Name                       string `json:"Name" validate:"required"`
	SortOrder                  int    `json:"SortOrder"`
	UseGuidedFailure           bool   `json:"UseGuidedFailure"`

	resources.Resource
}

func NewEnvironment(name string) *Environment {
	return &Environment{
		AllowDynamicInfrastructure: false,
		Name:                       name,
		SortOrder:                  0,
		UseGuidedFailure:           false,
		Resource:                   *resources.NewResource(),
	}
}

// Validate checks the state of the environment and returns an error if
// invalid.
func (e *Environment) Validate() error {
	return validator.New().Struct(e)
}

func (e *Environment) GetName() string {
	return e.Name
}
