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
	serviceName := serviceTenantService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *tenantService
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

func TestTenantServiceGetWithEmptyID(t *testing.T) {
	service := newTenantService(&sling.Sling{}, emptyString)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}
