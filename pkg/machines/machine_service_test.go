package machines

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createMachineService(t *testing.T) *MachineService {
	service := NewMachineService(nil, constants.TestURIMachines, constants.TestURIDiscoverMachine, constants.TestURIMachineOperatingSystems, constants.TestURIMachineShells)
	services.NewServiceTests(t, service, constants.TestURIMachines, constants.ServiceMachineService)
	return service
}

func TestMachineServiceAdd(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterResource))
	assert.Nil(t, resource)
}

func TestMachineServiceDelete(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	err := service.DeleteByID("")
	assert.Equal(t, internal.CreateInvalidParameterError("DeleteByID", "id"), err)

	err = service.DeleteByID(" ")
	assert.Equal(t, internal.CreateInvalidParameterError("DeleteByID", "id"), err)
}

func TestMachineServiceGetByID(t *testing.T) {
	service := createMachineService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID("")
	require.Equal(t, internal.CreateInvalidParameterError("GetByID", "id"), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(" ")
	require.Equal(t, internal.CreateInvalidParameterError("GetByID", "id"), err)
	require.Nil(t, resource)
}

func TestMachineServiceNew(t *testing.T) {
	ServiceFunction := NewMachineService
	client := &sling.Sling{}
	uriTemplate := ""
	discoverMachinePath := ""
	operatingSystemsPath := ""
	shellsPath := ""
	ServiceName := constants.ServiceMachineService

	testCases := []struct {
		name                 string
		f                    func(*sling.Sling, string, string, string, string) *MachineService
		client               *sling.Sling
		uriTemplate          string
		discoverMachinePath  string
		operatingSystemsPath string
		shellsPath           string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, discoverMachinePath, operatingSystemsPath, shellsPath},
		{"EmptyURITemplate", ServiceFunction, client, "", discoverMachinePath, operatingSystemsPath, shellsPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", discoverMachinePath, operatingSystemsPath, shellsPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.discoverMachinePath, tc.operatingSystemsPath, tc.shellsPath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}
