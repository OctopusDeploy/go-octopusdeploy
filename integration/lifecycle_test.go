package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go-octopusdeploy/octopusdeploy"
)

func init() {
	client = initTest()
}

func TestLifecycleAddAndDelete(t *testing.T) {
	lifecycleName := getRandomName()
	expected := getTestLifecycle(lifecycleName)
	actual := createTestLifecycle(t, lifecycleName)

	defer cleanLifecycle(t, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "lifecycle name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "lifecycle doesn't contain an ID from the octopus server")
}

func TestLifecycleAddGetAndDelete(t *testing.T) {
	lifecycle := createTestLifecycle(t, getRandomName())
	defer cleanLifecycle(t, lifecycle.ID)

	getLifecycle, err := client.Lifecycle.Get(lifecycle.ID)
	assert.Nil(t, err, "there was an error raised getting lifecycle when there should not be")
	assert.Equal(t, lifecycle.Name, getLifecycle.Name)
}

func TestLifecycleGetThatDoesNotExist(t *testing.T) {
	lifecycleID := "there-is-no-way-this-lifecycle-id-exists-i-hope"
	expected := octopusdeploy.ErrItemNotFound
	lifecycle, err := client.Lifecycle.Get(lifecycleID)

	assert.Error(t, err, "there should have been an error raised as this lifecycle should not be found")
	assert.Equal(t, expected, err, "a item not found error should have been raised")
	assert.Nil(t, lifecycle, "no lifecycle should have been returned")
}

func TestLifecycleGetAll(t *testing.T) {
	// create many lifecycles to test pagination
	lifecyclesToCreate := 32
	sum := 0
	for i := 0; i < lifecyclesToCreate; i++ {
		lifecycle := createTestLifecycle(t, getRandomName())
		defer cleanLifecycle(t, lifecycle.ID)
		sum += i
	}

	allLifecycles, err := client.Lifecycle.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all lifecycles failed when it shouldn't: %s", err)
	}

	numberOfLifecycles := len(*allLifecycles)

	// check there are greater than or equal to the amount of lifecycles requested to be created, otherwise pagination isn't working
	if numberOfLifecycles < lifecyclesToCreate {
		t.Fatalf("There should be at least %d lifecycles created but there was only %d. Pagination is likely not working.", lifecyclesToCreate, numberOfLifecycles)
	}

	additionalLifecycle := createTestLifecycle(t, getRandomName())
	defer cleanLifecycle(t, additionalLifecycle.ID)

	allLifecyclesAfterCreatingAdditional, err := client.Lifecycle.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all lifecycles failed when it shouldn't: %s", err)
	}

	assert.Nil(t, err, "error when looking for lifecycle when not expected")
	assert.Equal(t, len(*allLifecyclesAfterCreatingAdditional), numberOfLifecycles+1, "created an additional lifecycle and expected number of lifecycles to increase by 1")
}

func TestLifecycleUpdate(t *testing.T) {
	lifecycle := createTestLifecycle(t, getRandomName())
	defer cleanLifecycle(t, lifecycle.ID)

	newLifecycleName := getRandomName()
	const newDescription = "this should be updated"
	const newSkipMachineBehavior = "SkipUnavailableMachines"

	lifecycle.Name = newLifecycleName
	lifecycle.Description = newDescription

	updatedLifecycle, err := client.Lifecycle.Update(&lifecycle)
	assert.Nil(t, err, "error when updating lifecycle")
	assert.Equal(t, newLifecycleName, updatedLifecycle.Name, "lifecycle name was not updated")
	assert.Equal(t, newDescription, updatedLifecycle.Description, "lifecycle description was not updated")
}

func TestLifecycleGetByName(t *testing.T) {
	lifecycle := createTestLifecycle(t, getRandomName())
	defer cleanLifecycle(t, lifecycle.ID)

	foundLifecycle, err := client.Lifecycle.GetByName(lifecycle.Name)
	assert.Nil(t, err, "error when looking for lifecycle when not expected")
	assert.Equal(t, lifecycle.Name, foundLifecycle.Name, "lifecycle not found when searching by its name")
}

func createTestLifecycle(t *testing.T, lifecycleName string) octopusdeploy.Lifecycle {
	p := getTestLifecycle(lifecycleName)
	createdLifecycle, err := client.Lifecycle.Add(&p)

	if err != nil {
		t.Fatalf("creating lifecycle %s failed when it shouldn't: %s", lifecycleName, err)
	}

	return *createdLifecycle
}

func getTestLifecycle(lifecycleName string) octopusdeploy.Lifecycle {
	p := octopusdeploy.NewLifecycle(lifecycleName)

	return *p
}

func cleanLifecycle(t *testing.T, lifecycleID string) {
	err := client.Lifecycle.Delete(lifecycleID)

	if err == nil {
		return
	}

	if err == octopusdeploy.ErrItemNotFound {
		return
	}

	if err != nil {
		t.Fatalf("deleting lifecycle failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
