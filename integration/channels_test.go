package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

var (
	channelName        = getRandomName()
	channelDescription = getRandomName()
)

func init() {
	if octopusClient == nil {
		octopusClient = initTest()
	}
}

func TestAddNilChannel(t *testing.T) {
	channel, err := octopusClient.Channels.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, channel)
}

func TestAddAndDeleteAndGetValidChannel(t *testing.T) {
	projects, err := octopusClient.Projects.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, projects)

	if err != nil {
		return
	}

	for _, project := range projects {
		channel, err := model.NewChannel(channelName, channelDescription, project.ID)

		assert.NoError(t, err)
		assert.NotNil(t, channel)

		if err != nil {
			return
		}

		createdChannel, err := octopusClient.Channels.Add(channel)

		assert.NoError(t, err)
		assert.NotNil(t, createdChannel)

		err = octopusClient.Channels.Delete(createdChannel.ID)

		assert.NoError(t, err)
		assert.NotNil(t, createdChannel)

		deletedChannel, err := octopusClient.Channels.Get(createdChannel.ID)

		assert.Error(t, err)
		assert.Nil(t, deletedChannel)
	}
}

func TestAddAndDeleteValidChannel(t *testing.T) {
	projects, err := octopusClient.Projects.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, projects)

	if err != nil {
		return
	}

	for _, project := range projects {
		channel, err := model.NewChannel(channelName, channelDescription, project.ID)
		channel.IsDefault = true

		assert.NoError(t, err)
		assert.NotNil(t, channel)

		if err != nil {
			return
		}

		createdChannel, err := octopusClient.Channels.Add(channel)

		assert.NoError(t, err)
		assert.NotNil(t, createdChannel)

		if !channel.IsDefault {
			err = octopusClient.Channels.Delete(createdChannel.ID)

			assert.NoError(t, err)
			assert.NotNil(t, createdChannel)

			deletedChannel, err := octopusClient.Channels.Get(createdChannel.ID)

			assert.Error(t, err)
			assert.Nil(t, deletedChannel)
		}
	}
}

func TestGetReleasesForChannel(t *testing.T) {
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
	channels, err := octopusClient.Channels.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, channels)

	if len(channels) == 0 {
		return
	}

	for _, channel := range channels {
		project, err := octopusClient.Channels.GetProject(channel)

		assert.NoError(t, err)
		assert.NotNil(t, project)
	}
}

func TestGetAllChannels(t *testing.T) {
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
