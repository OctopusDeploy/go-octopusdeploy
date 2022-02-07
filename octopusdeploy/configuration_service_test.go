package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigurationService(t *testing.T) {
	ServiceFunction := newConfigurationService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	versionControlClearCachePath := services.emptyString
	ServiceName := ServiceConfigurationService

	testCases := []struct {
		name                         string
		f                            func(*sling.Sling, string, string) *configurationService
		client                       *sling.Sling
		uriTemplate                  string
		versionControlClearCachePath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, versionControlClearCachePath},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString, versionControlClearCachePath},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString, versionControlClearCachePath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.versionControlClearCachePath)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestConfigurationServiceGetOperations(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"All Operations", "go-"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createConfigurationService(t)

			assert.NotNil(t, service)
			if service == nil {
				return
			}

			// TODO: put GetBy operation here
		})
	}
}

func TestConfigurationServiceGetWithEmptyID(t *testing.T) {
	service := createConfigurationService(t)

	resource, err := service.GetByID(services.emptyString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(services.whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)
}

func createConfigurationService(t *testing.T) *configurationService {
	service := newConfigurationService(nil, TestURIConfiguration, TestURIVersionControlClearCache)
	services.testNewService(t, service, TestURIConfiguration, ServiceConfigurationService)
	return service
}
