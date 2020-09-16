package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewChannelServiceWithNil(t *testing.T) {
	service := NewChannelService(nil, "")
	assert.Nil(t, service)
}

func TestChannelServiceWithEmptyClient(t *testing.T) {
	service := NewChannelService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestChannelServiceGetWithEmptyID(t *testing.T) {
	service := NewChannelService(&sling.Sling{}, "")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
