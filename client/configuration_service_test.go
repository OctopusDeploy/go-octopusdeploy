package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigurationServiceWithNil(t *testing.T) {
	service := NewConfigurationService(nil, "")
	assert.Nil(t, service)
}

func TestConfigurationServiceWithEmptyClient(t *testing.T) {
	service := NewConfigurationService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestConfigurationServiceGetWithEmptyID(t *testing.T) {
	service := NewConfigurationService(&sling.Sling{}, "")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
