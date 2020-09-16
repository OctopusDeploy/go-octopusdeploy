package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewVariableServiceWithNil(t *testing.T) {
	service := NewVariableService(nil, "")
	assert.Nil(t, service)
}

func TestVariableServiceWithEmptyClient(t *testing.T) {
	service := NewVariableService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestVariableServiceGetAllWithEmptyID(t *testing.T) {
	service := NewVariableService(&sling.Sling{}, "")

	resource, err := service.GetAll("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.GetAll(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
