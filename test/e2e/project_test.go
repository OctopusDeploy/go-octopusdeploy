package e2e

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"net/url"
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

func TestConvertProjectToVcs(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	credentials := credentials.NewAnonymous()
	url, err := url.Parse("https://example.com/")
	gps := projects.NewGitPersistenceSettings(".octopus/foobar2", credentials, "master", nil, url)

	space := GetDefaultSpace(t, client)
	lifecycle := CreateTestLifecycle(t, client)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	defer DeleteTestProject(t, client, project)

	client.Projects.ConvertToVcs(project, "Initial Commit", "X", gps)

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
