package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createLifecycleService(t *testing.T) *lifecycleService {
	service := newLifecycleService(nil, TestURILifecycles)
	testNewService(t, service, TestURILifecycles, serviceLifecycleService)
	return service
}

func CreateTestLifecycle(t *testing.T, service *lifecycleService) *Lifecycle {
	if service == nil {
		service = createLifecycleService(t)
	}
	require.NotNil(t, service)

	name := getRandomName()

	lifecycle := NewLifecycle(name)
	require.NotNil(t, lifecycle)

	createdLifecycle, err := service.Add(lifecycle)
	require.NoError(t, err)
	require.NotNil(t, createdLifecycle)
	require.NotEmpty(t, createdLifecycle.GetID())

	return createdLifecycle
}

func DeleteTestLifecycle(t *testing.T, service *lifecycleService, lifecycle *Lifecycle) error {
	require.NotNil(t, lifecycle)

	if service == nil {
		service = createLifecycleService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(lifecycle.GetID())
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
	service := createLifecycleService(t)
	require.NotNil(t, service)

	resourceList, err := service.GetAll()
	require.NoError(t, err)
	require.NotEmpty(t, resourceList)

	for _, resource := range resourceList {
		resourceToCompare, err := service.GetByID(resource.ID)
		assert.NoError(t, err)
		assert.Equal(t, resource, resourceToCompare)
	}
}

func TestLifecycleServiceGetAll(t *testing.T) {
	service := createLifecycleService(t)
	require.NotNil(t, service)

	resourceList, err := service.GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, resourceList)
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
			assert.NotNil(t, service)
			if service == nil {
				return
			}

			resource, err := service.GetByID(tc.parameter)

			assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
			assert.Nil(t, resource)

			resourceList, err := service.GetByPartialName(tc.parameter)

			assert.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
			assert.NotNil(t, resourceList)

			err = service.DeleteByID(tc.parameter)

			assert.Error(t, err)
			assert.Equal(t, err, createInvalidParameterError(operationDeleteByID, parameterID))
		})
	}
}

func TestLifecycleServiceAdd(t *testing.T) {
	service := createLifecycleService(t)
	resource, err := service.Add(nil)

	assert.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
	assert.Nil(t, resource)

	resource, err = service.Add(&Lifecycle{})

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource = NewLifecycle(getRandomName())
	require.NotNil(t, resource)

	resource, err = service.Add(resource)

	assert.NoError(t, err)
	assert.NotNil(t, resource)

	err = service.DeleteByID(resource.ID)

	assert.NoError(t, err)
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

	account, err := service.Update(Lifecycle{})

	assert.Error(t, err)
	assert.Nil(t, account)
}
