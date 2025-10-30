package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2/ephemeralenvironments"
	"github.com/stretchr/testify/require"
)

func TestEnvironmentServiceCreateEphemeralEnvironment_NewClient(t *testing.T) {
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

	channel := CreateTestChannel(t, client, project)
	require.NotNil(t, channel)
	defer DeleteTestChannel(t, client, channel)

	parentEnvironment := CreateParentEnvironment(t, client)
	require.NotNil(t, parentEnvironment)
	defer DeleteParentEnvironment(t, client, parentEnvironment)

	ephemeralChannel := CreateEphemeralTestChannel(t, client, project, parentEnvironment)
	require.NotNil(t, ephemeralChannel)
	defer DeleteTestChannel(t, client, ephemeralChannel)

	createdEnvironmentId := CreateEphemeralEnvironment(t, client, project.ID)
	//	defer DeleteTestEnvironment_NewClient(t, client, createdEnvironment)

	environments, err := ephemeralenvironments.GetAll(client, client.GetSpaceID())
	require.NoError(t, err)
	require.NotNil(t, environments)

	require.Equal(t, createdEnvironmentId, environments.Items[0].ID)
}

func CreateEphemeralEnvironment(t *testing.T, client *client.Client, projectId string) string {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	createdEnvironment, err := ephemeralenvironments.Create(client, client.GetSpaceID(), projectId, name)
	require.NoError(t, err)
	require.NotNil(t, createdEnvironment)
	require.NotEmpty(t, createdEnvironment.Id)

	return createdEnvironment.Id
}
