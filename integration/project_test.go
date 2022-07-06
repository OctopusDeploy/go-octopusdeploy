package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateActionTemplateParameter() octopusdeploy.ActionTemplateParameter {
	actionTemplateParameter := octopusdeploy.NewActionTemplateParameter()
	return *actionTemplateParameter
}

func TestAddNilProject(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	project, err := octopusClient.Projects.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, project)
}

func TestGetSummary(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	projects, err := octopusClient.Projects.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, projects)

	for _, project := range projects {
		summary, err := octopusClient.Projects.GetSummary(project)

		assert.NoError(t, err)
		assert.NotNil(t, summary)
	}
}

func TestGetReleasesForProject(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	projects, err := octopusClient.Projects.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, projects)

	for _, project := range projects {
		releases, err := octopusClient.Projects.GetReleases(project)
		assert.NoError(t, err)
		assert.NotNil(t, releases)
	}
}

func TestGetChannelsForProject(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	projects, err := octopusClient.Projects.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, projects)

	for _, project := range projects {
		channels, err := octopusClient.Projects.GetChannels(project)
		assert.NoError(t, err)
		assert.NotNil(t, channels)
	}
}

// TODO: fix test
// func TestProjectGet(t *testing.T) {
// 	client, err := octopusdeploy.GetFakeOctopusClient(t, "/api/projects/Projects-663", http.StatusOK, getProjectResponseJSON)
// 	require.NoError(t, err)
// 	require.NotNil(t, client)

// 	project, err := client.Projects.GetByID("Projects-663")
// 	require.NoError(t, err)

// 	assert.Equal(t, "Canary .NET Core 2.0", project.Name)
// }

const getProjectResponseJSON = `
{
  "Id": "Projects-663",
  "VariableSetId": "variableset-Projects-663",
  "DeploymentProcessId": "deploymentprocess-Projects-663",
  "DiscreteChannelRelease": false,
  "IncludedLibraryVariableSetIds": [
    "LibraryVariableSets-124",
    "LibraryVariableSets-41",
    "LibraryVariableSets-21",
    "LibraryVariableSets-81"
  ],
  "DefaultToSkipIfAlreadyInstalled": false,
  "TenantedDeploymentMode": "Untenanted",
  "DefaultGuidedFailureMode": "EnvironmentDefault",
  "VersioningStrategy": {
    "DonorPackageStepId": null,
    "Template": "#{Octopus.Version.LastMajor}.#{Octopus.Version.LastMinor}.#{Octopus.Version.NextPatch}"
  },
  "ReleaseCreationStrategy": {
    "ReleaseCreationPackageStepId": "",
    "ChannelId": null
  },
  "Templates": [],
  "AutoDeployReleaseOverrides": [],
  "Name": "Canary .NET Core 2.0",
  "Slug": "canary-net-core-2-0",
  "Description": "Deployment pipeline for the canary .NET Core 2.0 skeleton project",
  "IsDisabled": false,
  "ProjectGroupId": "ProjectGroups-261",
  "LifecycleId": "Lifecycles-21",
  "AutoCreateRelease": false,
  "ProjectConnectivityPolicy": {
    "SkipMachineBehavior": "SkipUnavailableMachines",
    "TargetRoles": [],
    "AllowDeploymentsToNoTargets": false
  },
  "Links": {
    "Self": "/api/projects/Projects-663",
    "Releases": "/api/projects/Projects-663/releases{/version}{?skip,take,searchByVersion}",
    "Channels": "/api/projects/Projects-663/channels{?skip,take,partialName}",
    "Triggers": "/api/projects/Projects-663/triggers{?skip,take,partialName}",
    "OrderChannels": "/api/projects/Projects-663/channels/order",
    "Variables": "/api/variables/variableset-Projects-663",
    "Progression": "/api/progression/Projects-663{?aggregate}",
    "DeploymentProcess": "/api/deploymentprocesses/deploymentprocess-Projects-663",
    "Web": "/app#/projects/Projects-663",
    "Logo": "/api/projects/Projects-663/logo?cb=2018.3.6"
  }
}`
