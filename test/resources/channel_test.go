package resources

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/channels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptyChannel(t *testing.T) {
	channel := &channels.Channel{}
	assert.Error(t, channel.Validate())
}

func TestChannelWithName(t *testing.T) {
	name := internal.GetRandomName()
	channel := &channels.Channel{Name: name}
	assert.Error(t, channel.Validate())
}

func TestNewChannelWithEmptyName(t *testing.T) {
	projectID := internal.GetRandomName()

	channel := channels.NewChannel("", projectID)
	require.Error(t, channel.Validate())

	channel = channels.NewChannel(" ", projectID)
	require.Error(t, channel.Validate())
}
