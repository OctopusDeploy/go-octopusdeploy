package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func createChannelService(t *testing.T) *channelService {
	service := newChannelService(nil, TestURIChannels)
	testNewService(t, service, TestURIChannels, serviceChannelService)
	return service
}

func TestChannelService(t *testing.T) {
	t.Run("Add", TestChannelServiceAdd)
	t.Run("Delete", TestChannelServiceDelete)
	t.Run("GetAll", TestChannelServiceGetAll)
	t.Run("GetByID", TestChannelServiceGetByID)
	t.Run("GetByPartialName", TestChannelServiceGetByPartialName)
	t.Run("New", TestChannelServiceNew)
}

func TestChannelServiceAdd(t *testing.T) {
	assert := assert.New(t)

	service := createChannelService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(err, createInvalidParameterError(operationAdd, parameterResource))
	assert.Nil(resource)

	invalidResource := &model.Channel{}
	resource, err = service.Add(invalidResource)
	assert.Equal(createValidationFailureError("Add", invalidResource.Validate()), err)
	assert.Nil(resource)
}

func TestChannelServiceDelete(t *testing.T) {
	assert := assert.New(t)

	service := createChannelService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(emptyString)
	assert.Equal(createInvalidParameterError(operationDeleteByID, parameterID), err)

	err = service.DeleteByID(whitespaceString)
	assert.Equal(createInvalidParameterError(operationDeleteByID, parameterID), err)

	id := getRandomName()
	err = service.DeleteByID(id)
	assert.Equal(createResourceNotFoundError("channel", "ID", id), err)
}

func TestChannelServiceGetAll(t *testing.T) {
	assert := assert.New(t)

	service := createChannelService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	assert.NoError(err)
	assert.NotNil(resources)

	for _, resource := range resources {
		assert.NotNil(resource)
		assert.NotEmpty(resource.ID)
	}
}

func TestChannelServiceGetByID(t *testing.T) {
	assert := assert.New(t)

	service := createChannelService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	assert.Equal(createInvalidParameterError(operationGetByID, parameterID), err)
	assert.Nil(resource)

	resource, err = service.GetByID(whitespaceString)
	assert.Equal(createInvalidParameterError(operationGetByID, parameterID), err)
	assert.Nil(resource)

	id := getRandomName()
	resource, err = service.GetByID(id)
	assert.Equal(createResourceNotFoundError("channel", "ID", id), err)
	assert.Nil(resource)

	resources, err := service.GetAll()
	assert.NoError(err)
	assert.NotNil(resources)

	if len(resources) > 0 {
		resourceToCompare, err := service.GetByID(resources[0].ID)
		assert.NoError(err)
		assert.EqualValues(resources[0], *resourceToCompare)
	}
}

func TestChannelServiceGetByPartialName(t *testing.T) {
	assert := assert.New(t)

	service := createChannelService(t)
	require.NotNil(t, service)

	resources, err := service.GetByPartialName(emptyString)
	assert.Equal(err, createInvalidParameterError(operationGetByPartialName, parameterName))
	assert.NotNil(resources)
	assert.Len(resources, 0)

	resources, err = service.GetByPartialName(whitespaceString)
	assert.Equal(err, createInvalidParameterError(operationGetByPartialName, parameterName))
	assert.NotNil(resources)
	assert.Len(resources, 0)

	name := getRandomName()
	resources, err = service.GetByPartialName(name)
	assert.NoError(err)
	assert.NotNil(resources)
	assert.Len(resources, 0)

	resources, err = service.GetAll()
	assert.NoError(err)
	assert.NotNil(resources)

	if len(resources) > 0 {
		resourcesToCompare, err := service.GetByPartialName(resources[0].Name)
		assert.NoError(err)
		assert.EqualValues(resourcesToCompare[0], resources[0])
	}
}

func TestChannelServiceNew(t *testing.T) {
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
