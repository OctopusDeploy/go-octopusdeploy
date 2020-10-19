package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
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
	testNewService(t, service, TestURIWorkerPools, serviceWorkerPoolService)
	return service
}

func CreateTestDynamicWorkerPool(t *testing.T, service *workerPoolService) model.IWorkerPool {
	if service == nil {
		service = createWorkerPoolService(t)
	}
	require.NotNil(t, service)

	name := getRandomName()
	workerType := "Ubuntu1804"

	dynamicWorkerPool := model.NewDynamicWorkerPool(name, workerType)
	require.NoError(t, dynamicWorkerPool.Validate())

	createdDynamicWorkerPool, err := service.Add(dynamicWorkerPool)
	require.NoError(t, err)
	require.NotNil(t, createdDynamicWorkerPool)
	require.NotEmpty(t, createdDynamicWorkerPool.GetID())

	return createdDynamicWorkerPool
}

func CreateTestStaticWorkerPool(t *testing.T, service *workerPoolService) model.IWorkerPool {
	if service == nil {
		service = createWorkerPoolService(t)
	}
	require.NotNil(t, service)

	name := getRandomName()

	staticWorkerPool := model.NewStaticWorkerPool(name)
	require.NoError(t, staticWorkerPool.Validate())

	createdStaticWorkerPool, err := service.Add(staticWorkerPool)
	require.NoError(t, err)
	require.NotNil(t, createdStaticWorkerPool)
	require.NotEmpty(t, createdStaticWorkerPool.GetID())

	return createdStaticWorkerPool
}

func DeleteTestWorkerPool(t *testing.T, service *workerPoolService, workerPool model.IWorkerPool) error {
	if service == nil {
		service = createWorkerPoolService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(workerPool.GetID())
}

func IsEqualWorkerPools(t *testing.T, expected model.IWorkerPool, actual model.IWorkerPool) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	// IResource
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, IsEqualLinks(expected.GetLinks(), actual.GetLinks()))

	// IWorkerPool
	assert.Equal(t, expected.GetName(), actual.GetName())
	assert.Equal(t, expected.GetWorkerPoolType(), actual.GetWorkerPoolType())
}

func UpdateWorkerPool(t *testing.T, service *workerPoolService, workerPool *model.WorkerPool) *model.WorkerPool {
	if service == nil {
		service = createWorkerPoolService(t)
	}
	require.NotNil(t, service)

	updatedWorkerPool, err := service.Update(workerPool)
	assert.NoError(t, err)
	require.NotNil(t, updatedWorkerPool)

	return updatedWorkerPool
}

func TestWorkerPoolServiceAdd(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterWorkerPool))
	require.Nil(t, resource)

	// TODO: test this call; it should NOT send anything via HTTP
	resource, err = service.Add(&model.WorkerPool{})
	require.Error(t, err)
	require.Nil(t, resource)

	dynamicWorkerPool := CreateTestDynamicWorkerPool(t, service)
	defer DeleteTestWorkerPool(t, service, dynamicWorkerPool)

	staticWorkerPool := CreateTestStaticWorkerPool(t, service)
	defer DeleteTestWorkerPool(t, service, staticWorkerPool)
}

func TestWorkerPoolServiceAddGetDelete(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterWorkerPool))
	require.Nil(t, resource)

	// TODO: test this call; it should NOT send anything via HTTP
	resource, err = service.Add(&model.WorkerPool{})
	require.Error(t, err)
	require.Nil(t, resource)

	dynamicWorkerPool := CreateTestDynamicWorkerPool(t, service)
	defer DeleteTestWorkerPool(t, service, dynamicWorkerPool)

	resource, err = service.GetByID(dynamicWorkerPool.GetID())
	require.NoError(t, err)
	require.NotNil(t, resource)

	staticWorkerPool := CreateTestStaticWorkerPool(t, service)
	defer DeleteTestWorkerPool(t, service, staticWorkerPool)

	resource, err = service.GetByID(staticWorkerPool.GetID())
	require.NoError(t, err)
	require.NotNil(t, resource)
}

func TestWorkerPoolServiceDelete(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(emptyString)
	require.Equal(t, createInvalidParameterError(operationDeleteByID, parameterID), err)

	err = service.DeleteByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(operationDeleteByID, parameterID), err)

	id := getRandomName()
	err = service.DeleteByID(id)
	require.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
}

