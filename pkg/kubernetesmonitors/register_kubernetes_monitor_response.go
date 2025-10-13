package kubernetesmonitors

import (
	"github.com/go-playground/validator/v10"
)

// RegisterKubernetesMonitorResponse represents the successful response to a completed registration command for a Kubernetes Monitor.
type RegisterKubernetesMonitorResponse struct {
	Resource              KubernetesMonitor `json:"Resource" validate:"required"`
	AuthenticationToken   *string           `json:"AuthenticationToken,omitempty"`
	CertificateThumbprint string            `json:"CertificateThumbprint" validate:"required"`
}

// NewRegisterKubernetesMonitorResponse creates a new Kubernetes monitor registration response with the specified parameters.
func NewRegisterKubernetesMonitorResponse(
	monitor KubernetesMonitor, authenticationToken *string, certificateThumbprint string,
) *RegisterKubernetesMonitorResponse {
	return &RegisterKubernetesMonitorResponse{
		Resource:              monitor,
		AuthenticationToken:   authenticationToken,
		CertificateThumbprint: certificateThumbprint,
	}
}

// Validate checks the state of the Kubernetes monitor registration command and returns an error if invalid.
func (k *RegisterKubernetesMonitorResponse) Validate() error {
	validate := validator.New()
	return validate.Struct(k)
}
