package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/channels"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/releases"
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
