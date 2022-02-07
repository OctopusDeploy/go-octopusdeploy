package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createMachineService(t *testing.T) *machineService {
	service := newMachineService(nil, TestURIMachines, TestURIDiscoverMachine, TestURIMachineOperatingSystems, TestURIMachineShells)
	services.testNewService(t, service, TestURIMachines, ServiceMachineService)
	return service
}

func TestMachineServiceAdd(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterResource))
	assert.Nil(t, resource)
}

func TestMachineServiceDelete(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(services.emptyString)
	assert.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)

	err = service.DeleteByID(services.whitespaceString)
	assert.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)
}

func TestMachineServiceGetByID(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(services.emptyString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(services.whitespaceString)
	require.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	require.Nil(t, resource)
}

func TestMachineServiceNew(t *testing.T) {
	ServiceFunction := newMachineService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	discoverMachinePath := services.emptyString
	operatingSystemsPath := services.emptyString
	shellsPath := services.emptyString
	ServiceName := ServiceMachineService

	testCases := []struct {
		name                 string
		f                    func(*sling.Sling, string, string, string, string) *machineService
		client               *sling.Sling
		uriTemplate          string
		discoverMachinePath  string
		operatingSystemsPath string
		shellsPath           string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, discoverMachinePath, operatingSystemsPath, shellsPath},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString, discoverMachinePath, operatingSystemsPath, shellsPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString, discoverMachinePath, operatingSystemsPath, shellsPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.discoverMachinePath, tc.operatingSystemsPath, tc.shellsPath)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
