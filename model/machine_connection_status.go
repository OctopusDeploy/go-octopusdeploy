package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type MachineConnectionStatus struct {
	CurrentTentacleVersion string                `json:"CurrentTentacleVersion,omitempty"`
	LastChecked            time.Time             `json:"LastChecked,omitempty"`
	Logs                   []*ActivityLogElement `json:"Logs"`
	MachineID              string                `json:"MachineId,omitempty"`
	Status                 string                `json:"Status,omitempty"`

	Resource
}

// GetID returns the ID value of the MachineConnectionStatus.
func (resource MachineConnectionStatus) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this MachineConnectionStatus.
func (resource MachineConnectionStatus) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this MachineConnectionStatus was changed.
func (resource MachineConnectionStatus) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this MachineConnectionStatus.
func (resource MachineConnectionStatus) GetLinks() map[string]string {
	return resource.Links
}

func (resource MachineConnectionStatus) SetID(id string) {
	resource.ID = id
}

func (resource MachineConnectionStatus) SetLastModifiedBy(name string) {
	resource.LastModifiedBy = name
}

func (resource MachineConnectionStatus) SetLastModifiedOn(time *time.Time) {
	resource.LastModifiedOn = time
}

// Validate checks the state of the MachineConnectionStatus and returns an error if invalid.
func (resource MachineConnectionStatus) Validate() error {
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

var _ ResourceInterface = &MachineConnectionStatus{}
