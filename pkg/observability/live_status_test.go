package observability

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLiveStatusRequest_Validate(t *testing.T) {
	// Test valid request
	validRequest := &GetLiveStatusRequest{
		SpaceID:       "Spaces-1",
		ProjectID:     "Projects-1",
		EnvironmentID: "Environments-1",
		SummaryOnly:   false,
	}

	err := validRequest.Validate()
	assert.NoError(t, err)

	// Test valid request with tenant
	tenantID := "Tenants-1"
	validRequestWithTenant := &GetLiveStatusRequest{
		SpaceID:       "Spaces-1",
		ProjectID:     "Projects-1",
		EnvironmentID: "Environments-1",
		TenantID:      &tenantID,
		SummaryOnly:   true,
	}

	err = validRequestWithTenant.Validate()
	assert.NoError(t, err)

	// Test invalid request (missing required fields)
	invalidRequest := &GetLiveStatusRequest{}

	err = invalidRequest.Validate()
	assert.Error(t, err)

	// Test invalid request (missing SpaceID)
	invalidRequestNoSpace := &GetLiveStatusRequest{
		ProjectID:     "Projects-1",
		EnvironmentID: "Environments-1",
	}

	err = invalidRequestNoSpace.Validate()
	assert.Error(t, err)

	// Test invalid request (missing ProjectID)
	invalidRequestNoProject := &GetLiveStatusRequest{
		SpaceID:       "Spaces-1",
		EnvironmentID: "Environments-1",
	}

	err = invalidRequestNoProject.Validate()
	assert.Error(t, err)

	// Test invalid request (missing EnvironmentID)
	invalidRequestNoEnvironment := &GetLiveStatusRequest{
		SpaceID:   "Spaces-1",
		ProjectID: "Projects-1",
	}

	err = invalidRequestNoEnvironment.Validate()
	assert.Error(t, err)
}

func TestGetLiveStatusResponse_Validate(t *testing.T) {
	// Test valid response
	summary := LiveStatusSummaryResource{
		Status:      "Healthy",
		LastUpdated: time.Now(),
	}

	validResponse := &GetLiveStatusResponse{
		MachineStatuses: []KubernetesMachineLiveStatusResource{},
		Summary:         summary,
	}

	err := validResponse.Validate()
	assert.NoError(t, err)

	// Test valid response with machine statuses
	machineStatus := KubernetesMachineLiveStatusResource{
		MachineID: "Machines-1",
		Status:    "Healthy",
		Resources: []KubernetesLiveStatusResource{},
	}

	validResponseWithMachines := &GetLiveStatusResponse{
		MachineStatuses: []KubernetesMachineLiveStatusResource{machineStatus},
		Summary:         summary,
	}

	err = validResponseWithMachines.Validate()
	assert.NoError(t, err)

	// Test invalid response (missing summary)
	invalidResponse := &GetLiveStatusResponse{
		MachineStatuses: []KubernetesMachineLiveStatusResource{},
	}

	err = invalidResponse.Validate()
	assert.Error(t, err)
}

func TestLiveStatusSummaryResource_Validate(t *testing.T) {
	// Test valid summary resource
	validSummary := &LiveStatusSummaryResource{
		Status:      "Healthy",
		LastUpdated: time.Now(),
	}

	err := validSummary.Validate()
	assert.NoError(t, err)

	// Test invalid summary resource (missing required fields)
	invalidSummary := &LiveStatusSummaryResource{}

	err = invalidSummary.Validate()
	assert.Error(t, err)

	// Test invalid summary resource (missing status)
	invalidSummaryNoStatus := &LiveStatusSummaryResource{
		LastUpdated: time.Now(),
	}

	err = invalidSummaryNoStatus.Validate()
	assert.Error(t, err)

	// Test invalid summary resource (missing lastUpdated)
	invalidSummaryNoTime := &LiveStatusSummaryResource{
		Status: "Healthy",
	}

	err = invalidSummaryNoTime.Validate()
	assert.Error(t, err)
}

func TestKubernetesMachineLiveStatusResource_Validate(t *testing.T) {
	// Test valid machine status resource
	validMachineStatus := &KubernetesMachineLiveStatusResource{
		MachineID: "Machines-1",
		Status:    "Healthy",
		Resources: []KubernetesLiveStatusResource{},
	}

	err := validMachineStatus.Validate()
	assert.NoError(t, err)

	// Test valid machine status resource with resources
	resource := KubernetesLiveStatusResource{
		Name:         "my-deployment",
		Kind:         "Deployment",
		HealthStatus: "Healthy",
		MachineID:    "Machines-1",
		Children:     []KubernetesLiveStatusResource{},
	}

	validMachineStatusWithResources := &KubernetesMachineLiveStatusResource{
		MachineID: "Machines-1",
		Status:    "Healthy",
		Resources: []KubernetesLiveStatusResource{resource},
	}

	err = validMachineStatusWithResources.Validate()
	assert.NoError(t, err)

	// Test invalid machine status resource (missing required fields)
	invalidMachineStatus := &KubernetesMachineLiveStatusResource{}

	err = invalidMachineStatus.Validate()
	assert.Error(t, err)

	// Test invalid machine status resource (missing MachineID)
	invalidMachineStatusNoID := &KubernetesMachineLiveStatusResource{
		Status:    "Healthy",
		Resources: []KubernetesLiveStatusResource{},
	}

	err = invalidMachineStatusNoID.Validate()
	assert.Error(t, err)

	// Test invalid machine status resource (missing Status)
	invalidMachineStatusNoStatus := &KubernetesMachineLiveStatusResource{
		MachineID: "Machines-1",
		Resources: []KubernetesLiveStatusResource{},
	}

	err = invalidMachineStatusNoStatus.Validate()
	assert.Error(t, err)
}

