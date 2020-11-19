package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createDeploymentProcessService(t *testing.T) *deploymentProcessService {
	service := newDeploymentProcessService(nil, TestURIDeploymentProcesses)
	testNewService(t, service, TestURIDeploymentProcesses, ServiceDeploymentProcessesService)
	return service
}

func TestNewDeploymentProcessService(t *testing.T) {
	ServiceFunction := newDeploymentProcessService
	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServiceDeploymentProcessesService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *deploymentProcessService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestDeploymentProcessServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", emptyString},
		{"Whitespace", whitespaceString},
		{"InvalidID", getRandomName()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createDeploymentProcessService(t)
			require.NotNil(t, service)

			if isEmpty(tc.parameter) {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
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
	service := newDeploymentProcessService(&sling.Sling{}, emptyString)

	resource, err := service.GetByID(emptyString)
	require.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	require.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	require.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	require.Nil(t, resource)
}
