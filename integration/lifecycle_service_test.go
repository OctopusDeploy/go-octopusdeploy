package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

func TestLifecycleAddAndDelete(t *testing.T) {
	lifecycleName := getRandomName()

	expected, err := getTestLifecycle(lifecycleName)
	assert.NoError(t, err)

	actual := createTestLifecycle(t, lifecycleName)

	defer cleanLifecycle(t, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "lifecycle name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "lifecycle doesn't contain an ID from the octopus server")
}

func TestLifecycleAddGetAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()

	lifecycle := createTestLifecycle(t, getRandomName())
	defer cleanLifecycle(t, lifecycle.ID)

	getLifecycle, err := octopusClient.Lifecycles.GetByID(lifecycle.ID)
	assert.NoError(t, err, "there was an error raised getting lifecycle when there should not be")
	assert.Equal(t, lifecycle.Name, getLifecycle.Name)
}

func TestLifecycleGetThatDoesNotExist(t *testing.T) {
	octopusClient := getOctopusClient()

	lifecycleID := "there-is-no-way-this-lifecycle-id-exists-i-hope"
	expected := client.ErrItemNotFound
	lifecycle, err := octopusClient.Lifecycles.GetByID(lifecycleID)

	assert.Error(t, err, "there should have been an error raised as this lifecycle should not be found")
	assert.Equal(t, expected, err, "a item not found error should have been raised")
	assert.Nil(t, lifecycle, "no lifecycle should have been returned")
}

func TestLifecycleGetAll(t *testing.T) {
	octopusClient := getOctopusClient()

	// create many lifecycles to test pagination
	lifecyclesToCreate := 32
	sum := 0
	for i := 0; i < lifecyclesToCreate; i++ {
		lifecycle := createTestLifecycle(t, getRandomName())
		defer cleanLifecycle(t, lifecycle.ID)
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

	additionalLifecycle := createTestLifecycle(t, getRandomName())
	defer cleanLifecycle(t, additionalLifecycle.ID)

	allLifecyclesAfterCreatingAdditional, err := octopusClient.Lifecycles.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all lifecycles failed when it shouldn't: %s", err)
	}

	assert.NoError(t, err, "error when looking for lifecycle when not expected")
	assert.Equal(t, len(allLifecyclesAfterCreatingAdditional), numberOfLifecycles+1, "created an additional lifecycle and expected number of lifecycles to increase by 1")
}

func TestLifecycleUpdate(t *testing.T) {
	octopusClient := getOctopusClient()

	lifecycle := createTestLifecycle(t, getRandomName())
	defer cleanLifecycle(t, lifecycle.ID)

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

	createdLifecycle := createTestLifecycle(t, getRandomName())
	defer cleanLifecycle(t, createdLifecycle.ID)

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

func createTestLifecycle(t *testing.T, lifecycleName string) model.Lifecycle {
	octopusClient := getOctopusClient()

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

func cleanLifecycle(t *testing.T, lifecycleID string) {
	octopusClient := getOctopusClient()

	err := octopusClient.Lifecycles.DeleteByID(lifecycleID)

	if err == nil {
		return
	}

	if err == client.ErrItemNotFound {
		return
	}

	if err != nil {
		t.Fatalf("deleting lifecycle failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
