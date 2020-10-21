package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChannelServiceAdd(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	emptyChannel := &octopusdeploy.Channel{}
	channel, err := octopusClient.Channels.Add(emptyChannel)
	assert.Equal(t, createValidationFailureError("Add", emptyChannel.Validate()), err)
	assert.Nil(t, channel)
}

func TestChannelServiceGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	channels, err := octopusClient.Channels.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, channels)

	for _, channel := range channels {
		assert.NotNil(t, channel)
		assert.NotEmpty(t, channel.GetID())
	}
}

func TestChannelServiceGetByID(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	id := getRandomName()
	channel, err := octopusClient.Channels.GetByID(id)
	assert.Equal(t, createResourceNotFoundError("ChannelService", "ID", id), err)
	assert.Nil(t, channel)

	channels, err := octopusClient.Channels.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, channels)

	for _, channel := range channels {
		channelToCompare, err := octopusClient.Channels.GetByID(channel.GetID())
		assert.NoError(t, err)
		assert.EqualValues(t, channel, channelToCompare)
	}
}

func TestChannelServiceGetByPartialName(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	name := getRandomName()
	channels, err := octopusClient.Channels.GetByPartialName(name)
	assert.NoError(t, err)
	assert.NotNil(t, channels)
	assert.Len(t, channels, 0)

	channels, err = octopusClient.Channels.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, channels)

	for _, channel := range channels {
		channelToCompare, err := octopusClient.Channels.GetByPartialName(channel.Name)
		assert.NoError(t, err)
		assert.EqualValues(t, channel, channelToCompare[0])
	}
}
