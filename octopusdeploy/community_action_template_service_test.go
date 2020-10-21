package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createCommunityActionTemplateService(t *testing.T) *communityActionTemplateService {
	service := newCommunityActionTemplateService(nil, TestURICommunityActionTemplates)
	testNewService(t, service, TestURICommunityActionTemplates, serviceCommunityActionTemplateService)
	return service
}

func TestCommunityActionTemplateServiceNew(t *testing.T) {
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

func TestCommunityActionTemplateServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", emptyString},
		{"Whitespace", whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createCommunityActionTemplateService(t)
			require.NotNil(t, service)

			resource, err := service.GetByID(tc.parameter)
			assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
			assert.Nil(t, resource)
		})
	}
}

func TestCommunityActionTemplateServiceInstall(t *testing.T) {
	service := createCommunityActionTemplateService(t)
	require.NotNil(t, service)

	resource, err := service.Install(CommunityActionTemplate{})
	require.Error(t, err)
	require.Nil(t, resource)
}
