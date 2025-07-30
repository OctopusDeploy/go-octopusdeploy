package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/livestatusservice"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/machines"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupClientForTest sets up the Octopus client and new client for testing
func setupClientForTest(t *testing.T) (*client.Client, newclient.Client) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient, "octopusClient should not be nil - check environment variables")

	// Validate the client has required methods
	httpSession := octopusClient.HttpSession()
	require.NotNil(t, httpSession, "HttpSession should not be nil")

	clientSpaceID := octopusClient.GetSpaceID()
	require.NotEmpty(t, clientSpaceID, "SpaceID should not be empty")

	// Create a new client instance for the livestatus service
	newClient := newclient.NewClientS(httpSession, clientSpaceID)
	require.NotNil(t, newClient, "newClient should not be nil")

	return octopusClient, newClient
}

func TestGetResourceManifestWithClient(t *testing.T) {

	octopusClient, newClient := setupClientForTest(t)

	// Create test environment
	environment := CreateTestEnvironment(t, octopusClient)
	require.NotNil(t, environment)
	defer DeleteTestEnvironment(t, octopusClient, environment)

	space := GetDefaultSpace(t, octopusClient)

	lifecycle := CreateTestLifecycle(t, octopusClient)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, octopusClient, lifecycle)

	projectGroup := CreateTestProjectGroup(t, octopusClient)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, octopusClient, projectGroup)

	// Create test project
	project := CreateTestProject(t, octopusClient, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, octopusClient, project)

	// Create test deployment target (machine)
	deploymentTarget := CreateTestDeploymentTarget(t, octopusClient, environment)
	require.NotNil(t, deploymentTarget)
	defer CleanResourceManifestDeploymentTarget(t, octopusClient, deploymentTarget)

	// Test with untenanted request
	t.Run("GetResourceManifest_Untenanted", func(t *testing.T) {
		request := &livestatusservice.GetResourceManifestRequest{
			SpaceID:                                octopusClient.GetSpaceID(),
			ProjectID:                              project.GetID(),
			EnvironmentID:                          environment.GetID(),
			MachineID:                              deploymentTarget.GetID(),
			DesiredOrKubernetesMonitoredResourceID: "test-resource-id",
		}

		// Validate the request is properly formed
		assert.True(t, request.IsUntenanted())
		assert.False(t, request.IsTenanted())
		err := request.Validate()
		assert.NoError(t, err)

		result, err := livestatusservice.GetResourceManifestWithClient(newClient, request)

		// We expect this to fail since we don't have actual kubernetes resources deployed,
		// but we want to verify it doesn't fail due to parameter validation or URI construction
		if err != nil {
			// Verify it's an HTTP error (not a parameter validation error)
			assert.NotContains(t, err.Error(), "parameter")
			assert.NotContains(t, err.Error(), "invalid")
		} else {
			// If successful, verify the response structure
			assert.NotNil(t, result)
			assert.NotEmpty(t, result.LiveManifest)
		}
	})

	// Test with tenanted request (if we have tenants available)
	t.Run("GetResourceManifest_Tenanted", func(t *testing.T) {
		request := &livestatusservice.GetResourceManifestRequest{
			SpaceID:                                octopusClient.GetSpaceID(),
			ProjectID:                              project.GetID(),
			EnvironmentID:                          environment.GetID(),
			MachineID:                              deploymentTarget.GetID(),
			DesiredOrKubernetesMonitoredResourceID: "test-resource-id",
			TenantID:                               "Tenants-1",
		}

		// Validate the request is properly formed
		assert.False(t, request.IsUntenanted())
		assert.True(t, request.IsTenanted())
		err := request.Validate()
		assert.NoError(t, err)

		result, err := livestatusservice.GetResourceManifestWithClient(newClient, request)

		if err != nil {
			assert.NotContains(t, err.Error(), "parameter")
			assert.NotContains(t, err.Error(), "invalid")
		} else {
			assert.NotNil(t, result)
			assert.NotEmpty(t, result.LiveManifest)
		}
	})
}

func TestGetResourceManifestWithClient_ErrorCases(t *testing.T) {

	octopusClient, newClient := setupClientForTest(t)

	t.Run("NilRequest", func(t *testing.T) {
		result, err := livestatusservice.GetResourceManifestWithClient(newClient, nil)
		assert.Nil(t, result)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request")
	})

	t.Run("InvalidIDs", func(t *testing.T) {
		request := &livestatusservice.GetResourceManifestRequest{
			SpaceID:                                octopusClient.GetSpaceID(),
			ProjectID:                              "invalid-project-id",
			EnvironmentID:                          "invalid-environment-id",
			MachineID:                              "invalid-machine-id",
			DesiredOrKubernetesMonitoredResourceID: "test-resource-id",
		}

		result, err := livestatusservice.GetResourceManifestWithClient(newClient, request)

		// Should fail with HTTP error (404 or similar), not parameter validation
		assert.Nil(t, result)
		assert.Error(t, err)
		// Error should be HTTP-related, not parameter validation
		assert.NotContains(t, err.Error(), "parameter")
		assert.NotContains(t, err.Error(), "invalid parameter")
	})
}

func CleanResourceManifestDeploymentTarget(t *testing.T, client *client.Client, deploymentTarget *machines.DeploymentTarget) {
	if client == nil || deploymentTarget == nil {
		return
	}

	err := client.Machines.DeleteByID(deploymentTarget.GetID())
	require.NoError(t, err)
}
