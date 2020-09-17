package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

const (
	TestAPIKeyServiceURITemplate = "api-key-service"
)

func TestNewAPIKeyService(t *testing.T) {
	service := NewAPIKeyService(nil, "")
	assert.Nil(t, service)
	createAPIKeyService(t)
}

func TestAPIKeyServiceGetWithEmptyID(t *testing.T) {
	service := createAPIKeyService(t)

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "userID"))
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "userID"))
	assert.Nil(t, resource)
}

func createAPIKeyService(t *testing.T) *APIKeyService {
	service := NewAPIKeyService(&sling.Sling{}, TestAPIKeyServiceURITemplate)

	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
	assert.Equal(t, service.path, TestAPIKeyServiceURITemplate)
	assert.Equal(t, service.name, "APIKeyService")

	return service
}
