package machines

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createWorkerService(t *testing.T) *WorkerService {
	service := NewWorkerService(nil, constants.TestURIWorkers, constants.TestURIDiscoverWorker, constants.TestURIWorkerOperatingSystems, constants.TestURIWorkerShells)
	services.NewServiceTests(t, service, constants.TestURIWorkers, constants.ServiceWorkerService)
	return service
}

func TestWorkerServiceAdd(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterWorker))
	require.Nil(t, resource)

	resource, err = service.Add(&Worker{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestWorkerServiceAddGetDelete(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterWorker))
	require.Nil(t, resource)

	resource, err = service.Add(&Worker{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestWorkerServiceDelete(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	err := service.DeleteByID("")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID), err)

	err = service.DeleteByID(" ")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID), err)
}

func TestWorkerServiceGetByID(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID("")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(" ")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)
}

func TestWorkerServiceGetByName(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	workers, err := service.GetByName("")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName))
	assert.NotNil(t, workers)

	workers, err = service.GetByName(" ")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByName, constants.ParameterName))
	assert.NotNil(t, workers)
}

func TestWorkerServiceGetByPartialName(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	workers, err := service.GetByPartialName("")
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName))
	require.NotNil(t, workers)
	require.Len(t, workers, 0)

	workers, err = service.GetByPartialName(" ")
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByPartialName, constants.ParameterPartialName))
	require.NotNil(t, workers)
	require.Len(t, workers, 0)
}

func TestWorkerServiceNew(t *testing.T) {
	ServiceFunction := NewWorkerService
	client := &sling.Sling{}
	uriTemplate := ""
	discoverWorkerPath := ""
	operatingSystemsPath := ""
	shellsPath := ""
	ServiceName := constants.ServiceWorkerService

	testCases := []struct {
		name                 string
		f                    func(*sling.Sling, string, string, string, string) *WorkerService
		client               *sling.Sling
		uriTemplate          string
		discoverWorkerPath   string
		operatingSystemsPath string
		shellsPath           string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, discoverWorkerPath, operatingSystemsPath, shellsPath},
		{"EmptyURITemplate", ServiceFunction, client, "", discoverWorkerPath, operatingSystemsPath, shellsPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", discoverWorkerPath, operatingSystemsPath, shellsPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.discoverWorkerPath, tc.operatingSystemsPath, tc.shellsPath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}
