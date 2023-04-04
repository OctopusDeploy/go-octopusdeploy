package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/actiontemplates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateActionTemplateParameter() actiontemplates.ActionTemplateParameter {
	actionTemplateParameter := actiontemplates.NewActionTemplateParameter()
	return *actionTemplateParameter
}

func TestAddNilProject(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	project, err := client.Projects.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, project)
}

func TestGetSummary(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	projects, err := client.Projects.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, projects)

	for _, project := range projects {
		summary, err := client.Projects.GetSummary(project)

		assert.NoError(t, err)
		assert.NotNil(t, summary)
	}
}

func TestGetReleasesForProject(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	projects, err := client.Projects.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, projects)

	for _, project := range projects {
		releases, err := client.Projects.GetReleases(project)
		assert.NoError(t, err)
		assert.NotNil(t, releases)
	}
}

func TestGetChannelsForProject(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	projects, err := client.Projects.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, projects)

	for _, project := range projects {
		channels, err := client.Projects.GetChannels(project)
		assert.NoError(t, err)
		assert.NotNil(t, channels)
	}
}
