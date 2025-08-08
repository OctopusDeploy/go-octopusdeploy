package e2e

import (
	"strings"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/kubernetesmonitors"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/machines"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRegisterMonitor(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := internal.GetRandomName()
	err := machines.DeleteByID(client, client.GetSpaceID(), id)
	require.Error(t, err)

	environment := CreateTestEnvironment_NewClient(t, client)
	require.NotNil(t, environment)
	defer DeleteTestEnvironment_NewClient(t, client, environment)

	createdDeploymentTarget := CreateTestKubernetesDeploymentTarget_NewClient(t, client, environment)
	require.NotNil(t, createdDeploymentTarget)
	defer DeleteTestDeploymentTarget_NewClient(t, client, createdDeploymentTarget)

	monitorId, err := uuid.NewUUID()
	require.NoError(t, err)

	// Create a register command
	registerMonitorCommand := kubernetesmonitors.NewRegisterKubernetesMonitorCommand(
		&monitorId,
		createdDeploymentTarget.GetID(),
	)

	// Register the monitor
	res, err := kubernetesmonitors.Register(client, registerMonitorCommand)
	require.NoError(t, err)
	require.EqualValues(t, *res.Resource.InstallationID, monitorId)

	// Verify the monitor can be retrieved
	retrievedMonitor, err := kubernetesmonitors.GetByID(client, client.GetSpaceID(), res.Resource.ID)
	require.NoError(t, err)
	require.NotNil(t, retrievedMonitor)
	require.Equal(t, res.Resource.ID, retrievedMonitor.Resource.ID)

	// Delete the monitor
	err = kubernetesmonitors.DeleteByID(client, client.GetSpaceID(), res.Resource.ID)
	require.NoError(t, err)

	// Verify the monitor is deleted
	_, err = kubernetesmonitors.GetByID(client, client.GetSpaceID(), res.Resource.ID)
	require.Error(t, err)
	require.Contains(t, err.Error(), "not found")

}

func CreateTestKubernetesDeploymentTarget_NewClient(
	t *testing.T, client *client.Client, environment *environments.Environment,
) *machines.DeploymentTarget {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	thumbprint := strings.ToUpper(internal.GetRandomThumbprint()[:16])

	environmentIDs := []string{environment.GetID()}
	roles := []string{"Prod"}

	// Create a polling URL
	pollingUrl := internal.GetRandomPollingAddress()

	endpoint := machines.NewKubernetesTentacleEndpoint(pollingUrl, thumbprint, false, "Polling", "default")

	deploymentTarget := machines.NewDeploymentTarget(name, endpoint, environmentIDs, roles)
	deploymentTarget.IsDisabled = true
	deploymentTarget.MachinePolicyID = "MachinePolicies-1"
	deploymentTarget.Status = "Disabled"
	deploymentTarget.Thumbprint = thumbprint

	require.NoError(t, deploymentTarget.Validate())

	createdDeploymentTarget, err := machines.Add(client, deploymentTarget)
	require.NoError(t, err)
	require.NotNil(t, createdDeploymentTarget)
	require.NotEmpty(t, createdDeploymentTarget.GetID())

	return createdDeploymentTarget
}
