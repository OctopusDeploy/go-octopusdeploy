package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeploymentServiceGetByIDs(t *testing.T) {
	service := createDeploymentService(t)
	assert := assert.New(t)

	ids := []string{"Accounts-285", "Accounts-286"}

	resources, err := service.GetByIDs(ids)

	assert.NoError(err)

	t.Log(resources)
}

func createDeploymentService(t *testing.T) *deploymentService {
	service := newDeploymentService(nil, TestURIDeployments)
	testNewService(t, service, TestURIDeployments, serviceDeploymentService)
	return service
}
