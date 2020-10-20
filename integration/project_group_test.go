package integration

import (
	// "fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectGroupAddAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	projectGroupName := getRandomName()
	expected := getTestProjectGroup(projectGroupName)
	actual := createTestProjectGroup(t, octopusClient, projectGroupName)

	defer cleanProjectGroup(t, octopusClient, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "projectgroup name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "projectgroup doesn't contain an ID from the octopus server")
}

func TestProjectGroupAddGetAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	projectGroup := createTestProjectGroup(t, octopusClient, getRandomName())
	defer cleanProjectGroup(t, octopusClient, projectGroup.ID)

	getProjectGroup, err := octopusClient.ProjectGroups.GetByID(projectGroup.ID)
	assert.NoError(t, err, "there was an error raised getting projectgroup when there should not be")
	assert.Equal(t, projectGroup.Name, getProjectGroup.Name)
}

func TestProjectGroupGetThatDoesNotExist(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	id := getRandomName()
	resource, err := octopusClient.ProjectGroups.GetByID(id)
	require.Equal(t, createResourceNotFoundError("ProjectGroupService", "ID", id), err)
	require.Nil(t, resource)
}

func TestProjectGroupGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	// create many projects to test pagination
	projectsGroupsToCreate := 32
	sum := 0
	for i := 0; i < projectsGroupsToCreate; i++ {
		projectGroup := createTestProjectGroup(t, octopusClient, getRandomName())
		defer cleanProjectGroup(t, octopusClient, projectGroup.ID)
		sum += i
	}

	allProjectGroups, err := octopusClient.ProjectGroups.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	numberOfProjectGroups := len(allProjectGroups)

	// check there are greater than or equal to the amount of projects requested to be created, otherwise pagination isn't working
	if numberOfProjectGroups < projectsGroupsToCreate {
		t.Fatalf("There should be at least %d project groups created but there was only %d. Pagination is likely not working.", projectsGroupsToCreate, numberOfProjectGroups)
	}

	additionalProjectGroup := createTestProjectGroup(t, octopusClient, getRandomName())
	defer cleanProjectGroup(t, octopusClient, additionalProjectGroup.ID)

	allProjectGroupsAfterCreatingAdditional, err := octopusClient.ProjectGroups.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	assert.NoError(t, err, "error when looking for project when not expected")
	assert.Equal(t, len(allProjectGroupsAfterCreatingAdditional), numberOfProjectGroups+1, "created an additional projectgroup and expected number of projectgroups to increase by 1")
}

func TestProjectGroupUpdate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	projectGroup := createTestProjectGroup(t, octopusClient, getRandomName())
	defer cleanProjectGroup(t, octopusClient, projectGroup.ID)

	newProjectGroupName := getRandomName()
	const newDescription = "this should be updated"

	projectGroup.Name = newProjectGroupName
	projectGroup.Description = newDescription

	updatedProjectGroup, err := octopusClient.ProjectGroups.Update(*projectGroup)
	require.NoError(t, err, "error when updating projectgroup")
	require.Equal(t, newProjectGroupName, updatedProjectGroup.Name, "projectgroup name was not updated")
	require.Equal(t, newDescription, updatedProjectGroup.Description, "projectgroup description was not updated")
}

func createTestProjectGroup(t *testing.T, octopusClient *octopusdeploy.Client, projectGroupName string) *octopusdeploy.ProjectGroup {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	p := getTestProjectGroup(projectGroupName)
	createdProjectGroup, err := octopusClient.ProjectGroups.Add(&p)

	if err != nil {
		t.Fatalf("creating projectgroup %s failed when it shouldn't: %s", projectGroupName, err)
	}

	return createdProjectGroup
}

func getTestProjectGroup(projectGroupName string) octopusdeploy.ProjectGroup {
	p := octopusdeploy.NewProjectGroup(projectGroupName)

	return *p
}

func cleanProjectGroup(t *testing.T, octopusClient *octopusdeploy.Client, projectGroupID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	err := octopusClient.ProjectGroups.DeleteByID(projectGroupID)
	assert.NoError(t, err)
}
