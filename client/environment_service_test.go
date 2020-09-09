package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewEnvironmentServiceWithNil(t *testing.T) {
	service := NewEnvironmentService(nil)
	assert.Nil(t, service)
}

func TestEnvironmentServiceWithEmptyClient(t *testing.T) {
	service := NewEnvironmentService(&sling.Sling{})
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestEnvironmentServiceGetWithEmptyID(t *testing.T) {
	service := NewEnvironmentService(&sling.Sling{})

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
