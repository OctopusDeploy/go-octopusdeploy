package kubernetesmonitors

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// KubernetesMonitor represents an installation of the Kubernetes Monitor.
type KubernetesMonitor struct {
	ID             string     `json:"Id" validate:"required"`
	SpaceID        string     `json:"SpaceId" validate:"required"`
	InstallationID *uuid.UUID `json:"InstallationId" validate:"required"`
	MachineID      string     `json:"MachineId" validate:"required"`

	resources.Resource
}

// NewKubernetesMonitor creates a new Kubernetes Monitor with the specified parameters.
func NewKubernetesMonitor(spaceID string, installationID *uuid.UUID, machineID string) *KubernetesMonitor {
	return &KubernetesMonitor{
		SpaceID:        spaceID,
		InstallationID: installationID,
		MachineID:      machineID,
		Resource:       *resources.NewResource(),
	}
}

// Validate checks the state of the Kubernetes monitor and returns an error if invalid.
func (k *KubernetesMonitor) Validate() error {
	validate := validator.New()
	return validate.Struct(k)
}
