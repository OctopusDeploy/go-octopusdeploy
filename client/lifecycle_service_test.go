package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewLifecycleServiceWithNil(t *testing.T) {
	service := NewLifecycleService(nil, "")
	assert.Nil(t, service)
}

func TestLifecycleServiceWithEmptyClient(t *testing.T) {
	service := NewLifecycleService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestLifecycleServiceGetWithEmptyID(t *testing.T) {
	service := NewLifecycleService(&sling.Sling{}, "")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
