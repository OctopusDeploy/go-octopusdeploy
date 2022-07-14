package environments

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createEnvironmentService(t *testing.T) *EnvironmentService {
	service := NewEnvironmentService(nil, constants.TestURIEnvironments, constants.TestURIEnvironmentSortOrder, constants.TestURIEnvironmentsSummary)
	services.NewServiceTests(t, service, constants.TestURIEnvironments, constants.ServiceEnvironmentService)
	return service
}

func TestEnvironmentServiceAdd(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterEnvironment))
	require.Nil(t, resource)

	resource, err = service.Add(&Environment{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestEnvironmentServiceDelete(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	err := service.DeleteByID("")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID), err)

	err = service.DeleteByID(" ")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID), err)
}

func TestEnvironmentServiceGetByID(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID("")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(" ")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)
}

func TestEnvironmentServiceGetBy(t *testing.T) {
	service := createEnvironmentService(t)
	require.NotNil(t, service)

	environments, err := service.GetByName("")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName))
	assert.NotNil(t, environments)

	environments, err = service.GetByName(" ")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName))
	assert.NotNil(t, environments)

	environments, err = service.GetByPartialName("")
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName))
	require.NotNil(t, environments)
	require.Len(t, environments, 0)

	environments, err = service.GetByPartialName(" ")
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName))
	require.NotNil(t, environments)
	require.Len(t, environments, 0)
}

func TestEnvironmentServiceNew(t *testing.T) {
	ServiceFunction := NewEnvironmentService
	client := &sling.Sling{}
	uriTemplate := ""
	sortOrderPath := ""
	summaryPath := ""
	ServiceName := constants.ServiceEnvironmentService

	testCases := []struct {
		name          string
		f             func(*sling.Sling, string, string, string) *EnvironmentService
		client        *sling.Sling
		uriTemplate   string
		sortOrderPath string
		summaryPath   string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, sortOrderPath, summaryPath},
		{"EmptyURITemplate", ServiceFunction, client, "", sortOrderPath, summaryPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", sortOrderPath, summaryPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.sortOrderPath, tc.summaryPath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}
