package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	channelName        = "Channel Name"
	channelDescription = "This is the channel description."
	channelProjectID   = "ProjectID-123"
)

func TestEmptyChannel(t *testing.T) {
	channel := &Channel{}
	assert.Error(t, channel.Validate())
}

func TestChannelWithName(t *testing.T) {
	channel := &Channel{Name: channelName}
	assert.Error(t, channel.Validate())
}

func TestNewChannelWithEmptyName(t *testing.T) {
	channel := NewChannel(emptyString, channelDescription, channelProjectID)
	require.Error(t, channel.Validate())

	channel = NewChannel(whitespaceString, channelDescription, channelProjectID)
	require.Error(t, channel.Validate())
}
