package integration

import (
	// "fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"

	"github.com/stretchr/testify/assert"
)

func init() {
	client = initTest()
}

func TestProjectGroupAddAndDelete(t *testing.T) {
	projectGroupName := getRandomName()
	expected := getTestProjectGroup(projectGroupName)
	actual := createTestProjectGroup(t, projectGroupName)

	defer cleanProjectGroup(t, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "projectgroup name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "projectgroup doesn't contain an ID from the octopus server")
}

func TestProjectGroupAddGetAndDelete(t *testing.T) {
	projectGroup := createTestProjectGroup(t, getRandomName())
	defer cleanProjectGroup(t, projectGroup.ID)

	getProjectGroup, err := client.ProjectGroup.Get(projectGroup.ID)
	assert.Nil(t, err, "there was an error raised getting projectgroup when there should not be")
	assert.Equal(t, projectGroup.Name, getProjectGroup.Name)
}

func TestProjectGroupGetThatDoesNotExist(t *testing.T) {
	projectGroupID := "there-is-no-way-this-projectgroup-id-exists-i-hope"
	expected := octopusdeploy.ErrItemNotFound
	projectGroup, err := client.ProjectGroup.Get(projectGroupID)

	assert.Error(t, err, "there should have been an error raised as this projectgroup should not be found")
	assert.Equal(t, expected, err, "a item not found error should have been raised")
	assert.Nil(t, projectGroup, "no projectgroup should have been returned")
}

func TestProjectGroupGetAll(t *testing.T) {
	// create many projects to test pagination
	projectsGroupsToCreate := 32
	sum := 0
	for i := 0; i < projectsGroupsToCreate; i++ {
		projectGroup := createTestProjectGroup(t, getRandomName())
		defer cleanProjectGroup(t, projectGroup.ID)
		sum += i
	}

	allProjectGroups, err := client.ProjectGroup.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	numberOfProjectGroups := len(*allProjectGroups)

	// check there are greater than or equal to the amount of projects requested to be created, otherwise pagination isn't working
	if numberOfProjectGroups < projectsGroupsToCreate {
		t.Fatalf("There should be at least %d project groups created but there was only %d. Pagination is likely not working.", projectsGroupsToCreate, numberOfProjectGroups)
	}

	additionalProjectGroup := createTestProjectGroup(t, getRandomName())
	defer cleanProjectGroup(t, additionalProjectGroup.ID)

	allProjectGroupsAfterCreatingAdditional, err := client.ProjectGroup.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	assert.Nil(t, err, "error when looking for project when not expected")
	assert.Equal(t, len(*allProjectGroupsAfterCreatingAdditional), numberOfProjectGroups+1, "created an additional projectgroup and expected number of projectgroups to increase by 1")
}

func TestProjectGroupUpdate(t *testing.T) {
	projectGroup := createTestProjectGroup(t, getRandomName())
	defer cleanProjectGroup(t, projectGroup.ID)

	newProjectGroupName := getRandomName()
	const newDescription = "this should be updated"

	projectGroup.Name = newProjectGroupName
	projectGroup.Description = newDescription

	updatedProjectGroup, err := client.ProjectGroup.Update(&projectGroup)
	assert.Nil(t, err, "error when updating projectgroup")
	assert.Equal(t, newProjectGroupName, updatedProjectGroup.Name, "projectgroup name was not updated")
	assert.Equal(t, newDescription, updatedProjectGroup.Description, "projectgroup description was not updated")
}

func createTestProjectGroup(t *testing.T, projectGroupName string) octopusdeploy.ProjectGroup {
	p := getTestProjectGroup(projectGroupName)
	createdProjectGroup, err := client.ProjectGroup.Add(&p)

	if err != nil {
		t.Fatalf("creating projectgroup %s failed when it shouldn't: %s", projectGroupName, err)
	}

	return *createdProjectGroup
}

func getTestProjectGroup(projectGroupName string) octopusdeploy.ProjectGroup {
	p := octopusdeploy.NewProjectGroup(projectGroupName)

	return *p
}

func cleanProjectGroup(t *testing.T, projectGroupID string) {
	err := client.ProjectGroup.Delete(projectGroupID)

	if err == nil {
		return
	}

	if err == octopusdeploy.ErrItemNotFound {
		return
	}

	if err != nil {
		t.Fatalf("deleting projectgroup failed when it shouldn't. manual cleanup may be needed. (%s)", err.Error())
	}
}
