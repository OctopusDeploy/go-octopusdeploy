package integration

import (
	"net/http"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateLibraryVariableSet(t *testing.T, client *octopusdeploy.Client) *octopusdeploy.LibraryVariableSet {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := "Name " + getShortRandomName()
	libraryVariableSet := octopusdeploy.NewLibraryVariableSet(name)
	libraryVariableSet.Description = "Description " + getShortRandomName()

	actionTemplateParameter := octopusdeploy.NewActionTemplateParameter()
	propertyValue := octopusdeploy.NewPropertyValue(getShortRandomName(), false)
	actionTemplateParameter.DefaultValue = &propertyValue
	actionTemplateParameter.DisplaySettings = map[string]string{
		"Octopus.ControlType": "SingleLineText",
	}
	actionTemplateParameter.HelpText = "Help Text " + getShortRandomName()
	actionTemplateParameter.Label = "Label " + getShortRandomName()
	actionTemplateParameter.Name = "Name " + getShortRandomName()
	libraryVariableSet.Templates = append(libraryVariableSet.Templates, *actionTemplateParameter)

	actionTemplateParameter = octopusdeploy.NewActionTemplateParameter()
	propertyValue = octopusdeploy.NewPropertyValue(getShortRandomName(), false)
	actionTemplateParameter.DefaultValue = &propertyValue
	actionTemplateParameter.DisplaySettings = map[string]string{
		"Octopus.ControlType": "SingleLineText",
	}
	actionTemplateParameter.HelpText = "Help Text " + getShortRandomName()
	actionTemplateParameter.Label = "Label " + getShortRandomName()
	actionTemplateParameter.Name = "Name " + getShortRandomName()
	libraryVariableSet.Templates = append(libraryVariableSet.Templates, *actionTemplateParameter)

	createdLibraryVariableSet, err := client.LibraryVariableSets.Add(libraryVariableSet)
	require.NoError(t, err)
	require.NotNil(t, createdLibraryVariableSet)

	name = "Name " + getShortRandomName()
	variable := octopusdeploy.NewVariable(name)
	variable.Description = "Description " + getShortRandomName()
	variable.IsEditable = false
	variable.IsSensitive = true
	variable.Value = "Value " + getShortRandomName()

	variableSet, err := client.Variables.AddSingle(createdLibraryVariableSet.GetID(), variable)
	require.NoError(t, err)
	require.NotNil(t, variableSet)

	name = "Name " + getShortRandomName()
	variable = octopusdeploy.NewVariable(name)
	variable.Description = "Description " + getShortRandomName()
	variable.Value = "Value " + getShortRandomName()

	variableSet, err = client.Variables.AddSingle(createdLibraryVariableSet.GetID(), variable)
	require.NoError(t, err)
	require.NotNil(t, variableSet)

	return createdLibraryVariableSet
}

func DeleteLibraryVariableSet(t *testing.T, client *octopusdeploy.Client, libraryVariableSet *octopusdeploy.LibraryVariableSet) {
	require.NotNil(t, libraryVariableSet)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.LibraryVariableSets.DeleteByID(libraryVariableSet.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedLibraryVariableSet, err := client.LibraryVariableSets.GetByID(libraryVariableSet.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedLibraryVariableSet)
}

func TestLibraryVariableSetServiceAddDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	libraryVariableSet := CreateLibraryVariableSet(t, client)
	require.NotNil(t, libraryVariableSet)
	defer DeleteLibraryVariableSet(t, client, libraryVariableSet)

	name := libraryVariableSet.Name
	query := octopusdeploy.LibraryVariablesQuery{
		PartialName: name,
		Take:        1,
	}
	namedLibraryVariableSets, err := client.LibraryVariableSets.Get(query)
	require.NoError(t, err)
	require.NotNil(t, namedLibraryVariableSets)

	query = octopusdeploy.LibraryVariablesQuery{
		IDs:  []string{libraryVariableSet.GetID()},
		Take: 1,
	}
	namedLibraryVariableSets, err = client.LibraryVariableSets.Get(query)
	require.NoError(t, err)
	require.NotNil(t, namedLibraryVariableSets)
}

func TestLibraryVariablesGet(t *testing.T) {
	octopusClient, err := octopusdeploy.GetFakeOctopusClient(t, "/api/libraryvariablesets/LibraryVariables-41", http.StatusOK, getLibraryVariablesResponseJSON)
	require.NoError(t, err)
	require.NotNil(t, octopusClient)

	libraryVariables, err := octopusClient.LibraryVariableSets.GetByID("LibraryVariables-41")
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
	libraryVariables := octopusdeploy.NewLibraryVariableSet("My Set")
	assert.Nil(t, octopusdeploy.ValidateLibraryVariableSetValues(libraryVariables))
}

func TestValidateLibraryVariablesValuesMissingNameFails(t *testing.T) {
	libraryVariables := &octopusdeploy.LibraryVariableSet{}
	assert.Error(t, octopusdeploy.ValidateLibraryVariableSetValues(libraryVariables))
}
