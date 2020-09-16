package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectGroupServiceWithNil(t *testing.T) {
	service := NewProjectGroupService(nil, "")
	assert.Nil(t, service)
}

func TestProjectGroupServiceWithEmptyClient(t *testing.T) {
	service := NewProjectGroupService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestProjectGroupServiceGetWithEmptyID(t *testing.T) {
	service := NewProjectGroupService(&sling.Sling{}, "")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
