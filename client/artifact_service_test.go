package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewArtifactService(t *testing.T) {
	serviceFunction := newArtifactService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceArtifactService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *artifactService
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

func TestArtifactServiceOperationsWithStringParameter(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"EmptyParameter", emptyString},
		{"WhitespaceParameter", whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("[INFO] Parameter: %q", tc.parameter)

			service := createArtifactService(t)

			assert.NotNil(t, service)
			if service == nil {
				return
			}

			resource, err := service.GetByID(tc.parameter)

			assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
			assert.Nil(t, resource)

			resourceList, err := service.GetByPartialName(tc.parameter)

			assert.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
			assert.Nil(t, resourceList)

			err = service.DeleteByID(tc.parameter)

			assert.Error(t, err)
			assert.Equal(t, err, createInvalidParameterError(operationDeleteByID, parameterID))
		})
	}
}

func createArtifactService(t *testing.T) *artifactService {
	service := newArtifactService(&sling.Sling{}, TestURIArtifacts)
	testNewService(t, service, TestURIArtifacts, serviceArtifactService)
	return service
}
