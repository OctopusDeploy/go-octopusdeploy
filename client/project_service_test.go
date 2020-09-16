package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectServiceWithNil(t *testing.T) {
	service := NewProjectService(nil, "")
	assert.Nil(t, service)
}

func TestProjectServiceWithEmptyClient(t *testing.T) {
	service := NewProjectService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestProjectServiceGetWithEmptyID(t *testing.T) {
	service := NewProjectService(&sling.Sling{}, "")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
