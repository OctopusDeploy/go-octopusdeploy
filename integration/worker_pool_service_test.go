package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualWorkerPools(t *testing.T, expected octopusdeploy.IWorkerPool, actual octopusdeploy.IWorkerPool) {
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

	// TODO: add more validation
}

func CreateTestDynamicWorkerPool(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IWorkerPool {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()
	workerType := octopusdeploy.WorkerTypeUbuntu1804

	dynamicWorkerPool := octopusdeploy.NewDynamicWorkerPool(name, workerType)
	require.NotNil(t, dynamicWorkerPool)
	require.NoError(t, dynamicWorkerPool.Validate())

	createdDynamicWorkerPool, err := client.WorkerPools.Add(dynamicWorkerPool)
	require.NoError(t, err)
	require.NotNil(t, createdDynamicWorkerPool)
	require.NotEmpty(t, createdDynamicWorkerPool.GetID())

	// verify the add operation was successful
	dynamicWorkerPoolToCompare, err := client.WorkerPools.GetByID(createdDynamicWorkerPool.GetID())
	require.NoError(t, err)
	require.NotNil(t, dynamicWorkerPoolToCompare)
	AssertEqualWorkerPools(t, createdDynamicWorkerPool, dynamicWorkerPoolToCompare)

	return createdDynamicWorkerPool
}

func CreateTestStaticWorkerPool(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IWorkerPool {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()

	staticWorkerPool := octopusdeploy.NewStaticWorkerPool(name)
	require.NotNil(t, staticWorkerPool)
	require.NoError(t, staticWorkerPool.Validate())

	createdStaticWorkerPool, err := client.WorkerPools.Add(staticWorkerPool)
	require.NoError(t, err)
	require.NotNil(t, createdStaticWorkerPool)
	require.NotEmpty(t, createdStaticWorkerPool.GetID())

	// verify the add operation was successful
	staticWorkerPoolToCompare, err := client.WorkerPools.GetByID(createdStaticWorkerPool.GetID())
	require.NoError(t, err)
	require.NotNil(t, staticWorkerPoolToCompare)
	AssertEqualWorkerPools(t, createdStaticWorkerPool, staticWorkerPoolToCompare)

	return createdStaticWorkerPool
}

func DeleteTestWorkerPool(t *testing.T, client *octopusdeploy.Client, workerPool octopusdeploy.IWorkerPool) {
	require.NotNil(t, workerPool)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.WorkerPools.DeleteByID(workerPool.GetID())
	require.NoError(t, err)

	// verify the delete operation was successful
	deletedWorkerPool, err := client.WorkerPools.GetByID(workerPool.GetID())
	require.Error(t, err)
	require.Nil(t, deletedWorkerPool)
}

func UpdateWorkerPool(t *testing.T, client *octopusdeploy.Client, workerPool octopusdeploy.IWorkerPool) octopusdeploy.IWorkerPool {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	updatedWorkerPool, err := client.WorkerPools.Update(workerPool)
	require.NoError(t, err)
	require.NotNil(t, updatedWorkerPool)

	// verify the update operation was successful
	workerPoolToCompare, err := client.WorkerPools.GetByID(updatedWorkerPool.GetID())
	require.NoError(t, err)
	require.NotNil(t, workerPoolToCompare)
	AssertEqualWorkerPools(t, updatedWorkerPool, workerPoolToCompare)

	return updatedWorkerPool
}

func TestWorkerPoolServiceAdd(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	dynamicWorkerPool := CreateTestDynamicWorkerPool(t, client)
	require.NotNil(t, dynamicWorkerPool)
	defer DeleteTestWorkerPool(t, client, dynamicWorkerPool)

	staticWorkerPool := CreateTestStaticWorkerPool(t, client)
	require.NotNil(t, staticWorkerPool)
	defer DeleteTestWorkerPool(t, client, staticWorkerPool)
}

func TestWorkerPoolServiceGet(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	workerPools, err := client.WorkerPools.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)

	for _, workerPool := range workerPools {
		name := workerPool.GetName()
		query := octopusdeploy.WorkerPoolsQuery{
			PartialName: name,
			Take:        1,
		}
		namedWorkerPools, err := client.WorkerPools.Get(query)
		require.NoError(t, err)
		require.NotNil(t, namedWorkerPools)
		AssertEqualWorkerPools(t, workerPool, namedWorkerPools.Items[0])

		query = octopusdeploy.WorkerPoolsQuery{
			IDs:  []string{workerPool.GetID()},
			Take: 1,
		}
		namedWorkerPools, err = client.WorkerPools.Get(query)
		require.NoError(t, err)
		require.NotNil(t, namedWorkerPools)
		AssertEqualWorkerPools(t, workerPool, namedWorkerPools.Items[0])
	}
}

func TestWorkerPoolServiceCRUD(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	dynamicWorkerPool := CreateTestDynamicWorkerPool(t, client)
	require.NotNil(t, dynamicWorkerPool)
	defer DeleteTestWorkerPool(t, client, dynamicWorkerPool)

	dynamicWorkerPoolToCompare, err := client.WorkerPools.GetByID(dynamicWorkerPool.GetID())
	require.NoError(t, err)
	require.NotNil(t, dynamicWorkerPoolToCompare)
	AssertEqualWorkerPools(t, dynamicWorkerPool, dynamicWorkerPoolToCompare)

	staticWorkerPool := CreateTestStaticWorkerPool(t, client)
	require.NotNil(t, staticWorkerPool)
	defer DeleteTestWorkerPool(t, client, staticWorkerPool)

	updatedName := getRandomName()

	staticWorkerPool.SetName(updatedName)

	updatedStaticWorkerPool := UpdateWorkerPool(t, client, staticWorkerPool)
	require.NotNil(t, updatedStaticWorkerPool)

	staticWorkerPoolToCompare, err := client.WorkerPools.GetByID(updatedStaticWorkerPool.GetID())
	require.NoError(t, err)
	require.NotNil(t, staticWorkerPoolToCompare)
	AssertEqualWorkerPools(t, updatedStaticWorkerPool, staticWorkerPoolToCompare)
}

func TestWorkerPoolServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	workerPools, err := client.WorkerPools.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)

	for _, workerPool := range workerPools {
		if !workerPool.GetIsDefault() {
			defer DeleteTestWorkerPool(t, client, workerPool)
		}
	}
}

func TestWorkerPoolServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// create 10 test dynamic worker pools (to be deleted)
	for i := 0; i < 10; i++ {
		dynamicWorkerPool := CreateTestDynamicWorkerPool(t, client)
		require.NotNil(t, dynamicWorkerPool)
		defer DeleteTestWorkerPool(t, client, dynamicWorkerPool)
	}

	workerPools, err := client.WorkerPools.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)
	require.True(t, len(workerPools) > 10)

	for _, workerPool := range workerPools {
		require.NotNil(t, workerPool)
		require.NotEmpty(t, workerPool.GetID())
	}

	// create 10 test static worker pools (to be deleted)
	for i := 0; i < 10; i++ {
		staticWorkerPool := CreateTestStaticWorkerPool(t, client)
		require.NotNil(t, staticWorkerPool)
		defer DeleteTestWorkerPool(t, client, staticWorkerPool)
	}

	workerPools, err = client.WorkerPools.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)
	require.True(t, len(workerPools) > 10)

	for _, workerPool := range workerPools {
		require.NotNil(t, workerPool)
		require.NotEmpty(t, workerPool.GetID())
	}
}
