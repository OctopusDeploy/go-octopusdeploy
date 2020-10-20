package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVarAddAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	varProj := createVarTestProject(t, octopusClient, getRandomName())
	defer cleanProject(t, octopusClient, varProj.ID)
	varName := getRandomVarName()
	expected := getTestVariable(varName)
	actual := createTestVariable(t, varProj.ID, varName)
	defer cleanVar(t, octopusClient, actual.ID, varProj.ID)

	assert.Equal(t, expected.Name, actual.Name, "variable name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "variable doesn't contain an ID from the octopus server")
}

func createTestVariable(t *testing.T, projectID, variableName string) octopusdeploy.Variable {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	v := getTestVariable(variableName)
	variableSet, err := octopusClient.Variables.AddSingle(projectID, &v)
	if err != nil {
		t.Fatalf("creating variable %s failed when it shouldn't: %s", variableName, err)
	}

	for _, newV := range variableSet.Variables {
		if newV.Name == variableName {
			return newV
		}
	}

	t.Fatalf("Unable to locate variable named %s after creationg", variableName)
	return octopusdeploy.Variable{} //Blank variable to return
}

func getTestVariable(variableName string) octopusdeploy.Variable {
	v := octopusdeploy.NewVariable(variableName, "string", "octo-test value", "octo-test description", nil, false)

	return *v
}

func createVarTestProject(t *testing.T, octopusClient *octopusdeploy.Client, projectName string) octopusdeploy.Project {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	p := octopusdeploy.NewProject(projectName, "Lifecycles-1", "ProjectGroups-1")
	createdProject, err := octopusClient.Projects.Add(p)

	if err != nil {
		t.Fatalf("creating project %s failed when it shouldn't: %s", projectName, err)
	}

	return *createdProject
}

func cleanVar(t *testing.T, octopusClient *octopusdeploy.Client, varID string, projID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	_, err := octopusClient.Variables.DeleteSingle(projID, varID)
	assert.NoError(t, err)
}
