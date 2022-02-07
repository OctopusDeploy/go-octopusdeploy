package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createChannelService(t *testing.T) *channelService {
	service := newChannelService(nil, TestURIChannels, TestURIVersionRuleTest)
	services.testNewService(t, service, TestURIChannels, ServiceChannelService)
	return service
}

func TestChannelServiceAdd(t *testing.T) {
	service := createChannelService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterChannel))
	assert.Nil(t, resource)

	invalidResource := &Channel{}
	resource, err = service.Add(invalidResource)
	assert.Equal(t, createValidationFailureError(OperationAdd, invalidResource.Validate()), err)
	assert.Nil(t, resource)
}

func TestChannelServiceGetByID(t *testing.T) {
	service := createChannelService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(services.emptyString)
	assert.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	assert.Nil(t, resource)

	resource, err = service.GetByID(services.whitespaceString)
	assert.Equal(t, createInvalidParameterError(OperationGetByID, ParameterID), err)
	assert.Nil(t, resource)
}

func TestChannelServiceNew(t *testing.T) {
	ServiceFunction := newChannelService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	versionRuleTestPath := services.emptyString
	ServiceName := ServiceChannelService

	testCases := []struct {
		name                string
		f                   func(*sling.Sling, string, string) *channelService
		client              *sling.Sling
		uriTemplate         string
		versionRuleTestPath string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, versionRuleTestPath},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString, versionRuleTestPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString, versionRuleTestPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.versionRuleTestPath)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
