package kubernetesmonitors

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// KubernetesMonitor represents an installation of the Kubernetes Monitor.
type KubernetesMonitor struct {
	ID             string     `json:"Id" validate:"required"`
	SpaceID        string     `json:"SpaceId" validate:"required"`
	InstallationID *uuid.UUID `json:"InstallationId" validate:"required"`
	MachineID      string     `json:"MachineId" validate:"required"`
}

// Validate checks the state of the Kubernetes monitor and returns an error if invalid.
func (k *KubernetesMonitor) Validate() error {
	validate := validator.New()
	return validate.Struct(k)
}
