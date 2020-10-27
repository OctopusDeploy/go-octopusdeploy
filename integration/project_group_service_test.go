package integration

import (
	// "fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualProjectGroups(t *testing.T, expected *octopusdeploy.ProjectGroup, actual *octopusdeploy.ProjectGroup) {
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

	// ProjectGroup
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.EnvironmentIDs, actual.EnvironmentIDs)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.RetentionPolicyID, actual.RetentionPolicyID)
}

func CreateTestProjectGroup(t *testing.T, client *octopusdeploy.Client) (*octopusdeploy.ProjectGroup, error) {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()

	projectGroup := octopusdeploy.NewProjectGroup(name)
	require.NotNil(t, projectGroup)

	createdProjectGroup, err := client.ProjectGroups.Add(projectGroup)
	require.NoError(t, err)
	require.NotNil(t, createdProjectGroup)
	require.NotEmpty(t, createdProjectGroup.GetID())

	// verify the add operation was successful
	projectGroupToCompare, err := client.ProjectGroups.GetByID(createdProjectGroup.GetID())
	require.NoError(t, err)
	require.NotNil(t, projectGroupToCompare)
	AssertEqualProjectGroups(t, createdProjectGroup, projectGroupToCompare)

	return createdProjectGroup, nil
}

func DeleteTestProjectGroup(t *testing.T, client *octopusdeploy.Client, projectGroup *octopusdeploy.ProjectGroup) {
	require.NotNil(t, projectGroup)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.ProjectGroups.DeleteByID(projectGroup.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedProjectGroup, err := client.ProjectGroups.GetByID(projectGroup.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedProjectGroup)
}

func TestProjectGroupAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	projectGroup, err := CreateTestProjectGroup(t, client)
	require.NoError(t, err)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)
}

func TestProjectGroupGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := getRandomName()
	resource, err := client.ProjectGroups.GetByID(id)
	require.Equal(t, createResourceNotFoundError(octopusdeploy.ServiceProjectGroupService, "ID", id), err)
	require.Nil(t, resource)
}

func TestProjectGroupGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// create many projects to test pagination
	projectsGroupsToCreate := 32
	sum := 0
	for i := 0; i < projectsGroupsToCreate; i++ {
		projectGroup, err := CreateTestProjectGroup(t, client)
		require.NoError(t, err)
		require.NotNil(t, projectGroup)
		defer DeleteTestProjectGroup(t, client, projectGroup)
		sum += i
	}

	allProjectGroups, err := client.ProjectGroups.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	numberOfProjectGroups := len(allProjectGroups)

	// check there are greater than or equal to the amount of projects requested to be created, otherwise pagination isn't working
	if numberOfProjectGroups < projectsGroupsToCreate {
		t.Fatalf("There should be at least %d project groups created but there was only %d. Pagination is likely not working.", projectsGroupsToCreate, numberOfProjectGroups)
	}

	additionalProjectGroup, err := CreateTestProjectGroup(t, client)
	require.NoError(t, err)
	require.NotNil(t, additionalProjectGroup)
	defer DeleteTestProjectGroup(t, client, additionalProjectGroup)

	allProjectGroupsAfterCreatingAdditional, err := client.ProjectGroups.GetAll()
	if err != nil {
		t.Fatalf("Retrieving all projects failed when it shouldn't: %s", err)
	}

	assert.NoError(t, err, "error when looking for project when not expected")
	assert.Equal(t, len(allProjectGroupsAfterCreatingAdditional), numberOfProjectGroups+1, "created an additional projectgroup and expected number of projectgroups to increase by 1")
}

func TestProjectGroupUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	projectGroup, err := CreateTestProjectGroup(t, client)
	require.NoError(t, err)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	newProjectGroupName := getRandomName()
	const newDescription = "this should be updated"

	projectGroup.Name = newProjectGroupName
	projectGroup.Description = newDescription

	updatedProjectGroup, err := client.ProjectGroups.Update(*projectGroup)
	require.NoError(t, err, "error when updating projectgroup")
	require.Equal(t, newProjectGroupName, updatedProjectGroup.Name, "projectgroup name was not updated")
	require.Equal(t, newDescription, updatedProjectGroup.Description, "projectgroup description was not updated")
}
