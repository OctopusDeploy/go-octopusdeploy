package observability

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetResourceRequest_IsTenanted(t *testing.T) {
	t.Run("should return untenanted resource", func(t *testing.T) {
		// Test untenanted request
		untenanteRequest := &GetResourceRequest{
			SpaceID:                                "Spaces-1",
			ProjectID:                              "Projects-1",
			EnvironmentID:                          "Environments-1",
			MachineID:                              "Machines-1",
			TenantID:                               "",
			DesiredOrKubernetesMonitoredResourceID: "Resources-1",
		}

		assert.False(t, untenanteRequest.IsTenanted())
	})

	t.Run("should return tenanted resource", func(t *testing.T) {
		// Test tenanted request
		tenantedRequest := &GetResourceRequest{
			SpaceID:                                "Spaces-1",
			ProjectID:                              "Projects-1",
			EnvironmentID:                          "Environments-1",
			MachineID:                              "Machines-1",
			TenantID:                               "Tenants-1",
			DesiredOrKubernetesMonitoredResourceID: "Resources-1",
		}

		assert.True(t, tenantedRequest.IsTenanted())
	})
}

func TestGetResourceRequest_Validate(t *testing.T) {
	// Test valid request
	validRequest := &GetResourceRequest{
		SpaceID:                                "Spaces-1",
		ProjectID:                              "Projects-1",
		EnvironmentID:                          "Environments-1",
		MachineID:                              "Machines-1",
		DesiredOrKubernetesMonitoredResourceID: "Resources-1",
	}

	err := validRequest.Validate()
	assert.NoError(t, err)

	// Test invalid request (missing required fields)
	invalidRequest := &GetResourceRequest{}

	err = invalidRequest.Validate()
	assert.Error(t, err)
}

func TestGetResourceResponse_Validate(t *testing.T) {
	// Test valid response
	resource := &KubernetesLiveStatusDetailedResource{
		Kind:      "Pod",
		Name:      "test-pod",
		Namespace: "default",
	}
	validResponse := &GetResourceResponse{
		Resource: resource,
	}

	err := validResponse.Validate()
	assert.NoError(t, err)

	// Test invalid response (missing required Resource)
	invalidResponse := &GetResourceResponse{}

	err = invalidResponse.Validate()
	assert.Error(t, err)
}

func TestKubernetesLiveStatusDetailedResource_Validate(t *testing.T) {
	// Test valid detailed resource
	validResource := &KubernetesLiveStatusDetailedResource{
		Kind:         "Pod",
		Name:         "test-pod",
		Namespace:    "default",
		ResourceID:   "12345-67890",
		HealthStatus: "Healthy",
		LastUpdated:  "2024-01-01T00:00:00Z",
		ManifestSummary: &ManifestSummary{
			Labels: map[string]string{
				"app": "test",
			},
			Annotations: map[string]string{
				"kubectl.kubernetes.io/last-applied-configuration": "{}",
			},
			CreationTimestamp: "2024-01-01T00:00:00Z",
		},
	}

	err := validResource.Validate()
	assert.NoError(t, err)

	// Test empty resource (should still be valid since all fields are optional)
	emptyResource := &KubernetesLiveStatusDetailedResource{}

	err = emptyResource.Validate()
	assert.NoError(t, err)
}
