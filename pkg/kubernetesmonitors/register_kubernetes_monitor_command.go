package kubernetesmonitors

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// RegisterKubernetesMonitorCommand represents a command to register a Kubernetes monitor.
type RegisterKubernetesMonitorCommand struct {
	InstallationID *uuid.UUID `json:"InstallationId" validate:"required"`
	MachineID      string     `json:"MachineId" validate:"required"`
}

// NewRegisterKubernetesMonitorCommand creates a new Kubernetes monitor registration command with the specified parameters.
func NewRegisterKubernetesMonitorCommand(
	installationID *uuid.UUID, spaceID string, machineID string,
) *RegisterKubernetesMonitorCommand {
	return &RegisterKubernetesMonitorCommand{
		InstallationID: installationID,
		MachineID:      machineID,
	}
}

// Validate checks the state of the Kubernetes monitor registration command and returns an error if invalid.
func (k *RegisterKubernetesMonitorCommand) Validate() error {
	validate := validator.New()
	return validate.Struct(k)
}
