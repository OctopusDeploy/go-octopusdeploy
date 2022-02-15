package service

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewTenantService(t *testing.T) {
	ServiceFunction := newTenantService
	client := &sling.Sling{}
	uriTemplate := emptyString
	missingVariablesPath := emptyString
	statusPath := emptyString
	tagTestPath := emptyString
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
		{"EmptyURITemplate", ServiceFunction, client, emptyString, missingVariablesPath, statusPath, tagTestPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, missingVariablesPath, statusPath, tagTestPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.missingVariablesPath, tc.statusPath, tc.tagTestPath)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestTenantServiceGetWithEmptyID(t *testing.T) {
	service := newTenantService(nil, emptyString, TestURITenantsMissingVariables, TestURITenantsStatus, TestURITenantTagTest)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)
}
