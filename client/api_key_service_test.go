package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewAPIKeyServiceWithNil(t *testing.T) {
	service := NewAPIKeyService(nil, "")
	assert.Nil(t, service)
}

func TestNewAPIKeyServiceWithEmptyClient(t *testing.T) {
	service := NewAPIKeyService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestAPIKeyServiceGetWithEmptyID(t *testing.T) {
	service := NewAPIKeyService(&sling.Sling{}, "")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
