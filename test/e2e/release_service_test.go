package e2e

import (
	"net/http"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/channels"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/configuration"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/releases"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualReleases(t *testing.T, expected *releases.Release, actual *releases.Release) {
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

	// Release
	assert.Equal(t, expected.Assembled, actual.Assembled)
	assert.Equal(t, expected.BuildInformation, actual.BuildInformation)
	assert.Equal(t, expected.ChannelID, actual.ChannelID)
	assert.Equal(t, expected.IgnoreChannelRules, actual.IgnoreChannelRules)
	assert.Equal(t, expected.LibraryVariableSetSnapshotIDs, actual.LibraryVariableSetSnapshotIDs)
	assert.Equal(t, expected.ProjectDeploymentProcessSnapshotID, actual.ProjectDeploymentProcessSnapshotID)
	assert.Equal(t, expected.ProjectID, actual.ProjectID)
	assert.Equal(t, expected.ProjectVariableSetSnapshotID, actual.ProjectVariableSetSnapshotID)
	assert.Equal(t, expected.ReleaseNotes, actual.ReleaseNotes)
	assert.Equal(t, expected.SelectedPackages, actual.SelectedPackages)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
	assert.Equal(t, expected.Version, actual.Version)
}

func CreateTestRelease(t *testing.T, client *client.Client, channel *channels.Channel, project *projects.Project) *releases.Release {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	version := "0.0.1"

	release := releases.NewRelease(channel.GetID(), project.GetID(), version)

	require.NotNil(t, release)
	require.NoError(t, release.Validate())

	createdRelease, err := client.Releases.Add(release)
	require.NoError(t, err)
	require.NotNil(t, createdRelease)
	require.NotEmpty(t, createdRelease.GetID())

	// verify the add operation was successful
	releaseToCompare, err := client.Releases.GetByID(createdRelease.GetID())
	require.NoError(t, err)
	require.NotNil(t, releaseToCompare)
	AssertEqualReleases(t, createdRelease, releaseToCompare)

	return createdRelease
}

func DeleteTestRelease(t *testing.T, client *client.Client, release *releases.Release) {
	require.NotNil(t, release)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Releases.DeleteByID(release.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedRelease, err := client.Releases.GetByID(release.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedRelease)
}

func TestReleaseServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	channel := CreateTestChannel(t, client, project)
	require.NotNil(t, channel)
	defer DeleteTestChannel(t, client, channel)

	release := CreateTestRelease(t, client, channel, project)
	require.NotNil(t, release)
	defer DeleteTestRelease(t, client, release)

	releaseToCompare, err := client.Releases.GetByID(release.GetID())
	require.NoError(t, err)
	require.NotNil(t, releaseToCompare)
	AssertEqualReleases(t, release, releaseToCompare)
}

func TestReleaseServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	query := releases.ReleasesQuery{
		Take: 50,
	}

	releases, err := client.Releases.Get(query)
	require.NoError(t, err)
	require.NotNil(t, releases)

	for _, release := range releases.Items {
		defer DeleteTestRelease(t, client, release)
	}
}

func TestReleaseServiceCreateV1(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	channel := CreateTestChannel(t, client, project)
	require.NotNil(t, channel)
	defer DeleteTestChannel(t, client, channel)

	releaseCreate := releases.NewCreateReleaseCommandV1(space.Name, project.Name)
	createReleaseResponse, err := releases.CreateReleaseV1(client, releaseCreate)

	// if HTTP 404 response then Executions API is unavailable
	if err == nil {
		require.NotNil(t, createReleaseResponse)
		require.NotNil(t, createReleaseResponse.ReleaseID)
		require.NotNil(t, createReleaseResponse.ReleaseVersion)
	}
}

func TestReleaseServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := internal.GetRandomName()
	release, err := client.Releases.GetByID(id)
	require.Error(t, err)
	assert.Nil(t, release)

	query := releases.ReleasesQuery{
		Take: 50,
	}

	releases, err := client.Releases.Get(query)
	assert.NoError(t, err)
	assert.NotNil(t, releases)

	for _, release := range releases.Items {
		releaseToCompare, err := client.Releases.GetByID(release.GetID())
		assert.NoError(t, err)
		AssertEqualReleases(t, release, releaseToCompare)
	}

}

