package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createDeploymentProcessService(t *testing.T) *deploymentProcessService {
	service := newDeploymentProcessService(nil, TestURIDeploymentProcesses)
	testNewService(t, service, TestURIDeploymentProcesses, serviceDeploymentProcessesService)
	return service
}

func TestNewDeploymentProcessService(t *testing.T) {
	serviceFunction := newDeploymentProcessService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceDeploymentProcessesService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *deploymentProcessService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, emptyString},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, serviceName)
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
				require.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
				require.Nil(t, resource)
			} else {
				resource, err := service.GetByID(tc.parameter)
				require.Equal(t, err, createResourceNotFoundError(serviceDeploymentProcessesService, "ID", tc.parameter))
				require.Nil(t, resource)
			}
		})
	}
}

func TestDeploymentProcessServiceGetWithEmptyID(t *testing.T) {
	service := newDeploymentProcessService(&sling.Sling{}, emptyString)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}
