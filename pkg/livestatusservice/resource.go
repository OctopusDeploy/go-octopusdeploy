package livestatusservice

import (
	"github.com/go-playground/validator/v10"
)

// GetResourceRequest represents a request to get a detailed summary of a live Kubernetes resource
// Request for retrieving detailed information about either a top-level resource or a child resource
type GetResourceRequest struct {
	SpaceID                                string `json:"spaceId" validate:"required"`
	ProjectID                              string `json:"projectId" validate:"required" uri:"projectId" url:"projectId"`
	EnvironmentID                          string `json:"environmentId" validate:"required" uri:"environmentId" url:"environmentId"`
	MachineID                              string `json:"machineId" validate:"required" uri:"machineId" url:"machineId"`
	DesiredOrKubernetesMonitoredResourceID string `json:"desiredOrKubernetesMonitoredResourceId" validate:"required" uri:"desiredOrKubernetesMonitoredResourceId" url:"desiredOrKubernetesMonitoredResourceId"`
	TenantID                               string `json:"tenantId,omitempty" uri:"tenantId,omitempty" url:"tenantId,omitempty"`
}

// GetResourceResponse represents the response containing detailed summary of a live kubernetes resource
// Contains detailed information about either a top-level resource or a child resource  
type GetResourceResponse struct {
	Resource *KubernetesLiveStatusDetailedResource `json:"resource" validate:"required"`
}

// KubernetesLiveStatusDetailedResource represents detailed information about a live kubernetes resource
type KubernetesLiveStatusDetailedResource struct {
	Name               string                                   `json:"Name,omitempty"`
	Namespace          string                                   `json:"Namespace,omitempty"`
	Kind               string                                   `json:"Kind,omitempty"`
	HealthStatus       string                                   `json:"HealthStatus,omitempty"`
	SyncStatus         *string                                  `json:"SyncStatus,omitempty"`
	MachineID          string                                   `json:"MachineId,omitempty"`
	LastUpdated        string                                   `json:"LastUpdated,omitempty"`
	Details            *KubernetesResourceDetails               `json:"Details,omitempty"`
	Children           []KubernetesLiveStatusDetailedResource   `json:"Children,omitempty"`
	DesiredResourceID  *string                                  `json:"DesiredResourceId,omitempty"`
	ResourceID         string                                   `json:"ResourceId,omitempty"`
}

// KubernetesResourceDetails represents the detailed information about a kubernetes resource
type KubernetesResourceDetails struct {
	Labels            map[string]string      `json:"Labels,omitempty"`
	Annotations       map[string]string      `json:"Annotations,omitempty"`
	CreationTimestamp string                 `json:"CreationTimestamp,omitempty"`
	Spec              map[string]interface{} `json:"Spec,omitempty"`
	Status            map[string]interface{} `json:"Status,omitempty"`
	OwnerReferences   []interface{}          `json:"OwnerReferences,omitempty"`
	Events            []interface{}          `json:"Events,omitempty"`
	Logs              []string               `json:"Logs,omitempty"`
}

// NewGetResourceRequest creates a new GetResourceRequest for untenanted resources
func NewGetResourceRequest(spaceID, projectID, environmentID, machineID, resourceID string) *GetResourceRequest {
	return &GetResourceRequest{
		SpaceID:                                spaceID,
		ProjectID:                              projectID,
		EnvironmentID:                          environmentID,
		MachineID:                              machineID,
		DesiredOrKubernetesMonitoredResourceID: resourceID,
	}
}

// NewGetResourceRequestWithTenant creates a new GetResourceRequest for tenanted resources
func NewGetResourceRequestWithTenant(spaceID, projectID, environmentID, tenantID, machineID, resourceID string) *GetResourceRequest {
	return &GetResourceRequest{
		SpaceID:                                spaceID,
		ProjectID:                              projectID,
		EnvironmentID:                          environmentID,
		TenantID:                               tenantID,
		MachineID:                              machineID,
		DesiredOrKubernetesMonitoredResourceID: resourceID,
	}
}

// IsTenanted returns true if the request is for a tenanted resource
func (r *GetResourceRequest) IsTenanted() bool {
	return r.TenantID != ""
}

// Validate checks the state of the request and returns an error if invalid
func (r *GetResourceRequest) Validate() error {
	return validator.New().Struct(r)
}

// NewGetResourceResponse creates a new GetResourceResponse with the provided resource
func NewGetResourceResponse(resource *KubernetesLiveStatusDetailedResource) *GetResourceResponse {
	return &GetResourceResponse{
		Resource: resource,
	}
}

// Validate checks the state of the response and returns an error if invalid
func (r *GetResourceResponse) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the detailed resource and returns an error if invalid
func (r *KubernetesLiveStatusDetailedResource) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the resource details and returns an error if invalid
func (r *KubernetesResourceDetails) Validate() error {
	return validator.New().Struct(r)
}