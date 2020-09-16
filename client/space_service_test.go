package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewSpaceServiceWithNil(t *testing.T) {
	service := NewSpaceService(nil, "")
	assert.Nil(t, service)
}

func TestSpaceServiceWithEmptyClient(t *testing.T) {
	service := NewSpaceService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestSpaceServiceGetWithEmptyID(t *testing.T) {
	service := NewSpaceService(&sling.Sling{}, "")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
