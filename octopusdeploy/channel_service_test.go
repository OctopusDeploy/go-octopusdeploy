package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createChannelService(t *testing.T) *channelService {
	service := newChannelService(nil, TestURIChannels, TestURIVersionRuleTest)
	testNewService(t, service, TestURIChannels, serviceChannelService)
	return service
}

func TestChannelServiceAdd(t *testing.T) {
	service := createChannelService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
	assert.Nil(t, resource)

	invalidResource := &Channel{}
	resource, err = service.Add(invalidResource)
	assert.Equal(t, createValidationFailureError(operationAdd, invalidResource.Validate()), err)
	assert.Nil(t, resource)
}

func TestChannelServiceGetByID(t *testing.T) {
	service := createChannelService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	assert.Equal(t, createInvalidParameterError(operationGetByID, parameterID), err)
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	assert.Equal(t, createInvalidParameterError(operationGetByID, parameterID), err)
	assert.Nil(t, resource)
}

func TestChannelServiceNew(t *testing.T) {
	serviceFunction := newChannelService
	client := &sling.Sling{}
	uriTemplate := emptyString
	versionRuleTestPath := emptyString
	serviceName := serviceChannelService

	testCases := []struct {
		name                string
		f                   func(*sling.Sling, string, string) *channelService
		client              *sling.Sling
		uriTemplate         string
		versionRuleTestPath string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, versionRuleTestPath},
		{"EmptyURITemplate", serviceFunction, client, emptyString, versionRuleTestPath},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, versionRuleTestPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.versionRuleTestPath)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}