func TestWorkerPoolServiceDeleteAll(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	workerPools, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)

	for _, workerPool := range workerPools {
		err := DeleteTestWorkerPool(t, service, workerPool)

		if workerPool.IsDefault {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestWorkerPoolServiceGetAll(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	// create 30 test dynamic worker pools (to be deleted)
	for i := 0; i < 30; i++ {
		dynamicWorkerPool := CreateTestDynamicWorkerPool(t, service)
		require.NotNil(t, dynamicWorkerPool)
	}

	workerPools, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)

	for _, workerPool := range workerPools {
		require.NotNil(t, workerPool)
		require.NotEmpty(t, workerPool.GetID())
		err := DeleteTestWorkerPool(t, service, workerPool)

		if workerPool.IsDefault {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}

	// create 30 test static worker pools (to be deleted)
	for i := 0; i < 30; i++ {
		staticWorkerPool := CreateTestStaticWorkerPool(t, service)
		require.NotNil(t, staticWorkerPool)
	}

	workerPools, err = service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)

	for _, workerPool := range workerPools {
		require.NotNil(t, workerPool)
		require.NotEmpty(t, workerPool.GetID())
		err := DeleteTestWorkerPool(t, service, workerPool)

		if workerPool.IsDefault {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestWorkerPoolServiceGetByID(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
	assert.Nil(t, resource)

	id := getRandomName()
	resource, err = service.GetByID(id)
	require.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
	require.Nil(t, resource)

	workerPools, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)

	for _, workerPool := range workerPools {
		workerPoolToCompare, err := service.GetByID(workerPool.GetID())
		require.NoError(t, err)
		IsEqualWorkerPools(t, workerPool, workerPoolToCompare)
	}
}

func TestWorkerPoolServiceGetByName(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	workerPools, err := service.GetByName(emptyString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByName, parameterName))
	assert.NotNil(t, workerPools)

	workerPools, err = service.GetByName(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByName, parameterName))
	assert.NotNil(t, workerPools)

	workerPools, err = service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)

	for _, workerPool := range workerPools {
		namedWorkerPools, err := service.GetByName(workerPool.Name)
		require.NoError(t, err)
		require.NotNil(t, namedWorkerPools)
	}
}

func TestWorkerPoolServiceGetByPartialName(t *testing.T) {
	service := createWorkerPoolService(t)
	require.NotNil(t, service)

	workerPools, err := service.GetByPartialName(emptyString)
	require.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	require.NotNil(t, workerPools)
	require.Len(t, workerPools, 0)

	workerPools, err = service.GetByPartialName(whitespaceString)
	require.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	require.NotNil(t, workerPools)
	require.Len(t, workerPools, 0)

	workerPools, err = service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)

	for _, workerPool := range workerPools {
		namedWorkerPools, err := service.GetByPartialName(workerPool.Name)
		require.NoError(t, err)
		require.NotNil(t, namedWorkerPools)
	}
}

func TestWorkerPoolServiceNew(t *testing.T) {
	serviceFunction := newWorkerPoolService
	client := &sling.Sling{}
	uriTemplate := emptyString
	dynamicWorkerTypesPath := emptyString
	sortOrderPath := emptyString
	summaryPath := emptyString
	supportedTypesPath := emptyString
	serviceName := serviceWorkerPoolService

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
		{"NilClient", serviceFunction, nil, uriTemplate, dynamicWorkerTypesPath, sortOrderPath, summaryPath, supportedTypesPath},
		{"EmptyURITemplate", serviceFunction, client, emptyString, dynamicWorkerTypesPath, sortOrderPath, summaryPath, supportedTypesPath},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, dynamicWorkerTypesPath, sortOrderPath, summaryPath, supportedTypesPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.dynamicWorkerTypesPath, tc.sortOrderPath, tc.summaryPath, tc.supportedTypesPath)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

// func TestWorkerPoolServiceUpdate(t *testing.T) {
// 	service := createWorkerPoolService(t)
// 	require.NotNil(t, service)

// 	dynamicWorkerPool := CreateTestDynamicWorkerPool(t, service)
// 	defer DeleteTestWorkerPool(t, service, dynamicWorkerPool)

// 	newName := getRandomName()

// 	dynamicWorkerPool.SetName(newName)

// 	updatedWorkerPool := UpdateWorkerPool(t, service, dynamicWorkerPool)
// 	require.NotNil(t, updatedWorkerPool)

// 	require.NotEmpty(t, updatedWorkerPool.GetID())
// 	require.Equal(t, updatedWorkerPool.ID, updatedWorkerPool.GetID())
// 	require.Equal(t, newAllowDynamicInfrastructure, updatedWorkerPool.AllowDynamicInfrastructure)
// 	require.Equal(t, newDescription, updatedWorkerPool.Description)
// 	require.Equal(t, newName, updatedWorkerPool.Name)
// 	require.Equal(t, newSortOrder, updatedWorkerPool.SortOrder)
// 	require.Equal(t, newUseGuidedFailure, updatedWorkerPool.UseGuidedFailure)
// }
