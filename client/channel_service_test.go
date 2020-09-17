package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

const (
	TestChannelServiceURITemplate = "channel-service"
)

func TestNewChannelService(t *testing.T) {
	service := NewChannelService(nil, "")
	assert.Nil(t, service)
	createChannelService(t)
}

func TestChannelServiceGetWithEmptyID(t *testing.T) {
	service := createChannelService(t)

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "id"))
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "id"))
	assert.Nil(t, resource)
}

func TestChannelServiceDeleteWithEmptyID(t *testing.T) {
	service := createChannelService(t)

	err := service.Delete("")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Delete", "id"))

	err = service.Delete(" ")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Delete", "id"))
}

func createChannelService(t *testing.T) *ChannelService {
	service := NewChannelService(&sling.Sling{}, TestChannelServiceURITemplate)

	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
	assert.Equal(t, service.path, TestChannelServiceURITemplate)
	assert.Equal(t, service.name, "ChannelService")

	return service
}
