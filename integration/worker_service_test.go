package integration

import (
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestWorker(t *testing.T, client *octopusdeploy.Client, workerPool octopusdeploy.IWorkerPool) (*octopusdeploy.Worker, error) {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()
	thumbprint := getRandomName()

	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)
	require.NotNil(t, url)

	listeningTentacleEndpoint := octopusdeploy.NewListeningTentacleEndpoint(url, thumbprint)
	require.NotNil(t, listeningTentacleEndpoint)

	worker := octopusdeploy.NewWorker(name, listeningTentacleEndpoint)
	worker.WorkerPoolIDs = append(worker.WorkerPoolIDs, workerPool.GetID())
	require.NotNil(t, worker)
	require.NoError(t, worker.Validate())

	createdWorker, err := client.Workers.Add(worker)
	require.NoError(t, err)
	require.NotNil(t, createdWorker)
	require.NotEmpty(t, createdWorker.GetID())

	// verify the add operation was successful
	workerToCompare, err := client.Workers.GetByID(createdWorker.GetID())
	require.NoError(t, err)
	require.NotNil(t, workerToCompare)
	IsEqualWorkers(t, createdWorker, workerToCompare)

	return createdWorker, nil
}

func DeleteTestWorker(t *testing.T, client *octopusdeploy.Client, worker *octopusdeploy.Worker) {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Workers.DeleteByID(worker.GetID())
	require.NoError(t, err)

	// verify the delete operation was successful
	deletedWorker, err := client.Workers.GetByID(worker.GetID())
	require.Error(t, err)
	require.Nil(t, deletedWorker)
}

func IsEqualWorkers(t *testing.T, expected *octopusdeploy.Worker, actual *octopusdeploy.Worker) {
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

	// Worker
	assert.Equal(t, expected.Name, actual.Name)
}

func UpdateWorker(t *testing.T, client *octopusdeploy.Client, worker *octopusdeploy.Worker) *octopusdeploy.Worker {
	require.NotNil(t, worker)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	updatedWorker, err := client.Workers.Update(worker)
	assert.NoError(t, err)
	require.NotNil(t, updatedWorker)

	// verify the update operation was successful
	workerToCompare, err := client.Workers.GetByID(updatedWorker.GetID())
	require.NoError(t, err)
	require.NotNil(t, workerToCompare)
	IsEqualWorkers(t, updatedWorker, workerToCompare)

	return updatedWorker
}

func TestWorkerServiceAdd(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	staticWorkerPool := CreateTestStaticWorkerPool(t, client)
	require.NotNil(t, staticWorkerPool)
	defer DeleteTestWorkerPool(t, client, staticWorkerPool)

	worker, err := CreateTestWorker(t, client, staticWorkerPool)
	require.NoError(t, err)
	require.NotNil(t, worker)
	defer DeleteTestWorker(t, client, worker)
}

func TestWorkerServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	staticWorkerPool := CreateTestStaticWorkerPool(t, client)
	require.NotNil(t, staticWorkerPool)
	defer DeleteTestWorkerPool(t, client, staticWorkerPool)

	worker, err := CreateTestWorker(t, client, staticWorkerPool)
	require.NoError(t, err)
	require.NotNil(t, worker)
	defer DeleteTestWorker(t, client, worker)

	for _, id := range worker.WorkerPoolIDs {
		workerPool, err := client.WorkerPools.GetByID(id)
		require.NoError(t, err)
		require.NotNil(t, workerPool)
	}
}

func TestWorkerServiceDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := getRandomName()
	err := client.Workers.DeleteByID(id)
	require.Error(t, err)
}

func TestWorkerServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	workers, err := client.Workers.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workers)

	for _, worker := range workers {
		defer DeleteTestWorker(t, client, worker)
	}
}

// TODO: validation of DiscoverWorker is required

// func TestWorkerServiceDiscoverWorker(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	workers, err := client.Workers.DiscoverWorker()
// 	require.NoError(t, err)
// 	require.NotNil(t, workers)
// }

// TODO: fix test
// func TestWorkerServiceGetAll(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	staticWorkerPool := CreateTestStaticWorkerPool(t, client)
// 	require.NotNil(t, staticWorkerPool)
// 	defer DeleteTestWorkerPool(t, client, staticWorkerPool)

// 	// create 2 test workers (to be deleted)
// 	for i := 0; i < 2; i++ {
// 		worker, err := CreateTestWorker(t, client, staticWorkerPool)
// 		require.NoError(t, err)
// 		require.NotNil(t, worker)
// 	}

// 	workers, err := client.Workers.GetAll()
// 	require.NoError(t, err)
// 	require.NotNil(t, workers)

// 	for _, worker := range workers {
// 		require.NotNil(t, worker)
// 		require.NotEmpty(t, worker.GetID())
// 		defer DeleteTestWorker(t, client, worker)
// 	}
// }

func TestWorkerServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := getRandomName()
	invalidWorker, err := client.Workers.GetByID(id)
	require.Error(t, err)
	require.Nil(t, invalidWorker)

	workers, err := client.Workers.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workers)

	for _, worker := range workers {
		workerToCompare, err := client.Workers.GetByID(worker.GetID())
		require.NoError(t, err)
		IsEqualWorkers(t, worker, workerToCompare)
	}
}

func TestWorkerServiceGetByIDs(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	ids := []string{}
	resource, err := client.Workers.GetByIDs(ids)
	assert.NoError(t, err)
	assert.Equal(t, []*octopusdeploy.Worker{}, resource)
}

func TestWorkerServiceGetByName(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	workers, err := client.Workers.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workers)

	for _, worker := range workers {
		namedWorkers, err := client.Workers.GetByName(worker.Name)
		require.NoError(t, err)
		require.NotNil(t, namedWorkers)
	}
}

func TestWorkerServiceGetByPartialName(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	workers, err := client.Workers.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workers)

	for _, worker := range workers {
		namedWorkers, err := client.Workers.GetByPartialName(worker.Name)
		require.NoError(t, err)
		require.NotNil(t, namedWorkers)
	}
}

func TestWorkerServiceGetWorkerOperatingSystems(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	workerOperatingSystems, err := client.Workers.GetWorkerOperatingSystems()
	require.NoError(t, err)
	require.NotNil(t, workerOperatingSystems)

	for _, workerOperatingSystem := range workerOperatingSystems {
		t.Log(workerOperatingSystem)
	}
}

func TestWorkerServiceGetWorkerShells(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	workerShells, err := client.Workers.GetWorkerShells()
	require.NoError(t, err)
	require.NotNil(t, workerShells)

	for _, workerShell := range workerShells {
		t.Log(workerShell)
	}
}

// TODO: fix test
// func TestWorkerServiceUpdate(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	staticWorkerPool := CreateTestStaticWorkerPool(t, client)
// 	require.NotNil(t, staticWorkerPool)
// 	defer DeleteTestWorkerPool(t, client, staticWorkerPool)

// 	worker, err := CreateTestWorker(t, client, staticWorkerPool)
// 	require.NoError(t, err)
// 	require.NotNil(t, worker)
// 	defer DeleteTestWorker(t, client, worker)

// 	newName := getRandomName()

// 	worker.Name = newName

// 	updatedWorker := UpdateWorker(t, client, worker)
// 	require.NotNil(t, updatedWorker)

// 	require.NotEmpty(t, updatedWorker.GetID())
// 	require.Equal(t, updatedWorker.GetID(), updatedWorker.GetID())
// 	require.Equal(t, newName, updatedWorker.Name)
// }
