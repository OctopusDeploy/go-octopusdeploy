package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewTenantServiceWithNil(t *testing.T) {
	service := NewTenantService(nil)
	assert.Nil(t, service)
}

func TestTenantServiceWithEmptyClient(t *testing.T) {
	service := NewTenantService(&sling.Sling{})
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestTenantServiceGetWithEmptyID(t *testing.T) {
	service := NewTenantService(&sling.Sling{})

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
