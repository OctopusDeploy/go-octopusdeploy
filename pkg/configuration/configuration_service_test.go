package configuration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigurationService(t *testing.T) {
	ServiceFunction := NewConfigurationService
	client := &sling.Sling{}
	uriTemplate := ""
	versionControlClearCachePath := ""
	ServiceName := constants.ServiceConfigurationService

	testCases := []struct {
		name                         string
		f                            func(*sling.Sling, string, string) *ConfigurationService
		client                       *sling.Sling
		uriTemplate                  string
		versionControlClearCachePath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, versionControlClearCachePath},
		{"EmptyURITemplate", ServiceFunction, client, "", versionControlClearCachePath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", versionControlClearCachePath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.versionControlClearCachePath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
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

	resource, err := service.GetByID("")

	assert.Equal(t, err, internal.CreateInvalidParameterError("GetByID", "id"))
	assert.Nil(t, resource)

	resource, err = service.GetByID(" ")

	assert.Equal(t, err, internal.CreateInvalidParameterError("GetByID", "id"))
	assert.Nil(t, resource)
}

func createConfigurationService(t *testing.T) *ConfigurationService {
	service := NewConfigurationService(nil, constants.TestURIConfiguration, constants.TestURIVersionControlClearCache)
	services.NewServiceTests(t, service, constants.TestURIConfiguration, constants.ServiceConfigurationService)
	return service
}
