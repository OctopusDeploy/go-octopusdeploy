package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

const (
	TestBuildInformationServiceURITemplate = "build-information-service"
)

func TestNewBuildInformationService(t *testing.T) {
	service := NewBuildInformationService(nil, "")
	assert.Nil(t, service)
	createBuildInformationService(t)
}

func createBuildInformationService(t *testing.T) *BuildInformationService {
	service := NewBuildInformationService(&sling.Sling{}, TestBuildInformationServiceURITemplate)

	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
	assert.Equal(t, service.path, TestBuildInformationServiceURITemplate)
	assert.Equal(t, service.name, "BuildInformationService")

	return service
}
