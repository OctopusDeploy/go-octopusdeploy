package e2e

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/libraryvariableset"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actiontemplates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateLibraryVariableSet(t *testing.T, client *client.Client) *variables.LibraryVariableSet {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := "Name " + getShortRandomName()
	libraryVariableSet := variables.NewLibraryVariableSet(name)
	libraryVariableSet.Description = "Description " + getShortRandomName()

	actionTemplateParameter := actiontemplates.NewActionTemplateParameter()
	propertyValue := core.NewPropertyValue(getShortRandomName(), false)
	actionTemplateParameter.DefaultValue = &propertyValue
	actionTemplateParameter.DisplaySettings = map[string]string{
		"Octopus.ControlType": "SingleLineText",
	}
	actionTemplateParameter.HelpText = "Help Text " + getShortRandomName()
	actionTemplateParameter.Label = "Label " + getShortRandomName()
	actionTemplateParameter.Name = "Name " + getShortRandomName()
	libraryVariableSet.Templates = append(libraryVariableSet.Templates, *actionTemplateParameter)

	actionTemplateParameter = actiontemplates.NewActionTemplateParameter()
	propertyValue = core.NewPropertyValue(getShortRandomName(), false)
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
	variable := variables.NewVariable(name)
	variable.Description = "Description " + getShortRandomName()
	variable.IsEditable = false
	variable.IsSensitive = true
	variable.Value = "Value " + getShortRandomName()

	variableSet, err := client.Variables.AddSingle(createdLibraryVariableSet.GetID(), variable)
	require.NoError(t, err)
	require.NotNil(t, variableSet)

	name = "Name " + getShortRandomName()
	variable = variables.NewVariable(name)
	variable.Description = "Description " + getShortRandomName()
	variable.Value = "Value " + getShortRandomName()

	variableSet, err = client.Variables.AddSingle(createdLibraryVariableSet.GetID(), variable)
	require.NoError(t, err)
	require.NotNil(t, variableSet)

	return createdLibraryVariableSet
}

func DeleteLibraryVariableSet(t *testing.T, client *client.Client, libraryVariableSet *variables.LibraryVariableSet) {
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
	query := variables.LibraryVariablesQuery{
		PartialName: name,
		Take:        1,
	}
	namedLibraryVariableSets, err := client.LibraryVariableSets.Get(query)
	require.NoError(t, err)
	require.NotNil(t, namedLibraryVariableSets)

	query = variables.LibraryVariablesQuery{
		IDs:  []string{libraryVariableSet.GetID()},
		Take: 1,
	}
	namedLibraryVariableSets, err = client.LibraryVariableSets.Get(query)
	require.NoError(t, err)
	require.NotNil(t, namedLibraryVariableSets)
}

// TODO: fix test
// func TestLibraryVariablesGet(t *testing.T) {
// 	octopusClient, err := octopusdeploy.GetFakeOctopusClient(t, "/api/libraryvariablesets/LibraryVariables-41", http.StatusOK, getLibraryVariablesResponseJSON)
// 	require.NoError(t, err)
// 	require.NotNil(t, octopusClient)

// 	libraryVariables, err := octopusClient.LibraryVariableSets.GetByID("LibraryVariables-41")
// 	require.NoError(t, err)
// 	require.Equal(t, "MySet", libraryVariables.Name)
// 	require.Equal(t, "The Description", libraryVariables.Description)
// 	require.Equal(t, "variableset-LibraryVariables-41", libraryVariables.VariableSetID)
// 	require.Equal(t, "Variables", libraryVariables.ContentType)
// }

// const getLibraryVariablesResponseJSON = `
// {
//   "Id": "LibraryVariables-41",
//   "Name": "MySet",
//   "Description": "The Description",
//   "VariableSetId": "variableset-LibraryVariables-41",
//   "ContentType": "Variables",
//   "Templates": [],
//   "Links": {
//     "Self": "/api/libraryvariablesets/LibraryVariables-481",
//     "Variables": "/api/variables/variableset-LibraryVariables-481"
//   }
// }`

func TestValidateLibraryVariablesValuesJustANamePasses(t *testing.T) {
	libraryVariables := variables.NewLibraryVariableSet("My Set")
	assert.Nil(t, variables.ValidateLibraryVariableSetValues(libraryVariables))
}

func TestValidateLibraryVariablesValuesMissingNameFails(t *testing.T) {
	libraryVariables := &variables.LibraryVariableSet{}
	assert.Error(t, variables.ValidateLibraryVariableSetValues(libraryVariables))
}

// ----- new -----

