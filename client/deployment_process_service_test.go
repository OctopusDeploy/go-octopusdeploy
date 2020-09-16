package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewDeploymentProcessServiceWithNil(t *testing.T) {
	service := NewDeploymentProcessService(nil, "")
	assert.Nil(t, service)
}

func TestDeploymentProcessServiceWithEmptyClient(t *testing.T) {
	service := NewDeploymentProcessService(&sling.Sling{}, "")
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}

func TestDeploymentProcessServiceGetWithEmptyID(t *testing.T) {
	service := NewDeploymentProcessService(&sling.Sling{}, "")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}
