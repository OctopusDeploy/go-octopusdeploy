package resources

import (
	"github.com/go-playground/validator/v10"
)

// GetResourceManifestRequest represents a request to get a resource manifest
// Request for retrieving the manifest for a live kubernetes resource
type GetResourceManifestRequest struct {
	SpaceID                                string `json:"spaceId" validate:"required"`
	ProjectID                              string `json:"projectId" validate:"required" uri:"projectId" url:"projectId"`
	EnvironmentID                          string `json:"environmentId" validate:"required" uri:"environmentId" url:"environmentId"`
	MachineID                              string `json:"machineId" validate:"required" uri:"machineId" url:"machineId"`
	DesiredOrKubernetesMonitoredResourceID string `json:"desiredOrKubernetesMonitoredResourceId" validate:"required" uri:"desiredOrKubernetesMonitoredResourceId" url:"desiredOrKubernetesMonitoredResourceId"`
	TenantID                               string `json:"tenantId,omitempty" uri:"tenantId,omitempty" url:"tenantId,omitempty"`
}

// GetResourceManifestResponse represents the response from getting a resource manifest
// Contains the manifest for a live resource
type GetResourceManifestResponse struct {
	LiveManifest     string                `json:"liveManifest" validate:"required"`
	DesiredManifest  string                `json:"desiredManifest,omitempty"`
	Diff             *LiveResourceDiff     `json:"diff,omitempty"`
}

// LiveResourceDiff represents the diff between desired and live resource manifests
type LiveResourceDiff struct {
	Left  string `json:"left" validate:"required"`
	Right string `json:"right" validate:"required"`
	Diff  string `json:"diff" validate:"required"`
}

// NewGetResourceManifestRequest creates a new GetResourceManifestRequest for untenanted resources
func NewGetResourceManifestRequest(spaceID, projectID, environmentID, machineID, resourceID string) *GetResourceManifestRequest {
	return &GetResourceManifestRequest{
		SpaceID:                                spaceID,
		ProjectID:                              projectID,
		EnvironmentID:                          environmentID,
		MachineID:                              machineID,
		DesiredOrKubernetesMonitoredResourceID: resourceID,
	}
}

// NewGetResourceManifestRequestWithTenant creates a new GetResourceManifestRequest for tenanted resources
func NewGetResourceManifestRequestWithTenant(spaceID, projectID, environmentID, tenantID, machineID, resourceID string) *GetResourceManifestRequest {
	return &GetResourceManifestRequest{
		SpaceID:                                spaceID,
		ProjectID:                              projectID,
		EnvironmentID:                          environmentID,
		TenantID:                               tenantID,
		MachineID:                              machineID,
		DesiredOrKubernetesMonitoredResourceID: resourceID,
	}
}

// IsUntenanted returns true if the request is for an untenanted resource
func (r *GetResourceManifestRequest) IsUntenanted() bool {
	return r.TenantID == ""
}

// IsTenanted returns true if the request is for a tenanted resource
func (r *GetResourceManifestRequest) IsTenanted() bool {
	return r.TenantID != ""
}

// Validate checks the state of the request and returns an error if invalid
func (r *GetResourceManifestRequest) Validate() error {
	return validator.New().Struct(r)
}

// NewGetResourceManifestResponse creates a new GetResourceManifestResponse with required LiveManifest
func NewGetResourceManifestResponse(liveManifest string) *GetResourceManifestResponse {
	return &GetResourceManifestResponse{
		LiveManifest: liveManifest,
	}
}

// NewLiveResourceDiff creates a new LiveResourceDiff
func NewLiveResourceDiff(left, right, diff string) *LiveResourceDiff {
	return &LiveResourceDiff{
		Left:  left,
		Right: right,
		Diff:  diff,
	}
}

// Validate checks the state of the response and returns an error if invalid
func (r *GetResourceManifestResponse) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the diff and returns an error if invalid
func (d *LiveResourceDiff) Validate() error {
	return validator.New().Struct(d)
}