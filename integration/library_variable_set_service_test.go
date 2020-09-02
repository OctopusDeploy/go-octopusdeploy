package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
)

func init() {
	if octopusClient == nil {
		octopusClient = initTest()
	}
}

func TestLibraryVariableSetAddAndDelete(t *testing.T) {
	name := getRandomName()
	expected := getTestLibraryVariableSet(name)
	actual := createTestLibraryVariableSet(t, name)

	defer cleanLibraryVariableSet(t, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "libraryVariableSet name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "libraryVariableSet doesn't contain an ID from the octopus server")
}

func TestLibraryVariableSetAddGetAndDelete(t *testing.T) {
	libraryVariableSet := createTestLibraryVariableSet(t, getRandomName())
	defer cleanLibraryVariableSet(t, libraryVariableSet.ID)

	getLibraryVariableSet, err := octopusClient.LibraryVariableSets.Get(libraryVariableSet.ID)
	assert.Nil(t, err, "there was an error raised getting libraryVariableSet when there should not be")
	assert.Equal(t, libraryVariableSet.Name, getLibraryVariableSet.Name)
}

func TestLibraryVariableSetGetThatDoesNotExist(t *testing.T) {
	libraryVariableSetID := "there-is-no-way-this-libraryVariableSet-id-exists-i-hope"
	expected := client.ErrItemNotFound
	libraryVariableSet, err := octopusClient.LibraryVariableSets.Get(libraryVariableSetID)

	assert.Error(t, err, "there should have been an error raised as this libraryVariableSet should not be found")
	assert.Equal(t, expected, err, "a item not found error should have been raised")
	assert.Nil(t, libraryVariableSet, "no libraryVariableSet should have been returned")
}

func TestLibraryVariableSetGetAll(t *testing.T) {
	// create many libraryVariableSets to test pagination
	libraryVariableSetsToCreate := 32
	sum := 0
	for i := 0; i < libraryVariableSetsToCreate; i++ {
		libraryVariableSet := createTestLibraryVariableSet(t, getRandomName())
		defer cleanLibraryVariableSet(t, libraryVariableSet.ID)
		sum += i
	}

	allLibraryVariableSets, err := octopusClient.LibraryVariableSets.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all libraryVariableSets failed when it shouldn't: %s", err)
	}

	numberOfLibraryVariableSets := len(*allLibraryVariableSets)

	// check there are greater than or equal to the amount of libraryVariableSets requested to be created, otherwise pagination isn't working
	if numberOfLibraryVariableSets < libraryVariableSetsToCreate {
		t.Fatalf("There should be at least %d libraryVariableSets created but there was only %d. Pagination is likely not working.", libraryVariableSetsToCreate, numberOfLibraryVariableSets)
	}

	additionalLibraryVariableSet := createTestLibraryVariableSet(t, getRandomName())
	defer cleanLibraryVariableSet(t, additionalLibraryVariableSet.ID)

	allLibraryVariableSetsAfterCreatingAdditional, err := octopusClient.LibraryVariableSets.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all libraryVariableSets failed when it shouldn't: %s", err)
	}

	assert.Nil(t, err, "error when looking for libraryVariableSet when not expected")
	assert.Equal(t, len(*allLibraryVariableSetsAfterCreatingAdditional), numberOfLibraryVariableSets+1, "created an additional libraryVariableSet and expected number of libraryVariableSets to increase by 1")
}

func TestLibraryVariableSetUpdate(t *testing.T) {
	libraryVariableSet := createTestLibraryVariableSet(t, getRandomName())
	defer cleanLibraryVariableSet(t, libraryVariableSet.ID)

	newLibraryVariableSetName := getRandomName()
	const newDescription = "this should be updated"

	libraryVariableSet.Name = newLibraryVariableSetName
	libraryVariableSet.Description = newDescription

	updatedLibraryVariableSet, err := octopusClient.LibraryVariableSets.Update(&libraryVariableSet)
	assert.Nil(t, err, "error when updating libraryVariableSet")
	assert.Equal(t, newLibraryVariableSetName, updatedLibraryVariableSet.Name, "libraryVariableSet name was not updated")
	assert.Equal(t, newDescription, updatedLibraryVariableSet.Description, "libraryVariableSet description was not updated")
}

func TestLibraryVariableSetGetByName(t *testing.T) {
	libraryVariableSet := createTestLibraryVariableSet(t, getRandomName())
	defer cleanLibraryVariableSet(t, libraryVariableSet.ID)

	foundLibraryVariableSet, err := octopusClient.LibraryVariableSets.GetByName(libraryVariableSet.Name)
	assert.Nil(t, err, "error when looking for libraryVariableSet when not expected")
	assert.Equal(t, libraryVariableSet.Name, foundLibraryVariableSet.Name, "libraryVariableSet not found when searching by its name")
}

func createTestLibraryVariableSet(t *testing.T, libraryVariableSetName string) model.LibraryVariableSet {
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

func cleanLibraryVariableSet(t *testing.T, libraryVariableSetID string) {
	err := octopusClient.LibraryVariableSets.Delete(libraryVariableSetID)

	if err == nil {
		return
	}

	if err == client.ErrItemNotFound {
		return
	}

	if err != nil {
		t.Fatalf("deleting libraryVariableSet failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
