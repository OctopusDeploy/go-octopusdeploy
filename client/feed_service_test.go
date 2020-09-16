package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewFeedServiceWithNil(t *testing.T) {
	service := NewFeedService(nil, "")
	assert.Nil(t, service)
}

func TestFeedServiceWithEmptyClient(t *testing.T) {
	service := NewFeedService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestFeedServiceGetWithEmptyID(t *testing.T) {
	service := NewFeedService(&sling.Sling{}, "")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
