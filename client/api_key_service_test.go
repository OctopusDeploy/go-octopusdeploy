package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewAPIKeyService(t *testing.T) {
	serviceFunction := newAPIKeyService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceAPIKeyService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *apiKeyService
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

func TestAPIKeyServiceGetWithEmptyID(t *testing.T) {
	service := createAPIKeyService(t)

	resource, err := service.GetByUserID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByUserID, parameterUserID))
	assert.Nil(t, resource)

	resource, err = service.GetByUserID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByUserID, parameterUserID))
	assert.Nil(t, resource)
}

func createAPIKeyService(t *testing.T) *apiKeyService {
	service := newAPIKeyService(nil, TestURIAPIKeys)
	testNewService(t, service, TestURIAPIKeys, serviceAPIKeyService)
	return service
}
