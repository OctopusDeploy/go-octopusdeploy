package observability

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// GetLiveStatusRequest represents a request to get live statuses for a Project/Environment/Tenant
// Request the live statuses for a Project/Environment/Tenant
type GetLiveStatusRequest struct {
	SpaceID       string  `json:"spaceId" validate:"required"`
	ProjectID     string  `json:"projectId" validate:"required" uri:"projectId" url:"projectId"`
	EnvironmentID string  `json:"environmentId" validate:"required" uri:"environmentId" url:"environmentId"`
	TenantID      *string `json:"tenantId,omitempty" uri:"tenantId" url:"tenantId"`
	SummaryOnly   bool    `json:"summaryOnly"`
}

// GetLiveStatusResponse represents the live statuses for a given Project/Environment/Tenant
// Live statuses for a given Project/Environment/Tenant
type GetLiveStatusResponse struct {
	MachineStatuses []KubernetesMachineLiveStatusResource `json:"machineStatuses" validate:"required"`
	Summary         LiveStatusSummaryResource             `json:"summary" validate:"required"`
	Error           *MonitorErrorResource                 `json:"error,omitempty"`
}

// LiveStatusSummaryResource represents a summary of the live status
type LiveStatusSummaryResource struct {
	Status      string    `json:"status" validate:"required"`
	LastUpdated time.Time `json:"lastUpdated" validate:"required"`
}

// KubernetesMachineLiveStatusResource represents the live status for a specific machine
type KubernetesMachineLiveStatusResource struct {
	MachineID string                         `json:"machineId" validate:"required"`
	Status    string                         `json:"status" validate:"required"`
	Resources []KubernetesLiveStatusResource `json:"resources" validate:"required"`
}

// KubernetesLiveStatusResource represents a Kubernetes resource live status
type KubernetesLiveStatusResource struct {
	Name                string                         `json:"name" validate:"required"`
	Namespace           *string                        `json:"namespace,omitempty"`
	Kind                string                         `json:"kind" validate:"required"`
	HealthStatus        string                         `json:"healthStatus" validate:"required"`
	SyncStatus          *string                        `json:"syncStatus,omitempty"`
	MachineID           string                         `json:"machineId" validate:"required"`
	Children            []KubernetesLiveStatusResource `json:"children" validate:"required"`
	DesiredResourceID   *string                        `json:"desiredResourceId,omitempty"`
	MonitoredResourceID *string                        `json:"monitoredResourceId,omitempty"`
}

// NewGetLiveStatusRequest creates a new GetLiveStatusRequest
func NewGetLiveStatusRequest(spaceID, projectID, environmentID string) *GetLiveStatusRequest {
	return &GetLiveStatusRequest{
		SpaceID:       spaceID,
		ProjectID:     projectID,
		EnvironmentID: environmentID,
		SummaryOnly:   false,
	}
}

// NewGetLiveStatusRequestWithTenant creates a new GetLiveStatusRequest with tenant
func NewGetLiveStatusRequestWithTenant(spaceID, projectID, environmentID, tenantID string) *GetLiveStatusRequest {
	return &GetLiveStatusRequest{
		SpaceID:       spaceID,
		ProjectID:     projectID,
		EnvironmentID: environmentID,
		TenantID:      &tenantID,
		SummaryOnly:   false,
	}
}

// NewGetLiveStatusResponse creates a new GetLiveStatusResponse
func NewGetLiveStatusResponse(machineStatuses []KubernetesMachineLiveStatusResource, summary LiveStatusSummaryResource) *GetLiveStatusResponse {
	return &GetLiveStatusResponse{
		MachineStatuses: machineStatuses,
		Summary:         summary,
	}
}

// NewLiveStatusSummaryResource creates a new LiveStatusSummaryResource
func NewLiveStatusSummaryResource(status string, lastUpdated time.Time) *LiveStatusSummaryResource {
	return &LiveStatusSummaryResource{
		Status:      status,
		LastUpdated: lastUpdated,
	}
}

// NewKubernetesMachineLiveStatusResource creates a new KubernetesMachineLiveStatusResource
func NewKubernetesMachineLiveStatusResource(machineID, status string, resources []KubernetesLiveStatusResource) *KubernetesMachineLiveStatusResource {
	return &KubernetesMachineLiveStatusResource{
		MachineID: machineID,
		Status:    status,
		Resources: resources,
	}
}

// NewKubernetesLiveStatusResource creates a new KubernetesLiveStatusResource
func NewKubernetesLiveStatusResource(
	name string,
	kind string,
	healthStatus string,
	machineID string,
	children []KubernetesLiveStatusResource,
) *KubernetesLiveStatusResource {
	return &KubernetesLiveStatusResource{
		Name:         name,
		Kind:         kind,
		HealthStatus: healthStatus,
		MachineID:    machineID,
		Children:     children,
	}
}

// Validate checks the state of the request and returns an error if invalid
func (r *GetLiveStatusRequest) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the response and returns an error if invalid
func (r *GetLiveStatusResponse) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the summary resource and returns an error if invalid
func (s *LiveStatusSummaryResource) Validate() error {
	return validator.New().Struct(s)
}

// Validate checks the state of the machine status resource and returns an error if invalid
func (k *KubernetesMachineLiveStatusResource) Validate() error {
	return validator.New().Struct(k)
}

// Validate checks the state of the live status resource and returns an error if invalid
func (k *KubernetesLiveStatusResource) Validate() error {
	return validator.New().Struct(k)
}
