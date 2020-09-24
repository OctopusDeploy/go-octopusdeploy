package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewChannelService(t *testing.T) {
	serviceFunction := newChannelService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceChannelService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *channelService
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

func TestChannelServiceGetWithEmptyID(t *testing.T) {
	service := createChannelService(t)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}

func TestChannelServiceDeleteWithEmptyID(t *testing.T) {
	service := createChannelService(t)

	err := service.DeleteByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationDeleteByID, parameterID))

	err = service.DeleteByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationDeleteByID, parameterID))
}

func createChannelService(t *testing.T) *channelService {
	service := newChannelService(&sling.Sling{}, TestURIChannels)
	testNewService(t, service, TestURIChannels, serviceChannelService)
	return service
}
