package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createWorkerService(t *testing.T) *workerService {
	service := newWorkerService(nil, TestURIWorkers, TestURIDiscoverWorker, TestURIWorkerOperatingSystems, TestURIWorkerShells)
	testNewService(t, service, TestURIWorkers, ServiceWorkerService)
	return service
}

func TestWorkerServiceAdd(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterWorker))
	require.Nil(t, resource)

	resource, err = service.Add(&Worker{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestWorkerServiceAddGetDelete(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(OperationAdd, ParameterWorker))
	require.Nil(t, resource)

	resource, err = service.Add(&Worker{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestWorkerServiceDelete(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(emptyString)
	require.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)

	err = service.DeleteByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(OperationDeleteByID, ParameterID), err)
}

func TestWorkerServiceGetByID(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByID, ParameterID))
	assert.Nil(t, resource)
}

func TestWorkerServiceGetByName(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	workers, err := service.GetByName(emptyString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByName, ParameterName))
	assert.NotNil(t, workers)

	workers, err = service.GetByName(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(OperationGetByName, ParameterName))
	assert.NotNil(t, workers)
}

func TestWorkerServiceGetByPartialName(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	workers, err := service.GetByPartialName(emptyString)
	require.Equal(t, err, createInvalidParameterError(OperationGetByPartialName, ParameterName))
	require.NotNil(t, workers)
	require.Len(t, workers, 0)

	workers, err = service.GetByPartialName(whitespaceString)
	require.Equal(t, err, createInvalidParameterError(OperationGetByPartialName, ParameterName))
	require.NotNil(t, workers)
	require.Len(t, workers, 0)
}

func TestWorkerServiceNew(t *testing.T) {
	ServiceFunction := newWorkerService
	client := &sling.Sling{}
	uriTemplate := emptyString
	discoverWorkerPath := emptyString
	operatingSystemsPath := emptyString
	shellsPath := emptyString
	ServiceName := ServiceWorkerService

	testCases := []struct {
		name                 string
		f                    func(*sling.Sling, string, string, string, string) *workerService
		client               *sling.Sling
		uriTemplate          string
		discoverWorkerPath   string
		operatingSystemsPath string
		shellsPath           string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, discoverWorkerPath, operatingSystemsPath, shellsPath},
		{"EmptyURITemplate", ServiceFunction, client, emptyString, discoverWorkerPath, operatingSystemsPath, shellsPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString, discoverWorkerPath, operatingSystemsPath, shellsPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.discoverWorkerPath, tc.operatingSystemsPath, tc.shellsPath)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}
