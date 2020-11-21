package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createActionTemplate(t *testing.T) *octopusdeploy.ActionTemplate {
	resource := octopusdeploy.NewActionTemplate(getRandomName(), octopusdeploy.ActionTypeOctopusScript)
	require.NotNil(t, resource)

	resource.Properties = map[string]octopusdeploy.PropertyValue{}
	resource.Properties[octopusdeploy.ActionTypeOctopusActionScriptBody] = octopusdeploy.PropertyValue(getRandomName())

	return resource
}

func IsEqualActionTemplates(t *testing.T, expected *octopusdeploy.ActionTemplate, actual *octopusdeploy.ActionTemplate) {
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
	assert.True(t, IsEqualLinks(expected.GetLinks(), actual.GetLinks()))

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

	invalidResource := &octopusdeploy.ActionTemplate{}
	resource, err := client.ActionTemplates.Add(invalidResource)
	assert.NotNil(t, err)
	assert.Nil(t, resource)

	resource = createActionTemplate(t)
	require.NotNil(t, resource)

	resource, err = client.ActionTemplates.Add(resource)
	require.NoError(t, err)
	require.NotNil(t, resource)

	err = client.ActionTemplates.DeleteByID(resource.GetID())
	assert.NoError(t, err)
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

	id := getRandomName()
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

	resource, err := client.ActionTemplates.Search()
	assert.NoError(t, err)
	assert.NotEmpty(t, resource)
}
