package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewInterruptionsServiceWithNil(t *testing.T) {
	service := NewInterruptionsService(nil)
	assert.Nil(t, service)
}

func TestInterruptionsServiceWithEmptyClient(t *testing.T) {
	service := NewInterruptionsService(&sling.Sling{})
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestInterruptionsServiceGetWithEmptyID(t *testing.T) {
	service := NewInterruptionsService(&sling.Sling{})

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
