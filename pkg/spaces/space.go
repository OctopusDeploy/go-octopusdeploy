package spaces

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
)

type Space struct {
	Description              string   `json:"Description,omitempty"`
	Slug                     string   `json:"Slug"` // deliberately send empty string
	IsDefault                bool     `json:"IsDefault"`
	Name                     string   `json:"Name" validate:"required,max=20"`
	SpaceManagersTeamMembers []string `json:"SpaceManagersTeamMembers"` // deliberately send empty array
	SpaceManagersTeams       []string `json:"SpaceManagersTeams"`       // deliberately send empty array
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

var _ resources.IHasName = &Space{}
