package octopusdeploy

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type MachineCleanupPolicy struct {
	DeleteMachinesBehavior        string        `json:"DeleteMachinesBehavior" validate:"required,oneof=DeleteUnavailableMachines DoNotDelete"`
	DeleteMachinesElapsedTimeSpan time.Duration `json:"DeleteMachinesElapsedTimeSpan,omitempty"`
}

func NewMachineCleanupPolicy() *MachineCleanupPolicy {
	return &MachineCleanupPolicy{
		DeleteMachinesBehavior:        "DoNotDelete",
		DeleteMachinesElapsedTimeSpan: time.Hour,
	}
}

// MarshalJSON returns a machine policy as its JSON encoding.
func (m *MachineCleanupPolicy) MarshalJSON() ([]byte, error) {
	machineCleanupPolicy := struct {
		DeleteMachinesBehavior        string `json:"DeleteMachinesBehavior" validate:"required,oneof=DeleteUnavailableMachines DoNotDelete"`
		DeleteMachinesElapsedTimeSpan string `json:"DeleteMachinesElapsedTimeSpan,omitempty"`
	}{
		DeleteMachinesBehavior:        m.DeleteMachinesBehavior,
		DeleteMachinesElapsedTimeSpan: ToTimeSpan(m.DeleteMachinesElapsedTimeSpan),
	}

	return json.Marshal(machineCleanupPolicy)
}

// UnmarshalJSON sets this Kubernetes endpoint to its representation in JSON.
func (m *MachineCleanupPolicy) UnmarshalJSON(data []byte) error {
	var fields struct {
		DeleteMachinesBehavior        string `json:"DeleteMachinesBehavior" validate:"required,oneof=DeleteUnavailableMachines DoNotDelete"`
		DeleteMachinesElapsedTimeSpan string `json:"DeleteMachinesElapsedTimeSpan,omitempty"`
	}
	err := json.Unmarshal(data, &fields)
	if err != nil {
		return err
	}

	// validate JSON representation
	v := validator.New()
	err = v.Struct(fields)
	if err != nil {
		return err
	}

	if len(fields.DeleteMachinesElapsedTimeSpan) > 0 {
		m.DeleteMachinesElapsedTimeSpan = FromTimeSpan(fields.DeleteMachinesElapsedTimeSpan)
	}

	m.DeleteMachinesBehavior = fields.DeleteMachinesBehavior

	return nil
}

// Validate checks the state of the machine policy and returns an error if
// invalid.
func (m *MachineCleanupPolicy) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(m)
}
