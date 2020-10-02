package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvironments(t *testing.T) {
	t.Run("AddAndDelete", TestEnvironmentAddAndDelete)
	t.Run("AddGetAndDelete", TestEnvironmentAddGetAndDelete)
	t.Run("GetThatDoesNotExist", TestEnvironmentGetThatDoesNotExist)
	t.Run("GetAll", TestEnvironmentGetAll)
	t.Run("Update", TestEnvironmentUpdate)
	t.Run("GetByName", TestEnvironmentGetByName)
}

func TestEnvironmentAddAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	environmentName := getRandomName()
	expected := getTestEnvironment(environmentName)
	actual := createTestEnvironment(t, octopusClient, environmentName)

	defer cleanEnvironment(t, octopusClient, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "environment name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "environment doesn't contain an ID from the octopus server")
}

func TestEnvironmentAddGetAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	environment := createTestEnvironment(t, octopusClient, getRandomName())
	defer cleanEnvironment(t, octopusClient, environment.ID)

	getEnvironment, err := octopusClient.Environments.GetByID(environment.ID)
	require.NoError(t, err)
	require.Equal(t, environment.Name, getEnvironment.Name)
}

func TestEnvironmentGetThatDoesNotExist(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	id := getRandomName()
	environment, err := octopusClient.Environments.GetByID(id)
	require.Equal(t, createResourceNotFoundError("environment", "ID", id), err)
	require.Nil(t, environment)
}

func TestEnvironmentGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	// create many environments to test pagination
	environmentsToCreate := 32
	sum := 0
	for i := 0; i < environmentsToCreate; i++ {
		environment := createTestEnvironment(t, octopusClient, getRandomName())
		defer cleanEnvironment(t, octopusClient, environment.ID)
		sum += i
	}

	allEnvironments, err := octopusClient.Environments.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all environments failed when it shouldn't: %s", err)
	}

	numberOfEnvironments := len(allEnvironments)

	// check there are greater than or equal to the amount of environments requested to be created, otherwise pagination isn't working
	if numberOfEnvironments < environmentsToCreate {
		t.Fatalf("There should be at least %d environments created but there was only %d. Pagination is likely not working.", environmentsToCreate, numberOfEnvironments)
	}

	additionalEnvironment := createTestEnvironment(t, octopusClient, getRandomName())
	defer cleanEnvironment(t, octopusClient, additionalEnvironment.ID)

	allEnvironmentsAfterCreatingAdditional, err := octopusClient.Environments.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all environments failed when it shouldn't: %s", err)
	}

	assert.NoError(t, err, "error when looking for environment when not expected")
	assert.Equal(t, len(allEnvironmentsAfterCreatingAdditional), numberOfEnvironments+1, "created an additional environment and expected number of environments to increase by 1")
}

func TestEnvironmentUpdate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	environment := createTestEnvironment(t, octopusClient, getRandomName())
	defer cleanEnvironment(t, octopusClient, environment.ID)

	newEnvironmentName := getRandomName()
	const newDescription = "this should be updated"
	const newUseGuidedFailure = true

	environment.Name = newEnvironmentName
	environment.Description = newDescription
	environment.UseGuidedFailure = newUseGuidedFailure

	updatedEnvironment, err := octopusClient.Environments.Update(environment)
	assert.NoError(t, err, "error when updating environment")
	assert.Equal(t, newEnvironmentName, updatedEnvironment.Name, "environment name was not updated")
	assert.Equal(t, newDescription, updatedEnvironment.Description, "environment description was not updated")
	assert.Equal(t, newUseGuidedFailure, environment.UseGuidedFailure, "environment UseGuidedFailure was not updated")
}

func TestEnvironmentGetByName(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	expected := createTestEnvironment(t, octopusClient, getRandomName())
	defer cleanEnvironment(t, octopusClient, expected.ID)

	resources, err := octopusClient.Environments.GetByName(expected.Name)

	assert := assert.New(t)

	assert.NoError(err)
	assert.NotNil(resources)

	// equality cannot be determined through a direct comparison (below)
	// because GetByPartialName does not include the fields, LastModifiedBy and
	// LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	for _, actual := range resources {
		assert.Equal(expected.AllowDynamicInfrastructure, actual.AllowDynamicInfrastructure)
		assert.Equal(expected.Description, actual.Description)
		assert.Equal(expected.ID, actual.ID)
		assert.Equal(expected.Name, actual.Name)
		assert.Equal(expected.SortOrder, actual.SortOrder)
		assert.Equal(expected.UseGuidedFailure, actual.UseGuidedFailure)
	}
}

/*
func TestEnvironmentCleanup(t *testing.T) {
	allEnvironments, err := octopusClient.Environments.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all environments failed when it shouldn't: %s", err)
	}
	for _, e := range *allEnvironments {
		if len(e.Name) > 17 && e.Name[0:16] == "go-octopusdeploy" {
			octopusClient.Environments.Delete(e.ID)
		}
	}
}
*/

func createTestEnvironment(t *testing.T, octopusClient *client.Client, environmentName string) model.Environment {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	e := getTestEnvironment(environmentName)
	createdEnvironment, err := octopusClient.Environments.Add(&e)

	if err != nil {
		t.Fatalf("creating environment %s failed when it shouldn't: %s", environmentName, err)
	}

	return *createdEnvironment
}

func getTestEnvironment(environmentName string) model.Environment {
	e := model.NewEnvironment(environmentName, "Environment from testing suite", true)
	return *e
}

func cleanEnvironment(t *testing.T, octopusClient *client.Client, environmentID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	err := octopusClient.Environments.DeleteByID(environmentID)
	assert.NoError(t, err)
}
