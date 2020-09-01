package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

func init() {
	octopusClient = initTest()
}

func TestEnvironmentAddAndDelete(t *testing.T) {
	environmentName := getRandomName()
	expected := getTestEnvironment(environmentName)
	actual := createTestEnvironment(t, environmentName)

	defer cleanEnvironment(t, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "environment name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "environment doesn't contain an ID from the octopus server")
}

func TestEnvironmentAddGetAndDelete(t *testing.T) {
	environment := createTestEnvironment(t, getRandomName())
	defer cleanEnvironment(t, environment.ID)

	getEnvironment, err := octopusClient.Environments.Get(environment.ID)
	assert.Nil(t, err, "there was an error raised getting environment when there should not be")
	assert.Equal(t, environment.Name, getEnvironment.Name)
}

func TestEnvironmentGetThatDoesNotExist(t *testing.T) {
	environmentID := "there-is-no-way-this-environment-id-exists-i-hope"
	expected := client.ErrItemNotFound
	environment, err := octopusClient.Environments.Get(environmentID)

	assert.Error(t, err, "there should have been an error raised as this environment should not be found")
	assert.Equal(t, expected, err, "a item not found error should have been raised")
	assert.Nil(t, environment, "no environment should have been returned")
}

func TestEnvironmentGetAll(t *testing.T) {
	// create many environments to test pagination
	environmentsToCreate := 32
	sum := 0
	for i := 0; i < environmentsToCreate; i++ {
		environment := createTestEnvironment(t, getRandomName())
		defer cleanEnvironment(t, environment.ID)
		sum += i
	}

	allEnvironments, err := octopusClient.Environments.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all environments failed when it shouldn't: %s", err)
	}

	numberOfEnvironments := len(*allEnvironments)

	// check there are greater than or equal to the amount of environments requested to be created, otherwise pagination isn't working
	if numberOfEnvironments < environmentsToCreate {
		t.Fatalf("There should be at least %d environments created but there was only %d. Pagination is likely not working.", environmentsToCreate, numberOfEnvironments)
	}

	additionalEnvironment := createTestEnvironment(t, getRandomName())
	defer cleanEnvironment(t, additionalEnvironment.ID)

	allEnvironmentsAfterCreatingAdditional, err := octopusClient.Environments.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all environments failed when it shouldn't: %s", err)
	}

	assert.Nil(t, err, "error when looking for environment when not expected")
	assert.Equal(t, len(*allEnvironmentsAfterCreatingAdditional), numberOfEnvironments+1, "created an additional environment and expected number of environments to increase by 1")
}

func TestEnvironmentUpdate(t *testing.T) {
	environment := createTestEnvironment(t, getRandomName())
	defer cleanEnvironment(t, environment.ID)

	newEnvironmentName := getRandomName()
	const newDescription = "this should be updated"
	const newUseGuidedFailure = true

	environment.Name = newEnvironmentName
	environment.Description = newDescription
	environment.UseGuidedFailure = newUseGuidedFailure

	updatedEnvironment, err := octopusClient.Environments.Update(&environment)
	assert.Nil(t, err, "error when updating environment")
	assert.Equal(t, newEnvironmentName, updatedEnvironment.Name, "environment name was not updated")
	assert.Equal(t, newDescription, updatedEnvironment.Description, "environment description was not updated")
	assert.Equal(t, newUseGuidedFailure, environment.UseGuidedFailure, "environment UseGuidedFailure was not updated")
}

func TestEnvironmentGetByName(t *testing.T) {
	environment := createTestEnvironment(t, getRandomName())
	defer cleanEnvironment(t, environment.ID)

	foundEnvironment, err := octopusClient.Environments.GetByName(environment.Name)
	assert.Nil(t, err, "error when looking for environment when not expected")
	assert.Equal(t, environment.Name, foundEnvironment.Name, "environment not found when searching by its name")
}

/*
func TestEnvironmentCleanup(t *testing.T) {
	allEnvironments, err := octopusClient.Environments.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all environments failed when it shouldn't: %s", err)
	}
	for _, e := range *allEnvironments {
		if len(e.Name) > 17 && e.Name[0:16] == "go-octopusdeploy" {
			octopusClient.Environments.Delete(e.ID)
		}
	}
}
*/

func createTestEnvironment(t *testing.T, environmentName string) model.Environment {
	e := getTestEnvironment(environmentName)
	createdEnvironment, err := octopusClient.Environments.Add(&e)

	if err != nil {
		t.Fatalf("creating environment %s failed when it shouldn't: %s", environmentName, err)
	}

	return *createdEnvironment
}

func getTestEnvironment(environmentName string) model.Environment {
	e := model.NewEnvironment(environmentName, "Environment from testing suite", true)
	return *e
}

func cleanEnvironment(t *testing.T, environmentID string) {
	err := octopusClient.Environments.Delete(environmentID)

	if err == nil {
		return
	}

	if err == client.ErrItemNotFound {
		return
	}

	if err != nil {
		t.Fatalf("deleting environment failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
