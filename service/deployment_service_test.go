package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeploymentServiceGetByIDs(t *testing.T) {
	service := createDeploymentService(t)
	require.NotNil(t, service)

	ids := []string{"Accounts-285", "Accounts-286"}
	resources, err := service.GetByIDs(ids)

	assert.NoError(t, err)
	assert.NotNil(t, resources)
}

func createDeploymentService(t *testing.T) *deploymentService {
	service := newDeploymentService(nil, TestURIDeployments)
	testNewService(t, service, TestURIDeployments, ServiceDeploymentService)
	return service
}
