package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewActionTemplateServiceWithNil(t *testing.T) {
	service := NewActionTemplateService(nil)
	assert.Nil(t, service)
}

func TestNewActionTemplateServiceWithEmptyClient(t *testing.T) {
	service := NewActionTemplateService(&sling.Sling{})
	assert.NotNil(t, service)
	assert.Equal(t, service.path, "actiontemplates")
	assert.NotNil(t, service.sling)
}

func TestActionTemplateServiceGetWithEmptyID(t *testing.T) {
	service := NewActionTemplateService(&sling.Sling{})

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "actiontemplates")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
