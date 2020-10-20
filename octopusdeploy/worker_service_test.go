package octopusdeploy

import (
	"net/url"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createWorkerService(t *testing.T) *workerService {
	service := newWorkerService(nil, TestURIWorkers, TestURIDiscoverWorker, TestURIWorkerOperatingSystems, TestURIWorkerShells)
	testNewService(t, service, TestURIWorkers, serviceWorkerService)
	return service
}

func CreateTestWorker(t *testing.T, service *workerService) *Worker {
	if service == nil {
		service = createWorkerService(t)
	}
	require.NotNil(t, service)

	name := getRandomName()
	thumbprint := getRandomName()

	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)
	require.NotNil(t, url)

	staticWorkerPool := CreateTestStaticWorkerPool(t, nil)

	listeningTentacleEndpoint := NewListeningTentacleEndpoint(url, thumbprint)
	require.NotNil(t, listeningTentacleEndpoint)

	worker := NewWorker(name, listeningTentacleEndpoint)
	worker.WorkerPoolIDs = append(worker.WorkerPoolIDs, staticWorkerPool.GetID())
	require.NotNil(t, worker)
	require.NoError(t, worker.Validate())

	createdWorker, err := service.Add(worker)
	require.NoError(t, err)
	require.NotNil(t, createdWorker)
	require.NotEmpty(t, createdWorker.GetID())

	return createdWorker
}

func DeleteTestWorker(t *testing.T, service *workerService, worker *Worker) error {
	if service == nil {
		service = createWorkerService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(worker.GetID())
}

func IsEqualWorkers(t *testing.T, expected *Worker, actual *Worker) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	// IResource
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, IsEqualLinks(expected.GetLinks(), actual.GetLinks()))

	// Worker
	assert.Equal(t, expected.Name, actual.Name)
}

func UpdateWorker(t *testing.T, service *workerService, worker *Worker) *Worker {
	if service == nil {
		service = createWorkerService(t)
	}
	require.NotNil(t, service)

	updatedWorker, err := service.Update(worker)
	assert.NoError(t, err)
	require.NotNil(t, updatedWorker)

	return updatedWorker
}

func TestWorkerServiceAdd(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterWorker))
	require.Nil(t, resource)

	// TODO: test this call; it should NOT send anything via HTTP
	resource, err = service.Add(&Worker{})
	require.Error(t, err)
	require.Nil(t, resource)

	worker := CreateTestWorker(t, service)
	err = DeleteTestWorker(t, service, worker)
	require.NoError(t, err)
}

func TestWorkerServiceAddGetDelete(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterWorker))
	require.Nil(t, resource)

	resource, err = service.Add(&Worker{})
	require.Error(t, err)
	require.Nil(t, resource)

	resource = CreateTestWorker(t, service)

	worker, err := service.GetByID(resource.GetID())
	require.NoError(t, err)
	require.NotNil(t, worker)
	IsEqualWorkers(t, resource, worker)

	for _, id := range worker.WorkerPoolIDs {
		workerPoolService := createWorkerPoolService(t)
		require.NotNil(t, workerPoolService)

		workerPool, err := workerPoolService.GetByID(id)
		require.NoError(t, err)

		err = DeleteTestWorkerPool(t, workerPoolService, workerPool)
		require.NoError(t, err)
	}

	err = DeleteTestWorker(t, service, worker)
	require.NoError(t, err)
}

func TestWorkerServiceDelete(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	err := service.DeleteByID(emptyString)
	require.Equal(t, createInvalidParameterError(operationDeleteByID, parameterID), err)

	err = service.DeleteByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(operationDeleteByID, parameterID), err)

	id := getRandomName()
	err = service.DeleteByID(id)
	require.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
}

func TestWorkerServiceDeleteAll(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	workers, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workers)

	for _, worker := range workers {
		err = DeleteTestWorker(t, service, worker)
		require.NoError(t, err)
	}
}

func TestWorkerServiceDiscoverWorker(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	workers, err := service.DiscoverWorker()
	require.NoError(t, err)
	require.NotNil(t, workers)

	for _, worker := range workers {
		t.Log(worker)
	}
}

