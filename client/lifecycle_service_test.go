package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestLifecycleService(t *testing.T) {
	t.Run("New", TestNewLifecycleService)
	t.Run("Parameters", TestLifecycleServiceStringParameters)
	t.Run("GetAll", TestLifecycleServiceGetAll)
	t.Run("GetByID", TestLifecycleServiceGetByID)
	t.Run("Add", TestLifecycleServiceAdd)
	t.Run("Update", TestLifecycleServiceUpdateWithEmptyLifecycle)
}

func TestNewLifecycleService(t *testing.T) {
	serviceFunction := newLifecycleService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceLifecycleService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *lifecycleService
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

func TestLifecycleServiceGetByID(t *testing.T) {
	assert := assert.New(t)

	service := createLifecycleService(t)
	require.NotNil(t, service)

	resourceList, err := service.GetAll()
	assert.NoError(err)
	assert.NotEmpty(resourceList)

	for _, resource := range resourceList {
		resourceToCompare, err := service.GetByID(resource.ID)
		assert.NoError(err)
		assert.EqualValues(resource, *resourceToCompare)
	}
}

func TestLifecycleServiceGetAll(t *testing.T) {
	assert := assert.New(t)

	service := createLifecycleService(t)
	require.NotNil(t, service)

	resourceList, err := service.GetAll()
	assert.NoError(err)
	assert.NotEmpty(resourceList)
}

func TestLifecycleServiceStringParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", emptyString},
		{"Whitespace", whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createLifecycleService(t)
			assert := assert.New(t)

			assert.NotNil(service)
			if service == nil {
				return
			}

			resource, err := service.GetByID(tc.parameter)

			assert.Equal(err, createInvalidParameterError(operationGetByID, parameterID))
			assert.Nil(resource)

			resourceList, err := service.GetByPartialName(tc.parameter)

			assert.Equal(err, createInvalidParameterError(operationGetByPartialName, parameterName))
			assert.NotNil(resourceList)

			err = service.DeleteByID(tc.parameter)

			assert.Error(err)
			assert.Equal(err, createInvalidParameterError(operationDeleteByID, parameterID))
		})
	}
}

func TestLifecycleServiceAdd(t *testing.T) {
	service := createLifecycleService(t)
	assert := assert.New(t)

	resource, err := service.Add(nil)

	assert.Equal(err, createInvalidParameterError(operationAdd, parameterResource))
	assert.Nil(resource)

	resource, err = service.Add(&model.Lifecycle{})

	assert.Error(err)
	assert.Nil(resource)

	resource, err = model.NewLifecycle(getRandomName())

	assert.NoError(err)
	assert.NotNil(resource)

	if err != nil {
		return
	}

	resource, err = service.Add(resource)

	assert.NoError(err)
	assert.NotNil(resource)

	err = service.DeleteByID(resource.ID)

	assert.NoError(err)
}

func TestLifecycleServiceGetWithEmptyID(t *testing.T) {
	service := newLifecycleService(&sling.Sling{}, emptyString)

	resource, err := service.GetByID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)
}

func TestLifecycleServiceUpdateWithEmptyLifecycle(t *testing.T) {
	service := createLifecycleService(t)

	account, err := service.Update(model.Lifecycle{})

	assert.Error(t, err)
	assert.Nil(t, account)
}

func createLifecycleService(t *testing.T) *lifecycleService {
	service := newLifecycleService(nil, TestURILifecycles)
	testNewService(t, service, TestURILifecycles, serviceLifecycleService)
	return service
}
