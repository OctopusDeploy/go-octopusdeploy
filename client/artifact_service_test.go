package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

const (
	TestArtifactServiceURITemplate = "artifact-service"
)

func TestNewArtifactService(t *testing.T) {
	service := NewArtifactService(nil, "")
	assert.Nil(t, service)
	createArtifactService(t)
}

func TestArtifactServiceGetWithEmptyID(t *testing.T) {
	service := createArtifactService(t)

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "id"))
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "id"))
	assert.Nil(t, resource)
}

func createArtifactService(t *testing.T) *ArtifactService {
	service := NewArtifactService(&sling.Sling{}, TestArtifactServiceURITemplate)

	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
	assert.Equal(t, service.path, TestArtifactServiceURITemplate)
	assert.Equal(t, service.name, "ArtifactService")

	return service
}
