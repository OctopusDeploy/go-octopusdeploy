package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getTestVariable(name string) *octopusdeploy.Variable {
	variable := octopusdeploy.NewVariable(name)
	variable.Description = "octo-test description"
	variable.Value = "octo-test value"

	return variable
}

func CreateTestVariable(t *testing.T, ownerID string, name string) *octopusdeploy.Variable {
	client := getOctopusClient()
	require.NotNil(t, client)

	variable := getTestVariable(name)
	variableSet, err := client.Variables.AddSingle(ownerID, variable)
	require.NoError(t, err)
	require.Len(t, variableSet.Variables, 1)

	for _, v := range variableSet.Variables {
		if v.Name == name {
			createdVariable, err := client.Variables.GetByID(ownerID, v.GetID())
			require.NoError(t, err)
			require.NotNil(t, createdVariable)

			return createdVariable
		}
	}

	return nil
}

func DeleteTestVariable(t *testing.T, octopusClient *octopusdeploy.Client, variableID string, ownerID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	variableSet, err := octopusClient.Variables.DeleteSingle(ownerID, variableID)
	assert.NoError(t, err)
	assert.NotNil(t, variableSet)
}

func TestVariableService(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	name := getRandomVarName()
	expected := getTestVariable(name)
	createdVariable := CreateTestVariable(t, project.GetID(), name)
	defer DeleteTestVariable(t, client, createdVariable.GetID(), project.GetID())

	require.Equal(t, expected.Name, createdVariable.Name, "variable name doesn't match expected")
	require.NotEmpty(t, createdVariable.GetID(), "variable doesn't contain an ID from the octopus server")

	variable, err := client.Variables.GetByID(project.GetID(), createdVariable.GetID())
	require.NoError(t, err)
	require.Equal(t, expected.Name, variable.Name, "variable name doesn't match expected")
	require.NotEmpty(t, variable.GetID(), "variable doesn't contain an ID from the octopus server")
}

func TestVariableServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	ownerID := getShortRandomName()
	variableID := getShortRandomName()

	variable, err := client.Variables.GetByID(ownerID, variableID)
	require.Error(t, err)
	require.Nil(t, variable)

	apiError := err.(*octopusdeploy.APIError)
	require.Equal(t, apiError.StatusCode, 404)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	variable, err = client.Variables.GetByID(project.GetID(), variableID)
	require.Error(t, err)
	require.Nil(t, variable)

	apiError = err.(*octopusdeploy.APIError)
	require.Equal(t, apiError.StatusCode, 404)
}

func TestVariableServiceDeleteSingle(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	ownerID := getShortRandomName()
	variableID := getShortRandomName()

	variableSet, err := client.Variables.DeleteSingle(ownerID, variableID)
	require.Error(t, err)
	require.NotNil(t, variableSet)

	apiError := err.(*octopusdeploy.APIError)
	require.Equal(t, apiError.StatusCode, 404)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	variableSet, err = client.Variables.DeleteSingle(project.GetID(), variableID)
	require.Error(t, err)
	require.NotNil(t, variableSet)

	apiError = err.(*octopusdeploy.APIError)
	require.Equal(t, apiError.StatusCode, 404)

	expectedVariableSet, err := client.Variables.GetAll(project.GetID())
	require.NoError(t, err)
	require.NotNil(t, expectedVariableSet.ScopeValues)

	name := getRandomVarName()
	expected := getTestVariable(name)
	createdVariable := CreateTestVariable(t, project.GetID(), name)

	require.Equal(t, expected.Name, createdVariable.Name, "variable name doesn't match expected")
	require.NotEmpty(t, createdVariable.GetID(), "variable doesn't contain an ID from the octopus server")

	variableSet, err = client.Variables.DeleteSingle(project.GetID(), createdVariable.GetID())
	require.NoError(t, err)
	require.NotNil(t, variableSet)
	require.Equal(t, expectedVariableSet.ScopeValues, variableSet.ScopeValues)
}
