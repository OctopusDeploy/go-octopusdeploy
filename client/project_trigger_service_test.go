package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewProjectTriggerServiceWithNil(t *testing.T) {
	service := NewProjectTriggerService(nil, "")
	assert.Nil(t, service)
}

func TestProjectTriggerServiceWithEmptyClient(t *testing.T) {
	service := NewProjectTriggerService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestProjectTriggerServiceGetWithEmptyID(t *testing.T) {
	service := NewProjectTriggerService(&sling.Sling{}, "")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
