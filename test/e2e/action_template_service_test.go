package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actiontemplates"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createActionTemplate(t *testing.T) *actiontemplates.ActionTemplate {
	resource := actiontemplates.NewActionTemplate(internal.GetRandomName(), constants.ActionTypeOctopusScript)
	require.NotNil(t, resource)

	resource.Properties = map[string]core.PropertyValue{}
	resource.Properties[constants.ActionTypeOctopusActionScriptBody] = core.NewPropertyValue(internal.GetRandomName(), false)

	return resource
}

func IsEqualActionTemplates(t *testing.T, expected *actiontemplates.ActionTemplate, actual *actiontemplates.ActionTemplate) {
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

	// ActionTemplate
	assert.Equal(t, expected.ActionType, actual.ActionType)
	assert.Equal(t, expected.CommunityActionTemplateID, actual.CommunityActionTemplateID)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Packages, actual.Packages)
	assert.Equal(t, expected.Parameters, actual.Parameters)
	assert.Equal(t, expected.Properties, actual.Properties)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
	assert.Equal(t, expected.Version, actual.Version)
}

func TestActionTemplateServiceAdd(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	invalidResource := &actiontemplates.ActionTemplate{}
	resource, err := client.ActionTemplates.Add(invalidResource)
	assert.NotNil(t, err)
	assert.Nil(t, resource)

	resource = createActionTemplate(t)
	require.NotNil(t, resource)

	resource, err = client.ActionTemplates.Add(resource)
	require.NoError(t, err)
	require.NotNil(t, resource)
	defer client.ActionTemplates.DeleteByID(resource.GetID())
}

func TestActionTemplateServiceGetCategories(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	resource, err := client.ActionTemplates.GetCategories()
	assert.NoError(t, err)
	assert.NotEmpty(t, resource)
}

func TestActionTemplateServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := internal.GetRandomName()
	resource, err := client.ActionTemplates.GetByID(id)
	assert.NotNil(t, err)
	assert.Nil(t, resource)

	resources, err := client.ActionTemplates.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		resourceToCompare, err := client.ActionTemplates.GetByID(resource.GetID())
		require.NoError(t, err)
		IsEqualActionTemplates(t, resource, resourceToCompare)
	}
}

func TestActionTemplateServiceSearch(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	search := ""

	resource, err := client.ActionTemplates.Search(search)
	assert.NoError(t, err)
	assert.NotEmpty(t, resource)

	search = "Octopus.Script"

	resource, err = client.ActionTemplates.Search(search)
	assert.NoError(t, err)
	assert.NotEmpty(t, resource)
}

func AddActionTemplateTest(t *testing.T, client *client.Client) *actiontemplates.ActionTemplate {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()
	description := "Description for " + name + " (OK to Delete)"
	actiontemplate := actiontemplates.NewActionTemplate(name, "Octopus.Script")
	actiontemplate.Description = description
	actiontemplate.Properties = map[string]core.PropertyValue{
		"Octopus.Action.Script.ScriptBody":   core.NewPropertyValue("echo \"test\"", false),
		"Octopus.Action.Script.ScriptSource": core.NewPropertyValue("Inline", false),
		"Octopus.Action.Script.Syntax":       core.NewPropertyValue("PowerShell", false),
	}
	defaultValue := core.NewPropertyValue("Testing", false)
	actiontemplate.Parameters = []actiontemplates.ActionTemplateParameter{
		{
			DefaultValue: &defaultValue,
			DisplaySettings: map[string]string{
				"Octopus.ControlType": "SingleLineText",
			},
			HelpText: description,
			Label:    name,
		},
	}

	require.NoError(t, actiontemplate.Validate())

	createdActionTemplate, err := actiontemplates.Add(client, actiontemplate)
	require.NoError(t, err)
	require.NotNil(t, createdActionTemplate)
	require.NotEmpty(t, createdActionTemplate.GetID())

	return createdActionTemplate
}

func DeleteActionTemplateTest(t *testing.T, client *client.Client, actionTemplate *actiontemplates.ActionTemplate) {
	require.NotNil(t, actionTemplate)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := actiontemplates.DeleteByID(client, actionTemplate.SpaceID, actionTemplate.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedEnvironment, err := actiontemplates.GetByID(client, actionTemplate.SpaceID, actionTemplate.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedEnvironment)
}

func TestActionTemplateServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	createdActionTemplate := AddActionTemplateTest(t, client)
	defer DeleteActionTemplateTest(t, client, createdActionTemplate)

	actionTemplate, err := actiontemplates.GetByID(client, createdActionTemplate.SpaceID, createdActionTemplate.GetID())
	require.NoError(t, err)
	require.NotNil(t, actionTemplate)
}