func TestKubernetesLiveStatusResource_Validate(t *testing.T) {
	// Test valid resource
	validResource := &KubernetesLiveStatusResource{
		Name:         "my-deployment",
		Kind:         "Deployment",
		HealthStatus: "Healthy",
		MachineID:    "Machines-1",
		Children:     []KubernetesLiveStatusResource{},
	}

	err := validResource.Validate()
	assert.NoError(t, err)

	// Test valid resource with optional fields
	namespace := "default"
	syncStatus := "Synced"
	desiredResourceID := "desired-123"
	monitoredResourceID := "monitored-456"

	validResourceWithOptional := &KubernetesLiveStatusResource{
		Name:                "my-service",
		Namespace:           &namespace,
		Kind:                "Service",
		HealthStatus:        "Healthy",
		SyncStatus:          &syncStatus,
		MachineID:           "Machines-1",
		Children:            []KubernetesLiveStatusResource{},
		DesiredResourceID:   &desiredResourceID,
		MonitoredResourceID: &monitoredResourceID,
	}

	err = validResourceWithOptional.Validate()
	assert.NoError(t, err)

	// Test valid resource with children
	child := KubernetesLiveStatusResource{
		Name:         "my-pod",
		Kind:         "Pod",
		HealthStatus: "Healthy",
		MachineID:    "Machines-1",
		Children:     []KubernetesLiveStatusResource{},
	}

	validResourceWithChildren := &KubernetesLiveStatusResource{
		Name:         "my-deployment",
		Kind:         "Deployment",
		HealthStatus: "Healthy",
		MachineID:    "Machines-1",
		Children:     []KubernetesLiveStatusResource{child},
	}

	err = validResourceWithChildren.Validate()
	assert.NoError(t, err)

	// Test invalid resource (missing required fields)
	invalidResource := &KubernetesLiveStatusResource{}

	err = invalidResource.Validate()
	assert.Error(t, err)

	// Test invalid resource (missing Name)
	invalidResourceNoName := &KubernetesLiveStatusResource{
		Kind:         "Deployment",
		HealthStatus: "Healthy",
		MachineID:    "Machines-1",
		Children:     []KubernetesLiveStatusResource{},
	}

	err = invalidResourceNoName.Validate()
	assert.Error(t, err)

	// Test invalid resource (missing Kind)
	invalidResourceNoKind := &KubernetesLiveStatusResource{
		Name:         "my-deployment",
		HealthStatus: "Healthy",
		MachineID:    "Machines-1",
		Children:     []KubernetesLiveStatusResource{},
	}

	err = invalidResourceNoKind.Validate()
	assert.Error(t, err)

	// Test invalid resource (missing HealthStatus)
	invalidResourceNoHealth := &KubernetesLiveStatusResource{
		Name:      "my-deployment",
		Kind:      "Deployment",
		MachineID: "Machines-1",
		Children:  []KubernetesLiveStatusResource{},
	}

	err = invalidResourceNoHealth.Validate()
	assert.Error(t, err)

	// Test invalid resource (missing MachineID)
	invalidResourceNoMachine := &KubernetesLiveStatusResource{
		Name:         "my-deployment",
		Kind:         "Deployment",
		HealthStatus: "Healthy",
		Children:     []KubernetesLiveStatusResource{},
	}

	err = invalidResourceNoMachine.Validate()
	assert.Error(t, err)
}

func TestGetLiveStatusResponseWithError(t *testing.T) {
	summary := LiveStatusSummaryResource{
		Status:      "Error",
		LastUpdated: time.Now(),
	}

	errorResource := NewMonitorErrorResource("Connection failed", "ERR_CONNECTION")

	response := &GetLiveStatusResponse{
		MachineStatuses: []KubernetesMachineLiveStatusResource{},
		Summary:         summary,
		Error:           errorResource,
	}

	expected := &GetLiveStatusResponse{
		MachineStatuses: []KubernetesMachineLiveStatusResource{},
		Summary:         summary,
		Error: &MonitorErrorResource{
			Message: "Connection failed",
			Code:    "ERR_CONNECTION",
		},
	}

	assert.Equal(t, expected, response)
	assert.NoError(t, response.Validate())
}
