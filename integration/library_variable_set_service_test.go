package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLibraryVariableSetAddAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	name := getRandomName()
	expected := getTestLibraryVariableSet(name)
	actual := createTestLibraryVariableSet(t, octopusClient, name)

	defer cleanLibraryVariableSet(t, octopusClient, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "libraryVariableSet name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "libraryVariableSet doesn't contain an ID from the octopus server")
}

func TestLibraryVariableSetAddGetAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	libraryVariableSet := createTestLibraryVariableSet(t, octopusClient, getRandomName())
	defer cleanLibraryVariableSet(t, octopusClient, libraryVariableSet.ID)

	getLibraryVariableSet, err := octopusClient.LibraryVariableSets.GetByID(libraryVariableSet.ID)
	assert.NoError(t, err, "there was an error raised getting libraryVariableSet when there should not be")
	assert.Equal(t, libraryVariableSet.Name, getLibraryVariableSet.Name)
}

func TestLibraryVariableSetGetThatDoesNotExist(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	id := getRandomName()
	resource, err := octopusClient.LibraryVariableSets.GetByID(id)
	require.Equal(t, createResourceNotFoundError("library variable set", "ID", id), err)
	require.Nil(t, resource)
}

func TestLibraryVariableSetGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	// create many libraryVariableSets to test pagination
	libraryVariableSetsToCreate := 32
	sum := 0
	for i := 0; i < libraryVariableSetsToCreate; i++ {
		libraryVariableSet := createTestLibraryVariableSet(t, octopusClient, getRandomName())
		defer cleanLibraryVariableSet(t, octopusClient, libraryVariableSet.ID)
		sum += i
	}

	allLibraryVariableSets, err := octopusClient.LibraryVariableSets.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all libraryVariableSets failed when it shouldn't: %s", err)
	}

	numberOfLibraryVariableSets := len(allLibraryVariableSets)

	// check there are greater than or equal to the amount of libraryVariableSets requested to be created, otherwise pagination isn't working
	if numberOfLibraryVariableSets < libraryVariableSetsToCreate {
		t.Fatalf("There should be at least %d libraryVariableSets created but there was only %d. Pagination is likely not working.", libraryVariableSetsToCreate, numberOfLibraryVariableSets)
	}

	additionalLibraryVariableSet := createTestLibraryVariableSet(t, octopusClient, getRandomName())
	defer cleanLibraryVariableSet(t, octopusClient, additionalLibraryVariableSet.ID)

	allLibraryVariableSetsAfterCreatingAdditional, err := octopusClient.LibraryVariableSets.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all libraryVariableSets failed when it shouldn't: %s", err)
	}

	assert.NoError(t, err, "error when looking for libraryVariableSet when not expected")
	assert.Equal(t, len(allLibraryVariableSetsAfterCreatingAdditional), numberOfLibraryVariableSets+1, "created an additional libraryVariableSet and expected number of libraryVariableSets to increase by 1")
}

func TestLibraryVariableSetUpdate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	libraryVariableSet := createTestLibraryVariableSet(t, octopusClient, getRandomName())
	defer cleanLibraryVariableSet(t, octopusClient, libraryVariableSet.ID)

	newLibraryVariableSetName := getRandomName()
	const newDescription = "this should be updated"

	libraryVariableSet.Name = newLibraryVariableSetName
	libraryVariableSet.Description = newDescription

	updatedLibraryVariableSet, err := octopusClient.LibraryVariableSets.Update(libraryVariableSet)
	assert.NoError(t, err, "error when updating libraryVariableSet")
	assert.Equal(t, newLibraryVariableSetName, updatedLibraryVariableSet.Name, "libraryVariableSet name was not updated")
	assert.Equal(t, newDescription, updatedLibraryVariableSet.Description, "libraryVariableSet description was not updated")
}

func TestLibraryVariableSetGetByName(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	libraryVariableSet := createTestLibraryVariableSet(t, octopusClient, getRandomName())
	defer cleanLibraryVariableSet(t, octopusClient, libraryVariableSet.ID)

	foundLibraryVariableSets, err := octopusClient.LibraryVariableSets.GetByPartialName(libraryVariableSet.Name)
	assert.NoError(t, err, "error when looking for libraryVariableSet when not expected")
	assert.Equal(t, libraryVariableSet.Name, foundLibraryVariableSets[0].Name, "libraryVariableSet not found when searching by its name")
}

func createTestLibraryVariableSet(t *testing.T, octopusClient *client.Client, libraryVariableSetName string) model.LibraryVariableSet {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	p := getTestLibraryVariableSet(libraryVariableSetName)
	createdLibraryVariableSet, err := octopusClient.LibraryVariableSets.Add(&p)

	if err != nil {
		t.Fatalf("creating libraryVariableSet %s failed when it shouldn't: %s", libraryVariableSetName, err)
	}

	return *createdLibraryVariableSet
}

func getTestLibraryVariableSet(libraryVariableSetName string) model.LibraryVariableSet {
	p := model.NewLibraryVariableSet(libraryVariableSetName)

	return *p
}

func cleanLibraryVariableSet(t *testing.T, octopusClient *client.Client, libraryVariableSetID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	err := octopusClient.LibraryVariableSets.DeleteByID(libraryVariableSetID)
	assert.NoError(t, err)
}
