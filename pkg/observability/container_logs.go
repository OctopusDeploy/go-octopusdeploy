package observability

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// ContainerLogSessionId represents a container logs session identifier
type ContainerLogSessionId string

// GetContainerLogsRequest represents a request to get container logs for a session
// Request for retrieving all the logs for the specified session
type GetContainerLogsRequest struct {
	SpaceID   string                `json:"spaceId" validate:"required"`
	SessionID ContainerLogSessionId `json:"sessionId" validate:"required" uri:"sessionId" url:"sessionId"`
}

// GetContainerLogsResponse represents the response containing logs for a sessionID
// Response containing the logs for a sessionID
type GetContainerLogsResponse struct {
	Logs               []ContainerLogLineResource `json:"logs" validate:"required"`
	IsSessionCompleted bool                       `json:"isSessionCompleted"`
	Error              *MonitorErrorResource      `json:"error,omitempty"`
}

// ContainerLogLineResource represents a single container log line
type ContainerLogLineResource struct {
	Timestamp time.Time `json:"timestamp" validate:"required"`
	Message   string    `json:"message" validate:"required"`
}

// NewGetContainerLogsRequest creates a new GetContainerLogsRequest
func NewGetContainerLogsRequest(spaceID string, sessionID ContainerLogSessionId) *GetContainerLogsRequest {
	return &GetContainerLogsRequest{
		SpaceID:   spaceID,
		SessionID: sessionID,
	}
}

// NewGetContainerLogsResponse creates a new GetContainerLogsResponse
func NewGetContainerLogsResponse(logs []ContainerLogLineResource, isSessionCompleted bool) *GetContainerLogsResponse {
	return &GetContainerLogsResponse{
		Logs:               logs,
		IsSessionCompleted: isSessionCompleted,
	}
}

// NewContainerLogLineResource creates a new ContainerLogLineResource
func NewContainerLogLineResource(timestamp time.Time, message string) *ContainerLogLineResource {
	return &ContainerLogLineResource{
		Timestamp: timestamp,
		Message:   message,
	}
}

// Validate checks the state of the request and returns an error if invalid
func (r *GetContainerLogsRequest) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the response and returns an error if invalid
func (r *GetContainerLogsResponse) Validate() error {
	return validator.New().Struct(r)
}

// Validate checks the state of the log line resource and returns an error if invalid
func (c *ContainerLogLineResource) Validate() error {
	return validator.New().Struct(c)
}