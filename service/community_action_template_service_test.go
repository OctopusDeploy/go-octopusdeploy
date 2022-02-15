package service

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createCommunityActionTemplateService(t *testing.T) *communityActionTemplateService {
	service := newCommunityActionTemplateService(nil, TestURICommunityActionTemplates)
	testNewService(t, service, TestURICommunityActionTemplates, ServiceCommunityActionTemplateService)
	return service
}

func TestCommunityActionTemplateServiceNew(t *testing.T) {
	ServiceFunction := newCommunityActionTemplateService
	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServiceCommunityActionTemplateService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *communityActionTemplateService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, ServiceName)
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
			assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
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
