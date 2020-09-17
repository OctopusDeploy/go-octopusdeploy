package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

const (
	TestAzureDevOpsServiceURITemplate = "azure-devops-service"
)

func TestNewAzureDevOpsService(t *testing.T) {
	service := NewAzureDevOpsService(nil, "")
	assert.Nil(t, service)
	createAzureDevOpsService(t)
}

func createAzureDevOpsService(t *testing.T) *AzureDevOpsService {
	service := NewAzureDevOpsService(&sling.Sling{}, TestAzureDevOpsServiceURITemplate)

	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
	assert.Equal(t, service.path, TestAzureDevOpsServiceURITemplate)
	assert.Equal(t, service.name, "AzureDevOpsService")

	return service
}
