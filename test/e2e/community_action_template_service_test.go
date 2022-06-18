package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/actions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func IsEqualCommunityActionTemplates(t *testing.T, expected *actions.CommunityActionTemplate, actual *actions.CommunityActionTemplate) {
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
	assert.True(t, internal.IsEqualLinks(expected.GetLinks(), actual.GetLinks()))

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

	invalidID := internal.GetRandomName()
	communityActionTemplate, err := octopusClient.CommunityActionTemplates.GetByID(invalidID)
	require.Error(t, err)
	require.Nil(t, communityActionTemplate)

	invalidName := internal.GetRandomName()
	communityActionTemplate, err = octopusClient.CommunityActionTemplates.GetByName(invalidName)
	require.Error(t, err)
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

func TestCommunityActionTemplateServiceGet(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	query := actions.CommunityActionTemplatesQuery{
		Skip: 1,
		Take: 20,
	}

	communityActionTemplates, err := octopusClient.CommunityActionTemplates.Get(query)
	require.NoError(t, err)
	require.NotNil(t, communityActionTemplates)
}

func TestCommunityActionTemplateServiceGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	communityActionTemplates, err := octopusClient.CommunityActionTemplates.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, communityActionTemplates)
}

// TODO: fix test
// func TestCommunityActionTemplateServiceGetByIDs(t *testing.T) {
// 	octopusClient := getOctopusClient()
// 	require.NotNil(t, octopusClient)

// 	resources, err := octopusClient.CommunityActionTemplates.GetAll()
// 	require.NoError(t, err)
// 	require.NotNil(t, resources)

// 	ids := []string{}
// 	for _, resource := range resources {
// 		ids = append(ids, resource.GetID())
// 	}

// 	// no need to test if ID collection size is less than 2
// 	if len(ids) < 2 {
// 		return
// 	}

// 	resourceListToCompare, err := octopusClient.CommunityActionTemplates.GetByIDs(ids[0:2])
// 	assert.NoError(t, err)
// 	assert.NotNil(t, resourceListToCompare)
// 	assert.Equal(t, 2, len(resourceListToCompare))
// }

func TestCommunityActionTemplateServiceInstall(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	resource, err := octopusClient.CommunityActionTemplates.Install(actions.CommunityActionTemplate{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = actions.NewCommunityActionTemplate(internal.GetRandomName(), internal.GetRandomName())
	require.NotNil(t, resource)
}
