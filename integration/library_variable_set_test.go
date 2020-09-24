package integration

import (
	"net/http"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

func TestLibraryVariableSetGet(t *testing.T) {
	client, err := client.GetFakeOctopusClient(t, "/api/libraryvariablesets/LibraryVariableSets-41", http.StatusOK, getLibraryVariableSetResponseJSON)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	if err != nil {
		return
	}

	libraryVariableSet, err := client.LibraryVariableSets.GetByID("LibraryVariableSets-41")

	assert.NoError(t, err)

	if err != nil {
		return
	}

	assert.Equal(t, "MySet", libraryVariableSet.Name)
	assert.Equal(t, "The Description", libraryVariableSet.Description)
	assert.Equal(t, "variableset-LibraryVariableSets-41", libraryVariableSet.VariableSetID)
	assert.Equal(t, enum.Variables, libraryVariableSet.ContentType)
}

const getLibraryVariableSetResponseJSON = `
{
  "Id": "LibraryVariableSets-41",
  "Name": "MySet",
  "Description": "The Description",
  "VariableSetId": "variableset-LibraryVariableSets-41",
  "ContentType": "Variables",
  "Templates": [],
  "Links": {
    "Self": "/api/libraryvariablesets/LibraryVariableSets-481",
    "Variables": "/api/variables/variableset-LibraryVariableSets-481"
  }
}`

func TestValidateLibraryVariableSetValuesJustANamePasses(t *testing.T) {

	libraryVariableSet := model.NewLibraryVariableSet("My Set")

	assert.Nil(t, model.ValidateLibraryVariableSetValues(libraryVariableSet))
}

func TestValidateLibraryVariableSetValuesMissingNameFails(t *testing.T) {

	libraryVariableSet := &model.LibraryVariableSet{}

	assert.Error(t, model.ValidateLibraryVariableSetValues(libraryVariableSet))
}
