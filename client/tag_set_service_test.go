package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewTagSetServiceWithNil(t *testing.T) {
	service := NewTagSetService(nil, "")
	assert.Nil(t, service)
}

func TestTagSetServiceWithEmptyClient(t *testing.T) {
	service := NewTagSetService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestTagSetServiceGetWithEmptyID(t *testing.T) {
	service := NewTagSetService(&sling.Sling{}, "")

	resource, err := service.Get("")

	if err != nil {
		return
	}

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
