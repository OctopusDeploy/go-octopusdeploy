package deployments

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
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

func createDeploymentService(t *testing.T) *DeploymentService {
	service := NewDeploymentService(nil, constants.TestURIDeployments)
	services.NewServiceTests(t, service, constants.TestURIDeployments, constants.ServiceDeploymentService)
	return service
}
