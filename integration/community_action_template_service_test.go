package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func IsEqualCommunityActionTemplates(t *testing.T, expected *octopusdeploy.CommunityActionTemplate, actual *octopusdeploy.CommunityActionTemplate) {
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

	// community action template
	assert.Equal(t, expected.Author, actual.Author)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.HistoryURL, actual.HistoryURL)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Type, actual.Type)
	assert.Equal(t, expected.Version, actual.Version)
	assert.Equal(t, expected.Website, actual.Website)
}

func TestCommunityActionTemplateServiceGetBy(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	invalidID := getRandomName()
	communityActionTemplate, err := octopusClient.CommunityActionTemplates.GetByID(invalidID)
	require.Equal(t, err, createResourceNotFoundError(serviceCommunityActionTemplateService, "ID", invalidID))
	require.Nil(t, communityActionTemplate)

	invalidName := getRandomName()
	communityActionTemplate, err = octopusClient.CommunityActionTemplates.GetByName(invalidName)
	require.Equal(t, err, createResourceNotFoundError(serviceCommunityActionTemplateService, "name", invalidName))
	require.Nil(t, communityActionTemplate)

	communityActionTemplates, err := octopusClient.CommunityActionTemplates.GetAll()
	require.NoError(t, err)
	require.NotNil(t, communityActionTemplates)

	if len(communityActionTemplates) > 10 {
		for _, communityActionTemplate := range communityActionTemplates[0:10] {
			communityActionTemplateToCompare, err := octopusClient.CommunityActionTemplates.GetByID(communityActionTemplate.GetID())
			require.NoError(t, err)
			IsEqualCommunityActionTemplates(t, communityActionTemplate, communityActionTemplateToCompare)

			communityActionTemplateToCompare, err = octopusClient.CommunityActionTemplates.GetByName(communityActionTemplate.Name)
			require.NoError(t, err)
			IsEqualCommunityActionTemplates(t, communityActionTemplate, communityActionTemplateToCompare)
		}
	}
}

func TestCommunityActionTemplateServiceGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	communityActionTemplates, err := octopusClient.CommunityActionTemplates.GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, communityActionTemplates)
}

func TestCommunityActionTemplateServiceGetByIDs(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	resources, err := octopusClient.CommunityActionTemplates.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	ids := []string{}
	for _, resource := range resources {
		ids = append(ids, resource.GetID())
	}

	// no need to test if ID collection size is less than 2
	if len(ids) < 2 {
		return
	}

	resourceListToCompare, err := octopusClient.CommunityActionTemplates.GetByIDs(ids[0:2])
	assert.NoError(t, err)
	assert.NotNil(t, resourceListToCompare)
	assert.Equal(t, 2, len(resourceListToCompare))
}

func TestCommunityActionTemplateServiceInstall(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	resource, err := octopusClient.CommunityActionTemplates.Install(octopusdeploy.CommunityActionTemplate{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = octopusdeploy.NewCommunityActionTemplate(getRandomName())
	require.NotNil(t, resource)
}
