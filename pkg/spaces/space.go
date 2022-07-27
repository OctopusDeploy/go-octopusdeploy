package spaces

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type Spaces struct {
	Items []*Space `json:"Items"`
	resources.PagedResults
}

type Space struct {
	Description              string   `json:"Description,omitempty"`
	IsDefault                bool     `json:"IsDefault,omitempty"`
	Name                     string   `json:"Name" validate:"required,max=20"`
	SpaceManagersTeamMembers []string `json:"SpaceManagersTeamMembers,omitempty"`
	SpaceManagersTeams       []string `json:"SpaceManagersTeams,omitempty"`
	TaskQueueStopped         bool     `json:"TaskQueueStopped,omitempty"`

	resources.Resource
}

// NewSpace initializes a Space with a name.
func NewSpace(name string) *Space {
	return &Space{
		Name:     name,
		Resource: *resources.NewResource(),
	}
}

// GetName returns the name of the space.
func (s *Space) GetName() string {
	return s.Name
}

// SetName sets the name of the space.
func (s *Space) SetName(name string) {
	s.Name = name
}

// Validate checks the state of the space and returns an error if invalid.
func (s *Space) Validate() error {
	return validator.New().Struct(s)
}

func (s *Space) GetName() string {
	return s.Name
}
