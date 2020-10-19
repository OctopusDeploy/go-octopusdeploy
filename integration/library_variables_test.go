package integration

import (
	"net/http"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLibraryVariablesGet(t *testing.T) {
	client, err := client.GetFakeOctopusClient(t, "/api/libraryvariablesets/LibraryVariables-41", http.StatusOK, getLibraryVariablesResponseJSON)
	require.NoError(t, err)
	require.NotNil(t, client)

	libraryVariables, err := client.LibraryVariables.GetByID("LibraryVariables-41")
	require.NoError(t, err)
	require.Equal(t, "MySet", libraryVariables.Name)
	require.Equal(t, "The Description", libraryVariables.Description)
	require.Equal(t, "variableset-LibraryVariables-41", libraryVariables.VariableSetID)
	require.Equal(t, "Variables", libraryVariables.ContentType)
}

const getLibraryVariablesResponseJSON = `
{
  "Id": "LibraryVariables-41",
  "Name": "MySet",
  "Description": "The Description",
  "VariableSetId": "variableset-LibraryVariables-41",
  "ContentType": "Variables",
  "Templates": [],
  "Links": {
    "Self": "/api/libraryvariablesets/LibraryVariables-481",
    "Variables": "/api/variables/variableset-LibraryVariables-481"
  }
}`

func TestValidateLibraryVariablesValuesJustANamePasses(t *testing.T) {

	libraryVariables := model.NewLibraryVariableSet("My Set")

	assert.Nil(t, model.ValidateLibraryVariableSetValues(libraryVariables))
}

func TestValidateLibraryVariablesValuesMissingNameFails(t *testing.T) {

	libraryVariables := &model.LibraryVariableSet{}

	assert.Error(t, model.ValidateLibraryVariableSetValues(libraryVariables))
}