package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewLibraryVariableSetServiceWithNil(t *testing.T) {
	service := NewLibraryVariableSetService(nil)
	assert.Nil(t, service)
}

func TestLibraryVariableSetServiceWithEmptyClient(t *testing.T) {
	service := NewLibraryVariableSetService(&sling.Sling{})
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestLibraryVariableSetServiceGetWithEmptyID(t *testing.T) {
	service := NewLibraryVariableSetService(&sling.Sling{})

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
