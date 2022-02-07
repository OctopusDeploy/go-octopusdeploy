package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewTenantService(t *testing.T) {
	ServiceFunction := newTenantService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	missingVariablesPath := services.emptyString
	statusPath := services.emptyString
	tagTestPath := services.emptyString
	ServiceName := ServiceTenantService

	testCases := []struct {
		name                 string
		f                    func(*sling.Sling, string, string, string, string) *tenantService
		client               *sling.Sling
		uriTemplate          string
		missingVariablesPath string
		statusPath           string
		tagTestPath          string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, missingVariablesPath, statusPath, tagTestPath},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString, missingVariablesPath, statusPath, tagTestPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString, missingVariablesPath, statusPath, tagTestPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.missingVariablesPath, tc.statusPath, tc.tagTestPath)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestTenantServiceGetWithEmptyID(t *testing.T) {
	service := newTenantService(nil, services.emptyString, TestURITenantsMissingVariables, TestURITenantsStatus, TestURITenantTagTest)

	resource, err := service.GetByID(services.emptyString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(services.whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)
}
