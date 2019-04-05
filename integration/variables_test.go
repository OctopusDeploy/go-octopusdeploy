package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
)

func init() {
	client = initTest()
}

func TestVarAddAndDelete(t *testing.T) {
	varProj := createVarTestProject(t, getRandomName())
	defer cleanProject(t, varProj.ID)
	varName := getRandomVarName()
	expected := getTestVariable(varName)
	actual := createTestVariable(t, varProj.ID, varName)
	defer cleanVar(t, actual.ID, varProj.ID)

	assert.Equal(t, expected.Name, actual.Name, "variable name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "variable doesn't contain an ID from the octopus server")
}

func createTestVariable(t *testing.T, projectID, variableName string) octopusdeploy.Variable {
	v := getTestVariable(variableName)
	variableSet, err := client.Variable.AddSingle(projectID, &v)
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

func createVarTestProject(t *testing.T, projectName string) octopusdeploy.Project {
	p := octopusdeploy.NewProject(projectName, "Lifecycles-1", "ProjectGroups-1")
	createdProject, err := client.Project.Add(p)

	if err != nil {
		t.Fatalf("creating project %s failed when it shouldn't: %s", projectName, err)
	}

	return *createdProject
}
func cleanVar(t *testing.T, varID string, projID string) {
	_, err := client.Variable.DeleteSingle(projID, varID)
	if err == nil {
		return
	}
	if err == octopusdeploy.ErrItemNotFound {
		return
	}
	if err != nil {
		t.Fatalf("deleting variable failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
