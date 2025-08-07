package observability

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// ResourceEventsSessionId represents a resource events session identifier
type ResourceEventsSessionId string

// BeginResourceEventsSessionRequest represents a request to begin a resource events session
// Request to start monitoring events for a specific Kubernetes resource
type BeginResourceEventsSessionRequest struct {
	SpaceID                            string `json:"spaceId" validate:"required"`
	ProjectID                          string `json:"projectId" validate:"required"`
	EnvironmentID                      string `json:"environmentId" validate:"required"`
	TenantID                          *string `json:"tenantId,omitempty"`
	MachineID                         string `json:"machineId" validate:"required"`
	DesiredOrKubernetesMonitoredResourceID string `json:"desiredOrKubernetesMonitoredResourceId" validate:"required"`
}

// BeginResourceEventsSessionResponse represents the response for beginning a resource events session
// Response containing a session ID for the event monitoring session
type BeginResourceEventsSessionResponse struct {
	SessionID ResourceEventsSessionId `json:"sessionId" validate:"required"`
}

// GetResourceEventsRequest represents a request to get resource events for a session
// Request for retrieving all the events for the specified session
type GetResourceEventsRequest struct {
	SpaceID   string                  `json:"spaceId" validate:"required"`
	SessionID ResourceEventsSessionId `json:"sessionId" validate:"required" uri:"sessionId" url:"sessionId"`
}

// GetResourceEventsResponse represents the response containing events for a sessionID
// Response containing the events for a sessionID
type GetResourceEventsResponse struct {
	Events             []KubernetesEventResource `json:"events" validate:"required"`
	IsSessionCompleted bool                      `json:"isSessionCompleted"`
	Error              *MonitorErrorResource     `json:"error,omitempty"`
}

// KubernetesEventResource represents a Kubernetes event resource
type KubernetesEventResource struct {
	FirstObservedTime   time.Time `json:"firstObservedTime" validate:"required"`
	LastObservedTime    time.Time `json:"lastObservedTime" validate:"required"`
	Count               int       `json:"count" validate:"required"`
	Action              string    `json:"action" validate:"required"`
	Reason              string    `json:"reason" validate:"required"`
	Note                string    `json:"note" validate:"required"`
	ReportingController string    `json:"reportingController" validate:"required"`
	ReportingInstance   string    `json:"reportingInstance" validate:"required"`
	Type                string    `json:"type" validate:"required"`
	Manifest            string    `json:"manifest" validate:"required"`
}

// MonitorErrorResource represents an error resource for monitoring operations
type MonitorErrorResource struct {
	Message string `json:"message" validate:"required"`
	Code    string `json:"code,omitempty"`
}

// NewBeginResourceEventsSessionRequest creates a new BeginResourceEventsSessionRequest
func NewBeginResourceEventsSessionRequest(spaceID, projectID, environmentID, machineID, desiredOrKubernetesMonitoredResourceID string) *BeginResourceEventsSessionRequest {
	return &BeginResourceEventsSessionRequest{
		SpaceID:                            spaceID,
		ProjectID:                          projectID,
		EnvironmentID:                      environmentID,
		MachineID:                         machineID,
		DesiredOrKubernetesMonitoredResourceID: desiredOrKubernetesMonitoredResourceID,
	}
}

// NewBeginResourceEventsSessionResponse creates a new BeginResourceEventsSessionResponse
func NewBeginResourceEventsSessionResponse(sessionID ResourceEventsSessionId) *BeginResourceEventsSessionResponse {
	return &BeginResourceEventsSessionResponse{
		SessionID: sessionID,
	}
}

// NewGetResourceEventsRequest creates a new GetResourceEventsRequest
func NewGetResourceEventsRequest(spaceID string, sessionID ResourceEventsSessionId) *GetResourceEventsRequest {
	return &GetResourceEventsRequest{
		SpaceID:   spaceID,
		SessionID: sessionID,
	}
}

// NewGetResourceEventsResponse creates a new GetResourceEventsResponse
func NewGetResourceEventsResponse(events []KubernetesEventResource, isSessionCompleted bool) *GetResourceEventsResponse {
	return &GetResourceEventsResponse{
		Events:             events,
		IsSessionCompleted: isSessionCompleted,
	}
}

// NewKubernetesEventResource creates a new KubernetesEventResource
func NewKubernetesEventResource(
	firstObservedTime time.Time,
	lastObservedTime time.Time,
	count int,
	action string,
	reason string,
	note string,
	reportingController string,
	reportingInstance string,
	eventType string,
	manifest string,
) *KubernetesEventResource {
	return &KubernetesEventResource{
		FirstObservedTime:   firstObservedTime,
		LastObservedTime:    lastObservedTime,
		Count:               count,
		Action:              action,
		Reason:              reason,
		Note:                note,
		ReportingController: reportingController,
		ReportingInstance:   reportingInstance,
		Type:                eventType,
		Manifest:            manifest,
	}
}

// NewMonitorErrorResource creates a new MonitorErrorResource
func NewMonitorErrorResource(message string, code string) *MonitorErrorResource {
	return &MonitorErrorResource{
		Message: message,
		Code:    code,
	}
}

// Validate checks the state of the request and returns an error if invalid
func (r *BeginResourceEventsSessionRequest) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the response and returns an error if invalid
func (r *BeginResourceEventsSessionResponse) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the request and returns an error if invalid
func (r *GetResourceEventsRequest) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the response and returns an error if invalid
func (r *GetResourceEventsResponse) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the event resource and returns an error if invalid
func (k *KubernetesEventResource) Validate() error {
	return validator.New().Struct(k)
}

// Validate checks the state of the error resource and returns an error if invalid
func (e *MonitorErrorResource) Validate() error {
	return validator.New().Struct(e)
}
