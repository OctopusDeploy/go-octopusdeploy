package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Spaces struct {
	Items []Space `json:"Items"`
	PagedResults
}

type Space struct {
	Description              string   `json:"Description,omitempty"`
	IsDefault                bool     `json:"IsDefault,omitempty"`
	Name                     string   `json:"Name"`
	SpaceManagersTeamMembers []string `json:"SpaceManagersTeamMembers"`
	SpaceManagersTeams       []string `json:"SpaceManagersTeams"`
	TaskQueueStopped         bool     `json:"TaskQueueStopped,omitempty"`

	Resource
}

// NewSpace initializes a Space with a name.
func NewSpace(name string) *Space {
	return &Space{
		Name: name,
	}
}

// GetID returns the ID value of the Space.
func (resource Space) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this Space.
func (resource Space) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Space was changed.
func (resource Space) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Space.
func (resource Space) GetLinks() map[string]string {
	return resource.Links
}

func (resource Space) SetID(id string) {
	resource.ID = id
}

func (resource Space) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource Space) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the Space and returns an error if invalid.
func (resource Space) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &Space{}
