package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewMachineServiceWithNil(t *testing.T) {
	service := NewMachineService(nil)
	assert.Nil(t, service)
}

func TestMachineServiceWithEmptyClient(t *testing.T) {
	service := NewMachineService(&sling.Sling{})
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestMachineServiceGetWithEmptyID(t *testing.T) {
	service := NewMachineService(&sling.Sling{})

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
