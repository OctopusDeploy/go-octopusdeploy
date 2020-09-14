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

func (m *MachineConnectionStatus) GetID() string {
	return m.ID
}

// Validate returns a collection of validation errors against the machine
// connection status' internal values.
func (m *MachineConnectionStatus) Validate() error {
	validate := validator.New()
	err := validate.Struct(m)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}

		return err
	}

	return nil
}

var _ ResourceInterface = &MachineConnectionStatus{}
