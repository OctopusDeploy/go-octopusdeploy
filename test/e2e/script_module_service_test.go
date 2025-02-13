package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/scriptmodules"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScriptModuleAddGetAndDelete_NewClient(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	name := internal.GetRandomName()
	description := internal.GetRandomName()
	scriptBody := "function Say-Hello()\r\n{\r\n    Write-Output \"Hello, Octopus!\"\r\n}\r\n"
	syntax := "PowerShell"

	scriptModule := variables.NewScriptModule(name)
	scriptModule.Description = description
	scriptModule.ScriptBody = scriptBody
	scriptModule.Syntax = syntax
	require.NoError(t, scriptModule.Validate())

	createdScriptModule, err := scriptmodules.Add(client, scriptModule)
	require.NoError(t, err)
	require.NotNil(t, createdScriptModule)
	require.NotEmpty(t, createdScriptModule.GetID())
	require.Equal(t, description, createdScriptModule.Description)
	require.Equal(t, name, createdScriptModule.Name)
	require.Equal(t, scriptBody, createdScriptModule.ScriptBody)
	require.Equal(t, syntax, createdScriptModule.Syntax)

	allScriptModules, err := scriptmodules.Get(client, space.GetID(), variables.LibraryVariablesQuery{
		ContentType: "ScriptModules",
	})

	require.NoError(t, err)
	require.NotNil(t, allScriptModules)

	for _, resource := range allScriptModules.Items {
		resourceToCompare, err := scriptmodules.GetByID(client, space.GetID(), resource.GetID())
		require.NoError(t, err)
		IsEqualScriptModules(t, resource, resourceToCompare)
	}

	DeleteTestScriptModule(t, client, space.GetID(), createdScriptModule)
}

func IsEqualScriptModules(t *testing.T, expected *variables.ScriptModule, actual *variables.ScriptModule) {
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
	assert.True(t, internal.IsLinksEqual(expected.GetLinks(), actual.GetLinks()))

	// script module
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ScriptBody, actual.ScriptBody)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
	assert.Equal(t, expected.Syntax, actual.Syntax)
	assert.Equal(t, expected.VariableSetID, actual.VariableSetID)
}

func DeleteTestScriptModule(t *testing.T, client *client.Client, spaceID string, scriptModule *variables.ScriptModule) error {
	require.NotNil(t, scriptModule)

	return scriptmodules.DeleteByID(client, spaceID, scriptModule.GetID())
}
