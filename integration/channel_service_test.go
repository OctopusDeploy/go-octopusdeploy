package integration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/service"
	service2 "github.com/OctopusDeploy/go-octopusdeploy/service"
	"reflect"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualChannels(t *testing.T, expected *service.Channel, actual *service.Channel) {
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

func CreateTestChannel(t *testing.T, client *octopusdeploy.client, project *service.Project) *service.Channel {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()

	channel := service.NewChannel(name, project.GetID())
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

func DeleteTestChannel(t *testing.T, client *octopusdeploy.client, channel *service.Channel) {
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
		project, err := client.Channels.GetProject(channel)

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
		releases, err := client.Channels.GetReleases(channel)
		require.NotNil(t, releases)
		require.NoError(t, err)

		for _, release := range releases.Items {
			releaseToCompare, err := client.Releases.GetByID(release.GetID())
			require.NotNil(t, releaseToCompare)
			require.NoError(t, err)
			AssertEqualReleases(t, release, releaseToCompare)
		}

		releaseQuery := &service2.ReleaseQuery{
			SearchByVersion: "0.0.1",
			Skip:            1,
			Take:            1,
		}

		releases, err = client.Channels.GetReleases(channel, releaseQuery)
		require.NotNil(t, releases)
		require.NoError(t, err)

		for _, release := range releases.Items {
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

	id := getRandomName()
	channel, err := client.Channels.GetByID(id)
	assert.NotNil(t, err)
	assert.Nil(t, channel)

	apiError := err.(*service2.APIError)
	assert.Equal(t, 404, apiError.StatusCode)

	channels, err := client.Channels.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, channels)

	for _, channel := range channels {
		query := service2.ChannelsQuery{
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

	query := service2.ChannelsQuery{PartialName: getRandomName()}
	channels, err := client.Channels.Get(query)
	require.NoError(t, err)
	require.NotNil(t, channels)
	require.True(t, len(channels.Items) == 0)

	allChannels, err := client.Channels.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allChannels)

	for _, channel := range allChannels {
		query := service2.ChannelsQuery{PartialName: channel.Name}
		channelsToCompare, err := client.Channels.Get(query)
		require.NoError(t, err)
		require.NotNil(t, channelsToCompare)
		require.NotNil(t, channelsToCompare.Items)
		require.True(t, len(channelsToCompare.Items) > 0)
	}
}
