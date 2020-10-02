package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLifecycleAddAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	lifecycleName := getRandomName()

	expected, err := getTestLifecycle(lifecycleName)
	assert.NoError(t, err)

	actual := createTestLifecycle(t, octopusClient, lifecycleName)

	defer cleanLifecycle(t, octopusClient, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "lifecycle name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "lifecycle doesn't contain an ID from the octopus server")
}

func TestLifecycleAddGetAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	lifecycle := createTestLifecycle(t, octopusClient, getRandomName())
	defer cleanLifecycle(t, octopusClient, lifecycle.ID)

	getLifecycle, err := octopusClient.Lifecycles.GetByID(lifecycle.ID)
	assert.NoError(t, err, "there was an error raised getting lifecycle when there should not be")
	assert.Equal(t, lifecycle.Name, getLifecycle.Name)
}

func TestLifecycleGetThatDoesNotExist(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	id := getRandomName()
	resource, err := octopusClient.Lifecycles.GetByID(id)
	require.Equal(t, createResourceNotFoundError("lifecycle", "ID", id), err)
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
		defer cleanLifecycle(t, octopusClient, lifecycle.ID)
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
	defer cleanLifecycle(t, octopusClient, additionalLifecycle.ID)

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
	defer cleanLifecycle(t, octopusClient, lifecycle.ID)

	newLifecycleName := getRandomName()
	const newDescription = "this should be updated"
	// const newSkipMachineBehavior = "SkipUnavailableMachines"

	lifecycle.Name = newLifecycleName
	lifecycle.Description = newDescription

	updatedLifecycle, err := octopusClient.Lifecycles.Update(lifecycle)
	assert.NoError(t, err, "error when updating lifecycle")

	if err != nil {
		return
	}

	assert.Equal(t, newLifecycleName, updatedLifecycle.Name, "lifecycle name was not updated")
	assert.Equal(t, newDescription, updatedLifecycle.Description, "lifecycle description was not updated")
}

func TestLifecycleGetByPartialName(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	createdLifecycle := createTestLifecycle(t, octopusClient, getRandomName())
	defer cleanLifecycle(t, octopusClient, createdLifecycle.ID)

	lifecycleList, err := octopusClient.Lifecycles.GetByPartialName(createdLifecycle.Name)
	assert.NoError(t, err, "error when looking for lifecycle when not expected")

	if err != nil {
		return
	}

	for _, lifecycle := range lifecycleList {
		if lifecycle.Name == createdLifecycle.Name {
			return
		}
	}

	t.Errorf("lifecycle not found when searching by its name (%s)", createdLifecycle.Name)
}

func createTestLifecycle(t *testing.T, octopusClient *client.Client, lifecycleName string) model.Lifecycle {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	p, err := getTestLifecycle(lifecycleName)
	assert.NoError(t, err)

	createdLifecycle, err := octopusClient.Lifecycles.Add(p)

	if err != nil {
		t.Fatalf("creating lifecycle %s failed when it shouldn't: %s", lifecycleName, err)
	}

	return *createdLifecycle
}

func getTestLifecycle(name string) (*model.Lifecycle, error) {
	return model.NewLifecycle(name)
}

func cleanLifecycle(t *testing.T, octopusClient *client.Client, lifecycleID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	err := octopusClient.Lifecycles.DeleteByID(lifecycleID)
	assert.NoError(t, err)
}
