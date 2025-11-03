package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2/ephemeralenvironments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/stretchr/testify/require"
)

func TestEnvironmentServiceCreateEphemeralEnvironment(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	parentEnvironment := CreateParentEnvironment(t, client)
	require.NotNil(t, parentEnvironment)
	defer DeleteParentEnvironment(t, client, parentEnvironment)

	ephemeralChannel := CreateEphemeralTestChannel(t, client, project, parentEnvironment)
	require.NotNil(t, ephemeralChannel)
	defer DeleteTestChannel(t, client, ephemeralChannel)

	createdEnvironmentId := CreateEphemeralEnvironment(t, client, project)
	//	defer DeleteTestEnvironment_NewClient(t, client, createdEnvironment)

	environments, err := ephemeralenvironments.GetAll(client, client.GetSpaceID())
	require.NoError(t, err)
	require.NotNil(t, environments)
	require.NotEmpty(t, environments.Items)

	require.Equal(t, createdEnvironmentId, environments.Items[0].ID)
}

func TestEnvironmentServiceDeprovisionEphemeralEnvironmentForProject(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	runbook := CreateTestRunbook(t, client, lifecycle, projectGroup, project)
	require.NotNil(t, runbook)
	defer DeleteTestRunbook(t, client, runbook)

	parentEnvironment := CreateParentEnvironment(t, client)
	require.NotNil(t, parentEnvironment)
	defer DeleteParentEnvironment(t, client, parentEnvironment)

	ephemeralChannel := CreateEphemeralTestChannel(t, client, project, parentEnvironment)
	require.NotNil(t, ephemeralChannel)
	defer DeleteTestChannel(t, client, ephemeralChannel)

	createdEnvironmentId := CreateEphemeralEnvironment(t, client, project)
	DeprovisionEphemeralEnvironmentForProject(t, client, &createdEnvironmentId, project)

	environments, err := ephemeralenvironments.GetAll(client, client.GetSpaceID())
	require.NoError(t, err)
	require.NotNil(t, environments)
}

func CreateEphemeralEnvironment(t *testing.T, client *client.Client, project *projects.Project) string {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	createdEnvironment, err := ephemeralenvironments.Add(client, client.GetSpaceID(), project.GetID(), name)
	require.NoError(t, err)
	require.NotNil(t, createdEnvironment)
	require.NotEmpty(t, createdEnvironment.Id)

	return createdEnvironment.Id
}

func DeprovisionEphemeralEnvironmentForProject(t *testing.T, client *client.Client, environmentId *string, project *projects.Project) ephemeralenvironments.DeprovisioningRunbookRun {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	deprovisionResponse, err := ephemeralenvironments.DeprovisionForProject(client, client.GetSpaceID(), *environmentId, project.GetID())
	require.NoError(t, err)
	require.NotNil(t, deprovisionResponse)

	return deprovisionResponse.DeprovisioningRun
}
