package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestEnvironment(t *testing.T, client *octopusdeploy.client) *octopusdeploy.Environment {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	allowDynamicInfrastructure := createRandomBoolean()
	name := getRandomName()
	description := "Description for " + name + " (OK to Delete)"
	useGuidedFailure := createRandomBoolean()

	environment := octopusdeploy.NewEnvironment(name)
	environment.AllowDynamicInfrastructure = allowDynamicInfrastructure
	environment.Description = description
	environment.UseGuidedFailure = useGuidedFailure

	require.NoError(t, environment.Validate())

	createdEnvironment, err := client.Environments.Add(environment)
	require.NoError(t, err)
	require.NotNil(t, createdEnvironment)
	require.NotEmpty(t, createdEnvironment.GetID())

	return createdEnvironment
}

func DeleteTestEnvironment(t *testing.T, client *octopusdeploy.client, environment *octopusdeploy.Environment) {
	require.NotNil(t, environment)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Environments.DeleteByID(environment.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedEnvironment, err := client.Environments.GetByID(environment.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedEnvironment)
}

func IsEqualEnvironments(t *testing.T, expected *octopusdeploy.Environment, actual *octopusdeploy.Environment) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	// IResource
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, IsEqualLinks(expected.GetLinks(), actual.GetLinks()))

	// Environment
	assert.Equal(t, expected.AllowDynamicInfrastructure, actual.AllowDynamicInfrastructure)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.SortOrder, actual.SortOrder)
	assert.Equal(t, expected.UseGuidedFailure, actual.UseGuidedFailure)
}

func UpdateEnvironment(t *testing.T, client *octopusdeploy.client, environment *octopusdeploy.Environment) *octopusdeploy.Environment {
	require.NotNil(t, environment)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	updatedEnvironment, err := client.Environments.Update(environment)
	assert.NoError(t, err)
	require.NotNil(t, updatedEnvironment)

	return updatedEnvironment
}

func TestEnvironmentServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	createdEnvironment := CreateTestEnvironment(t, client)
	defer DeleteTestEnvironment(t, client, createdEnvironment)

	environment, err := client.Environments.GetByID(createdEnvironment.GetID())
	require.NoError(t, err)
	require.NotNil(t, environment)
}

func TestEnvironmentServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	environments, err := client.Environments.GetAll()
	require.NoError(t, err)
	require.NotNil(t, environments)

	for _, environment := range environments {
		defer DeleteTestEnvironment(t, client, environment)
	}
}

func TestEnvironmentServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// create 30 test environments (to be deleted)
	for i := 0; i < 30; i++ {
		environment := CreateTestEnvironment(t, client)
		require.NotNil(t, environment)
		defer DeleteTestEnvironment(t, client, environment)
	}

	allEnvironments, err := client.Environments.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allEnvironments)
	require.True(t, len(allEnvironments) >= 30)

	for _, environment := range allEnvironments {
		environmentToCompare, err := client.Environments.GetByID(environment.GetID())
		require.NoError(t, err)
		require.NotNil(t, environment)
		require.NotEmpty(t, environment.GetID())
		IsEqualEnvironments(t, environment, environmentToCompare)
	}
}

func TestEnvironmentServiceUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	environment := CreateTestEnvironment(t, client)
	defer DeleteTestEnvironment(t, client, environment)

	newAllowDynamicInfrastructure := createRandomBoolean()
	newDescription := getRandomName()
	newName := getRandomName()
	newSortOrder := environment.SortOrder + 1
	newUseGuidedFailure := createRandomBoolean()

	environment.AllowDynamicInfrastructure = newAllowDynamicInfrastructure
	environment.Description = newDescription
	environment.Name = newName
	environment.SortOrder = newSortOrder
	environment.UseGuidedFailure = newUseGuidedFailure

	updatedEnvironment := UpdateEnvironment(t, client, environment)
	require.NotNil(t, updatedEnvironment)

	require.NotEmpty(t, updatedEnvironment.GetID())
	require.Equal(t, updatedEnvironment.GetID(), updatedEnvironment.GetID())
	require.Equal(t, newAllowDynamicInfrastructure, updatedEnvironment.AllowDynamicInfrastructure)
	require.Equal(t, newDescription, updatedEnvironment.Description)
	require.Equal(t, newName, updatedEnvironment.Name)
	require.Equal(t, newSortOrder, updatedEnvironment.SortOrder)
	require.Equal(t, newUseGuidedFailure, updatedEnvironment.UseGuidedFailure)
}
