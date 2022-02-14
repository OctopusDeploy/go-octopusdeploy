package resources

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptyChannel(t *testing.T) {
	channel := &Channel{}
	assert.Error(t, channel.Validate())
}

func TestChannelWithName(t *testing.T) {
	name := octopusdeploy.getRandomName()
	channel := &Channel{Name: name}
	assert.Error(t, channel.Validate())
}

func TestNewChannelWithEmptyName(t *testing.T) {
	projectID := octopusdeploy.getRandomName()

	channel := NewChannel(emptyString, projectID)
	require.Error(t, channel.Validate())

	channel = NewChannel(whitespaceString, projectID)
	require.Error(t, channel.Validate())
}
