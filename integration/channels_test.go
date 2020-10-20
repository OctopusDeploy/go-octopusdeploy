package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	channelName        = getRandomName()
	channelDescription = getRandomName()
)

func TestAddNilChannel(t *testing.T) {
	octopusClient := getOctopusClient()

	channel, err := octopusClient.Channels.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, channel)
}

func TestAddAndDeleteAndGetValidChannel(t *testing.T) {
	octopusClient := getOctopusClient()

	projects, err := octopusClient.Projects.GetAll()
	require.NoError(t, err)
	require.NotNil(t, projects)

	for _, project := range projects {
		channel := octopusdeploy.NewChannel(channelName, channelDescription, project.ID)
		require.NotNil(t, channel)

		createdChannel, err := octopusClient.Channels.Add(channel)
		require.NoError(t, err)
		require.NotNil(t, createdChannel)

		err = octopusClient.Channels.DeleteByID(createdChannel.ID)
		require.NoError(t, err)
		require.NotNil(t, createdChannel)

		deletedChannel, err := octopusClient.Channels.GetByID(createdChannel.ID)
		require.Error(t, err)
		require.Nil(t, deletedChannel)
	}
}

func TestAddAndDeleteValidChannel(t *testing.T) {
	octopusClient := getOctopusClient()

	projects, err := octopusClient.Projects.GetAll()
	require.NoError(t, err)
	require.NotNil(t, projects)

	for _, project := range projects {
		channel := octopusdeploy.NewChannel(channelName, channelDescription, project.ID)
		channel.IsDefault = true
		require.NotNil(t, channel)

		createdChannel, err := octopusClient.Channels.Add(channel)
		require.NoError(t, err)
		require.NotNil(t, createdChannel)

		if !channel.IsDefault {
			err = octopusClient.Channels.DeleteByID(createdChannel.ID)
			require.NoError(t, err)
			require.NotNil(t, createdChannel)

			deletedChannel, err := octopusClient.Channels.GetByID(createdChannel.ID)
			require.Error(t, err)
			require.Nil(t, deletedChannel)
		}
	}
}

func TestGetReleasesForChannel(t *testing.T) {
	octopusClient := getOctopusClient()

	channels, err := octopusClient.Channels.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, channels)

	if len(channels) == 0 {
		return
	}

	for _, channel := range channels {
		releases, err := octopusClient.Channels.GetReleases(channel)

		assert.NoError(t, err)
		assert.NotNil(t, releases)
	}
}

func TestGetProject(t *testing.T) {
	octopusClient := getOctopusClient()

	channels, err := octopusClient.Channels.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, channels)

	if len(channels) == 0 {
		return
	}

	for _, channel := range channels {
		project, err := octopusClient.Channels.GetProject(channel)

		if err != nil {
			assert.Nil(t, project)
		} else {
			assert.NotNil(t, project)
		}
	}
}

func TestGetAllChannels(t *testing.T) {
	octopusClient := getOctopusClient()

	channels, err := octopusClient.Channels.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, channels)

	if len(channels) == 0 {
		return
	}

	for _, channel := range channels {
		assert.NoError(t, err)
		assert.NotEmpty(t, channel)
	}
}
