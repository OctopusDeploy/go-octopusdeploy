package observability

import (
	"encoding/json"
	"time"

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
	Name              string                                 `json:"Name,omitempty"`
	Namespace         string                                 `json:"Namespace,omitempty"`
	Kind              string                                 `json:"Kind,omitempty"`
	HealthStatus      string                                 `json:"HealthStatus,omitempty"`
	SyncStatus        *string                                `json:"SyncStatus,omitempty"`
	MachineID         string                                 `json:"MachineId,omitempty"`
	LastUpdated       string                                 `json:"LastUpdated,omitempty"`
	ManifestSummary   ManifestSummaryResource                `json:"ManifestSummary,omitempty"`
	Children          []KubernetesLiveStatusDetailedResource `json:"Children,omitempty"`
	DesiredResourceID *string                                `json:"DesiredResourceId,omitempty"`
	ResourceID        string                                 `json:"ResourceId,omitempty"`
}

// ManifestSummaryResource defines the interface for manifest summary resources
type ManifestSummaryResource interface {
	GetLabels() map[string]string
	GetAnnotations() map[string]string
	GetCreationTimestamp() time.Time
	Validate() error
}

// ManifestSummary represents the detailed information about a kubernetes resource
type ManifestSummary struct {
	Labels            map[string]string `json:"Labels,omitempty"`
	Annotations       map[string]string `json:"Annotations,omitempty"`
	CreationTimestamp time.Time         `json:"CreationTimestamp,omitempty"`
}

// GetLabels returns the labels map
func (m *ManifestSummary) GetLabels() map[string]string {
	return m.Labels
}

// GetAnnotations returns the annotations map
func (m *ManifestSummary) GetAnnotations() map[string]string {
	return m.Annotations
}

// GetCreationTimestamp returns the creation timestamp
func (m *ManifestSummary) GetCreationTimestamp() time.Time {
	return m.CreationTimestamp
}

// PodManifestSummary represents the detailed information about a kubernetes pod resource
type PodManifestSummary struct {
	ManifestSummary
	Containers []string `json:"Containers,omitempty"`
}

// GetContainers returns the list of container names
func (p *PodManifestSummary) GetContainers() []string {
	return p.Containers
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
func (r *ManifestSummary) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the pod manifest summary and returns an error if invalid
func (p *PodManifestSummary) Validate() error {
	return validator.New().Struct(p)
}

// UnmarshalJSON custom unmarshaling for KubernetesLiveStatusDetailedResource
func (r *KubernetesLiveStatusDetailedResource) UnmarshalJSON(data []byte) error {
	// Create a temporary struct with the same fields but using raw JSON for ManifestSummary
	type TempResource struct {
		Name              string                                 `json:"Name,omitempty"`
		Namespace         string                                 `json:"Namespace,omitempty"`
		Kind              string                                 `json:"Kind,omitempty"`
		HealthStatus      string                                 `json:"HealthStatus,omitempty"`
		SyncStatus        *string                                `json:"SyncStatus,omitempty"`
		MachineID         string                                 `json:"MachineId,omitempty"`
		LastUpdated       string                                 `json:"LastUpdated,omitempty"`
		ManifestSummary   json.RawMessage                        `json:"ManifestSummary,omitempty"`
		Children          []KubernetesLiveStatusDetailedResource `json:"Children,omitempty"`
		DesiredResourceID *string                                `json:"DesiredResourceId,omitempty"`
		ResourceID        string                                 `json:"ResourceId,omitempty"`
	}

	var temp TempResource
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Copy all fields except ManifestSummary
	r.Name = temp.Name
	r.Namespace = temp.Namespace
	r.Kind = temp.Kind
	r.HealthStatus = temp.HealthStatus
	r.SyncStatus = temp.SyncStatus
	r.MachineID = temp.MachineID
	r.LastUpdated = temp.LastUpdated
	r.Children = temp.Children
	r.DesiredResourceID = temp.DesiredResourceID
	r.ResourceID = temp.ResourceID

	// Handle ManifestSummary based on resource Kind
	if len(temp.ManifestSummary) > 0 {
		if r.Kind == "Pod" {
			var podSummary PodManifestSummary
			if err := json.Unmarshal(temp.ManifestSummary, &podSummary); err != nil {
				return err
			}
			r.ManifestSummary = &podSummary
		} else {
			var manifestSummary ManifestSummary
			if err := json.Unmarshal(temp.ManifestSummary, &manifestSummary); err != nil {
				return err
			}
			r.ManifestSummary = &manifestSummary
		}
	}

	return nil
}
