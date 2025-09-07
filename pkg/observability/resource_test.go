package observability

import (
	"encoding/json"
	"testing"
	"time"

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
			CreationTimestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	err := validResource.Validate()
	assert.NoError(t, err)

	// Test empty resource (should still be valid since all fields are optional)
	emptyResource := &KubernetesLiveStatusDetailedResource{}

	err = emptyResource.Validate()
	assert.NoError(t, err)
}

func TestManifestSummary_Interface(t *testing.T) {
	// Test that ManifestSummary implements ManifestSummaryResource interface
	var resource ManifestSummaryResource = &ManifestSummary{
		Labels: map[string]string{
			"app": "test",
		},
		Annotations: map[string]string{
			"version": "1.0",
		},
		CreationTimestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	assert.Equal(t, map[string]string{"app": "test"}, resource.GetLabels())
	assert.Equal(t, map[string]string{"version": "1.0"}, resource.GetAnnotations())
	assert.Equal(t, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), resource.GetCreationTimestamp())
	assert.NoError(t, resource.Validate())
}

func TestPodManifestSummary_Validate(t *testing.T) {
	// Test valid pod manifest summary
	validPodSummary := &PodManifestSummary{
		ManifestSummary: ManifestSummary{
			Labels: map[string]string{
				"app": "test-pod",
			},
			Annotations: map[string]string{
				"kubernetes.io/managed-by": "octopus",
			},
			CreationTimestamp: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Containers: []string{"nginx", "sidecar"},
	}

	err := validPodSummary.Validate()
	assert.NoError(t, err)

	// Test interface implementation
	var resource ManifestSummaryResource = validPodSummary
	assert.Equal(t, map[string]string{"app": "test-pod"}, resource.GetLabels())
	assert.Equal(t, []string{"nginx", "sidecar"}, validPodSummary.GetContainers())

	// Test empty pod summary (should still be valid since all fields are optional)
	emptyPodSummary := &PodManifestSummary{}
	err = emptyPodSummary.Validate()
	assert.NoError(t, err)
}

func TestKubernetesLiveStatusDetailedResource_UnmarshalJSON_Pod(t *testing.T) {
	// Test JSON unmarshaling for Pod resource with Containers field
	podJSON := `{
		"Name": "test-api-deploy-7bb97cf98-d7rb6",
		"Namespace": "default",
		"Kind": "Pod",
		"HealthStatus": "Healthy",
		"SyncStatus": null,
		"MachineId": "Machines-61",
		"LastUpdated": "2025-08-08T01:41:15.027298Z",
		"ManifestSummary": {
			"Containers": ["test-nginx"],
			"Kind": "Pod",
			"Labels": {
				"Octopus.Action.Id": "d7419a77-ff63-4072-a861-dd7ab4ec4229",
				"Octopus.Deployment.Id": "deployments-145",
				"pod-template-hash": "7bb97cf98"
			},
			"Annotations": {},
			"CreationTimestamp": "2025-08-01T05:54:32.000+00:00"
		},
		"Children": [],
		"DesiredResourceId": null,
		"ResourceId": "0992a29a-5070-4314-9b34-034a81200dc9"
	}`

	var resource KubernetesLiveStatusDetailedResource
	err := json.Unmarshal([]byte(podJSON), &resource)
	assert.NoError(t, err)

	// Verify basic fields
	assert.Equal(t, "test-api-deploy-7bb97cf98-d7rb6", resource.Name)
	assert.Equal(t, "default", resource.Namespace)
	assert.Equal(t, "Pod", resource.Kind)
	assert.Equal(t, "Healthy", resource.HealthStatus)
	assert.Equal(t, "Machines-61", resource.MachineID)
	assert.Equal(t, "0992a29a-5070-4314-9b34-034a81200dc9", resource.ResourceID)

	// Verify ManifestSummary is properly parsed as PodManifestSummary
	assert.NotNil(t, resource.ManifestSummary)
	
	// Type assert to PodManifestSummary to access Containers
	podSummary, ok := resource.ManifestSummary.(*PodManifestSummary)
	assert.True(t, ok, "ManifestSummary should be of type *PodManifestSummary for Pod resources")
	assert.NotNil(t, podSummary)
	
	// Verify Containers field is properly parsed
	assert.Equal(t, []string{"test-nginx"}, podSummary.GetContainers())
	
	// Verify basic ManifestSummary interface methods work
	assert.Contains(t, resource.ManifestSummary.GetLabels(), "Octopus.Action.Id")
	assert.Equal(t, "d7419a77-ff63-4072-a861-dd7ab4ec4229", resource.ManifestSummary.GetLabels()["Octopus.Action.Id"])
}

func TestKubernetesLiveStatusDetailedResource_UnmarshalJSON_NonPod(t *testing.T) {
	// Test JSON unmarshaling for non-Pod resource (should use base ManifestSummary)
	deploymentJSON := `{
		"Name": "test-deployment",
		"Namespace": "default",
		"Kind": "Deployment",
		"HealthStatus": "Healthy",
		"MachineId": "Machines-61",
		"LastUpdated": "2025-08-08T01:41:15.027298Z",
		"ManifestSummary": {
			"Labels": {
				"app": "test-app"
			},
			"Annotations": {
				"deployment.kubernetes.io/revision": "1"
			},
			"CreationTimestamp": "2025-08-01T05:54:32.000+00:00"
		},
		"ResourceId": "deployment-123"
	}`

	var resource KubernetesLiveStatusDetailedResource
	err := json.Unmarshal([]byte(deploymentJSON), &resource)
	assert.NoError(t, err)

	// Verify basic fields
	assert.Equal(t, "test-deployment", resource.Name)
	assert.Equal(t, "Deployment", resource.Kind)

	// Verify ManifestSummary is parsed as base ManifestSummary (not PodManifestSummary)
	assert.NotNil(t, resource.ManifestSummary)
	
	// Type assert to make sure it's the base type, not PodManifestSummary
	manifestSummary, ok := resource.ManifestSummary.(*ManifestSummary)
	assert.True(t, ok, "ManifestSummary should be of type *ManifestSummary for non-Pod resources")
	assert.NotNil(t, manifestSummary)
	
	// Verify interface methods work
	assert.Equal(t, "test-app", resource.ManifestSummary.GetLabels()["app"])
	assert.Equal(t, "1", resource.ManifestSummary.GetAnnotations()["deployment.kubernetes.io/revision"])
}

func TestGetResourceResponse_UnmarshalJSON_Pod(t *testing.T) {
	// Test that full GetResourceResponse unmarshaling works with Pod
	responseJSON := `{
		"Resource": {
			"Name": "test-api-deploy-7bb97cf98-d7rb6",
			"Namespace": "default",
			"Kind": "Pod",
			"HealthStatus": "Healthy",
			"SyncStatus": null,
			"MachineId": "Machines-61",
			"LastUpdated": "2025-08-08T01:41:15.027298Z",
			"ManifestSummary": {
				"Containers": ["test-nginx"],
				"Kind": "Pod",
				"Labels": {
					"Octopus.Action.Id": "d7419a77-ff63-4072-a861-dd7ab4ec4229",
					"pod-template-hash": "7bb97cf98"
				},
				"Annotations": {},
				"CreationTimestamp": "2025-08-01T05:54:32.000+00:00"
			},
			"Children": [],
			"DesiredResourceId": null,
			"ResourceId": "0992a29a-5070-4314-9b34-034a81200dc9"
		}
	}`

	var response GetResourceResponse
	err := json.Unmarshal([]byte(responseJSON), &response)
	assert.NoError(t, err)

	// Verify the resource is properly parsed
	assert.NotNil(t, response.Resource)
	assert.Equal(t, "Pod", response.Resource.Kind)
	
	// Verify ManifestSummary contains Containers for Pod
	podSummary, ok := response.Resource.ManifestSummary.(*PodManifestSummary)
	assert.True(t, ok, "ManifestSummary should be PodManifestSummary for Pod resources")
	assert.Equal(t, []string{"test-nginx"}, podSummary.GetContainers())
}
