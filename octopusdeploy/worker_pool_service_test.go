package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createWorkerPoolService(t *testing.T) *workerPoolService {
	service := newWorkerPoolService(nil,
		TestURIWorkerPools,
		TestURIWorkerPoolsDynamicWorkerTypes,
		TestURIWorkerPoolsSortOrder,
		TestURIWorkerPoolsSummary,
		TestURIWorkerPoolsSupportedTypes)
	services.testNewService(t, service, TestURIWorkerPools, ServiceWorkerPoolService)
	return service
}

func TestWorkerPoolServiceAdd(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterWorkerPool))
	require.Nil(t, resource)

	resource, err = service.Add(&WorkerPoolResource{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestWorkerPoolServiceAddGetDelete(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterWorkerPool))
	require.Nil(t, resource)

	resource, err = service.Add(&WorkerPoolResource{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestWorkerPoolServiceDelete(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(services.emptyString)
	require.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)

	err = service.DeleteByID(services.whitespaceString)
	require.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)
}

func TestWorkerPoolServiceGetByID(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(services.emptyString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(services.whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)
}

func TestWorkerPoolServiceNew(t *testing.T) {
	ServiceFunction := newWorkerPoolService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	dynamicWorkerTypesPath := services.emptyString
	sortOrderPath := services.emptyString
	summaryPath := services.emptyString
	supportedTypesPath := services.emptyString
	ServiceName := ServiceWorkerPoolService

	testCases := []struct {
		name                   string
		f                      func(*sling.Sling, string, string, string, string, string) *workerPoolService
		client                 *sling.Sling
		uriTemplate            string
		dynamicWorkerTypesPath string
		sortOrderPath          string
		summaryPath            string
		supportedTypesPath     string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, dynamicWorkerTypesPath, sortOrderPath, summaryPath, supportedTypesPath},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString, dynamicWorkerTypesPath, sortOrderPath, summaryPath, supportedTypesPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString, dynamicWorkerTypesPath, sortOrderPath, summaryPath, supportedTypesPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.dynamicWorkerTypesPath, tc.sortOrderPath, tc.summaryPath, tc.supportedTypesPath)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
