package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewCommunityActionTemplateService(t *testing.T) {
	serviceFunction := newCommunityActionTemplateService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceCommunityActionTemplateService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *communityActionTemplateService
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

func TestCommunityActionTemplateServiceGetWithEmptyID(t *testing.T) {
	service := createCommunityActionTemplateService(t)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}

func createCommunityActionTemplateService(t *testing.T) *communityActionTemplateService {
	service := newCommunityActionTemplateService(nil, TestURICommunityActionTemplates)
	testNewService(t, service, TestURICommunityActionTemplates, serviceCommunityActionTemplateService)
	return service
}
