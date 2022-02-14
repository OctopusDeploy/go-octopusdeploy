package services

import (
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
	testNewService(t, service, TestURIWorkerPools, ServiceWorkerPoolService)
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

	err := service.DeleteByID(emptyString)
	require.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)

	err = service.DeleteByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)
}

func TestWorkerPoolServiceGetByID(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)
}

func TestWorkerPoolServiceNew(t *testing.T) {
	ServiceFunction := newWorkerPoolService
	client := &sling.Sling{}
	uriTemplate := emptyString
	dynamicWorkerTypesPath := emptyString
	sortOrderPath := emptyString
	summaryPath := emptyString
	supportedTypesPath := emptyString
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
		{"EmptyURITemplate", ServiceFunction, client, emptyString, dynamicWorkerTypesPath, sortOrderPath, summaryPath, supportedTypesPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, dynamicWorkerTypesPath, sortOrderPath, summaryPath, supportedTypesPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.dynamicWorkerTypesPath, tc.sortOrderPath, tc.summaryPath, tc.supportedTypesPath)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
