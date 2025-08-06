package observability

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetContainerLogsRequest_Validate(t *testing.T) {
	// Test valid request
	validRequest := &GetContainerLogsRequest{
		SpaceID:   "Spaces-1",
		SessionID: ContainerLogSessionId("session-123"),
	}

	err := validRequest.Validate()
	assert.NoError(t, err)

	// Test invalid request (missing required fields)
	invalidRequest := &GetContainerLogsRequest{}

	err = invalidRequest.Validate()
	assert.Error(t, err)

	// Test invalid request (missing SessionID)
	invalidRequestNoSession := &GetContainerLogsRequest{
		SpaceID: "Spaces-1",
	}

	err = invalidRequestNoSession.Validate()
	assert.Error(t, err)

	// Test invalid request (missing SpaceID)
	invalidRequestNoSpace := &GetContainerLogsRequest{
		SessionID: ContainerLogSessionId("session-123"),
	}

	err = invalidRequestNoSpace.Validate()
	assert.Error(t, err)
}

func TestGetContainerLogsResponse_Validate(t *testing.T) {
	// Test valid response with empty logs
	validResponse := &GetContainerLogsResponse{
		Logs:               []ContainerLogLineResource{},
		IsSessionCompleted: true,
	}

	err := validResponse.Validate()
	assert.NoError(t, err)

	// Test valid response with logs
	logLine := &ContainerLogLineResource{
		Timestamp: time.Now(),
		Message:   "Application started successfully",
	}

	validResponseWithLogs := &GetContainerLogsResponse{
		Logs:               []ContainerLogLineResource{*logLine},
		IsSessionCompleted: false,
	}

	err = validResponseWithLogs.Validate()
	assert.NoError(t, err)
}

func TestContainerLogLineResource_Validate(t *testing.T) {
	// Test valid log line resource
	validLogLine := &ContainerLogLineResource{
		Timestamp: time.Now(),
		Message:   "Application started successfully",
	}

	err := validLogLine.Validate()
	assert.NoError(t, err)

	// Test invalid log line resource (missing required fields)
	invalidLogLine := &ContainerLogLineResource{}

	err = invalidLogLine.Validate()
	assert.Error(t, err)

	// Test invalid log line resource (missing message)
	invalidLogLineNoMessage := &ContainerLogLineResource{
		Timestamp: time.Now(),
	}

	err = invalidLogLineNoMessage.Validate()
	assert.Error(t, err)

	// Test invalid log line resource (missing timestamp)
	invalidLogLineNoTimestamp := &ContainerLogLineResource{
		Message: "Application started successfully",
	}

	err = invalidLogLineNoTimestamp.Validate()
	assert.Error(t, err)
}

func TestGetContainerLogsResponseWithError(t *testing.T) {
	logs := []ContainerLogLineResource{}
	isSessionCompleted := false
	errorResource := NewMonitorErrorResource("Session timeout", "ERR_TIMEOUT")

	response := &GetContainerLogsResponse{
		Logs:               logs,
		IsSessionCompleted: isSessionCompleted,
		Error:              errorResource,
	}

	expected := &GetContainerLogsResponse{
		Logs:               logs,
		IsSessionCompleted: isSessionCompleted,
		Error: &MonitorErrorResource{
			Message: "Session timeout",
			Code:    "ERR_TIMEOUT",
		},
	}

	assert.Equal(t, expected, response)
	assert.NoError(t, response.Validate())
}
