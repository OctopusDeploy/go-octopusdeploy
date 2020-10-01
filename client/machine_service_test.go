package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createMachineService(t *testing.T) *machineService {
	service := newMachineService(nil, TestURIMachines)
	testNewService(t, service, TestURIMachines, serviceMachineService)
	return service
}

func TestMachineService(t *testing.T) {
	t.Run("Add", TestMachineServiceAdd)
	t.Run("Delete", TestMachineServiceDelete)
	t.Run("GetByID", TestMachineServiceGetByID)
	t.Run("New", TestMachineServiceNew)
}

func TestMachineServiceAdd(t *testing.T) {
	assert := assert.New(t)

	service := createMachineService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	assert.Equal(err, createInvalidParameterError(operationAdd, parameterResource))
	assert.Nil(resource)

	invalidResource := &model.Machine{}
	resource, err = service.Add(invalidResource)
	assert.Equal(createValidationFailureError("Add", invalidResource.Validate()), err)
	assert.Nil(resource)
}

func TestMachineServiceDelete(t *testing.T) {
	assert := assert.New(t)

	service := createMachineService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(emptyString)
	assert.Equal(createInvalidParameterError(operationDeleteByID, parameterID), err)

	err = service.DeleteByID(whitespaceString)
	assert.Equal(createInvalidParameterError(operationDeleteByID, parameterID), err)

	id := getRandomName()
	err = service.DeleteByID(id)
	assert.Equal(createResourceNotFoundError("machine", "ID", id), err)
}

func TestMachineServiceGetByID(t *testing.T) {
	assert := assert.New(t)

	service := createMachineService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	assert.Equal(createInvalidParameterError(operationGetByID, parameterID), err)
	assert.Nil(resource)

	resource, err = service.GetByID(whitespaceString)
	assert.Equal(createInvalidParameterError(operationGetByID, parameterID), err)
	assert.Nil(resource)

	id := getRandomName()
	resource, err = service.GetByID(id)
	assert.Equal(createResourceNotFoundError("machine", "ID", id), err)
	assert.Nil(resource)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	if len(resources) > 0 {
		resourceToCompare, err := service.GetByID(resources[0].ID)
		assert.NoError(err)
		assert.EqualValues(resources[0], *resourceToCompare)
	}
}

func TestMachineServiceNew(t *testing.T) {
	serviceFunction := newMachineService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceMachineService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *machineService
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
