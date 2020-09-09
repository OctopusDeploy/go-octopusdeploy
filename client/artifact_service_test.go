package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewArtifactServiceWithNil(t *testing.T) {
	service := NewArtifactService(nil)
	assert.Nil(t, service)
}

func TestArtifactServiceWithEmptyClient(t *testing.T) {
	service := NewArtifactService(&sling.Sling{})
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestArtifactServiceGetWithEmptyID(t *testing.T) {
	service := NewArtifactService(&sling.Sling{})

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
