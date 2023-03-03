package lifecycles

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createLifecycleService(t *testing.T) *LifecycleService {
	service := NewLifecycleService(nil, constants.TestURILifecycles)
	services.NewServiceTests(t, service, constants.TestURILifecycles, constants.ServiceLifecycleService)
	return service
}

func CreateTestLifecycle(t *testing.T, service *LifecycleService) *Lifecycle {
	if service == nil {
		service = createLifecycleService(t)
	}
	require.NotNil(t, service)

	name := internal.GetRandomName()

	lifecycle := NewLifecycle(name)
	require.NotNil(t, lifecycle)

	createdLifecycle, err := service.Add(lifecycle)
	require.NoError(t, err)
	require.NotNil(t, createdLifecycle)
	require.NotEmpty(t, createdLifecycle.GetID())

	return createdLifecycle
}

func DeleteTestLifecycle(t *testing.T, service *LifecycleService, lifecycle *Lifecycle) error {
	require.NotNil(t, lifecycle)

	if service == nil {
		service = createLifecycleService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(lifecycle.GetID())
}

func TestNewLifecycleService(t *testing.T) {
	ServiceFunction := NewLifecycleService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceLifecycleService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *LifecycleService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, ""},
		{"URITemplateWithWhitespace", ServiceFunction, client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestLifecycleServiceGetByID(t *testing.T) {
	t.Skip("This test is not working.")
	service := createLifecycleService(t)
	require.NotNil(t, service)

	resourceList, err := service.GetAll()
	require.NoError(t, err)
	require.NotEmpty(t, resourceList)

	for _, resource := range resourceList {
		resourceToCompare, err := service.GetByID(resource.GetID())
		assert.NoError(t, err)
		assert.Equal(t, resource, resourceToCompare)
	}
}

func TestLifecycleServiceGetAll(t *testing.T) {
	t.Skip("This test is not working.")
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
		{"Empty", ""},
		{"Whitespace", " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createLifecycleService(t)
			require.NotNil(t, service)

			resource, err := service.GetByID(tc.parameter)
			assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
			assert.Nil(t, resource)

			resourceList, err := service.GetByPartialName(tc.parameter)
			assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName))
			assert.NotNil(t, resourceList)

			err = service.DeleteByID(tc.parameter)
			assert.Error(t, err)
			assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID))
		})
	}
}

func TestLifecycleServiceAdd(t *testing.T) {
	service := createLifecycleService(t)

	resource, err := service.Add(nil)
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, "lifecycle"))
	assert.Nil(t, resource)

	resource, err = service.Add(&Lifecycle{})
	assert.Error(t, err)
	assert.Nil(t, resource)

	resource = NewLifecycle(internal.GetRandomName())
	require.NotNil(t, resource)

	resource, err = service.Add(resource)
	assert.NoError(t, err)
	assert.NotNil(t, resource)
	defer service.DeleteByID(resource.GetID())
}

func TestLifecycleServiceGetWithEmptyID(t *testing.T) {
	service := NewLifecycleService(&sling.Sling{}, "")

	resource, err := service.GetByID("")

	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(" ")

	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)
}

func TestLifecycleServiceGetByNameWithEmptyID(t *testing.T) {
	service := NewLifecycleService(&sling.Sling{}, "")

	resource, err := service.GetByName("")

	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName))
	assert.Nil(t, resource)

	resource, err = service.GetByName(" ")

	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName))
	assert.Nil(t, resource)
}

func TestLifecycleServiceUpdateWithEmptyLifecycle(t *testing.T) {
	service := createLifecycleService(t)

	lifecycle, err := service.Update(&Lifecycle{})
	assert.Error(t, err)
	assert.Nil(t, lifecycle)
}