func TestReleaseServiceSnapshotVariablesByName(t *testing.T) {

	//Arrange
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	toggle, err := configuration.Get(octopusClient, &configuration.FeatureToggleConfigurationQuery{
		Name: "partial-updates-on-variables",
	})
	if err != nil {
		t.Skip("Could not get feature toggle configuration")
	} else if len(toggle.FeatureToggles) == 0 {
		t.Skip("PartialUpdatesOnVariables feature toggle is not present")
	} else if !toggle.FeatureToggles[0].IsEnabled {
		t.Skip("PartialUpdatesOnVariables feature toggle is not enabled")
	}

	space := GetDefaultSpace(t, octopusClient)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, octopusClient)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, octopusClient, lifecycle)

	projectGroup := CreateTestProjectGroup(t, octopusClient)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, octopusClient, projectGroup)

	project := CreateTestProject(t, octopusClient, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, octopusClient, project)

	variable := CreateTestVariable(t, project.ID, internal.GetRandomName())
	require.NotNil(t, variable)

	variable.Value = "oldValue"
	_, err = variables.UpdateSingle(octopusClient, space.ID, project.ID, variable)
	require.NoError(t, err)

	channel := CreateTestChannel(t, octopusClient, project)
	require.NotNil(t, channel)
	defer DeleteTestChannel(t, octopusClient, channel)

	release := CreateTestRelease(t, octopusClient, channel, project)
	require.NotNil(t, release)
	defer DeleteTestRelease(t, octopusClient, release)

	oldProjectSnapshotId := release.ProjectVariableSetSnapshotID

	// Act
	variable.Value = "newValue"
	_, err = variables.UpdateSingle(octopusClient, space.ID, project.ID, variable)
	require.NoError(t, err)

	variableIdentifier := core.VariableIdentifier{Name: variable.Name, OwnerID: project.ID}
	variableIdentifiers := []core.VariableIdentifier{variableIdentifier}

	updatedRelease, err := releases.SnapshotVariablesByName(octopusClient, release, variableIdentifiers)
	require.NoError(t, err)
	require.NotNil(t, updatedRelease)

	// Assert
	assert.NotEqual(t, oldProjectSnapshotId, updatedRelease.ProjectVariableSetSnapshotID)
}

func TestReleaseServiceSnapshotVariablesByNameConcurrency(t *testing.T) {

	//Arrange
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	toggle, err := configuration.Get(octopusClient, &configuration.FeatureToggleConfigurationQuery{
		Name: "partial-updates-on-variables",
	})
	if err != nil {
		t.Skip("Could not get feature toggle configuration")
	} else if len(toggle.FeatureToggles) == 0 {
		t.Skip("PartialUpdatesOnVariables feature toggle is not present")
	} else if !toggle.FeatureToggles[0].IsEnabled {
		t.Skip("PartialUpdatesOnVariables feature toggle is not enabled")
	}

	space := GetDefaultSpace(t, octopusClient)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, octopusClient)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, octopusClient, lifecycle)

	projectGroup := CreateTestProjectGroup(t, octopusClient)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, octopusClient, projectGroup)

	project := CreateTestProject(t, octopusClient, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, octopusClient, project)

	variableA := CreateTestVariable(t, project.ID, internal.GetRandomName())
	require.NotNil(t, variableA)
	variableB := CreateTestVariable(t, project.ID, internal.GetRandomName())
	require.NotNil(t, variableB)

	channel := CreateTestChannel(t, octopusClient, project)
	require.NotNil(t, channel)
	defer DeleteTestChannel(t, octopusClient, channel)

	release := CreateTestRelease(t, octopusClient, channel, project)
	require.NotNil(t, release)
	defer DeleteTestRelease(t, octopusClient, release)

	staleToken := release.VariableSnapshotConcurrencyToken

	variableA.Value = "updatedA"
	_, err = variables.UpdateSingle(octopusClient, space.ID, project.ID, variableA)
	require.NoError(t, err)

	// Act: refresh using the token captured before the release was fetched again advances the snapshot
	firstUpdate, err := releases.SnapshotVariablesByName(octopusClient, release, []core.VariableIdentifier{
		{Name: variableA.Name, OwnerID: project.ID},
	}, staleToken)
	require.NoError(t, err)
	require.NotNil(t, firstUpdate)

	// Assert: reusing the now-stale token is rejected with a conflict
	variableB.Value = "updatedB"
	_, err = variables.UpdateSingle(octopusClient, space.ID, project.ID, variableB)
	require.NoError(t, err)

	_, err = releases.SnapshotVariablesByName(octopusClient, release, []core.VariableIdentifier{
		{Name: variableB.Name, OwnerID: project.ID},
	}, staleToken)
	require.Error(t, err)
	apiError, ok := err.(*core.APIError)
	require.True(t, ok)
	assert.Equal(t, http.StatusConflict, apiError.StatusCode)

	// Assert: supplying the current token succeeds
	secondUpdate, err := releases.SnapshotVariablesByName(octopusClient, firstUpdate, []core.VariableIdentifier{
		{Name: variableB.Name, OwnerID: project.ID},
	}, firstUpdate.VariableSnapshotConcurrencyToken)
	require.NoError(t, err)
	require.NotNil(t, secondUpdate)
}
