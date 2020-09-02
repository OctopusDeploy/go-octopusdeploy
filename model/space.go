package model

import (
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

func (t *Space) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewSpace(name string) *Space {
	return &Space{
		Name: name,
	}
}