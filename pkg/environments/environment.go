package environments

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type Environment struct {
	AllowDynamicInfrastructure bool   `json:"AllowDynamicInfrastructure"`
	Description                string `json:"Description,omitempty"`
	Name                       string `json:"Name" validate:"required"`
	SortOrder                  int    `json:"SortOrder"`
	UseGuidedFailure           bool   `json:"UseGuidedFailure"`
	SpaceID                    string `json:"SpaceId"`
	Slug                       string `json:"Slug,omitempty"` // will be an empty string on older servers

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

// GetName returns the name of the environment.
func (e *Environment) GetName() string {
	return e.Name
}

// SetName sets the name of the environment.
func (e *Environment) SetName(name string) {
	e.Name = name
}

var _ resources.IHasName = &Environment{}
