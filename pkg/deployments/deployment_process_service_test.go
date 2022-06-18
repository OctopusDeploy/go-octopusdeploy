package deployments

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createDeploymentProcessService(t *testing.T) *DeploymentProcessService {
	service := NewDeploymentProcessService(nil, constants.TestURIDeploymentProcesses)
	services.NewServiceTests(t, service, constants.TestURIDeploymentProcesses, constants.ServiceDeploymentProcessesService)
	return service
}

func TestNewDeploymentProcessService(t *testing.T) {
	ServiceFunction := NewDeploymentProcessService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceDeploymentProcessesService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *DeploymentProcessService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, ""},
		{"URITemplateWithWhitespace", ServiceFunction, client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestDeploymentProcessServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", ""},
		{"Whitespace", " "},
		{"InvalidID", internal.GetRandomName()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createDeploymentProcessService(t)
			require.NotNil(t, service)

			if internal.IsEmpty(tc.parameter) {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, internal.CreateInvalidParameterError("GetByID", "id"))
				require.Nil(t, resource)
			} else {
				resource, err := service.GetByID(tc.parameter)
				require.Error(t, err)
				require.Nil(t, resource)
			}
		})
	}
}

func TestDeploymentProcessServiceGetWithEmptyID(t *testing.T) {
	service := NewDeploymentProcessService(&sling.Sling{}, "")

	resource, err := service.GetByID("")
	require.Equal(t, err, internal.CreateInvalidParameterError("GetByID", "id"))
	require.Nil(t, resource)

	resource, err = service.GetByID(" ")
	require.Equal(t, err, internal.CreateInvalidParameterError("GetByID", "id"))
	require.Nil(t, resource)
}
