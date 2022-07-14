package workerpools

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createWorkerPoolService(t *testing.T) *WorkerPoolService {
	service := NewWorkerPoolService(nil,
		constants.TestURIWorkerPools,
		constants.TestURIWorkerPoolsDynamicWorkerTypes,
		constants.TestURIWorkerPoolsSortOrder,
		constants.TestURIWorkerPoolsSummary,
		constants.TestURIWorkerPoolsSupportedTypes)
	services.NewServiceTests(t, service, constants.TestURIWorkerPools, constants.ServiceWorkerPoolService)
	return service
}

func TestWorkerPoolServiceAdd(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterWorkerPool))
	require.Nil(t, resource)

	resource, err = service.Add(&WorkerPoolResource{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestWorkerPoolServiceAddGetDelete(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterWorkerPool))
	require.Nil(t, resource)

	resource, err = service.Add(&WorkerPoolResource{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestWorkerPoolServiceDelete(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	err := service.DeleteByID("")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID), err)

	err = service.DeleteByID(" ")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID), err)
}

func TestWorkerPoolServiceGetByID(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID("")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(" ")
	assert.Equal(t, err, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID))
	assert.Nil(t, resource)
}

func TestWorkerPoolServiceNew(t *testing.T) {
	ServiceFunction := NewWorkerPoolService
	client := &sling.Sling{}
	uriTemplate := ""
	dynamicWorkerTypesPath := ""
	sortOrderPath := ""
	summaryPath := ""
	supportedTypesPath := ""
	ServiceName := constants.ServiceWorkerPoolService

	testCases := []struct {
		name                   string
		f                      func(*sling.Sling, string, string, string, string, string) *WorkerPoolService
		client                 *sling.Sling
		uriTemplate            string
		dynamicWorkerTypesPath string
		sortOrderPath          string
		summaryPath            string
		supportedTypesPath     string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate, dynamicWorkerTypesPath, sortOrderPath, summaryPath, supportedTypesPath},
		{"EmptyURITemplate", ServiceFunction, client, "", dynamicWorkerTypesPath, sortOrderPath, summaryPath, supportedTypesPath},
		{"URITemplateWithWhitespace", ServiceFunction, client, " ", dynamicWorkerTypesPath, sortOrderPath, summaryPath, supportedTypesPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.dynamicWorkerTypesPath, tc.sortOrderPath, tc.summaryPath, tc.supportedTypesPath)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}
