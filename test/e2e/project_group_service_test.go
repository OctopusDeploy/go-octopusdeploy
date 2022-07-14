package e2e

import (
	// "fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/projectgroups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualProjectGroups(t *testing.T, expected *projectgroups.ProjectGroup, actual *projectgroups.ProjectGroup) {
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
	assert.True(t, internal.IsLinksEqual(expected.GetLinks(), actual.GetLinks()))

	// ProjectGroup
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.EnvironmentIDs, actual.EnvironmentIDs)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.RetentionPolicyID, actual.RetentionPolicyID)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
}

func CreateTestProjectGroup(t *testing.T, client *client.Client) *projectgroups.ProjectGroup {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	projectGroup := projectgroups.NewProjectGroup(name)
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

	return createdProjectGroup
}

func DeleteTestProjectGroup(t *testing.T, client *client.Client, projectGroup *projectgroups.ProjectGroup) {
	require.NotNil(t, projectGroup)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	projectGroup, err := client.ProjectGroups.GetByID(projectGroup.GetID())
	require.NoError(t, err)
	require.NotNil(t, projectGroup)

	projects, err := client.ProjectGroups.GetProjects(projectGroup)
	require.NoError(t, err)
	require.NotNil(t, projects)

	// cannot delete project groups that contain projects
	if len(projects) > 0 {
		return
	}

	err = client.ProjectGroups.DeleteByID(projectGroup.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedProjectGroup, err := client.ProjectGroups.GetByID(projectGroup.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedProjectGroup)
}

func TestProjectGroupServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	projectGroups, err := client.ProjectGroups.GetAll()
	require.NoError(t, err)
	require.NotNil(t, projectGroups)

	for _, projectGroup := range projectGroups {
		defer DeleteTestProjectGroup(t, client, projectGroup)
	}
}

func TestProjectGroupAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)
}

func TestProjectGroupGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := internal.GetRandomName()
	resource, err := client.ProjectGroups.GetByID(id)
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestProjectGroupGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// create many projects to test pagination
	projectsGroupsToCreate := 32
	sum := 0
	for i := 0; i < projectsGroupsToCreate; i++ {
		projectGroup := CreateTestProjectGroup(t, client)
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

	additionalProjectGroup := CreateTestProjectGroup(t, client)
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

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	newProjectGroupName := internal.GetRandomName()
	const newDescription = "this should be updated"

	projectGroup.Name = newProjectGroupName
	projectGroup.Description = newDescription

	updatedProjectGroup, err := client.ProjectGroups.Update(*projectGroup)
	require.NoError(t, err, "error when updating projectgroup")
	require.Equal(t, newProjectGroupName, updatedProjectGroup.Name, "projectgroup name was not updated")
	require.Equal(t, newDescription, updatedProjectGroup.Description, "projectgroup description was not updated")
}
