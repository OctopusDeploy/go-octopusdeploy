package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/workerpools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualWorkerPools(t *testing.T, expected workerpools.IWorkerPool, actual workerpools.IWorkerPool) {
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
	assert.True(t, internal.IsLinksEqual(expected.GetLinks(), actual.GetLinks()))

	// TODO: add more validation
}

func CreateTestStaticWorkerPool(t *testing.T, client *client.Client) workerpools.IWorkerPool {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	staticWorkerPool := workerpools.NewStaticWorkerPool(name)
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

func DeleteTestWorkerPool(t *testing.T, client *client.Client, workerPoolId string) {

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.WorkerPools.DeleteByID(workerPoolId)
	require.NoError(t, err)

	// verify the delete operation was successful
	deletedWorkerPool, err := client.WorkerPools.GetByID(workerPoolId)
	require.Error(t, err)
	require.Nil(t, deletedWorkerPool)
}

func UpdateWorkerPool(t *testing.T, client *client.Client, workerPool workerpools.IWorkerPool) workerpools.IWorkerPool {
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

	staticWorkerPool := CreateTestStaticWorkerPool(t, client)
	require.NotNil(t, staticWorkerPool)
	defer DeleteTestWorkerPool(t, client, staticWorkerPool.GetID())
}

func TestWorkerPoolServiceGet(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	workerPools, err := client.WorkerPools.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)

	for _, workerPool := range workerPools {
		name := workerPool.Name
		query := workerpools.WorkerPoolsQuery{
			PartialName: name,
			Take:        1,
		}
		namedWorkerPools, err := client.WorkerPools.Get(query)
		require.NoError(t, err)
		require.NotNil(t, namedWorkerPools)

		wp, err := client.WorkerPools.GetByID(workerPool.ID)
		AssertEqualWorkerPools(t, wp, namedWorkerPools.Items[0])

		query = workerpools.WorkerPoolsQuery{
			IDs:  []string{workerPool.ID},
			Take: 1,
		}
		namedWorkerPools, err = client.WorkerPools.Get(query)
		require.NoError(t, err)
		require.NotNil(t, namedWorkerPools)
		wp, err = client.WorkerPools.GetByID(workerPool.ID)
		AssertEqualWorkerPools(t, wp, namedWorkerPools.Items[0])
	}
}

func TestWorkerPoolServiceCRUD(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	staticWorkerPool := CreateTestStaticWorkerPool(t, client)
	require.NotNil(t, staticWorkerPool)
	defer DeleteTestWorkerPool(t, client, staticWorkerPool.GetID())

	updatedName := internal.GetRandomName()

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
		if !workerPool.IsDefault {
			defer DeleteTestWorkerPool(t, client, workerPool.ID)
		}
	}
}

func TestWorkerPoolServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// create 10 test static worker pools (to be deleted)
	for i := 0; i < 10; i++ {
		staticWorkerPool := CreateTestStaticWorkerPool(t, client)
		require.NotNil(t, staticWorkerPool)
		defer DeleteTestWorkerPool(t, client, staticWorkerPool.GetID())
	}

	workerPools, err := client.WorkerPools.GetAll()
	require.NoError(t, err)
	require.NotNil(t, workerPools)
	require.True(t, len(workerPools) > 10)

	for _, workerPool := range workerPools {
		require.NotNil(t, workerPool)
		require.NotEmpty(t, workerPool.ID)
	}
}

func TestWorkerPoolServiceGetDynamicWorkerTypes(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	types, err := client.WorkerPools.GetDynamicWorkerTypes()
	require.NoError(t, err)
	require.NotNil(t, types)
}

// === NEW ===

func TestWorkerPoolServiceCRUD_NewClient(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	staticWorkerPool := CreateTestStaticWorkerPool_NewClient(t, client)
	require.NotNil(t, staticWorkerPool)
	defer DeleteTestWorkerPool_NewClient(t, client, staticWorkerPool)

	updatedName := internal.GetRandomName()

	staticWorkerPool.SetName(updatedName)

	updatedStaticWorkerPool := UpdateWorkerPool(t, client, staticWorkerPool)
	require.NotNil(t, updatedStaticWorkerPool)

	staticWorkerPoolToCompare, err := workerpools.GetByID(client, updatedStaticWorkerPool.GetSpaceID(), updatedStaticWorkerPool.GetID())
	require.NoError(t, err)
	require.NotNil(t, staticWorkerPoolToCompare)
	AssertEqualWorkerPools(t, updatedStaticWorkerPool, staticWorkerPoolToCompare)
}

func CreateTestStaticWorkerPool_NewClient(t *testing.T, client *client.Client) workerpools.IWorkerPool {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	staticWorkerPool := workerpools.NewStaticWorkerPool(name)
	require.NotNil(t, staticWorkerPool)
	require.NoError(t, staticWorkerPool.Validate())

	createdStaticWorkerPool, err := workerpools.Add(client, staticWorkerPool)
	require.NoError(t, err)
	require.NotNil(t, createdStaticWorkerPool)
	require.NotEmpty(t, createdStaticWorkerPool.GetID())

	// verify the add operation was successful
	staticWorkerPoolToCompare, err := workerpools.GetByID(client, createdStaticWorkerPool.GetSpaceID(), createdStaticWorkerPool.GetID())
	require.NoError(t, err)
	require.NotNil(t, staticWorkerPoolToCompare)
	AssertEqualWorkerPools(t, createdStaticWorkerPool, staticWorkerPoolToCompare)

	return createdStaticWorkerPool
}

func DeleteTestWorkerPool_NewClient(t *testing.T, client *client.Client, workerPool workerpools.IWorkerPool) {

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := workerpools.DeleteByID(client, workerPool.GetSpaceID(), workerPool.GetID())
	require.NoError(t, err)

	// verify the delete operation was successful
	deletedWorkerPool, err := workerpools.GetByID(client, workerPool.GetSpaceID(), workerPool.GetID())
	require.Error(t, err)
	require.Nil(t, deletedWorkerPool)
}
