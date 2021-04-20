package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVarAddAndDelete(t *testing.T) {
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

	varName := getRandomVarName()
	expected := getTestVariable(varName)
	actual := createTestVariable(t, project.GetID(), varName)
	defer cleanVar(t, client, actual.GetID(), project.GetID())

	assert.Equal(t, expected.Name, actual.Name, "variable name doesn't match expected")
	assert.NotEmpty(t, actual.GetID(), "variable doesn't contain an ID from the octopus server")
}

func createTestVariable(t *testing.T, projectID, variableName string) *octopusdeploy.Variable {
	client := getOctopusClient()
	require.NotNil(t, client)

	v := getTestVariable(variableName)
	variableSet, err := client.Variables.AddSingle(projectID, v)
	if err != nil {
		t.Fatalf("creating variable %s failed when it shouldn't: %s", variableName, err)
	}

	for _, newV := range variableSet.Variables {
		if newV.Name == variableName {
			return newV
		}
	}

	t.Fatalf("Unable to locate variable, %s after creation", variableName)
	return nil //Blank variable to return
}

func getTestVariable(variableName string) *octopusdeploy.Variable {
	v := octopusdeploy.NewVariable(variableName)
	v.Description = "octo-test description"
	v.Value = "octo-test value"

	return v
}

func cleanVar(t *testing.T, octopusClient *octopusdeploy.Client, varID string, projID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	_, err := octopusClient.Variables.DeleteSingle(projID, varID)
	assert.NoError(t, err)
}
