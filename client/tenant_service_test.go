package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewTenantService(t *testing.T) {
	serviceFunction := newTenantService
	client := &sling.Sling{}
	uriTemplate := emptyString
	missingVariablesPath := emptyString
	statusPath := emptyString
	tagTestPath := emptyString
	serviceName := serviceTenantService

	testCases := []struct {
		name                 string
		f                    func(*sling.Sling, string, string, string, string) *tenantService
		client               *sling.Sling
		uriTemplate          string
		missingVariablesPath string
		statusPath           string
		tagTestPath          string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, missingVariablesPath, statusPath, tagTestPath},
		{"EmptyURITemplate", serviceFunction, client, emptyString, missingVariablesPath, statusPath, tagTestPath},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, missingVariablesPath, statusPath, tagTestPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.missingVariablesPath, tc.statusPath, tc.tagTestPath)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestTenantServiceGetWithEmptyID(t *testing.T) {
	service := newTenantService(nil, emptyString, TestURITenantsMissingVariables, TestURITenantsStatus, TestURITenantTagTest)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}
