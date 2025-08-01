package livestatusservice

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetResourceEventsRequest_Validate(t *testing.T) {
	// Test valid request
	validRequest := &GetResourceEventsRequest{
		SpaceID:   "Spaces-1",
		SessionID: ResourceEventsSessionId("session-123"),
	}

	err := validRequest.Validate()
	assert.NoError(t, err)

	// Test invalid request (missing required fields)
	invalidRequest := &GetResourceEventsRequest{}

	err = invalidRequest.Validate()
	assert.Error(t, err)

	// Test invalid request (missing SessionID)
	invalidRequestNoSession := &GetResourceEventsRequest{
		SpaceID: "Spaces-1",
	}

	err = invalidRequestNoSession.Validate()
	assert.Error(t, err)

	// Test invalid request (missing SpaceID)
	invalidRequestNoSpace := &GetResourceEventsRequest{
		SessionID: ResourceEventsSessionId("session-123"),
	}

	err = invalidRequestNoSpace.Validate()
	assert.Error(t, err)
}

func TestGetResourceEventsResponse_Validate(t *testing.T) {
	// Test valid response
	validResponse := &GetResourceEventsResponse{
		Events:             []KubernetesEventResource{},
		IsSessionCompleted: true,
	}

	err := validResponse.Validate()
	assert.NoError(t, err)

	// Test valid response with events
	event := &KubernetesEventResource{
		FirstObservedTime:   time.Now(),
		LastObservedTime:    time.Now(),
		Count:               1,
		Action:              "Created",
		Reason:              "SuccessfulCreate",
		Note:                "Created pod successfully",
		ReportingController: "replicaset-controller",
		ReportingInstance:   "replicaset-controller-xyz",
		Type:                "Normal",
		Manifest:            "{}",
	}

	validResponseWithEvents := &GetResourceEventsResponse{
		Events:             []KubernetesEventResource{*event},
		IsSessionCompleted: false,
	}

	err = validResponseWithEvents.Validate()
	assert.NoError(t, err)
}

func TestKubernetesEventResource_Validate(t *testing.T) {
	// Test valid event resource
	validEvent := &KubernetesEventResource{
		FirstObservedTime:   time.Now(),
		LastObservedTime:    time.Now(),
		Count:               1,
		Action:              "Created",
		Reason:              "SuccessfulCreate",
		Note:                "Created pod successfully",
		ReportingController: "replicaset-controller",
		ReportingInstance:   "replicaset-controller-xyz",
		Type:                "Normal",
		Manifest:            "{}",
	}

	err := validEvent.Validate()
	assert.NoError(t, err)

	// Test invalid event resource (missing required fields)
	invalidEvent := &KubernetesEventResource{}

	err = invalidEvent.Validate()
	assert.Error(t, err)
}

func TestMonitorErrorResource_Validate(t *testing.T) {
	// Test valid error resource
	validError := &MonitorErrorResource{
		Message: "Something went wrong",
		Code:    "ERR001",
	}

	err := validError.Validate()
	assert.NoError(t, err)

	// Test valid error resource without code
	validErrorNoCode := &MonitorErrorResource{
		Message: "Something went wrong",
	}

	err = validErrorNoCode.Validate()
	assert.NoError(t, err)

	// Test invalid error resource (missing required message)
	invalidError := &MonitorErrorResource{}

	err = invalidError.Validate()
	assert.Error(t, err)
}

func TestGetResourceEventsResponseWithError(t *testing.T) {
	events := []KubernetesEventResource{}
	isSessionCompleted := false
	errorResource := NewMonitorErrorResource("Session timeout", "ERR_TIMEOUT")

	response := &GetResourceEventsResponse{
		Events:             events,
		IsSessionCompleted: isSessionCompleted,
		Error:              errorResource,
	}

	expected := &GetResourceEventsResponse{
		Events:             events,
		IsSessionCompleted: isSessionCompleted,
		Error: &MonitorErrorResource{
			Message: "Session timeout",
			Code:    "ERR_TIMEOUT",
		},
	}

	assert.Equal(t, expected, response)
	assert.NoError(t, response.Validate())
}
