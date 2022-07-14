package e2e

import (
	"reflect"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/channels"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/releases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualChannels(t *testing.T, expected *channels.Channel, actual *channels.Channel) {
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
	assert.True(t, reflect.DeepEqual(expected.GetLinks(), actual.GetLinks()))

	// channel
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.IsDefault, actual.IsDefault)
	assert.Equal(t, expected.LifecycleID, actual.LifecycleID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ProjectID, actual.ProjectID)
	assert.True(t, reflect.DeepEqual(expected.Rules, actual.Rules))
	assert.True(t, reflect.DeepEqual(expected.TenantTags, actual.TenantTags))
}

func CreateTestChannel(t *testing.T, client *client.Client, project *projects.Project) *channels.Channel {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	channel := channels.NewChannel(name, project.GetID())
	require.NotNil(t, channel)
	require.NoError(t, channel.Validate())

	createdChannel, err := client.Channels.Add(channel)
	require.NoError(t, err)
	require.NotNil(t, createdChannel)
	require.NotEmpty(t, createdChannel.GetID())

	// verify the add operation was successful
	channelToCompare, err := client.Channels.GetByID(createdChannel.GetID())
	require.NoError(t, err)
	require.NotNil(t, channelToCompare)
	AssertEqualChannels(t, createdChannel, channelToCompare)

	return createdChannel
}

func DeleteTestChannel(t *testing.T, client *client.Client, channel *channels.Channel) {
	require.NotNil(t, channel)

	if channel.IsDefault {
		return
	}

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Channels.DeleteByID(channel.GetID())
	require.NoError(t, err)

	// verify the delete operation was successful
	deletedChannel, err := client.Channels.GetByID(channel.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedChannel)
}

func TestChannelServiceAddGetDelete(t *testing.T) {
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

	channelToCompare, err := client.Channels.GetByID(channel.GetID())
	require.NotNil(t, channelToCompare)
	require.NoError(t, err)
	AssertEqualChannels(t, channel, channelToCompare)
}

func TestChannelServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	channels, err := client.Channels.GetAll()
	require.NoError(t, err)
	require.NotNil(t, channels)

	for _, channel := range channels {
		if !channel.IsDefault {
			defer DeleteTestChannel(t, client, channel)
		}
	}
}

func TestGetAllChannels(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	channels, err := client.Channels.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, channels)

	for _, channel := range channels {
		assert.NoError(t, err)
		assert.NotEmpty(t, channel)
	}
}

func TestGetProject(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	channels, err := client.Channels.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, channels)

	for _, channel := range channels {
		project, err := client.Projects.GetProject(channel)

		if err != nil {
			assert.Nil(t, project)
		} else {
			assert.NotNil(t, project)
		}
	}
}

func TestChannelServiceGetReleases(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	channels, err := client.Channels.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, channels)

	for _, channel := range channels {
		releasesToTest, err := client.Releases.GetReleases(channel)
		require.NotNil(t, releasesToTest)
		require.NoError(t, err)

		for _, release := range releasesToTest.Items {
			releaseToCompare, err := client.Releases.GetByID(release.GetID())
			require.NotNil(t, releaseToCompare)
			require.NoError(t, err)
			AssertEqualReleases(t, release, releaseToCompare)
		}

		releaseQuery := &releases.ReleaseQuery{
			SearchByVersion: "0.0.1",
			Skip:            1,
			Take:            1,
		}

		releasesToTest, err = client.Releases.GetReleases(channel, releaseQuery)
		require.NotNil(t, releasesToTest)
		require.NoError(t, err)

		for _, release := range releasesToTest.Items {
			releaseToCompare, err := client.Releases.GetByID(release.GetID())
			require.NotNil(t, releaseToCompare)
			require.NoError(t, err)
			AssertEqualReleases(t, release, releaseToCompare)
		}
	}
}

func TestChannelServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := internal.GetRandomName()
	channel, err := client.Channels.GetByID(id)
	assert.NotNil(t, err)
	assert.Nil(t, channel)

	apiError := err.(*core.APIError)
	assert.Equal(t, 404, apiError.StatusCode)

	allChannels, err := client.Channels.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, allChannels)

	for _, channel := range allChannels {
		query := channels.Query{
			IDs:  []string{channel.GetID()},
			Take: 1,
		}
		channelsToVerify, err := client.Channels.Get(query)
		require.NoError(t, err)
		require.NotNil(t, channelsToVerify)
		AssertEqualChannels(t, channel, channelsToVerify.Items[0])

		channelToCompare, err := client.Channels.GetByID(channel.GetID())
		assert.NoError(t, err)
		AssertEqualChannels(t, channel, channelToCompare)
	}
}

func TestChannelServiceGetByPartialName(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	query := channels.Query{PartialName: internal.GetRandomName()}
	channelsToTest, err := client.Channels.Get(query)
	require.NoError(t, err)
	require.NotNil(t, channelsToTest)
	require.True(t, len(channelsToTest.Items) == 0)

	channelsSlice, err := client.Channels.GetAll()
	require.NoError(t, err)
	require.NotNil(t, channelsToTest)

	for _, channel := range channelsSlice {
		query := channels.Query{PartialName: channel.Name}
		channelsToCompare, err := client.Channels.Get(query)
		require.NoError(t, err)
		require.NotNil(t, channelsToCompare)
		require.NotNil(t, channelsToCompare.Items)
		require.True(t, len(channelsToCompare.Items) > 0)
	}
}
