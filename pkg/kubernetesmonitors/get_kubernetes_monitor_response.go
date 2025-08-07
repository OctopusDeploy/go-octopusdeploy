package kubernetesmonitors

import "github.com/go-playground/validator/v10"

// GetKubernetesMonitorResponse represents the response to a request for a Kubernetes monitor.
type GetKubernetesMonitorResponse struct {
	Resource KubernetesMonitor `json:"Resource" validate:"required"`
}

// Validate checks the state of the Kubernetes Monitor response and returns an error if invalid.
func (k *GetKubernetesMonitorResponse) Validate() error {
	validate := validator.New()
	return validate.Struct(k)
}
