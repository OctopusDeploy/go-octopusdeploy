package integration

import (
	"reflect"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func DeleteTestChannel(t *testing.T, client *octopusdeploy.Client, channel *octopusdeploy.Channel) {
	require.NotNil(t, channel)

	if channel.IsDefault {
		return
	}

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Channels.DeleteByID(channel.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	channels, err := client.Channels.GetByID(channel.GetID())
	assert.Error(t, err)
	assert.Nil(t, channels)
}

func IsEqualChannels(t *testing.T, expected *octopusdeploy.Channel, actual *octopusdeploy.Channel) {
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

func TestChannelServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	channels, err := client.Channels.GetAll()
	require.NoError(t, err)
	require.NotNil(t, channels)

	for _, channel := range channels {
		defer DeleteTestChannel(t, client, channel)
	}
}

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
		IsEqualChannels(t, channel, channelToCompare)
	}
}

func TestChannelServiceGetByPartialName(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	query := octopusdeploy.ChannelsQuery{PartialName: getRandomName()}
	channels, err := octopusClient.Channels.Get(query)
	assert.NoError(t, err)
	assert.NotNil(t, channels)
	assert.Len(t, channels.Items, 0)

	allChannels, err := octopusClient.Channels.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, allChannels)

	for _, channel := range allChannels {
		query := octopusdeploy.ChannelsQuery{PartialName: channel.Name}
		channelsToCompare, err := octopusClient.Channels.Get(query)
		assert.NoError(t, err)
		assert.NotNil(t, channelsToCompare)
		assert.NotNil(t, channelsToCompare.Items)
		assert.True(t, len(channelsToCompare.Items) > 0)
	}
}
