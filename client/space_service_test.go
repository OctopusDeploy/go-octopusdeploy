package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewSpaceService(t *testing.T) {
	serviceFunction := newSpaceService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceSpaceService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *spaceService
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

func TestSpaceServiceGetWithEmptyID(t *testing.T) {
	service := newSpaceService(&sling.Sling{}, emptyString)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}
