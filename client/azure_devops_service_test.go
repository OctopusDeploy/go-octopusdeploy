package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewAzureDevOpsServiceWithNil(t *testing.T) {
	service := NewAzureDevOpsService(nil)
	assert.Nil(t, service)
}

func TestAzureDevOpsServiceWithEmptyClient(t *testing.T) {
	service := NewAzureDevOpsService(&sling.Sling{})
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}