func TestWorkerServiceGetAll(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	// create 30 test workers (to be deleted)
	for i := 0; i < 30; i++ {
		worker := CreateTestWorker(t, service)
		require.NotNil(t, worker)
	}

	workers, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workers)

	for _, worker := range workers {
		require.NotNil(t, worker)
		require.NotEmpty(t, worker.GetID())
		err = DeleteTestWorker(t, service, worker)
		require.NoError(t, err)
	}
}

func TestWorkerServiceGetByID(t *testing.T) {
	service := createWorkerService(t)
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

	workers, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workers)

	for _, worker := range workers {
		workerToCompare, err := service.GetByID(worker.GetID())
		require.NoError(t, err)
		IsEqualWorkers(t, worker, workerToCompare)
	}
}

func TestWorkerServiceGetByIDs(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	ids := []string{}
	resource, err := service.GetByIDs(ids)
	assert.NoError(t, err)
	assert.Equal(t, []*Worker{}, resource)
}

func TestWorkerServiceGetByName(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	workers, err := service.GetByName(emptyString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByName, parameterName))
	assert.NotNil(t, workers)

	workers, err = service.GetByName(whitespaceString)
	assert.Equal(t, err, createInvalidParameterError(operationGetByName, parameterName))
	assert.NotNil(t, workers)

	workers, err = service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workers)

	for _, worker := range workers {
		namedWorkers, err := service.GetByName(worker.Name)
		require.NoError(t, err)
		require.NotNil(t, namedWorkers)
	}
}

func TestWorkerServiceGetByPartialName(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	workers, err := service.GetByPartialName(emptyString)
	require.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	require.NotNil(t, workers)
	require.Len(t, workers, 0)

	workers, err = service.GetByPartialName(whitespaceString)
	require.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	require.NotNil(t, workers)
	require.Len(t, workers, 0)

	workers, err = service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workers)

	for _, worker := range workers {
		namedWorkers, err := service.GetByPartialName(worker.Name)
		require.NoError(t, err)
		require.NotNil(t, namedWorkers)
	}
}

func TestWorkerServiceGetWorkerOperatingSystems(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	workerOperatingSystems, err := service.GetWorkerOperatingSystems()
	require.NoError(t, err)
	require.NotNil(t, workerOperatingSystems)

	for _, workerOperatingSystem := range workerOperatingSystems {
		t.Log(workerOperatingSystem)
	}
}

func TestWorkerServiceGetWorkerShells(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	workerShells, err := service.GetWorkerShells()
	require.NoError(t, err)
	require.NotNil(t, workerShells)

	for _, workerShell := range workerShells {
		t.Log(workerShell)
	}
}

func TestWorkerServiceNew(t *testing.T) {
	serviceFunction := newWorkerService
	client := &sling.Sling{}
	uriTemplate := emptyString
	discoverWorkerPath := emptyString
	operatingSystemsPath := emptyString
	shellsPath := emptyString
	serviceName := serviceWorkerService

	testCases := []struct {
		name                 string
		f                    func(*sling.Sling, string, string, string, string) *workerService
		client               *sling.Sling
		uriTemplate          string
		discoverWorkerPath   string
		operatingSystemsPath string
		shellsPath           string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate, discoverWorkerPath, operatingSystemsPath, shellsPath},
		{"EmptyURITemplate", serviceFunction, client, emptyString, discoverWorkerPath, operatingSystemsPath, shellsPath},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString, discoverWorkerPath, operatingSystemsPath, shellsPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate, tc.discoverWorkerPath, tc.operatingSystemsPath, tc.shellsPath)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestWorkerServiceUpdate(t *testing.T) {
	service := createWorkerService(t)
	require.NotNil(t, service)

	worker := CreateTestWorker(t, service)
	defer DeleteTestWorker(t, service, worker)

	newName := getRandomName()

	worker.Name = newName

	updatedWorker := UpdateWorker(t, service, worker)
	require.NotNil(t, updatedWorker)

	require.NotEmpty(t, updatedWorker.GetID())
	require.Equal(t, updatedWorker.GetID(), updatedWorker.GetID())
	require.Equal(t, newName, updatedWorker.Name)
}
