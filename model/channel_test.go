package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	channelName        = "Channel Name"
	channelDescription = "This is the channel description."
	channelProjectID   = "ProjectID-123"
)

func TestEmptyChannel(t *testing.T) {
	channel := &Channel{}

	assert.NotNil(t, channel)
	assert.Error(t, channel.Validate())
}

func TestChannelWithName(t *testing.T) {
	channel := &Channel{Name: channelName}

	assert.NotNil(t, channel)
	assert.Error(t, channel.Validate())
}

func TestNewChannelWithEmptyName(t *testing.T) {
	channel, err := NewChannel(emptyString, channelDescription, channelProjectID)

	assert.Error(t, err)
	assert.Nil(t, channel)

	channel, err = NewChannel(whitespaceString, channelDescription, channelProjectID)

	assert.Error(t, err)
	assert.Nil(t, channel)
}
