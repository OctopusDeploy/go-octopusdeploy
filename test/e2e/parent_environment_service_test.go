package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments/v2/parentenvironments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateParentEnvironment(t *testing.T, client *client.Client) *parentenvironments.ParentEnvironment {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()
	description := "Description for " + name + " (OK to Delete)"
	useGuidedFailure := createRandomBoolean()

	environment := parentenvironments.NewParentEnvironment(name, client.GetSpaceID())
	environment.Description = description
	environment.UseGuidedFailure = useGuidedFailure

	require.NoError(t, environment.Validate())

	createdEnvironment, err := parentenvironments.Add(client, environment)
	require.NoError(t, err)
	require.NotNil(t, createdEnvironment)
	require.NotEmpty(t, createdEnvironment.ID)

	return createdEnvironment
}

func DeleteParentEnvironment(t *testing.T, client *client.Client, environment *parentenvironments.ParentEnvironment) {
	require.NotNil(t, environment)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := parentenvironments.DeleteByID(client, environment.SpaceID, environment.ID)
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedEnvironment, err := parentenvironments.GetByID(client, environment.SpaceID, environment.ID)
	assert.Error(t, err)
	assert.Nil(t, deletedEnvironment)
}
