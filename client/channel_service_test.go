package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
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

	invalidResource := &model.Channel{}
	resource, err = service.Add(invalidResource)
	assert.Equal(t, createValidationFailureError("Add", invalidResource.Validate()), err)
	assert.Nil(t, resource)
}

func TestChannelServiceGetAll(t *testing.T) {
	service := createChannelService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, resources)

	for _, resource := range resources {
		assert.NotNil(t, resource)
		assert.NotEmpty(t, resource.ID)
	}
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

	id := getRandomName()
	resource, err = service.GetByID(id)
	assert.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
	assert.Nil(t, resource)

	resources, err := service.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, resources)

	if len(resources) > 0 {
		resourceToCompare, err := service.GetByID(resources[0].ID)
		assert.NoError(t, err)
		assert.EqualValues(t, resources[0], *resourceToCompare)
	}
}

func TestChannelServiceGetByPartialName(t *testing.T) {
	service := createChannelService(t)
	require.NotNil(t, service)

	resources, err := service.GetByPartialName(emptyString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	assert.NotNil(t, resources)
	assert.Len(t, resources, 0)

	resources, err = service.GetByPartialName(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	assert.NotNil(t, resources)
	assert.Len(t, resources, 0)

	name := getRandomName()
	resources, err = service.GetByPartialName(name)
	assert.NoError(t, err)
	assert.NotNil(t, resources)
	assert.Len(t, resources, 0)

	resources, err = service.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, resources)

	if len(resources) > 0 {
		resourcesToCompare, err := service.GetByPartialName(resources[0].Name)
		assert.NoError(t, err)
		assert.EqualValues(t, resourcesToCompare[0], resources[0])
	}
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
