package integration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/service"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualLifecycles(t *testing.T, expected *service.Lifecycle, actual *service.Lifecycle) {
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

	// Lifecycle
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.ReleaseRetentionPolicy, actual.ReleaseRetentionPolicy)
	assert.Equal(t, expected.TentacleRetentionPolicy, actual.TentacleRetentionPolicy)
	assert.Equal(t, expected.Phases, actual.Phases)
}

func CreateTestLifecycle(t *testing.T, client *octopusdeploy.client) *service.Lifecycle {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()

	lifecycle := service.NewLifecycle(name)
	require.NotNil(t, lifecycle)

	createdLifecycle, err := client.Lifecycles.Add(lifecycle)
	require.NoError(t, err)
	require.NotNil(t, createdLifecycle)
	require.NotEmpty(t, createdLifecycle.GetID())

	// verify the add operation was successful
	lifecycleToCompare, err := client.Lifecycles.GetByID(createdLifecycle.GetID())
	require.NoError(t, err)
	require.NotNil(t, lifecycleToCompare)
	AssertEqualLifecycles(t, createdLifecycle, lifecycleToCompare)

	return createdLifecycle
}

func DeleteTestLifecycle(t *testing.T, client *octopusdeploy.client, lifecycle *service.Lifecycle) {
	require.NotNil(t, lifecycle)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Lifecycles.DeleteByID(lifecycle.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedLifecycle, err := client.Lifecycles.GetByID(lifecycle.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedLifecycle)
}

func TestLifecycleAddAndDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	lifecycleName := getRandomName()

	expected := getTestLifecycle(lifecycleName)
	require.NotNil(t, expected)

	actual := createTestLifecycle(t, client, lifecycleName)

	defer cleanLifecycle(t, client, actual.GetID())

	assert.Equal(t, expected.Name, actual.Name, "lifecycle name doesn't match expected")
	assert.NotEmpty(t, actual.GetID(), "lifecycle doesn't contain an ID from the octopus server")
}

func TestLifecycleAddGetAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	lifecycle := createTestLifecycle(t, octopusClient, getRandomName())
	defer cleanLifecycle(t, octopusClient, lifecycle.GetID())

	getLifecycle, err := octopusClient.Lifecycles.GetByID(lifecycle.GetID())
	assert.NoError(t, err, "there was an error raised getting lifecycle when there should not be")
	assert.Equal(t, lifecycle.Name, getLifecycle.Name)
}

func TestLifecycleServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	lifecycles, err := client.Lifecycles.GetAll()
	require.NoError(t, err)
	require.NotNil(t, lifecycles)

	for _, lifecycle := range lifecycles {
		defer DeleteTestLifecycle(t, client, lifecycle)
	}
}

func TestLifecycleGetThatDoesNotExist(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	id := getRandomName()
	resource, err := octopusClient.Lifecycles.GetByID(id)
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestLifecycleGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	// create many lifecycles to test pagination
	lifecyclesToCreate := 32
	sum := 0
	for i := 0; i < lifecyclesToCreate; i++ {
		lifecycle := createTestLifecycle(t, octopusClient, getRandomName())
		defer cleanLifecycle(t, octopusClient, lifecycle.GetID())
		sum += i
	}

	allLifecycles, err := octopusClient.Lifecycles.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all lifecycles failed when it shouldn't: %s", err)
	}

	numberOfLifecycles := len(allLifecycles)

	// check there are greater than or equal to the amount of lifecycles requested to be created, otherwise pagination isn't working
	if numberOfLifecycles < lifecyclesToCreate {
		t.Fatalf("There should be at least %d lifecycles created but there was only %d. Pagination is likely not working.", lifecyclesToCreate, numberOfLifecycles)
	}

	additionalLifecycle := createTestLifecycle(t, octopusClient, getRandomName())
	defer cleanLifecycle(t, octopusClient, additionalLifecycle.GetID())

	allLifecyclesAfterCreatingAdditional, err := octopusClient.Lifecycles.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all lifecycles failed when it shouldn't: %s", err)
	}

	assert.NoError(t, err, "error when looking for lifecycle when not expected")
	assert.Equal(t, len(allLifecyclesAfterCreatingAdditional), numberOfLifecycles+1, "created an additional lifecycle and expected number of lifecycles to increase by 1")
}

func TestLifecycleUpdate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	lifecycle := createTestLifecycle(t, octopusClient, getRandomName())
	defer cleanLifecycle(t, octopusClient, lifecycle.GetID())

	newLifecycleName := getRandomName()
	const newDescription = "this should be updated"
	// const newSkipMachineBehavior = "SkipUnavailableMachines"

	lifecycle.Name = newLifecycleName
	lifecycle.Description = newDescription

	updatedLifecycle, err := octopusClient.Lifecycles.Update(lifecycle)
	require.NoError(t, err, "error when updating lifecycle")
	require.Equal(t, newLifecycleName, updatedLifecycle.Name, "lifecycle name was not updated")
	require.Equal(t, newDescription, updatedLifecycle.Description, "lifecycle description was not updated")
}

func TestLifecycleGetByPartialName(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	createdLifecycle := createTestLifecycle(t, octopusClient, getRandomName())
	defer cleanLifecycle(t, octopusClient, createdLifecycle.GetID())

	lifecycleList, err := octopusClient.Lifecycles.GetByPartialName(createdLifecycle.Name)
	require.NoError(t, err, "error when looking for lifecycle when not expected")

	for _, lifecycle := range lifecycleList {
		if lifecycle.Name == createdLifecycle.Name {
			return
		}
	}

	t.Errorf("lifecycle not found when searching by its name (%s)", createdLifecycle.Name)
}

func createTestLifecycle(t *testing.T, octopusClient *octopusdeploy.client, lifecycleName string) *service.Lifecycle {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	p := getTestLifecycle(lifecycleName)
	require.NotNil(t, p)

	createdLifecycle, err := octopusClient.Lifecycles.Add(p)
	require.NoError(t, err)

	return createdLifecycle
}

func getTestLifecycle(name string) *service.Lifecycle {
	return service.NewLifecycle(name)
}

func cleanLifecycle(t *testing.T, octopusClient *octopusdeploy.client, lifecycleID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	err := octopusClient.Lifecycles.DeleteByID(lifecycleID)
	assert.NoError(t, err)
}