func CreateLibraryVariableSet_NewClient(t *testing.T, client *client.Client) *variables.LibraryVariableSet {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := "Name " + getShortRandomName()
	libraryVariableSet := variables.NewLibraryVariableSet(name)
	libraryVariableSet.Description = "Description " + getShortRandomName()

	actionTemplateParameter := actiontemplates.NewActionTemplateParameter()
	propertyValue := core.NewPropertyValue(getShortRandomName(), false)
	actionTemplateParameter.DefaultValue = &propertyValue
	actionTemplateParameter.DisplaySettings = map[string]string{
		"Octopus.ControlType": "SingleLineText",
	}
	actionTemplateParameter.HelpText = "Help Text " + getShortRandomName()
	actionTemplateParameter.Label = "Label " + getShortRandomName()
	actionTemplateParameter.Name = "Name " + getShortRandomName()
	libraryVariableSet.Templates = append(libraryVariableSet.Templates, *actionTemplateParameter)

	actionTemplateParameter = actiontemplates.NewActionTemplateParameter()
	propertyValue = core.NewPropertyValue(getShortRandomName(), false)
	actionTemplateParameter.DefaultValue = &propertyValue
	actionTemplateParameter.DisplaySettings = map[string]string{
		"Octopus.ControlType": "SingleLineText",
	}
	actionTemplateParameter.HelpText = "Help Text " + getShortRandomName()
	actionTemplateParameter.Label = "Label " + getShortRandomName()
	actionTemplateParameter.Name = "Name " + getShortRandomName()
	libraryVariableSet.Templates = append(libraryVariableSet.Templates, *actionTemplateParameter)

	createdLibraryVariableSet, err := libraryvariableset.Add(client, libraryVariableSet)
	require.NoError(t, err)
	require.NotNil(t, createdLibraryVariableSet)

	name = "Name " + getShortRandomName()
	variable := variables.NewVariable(name)
	variable.Description = "Description " + getShortRandomName()
	variable.IsEditable = false
	variable.IsSensitive = true
	variable.Value = "Value " + getShortRandomName()

	variableSet, err := variables.AddSingle(client, createdLibraryVariableSet.SpaceID, createdLibraryVariableSet.GetID(), variable)
	require.NoError(t, err)
	require.NotNil(t, variableSet)

	name = "Name " + getShortRandomName()
	variable = variables.NewVariable(name)
	variable.Description = "Description " + getShortRandomName()
	variable.Value = "Value " + getShortRandomName()

	variableSet, err = variables.AddSingle(client, createdLibraryVariableSet.SpaceID, createdLibraryVariableSet.GetID(), variable)
	require.NoError(t, err)
	require.NotNil(t, variableSet)

	return createdLibraryVariableSet
}

func DeleteLibraryVariableSet_NewClient(t *testing.T, client *client.Client, libraryVariableSet *variables.LibraryVariableSet) {
	require.NotNil(t, libraryVariableSet)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := libraryvariableset.DeleteByID(client, libraryVariableSet.SpaceID, libraryVariableSet.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedLibraryVariableSet, err := libraryvariableset.GetByID(client, libraryVariableSet.SpaceID, libraryVariableSet.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedLibraryVariableSet)
}

func TestLibraryVariableSetAddGetDelete_NewClient(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	libraryVariableSet := CreateLibraryVariableSet_NewClient(t, client)
	require.NotNil(t, libraryVariableSet)
	defer DeleteLibraryVariableSet_NewClient(t, client, libraryVariableSet)

	name := libraryVariableSet.Name
	query := variables.LibraryVariablesQuery{
		PartialName: name,
		Take:        1,
	}
	namedLibraryVariableSets, err := client.LibraryVariableSets.Get(query)
	require.NoError(t, err)
	require.NotNil(t, namedLibraryVariableSets)

	query = variables.LibraryVariablesQuery{
		IDs:  []string{libraryVariableSet.GetID()},
		Take: 1,
	}
	namedLibraryVariableSets, err = client.LibraryVariableSets.Get(query)
	require.NoError(t, err)
	require.NotNil(t, namedLibraryVariableSets)
}

func TestLibraryVariableSetUpdate_NewClient(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	libraryVariableSet := CreateLibraryVariableSet_NewClient(t, client)
	require.NotNil(t, libraryVariableSet)
	defer DeleteLibraryVariableSet_NewClient(t, client, libraryVariableSet)

	libraryVariableSet.Description = "updated description"

	expectedProjectName := libraryVariableSet.Name
	expectedContentType := libraryVariableSet.ContentType
	expectedDescription := libraryVariableSet.Description

	updatedLibraryVariableSet, err := libraryvariableset.Update(client, libraryVariableSet)

	require.NoError(t, err)
	require.Equal(t, expectedProjectName, updatedLibraryVariableSet.Name, "libraryvariableset name was not updated")
	require.Equal(t, expectedContentType, updatedLibraryVariableSet.ContentType, "libraryvariableset contentType was not updated")
	require.Equal(t, expectedDescription, updatedLibraryVariableSet.Description, "libraryvariableset description was updated")
}
