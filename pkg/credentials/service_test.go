package credentials_test

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createService(t *testing.T) *credentials.Service {
	service := credentials.NewService(nil, constants.TestURIGitCredentials)
	services.NewServiceTests(t, service, constants.TestURIGitCredentials, constants.ServiceGitCredentialService)
	return service
}

func TestServiceAdd(t *testing.T) {
	service := createService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterGitCredential))
	require.Nil(t, resource)

	resource, err = service.Add(&credentials.Resource{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestServiceDelete(t *testing.T) {
	service := createService(t)
	require.NotNil(t, service)

	err := service.DeleteByID("")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID), err)

	err = service.DeleteByID(" ")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID), err)
}

func TestServiceGetByID(t *testing.T) {
	service := createService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID("")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(" ")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)
}

func TestService_GetByName(t *testing.T) {
	service := createService(t)
	require.NotNil(t, service)

	resource, err := service.GetByName("")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName))
	assert.Nil(t, resource)

	resource, err = service.GetByName(" ")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName))
	assert.Nil(t, resource)
}
