package observability

import (
	"github.com/go-playground/validator/v10"
)

// KubernetesMonitorInstallationId represents a Kubernetes Monitor installation identifier
type KubernetesMonitorInstallationId string

// KubernetesMonitorId represents a Kubernetes Monitor identifier
type KubernetesMonitorId string

// RegisterKubernetesMonitorCommand represents a request to register a new Kubernetes Monitor
// Registers and trusts new Kubernetes Monitor.
type RegisterKubernetesMonitorCommand struct {
	InstallationID KubernetesMonitorInstallationId `json:"installationId" validate:"required"`
	SpaceID        string                          `json:"spaceId" validate:"required"`
	MachineID      string                          `json:"machineId" validate:"required"`
}

// RegisterKubernetesMonitorResponse represents the response from registering a Kubernetes Monitor
// Response containing the registered agent.
type RegisterKubernetesMonitorResponse struct {
	Resource              *KubernetesMonitorResource `json:"resource" validate:"required"`
	AuthenticationToken   string                     `json:"authenticationToken" validate:"required"`
	CertificateThumbprint string                     `json:"certificateThumbprint" validate:"required"`
}

// KubernetesMonitorResource represents an installation of the Kubernetes Monitor
// Represents an installation of the Kubernetes Monitor.
type KubernetesMonitorResource struct {
	ID             KubernetesMonitorId             `json:"id" validate:"required"`
	SpaceID        string                          `json:"spaceId" validate:"required"`
	InstallationID KubernetesMonitorInstallationId `json:"installationId" validate:"required"`
	MachineID      string                          `json:"machineId" validate:"required"`
}

// NewRegisterKubernetesMonitorCommand creates a new RegisterKubernetesMonitorCommand
func NewRegisterKubernetesMonitorCommand(installationID KubernetesMonitorInstallationId, spaceID string, machineID string) *RegisterKubernetesMonitorCommand {
	return &RegisterKubernetesMonitorCommand{
		InstallationID: installationID,
		SpaceID:        spaceID,
		MachineID:      machineID,
	}
}

// NewRegisterKubernetesMonitorResponse creates a new RegisterKubernetesMonitorResponse
func NewRegisterKubernetesMonitorResponse(resource *KubernetesMonitorResource, authenticationToken string, certificateThumbprint string) *RegisterKubernetesMonitorResponse {
	return &RegisterKubernetesMonitorResponse{
		Resource:              resource,
		AuthenticationToken:   authenticationToken,
		CertificateThumbprint: certificateThumbprint,
	}
}

// NewKubernetesMonitorResource creates a new KubernetesMonitorResource
func NewKubernetesMonitorResource(id KubernetesMonitorId, spaceID string, installationID KubernetesMonitorInstallationId, machineID string) *KubernetesMonitorResource {
	return &KubernetesMonitorResource{
		ID:             id,
		SpaceID:        spaceID,
		InstallationID: installationID,
		MachineID:      machineID,
	}
}

// Validate checks the state of the command and returns an error if invalid
func (c *RegisterKubernetesMonitorCommand) Validate() error {
	return validator.New().Struct(c)
}

// Validate checks the state of the response and returns an error if invalid
func (r *RegisterKubernetesMonitorResponse) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the resource and returns an error if invalid
func (k *KubernetesMonitorResource) Validate() error {
	return validator.New().Struct(k)
}