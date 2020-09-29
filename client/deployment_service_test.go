package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeploymentServiceGetByIDs(t *testing.T) {
	assert := assert.New(t)

	service := createDeploymentService(t)
	assert.NotNil(service)
	if service == nil {
		return
	}

	ids := []string{"Accounts-285", "Accounts-286"}
	resources, err := service.GetByIDs(ids)

	assert.NoError(err)
	assert.NotNil(resources)
}

func createDeploymentService(t *testing.T) *deploymentService {
	service := newDeploymentService(nil, TestURIDeployments)
	testNewService(t, service, TestURIDeployments, serviceDeploymentService)
	return service
}
