package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestSpace(t *testing.T, client *octopusdeploy.Client) *octopusdeploy.Space {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	user := CreateTestUser(t, client)

	name := getShortRandomName()

	space := octopusdeploy.NewSpace(name)
	require.NoError(t, space.Validate())

	space.SpaceManagersTeamMembers = []string{user.GetID()}

	createdSpace, err := client.Spaces.Add(space)
	require.NoError(t, err)
	require.NotNil(t, createdSpace)
	require.NotEmpty(t, createdSpace.GetID())
	require.Equal(t, name, createdSpace.Name)

	return createdSpace
}

func DeleteTestSpace(t *testing.T, client *octopusdeploy.Client, space *octopusdeploy.Space) {
	require.NotNil(t, space)

	// if space.IsDefault {
	// 	return
	// }

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	spaceID := space.GetID()
	assert.NotEmpty(t, spaceID)

	if !space.TaskQueueStopped {
		space.TaskQueueStopped = true
		updatedSpace, err := client.Spaces.Update(space)
		assert.NoError(t, err)

		spaceID = updatedSpace.GetID()
		assert.NotEmpty(t, spaceID)
	}

	err := client.Spaces.DeleteByID(spaceID)
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedAccount, err := client.Spaces.GetByID(spaceID)
	assert.Error(t, err)
	assert.Nil(t, deletedAccount)
}

func IsEqualSpaces(t *testing.T, expected *octopusdeploy.Space, actual *octopusdeploy.Space) {
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

	// TODO: complete space comparison
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
}

func UpdateTestSpace(t *testing.T, client *octopusdeploy.Client, space *octopusdeploy.Space) *octopusdeploy.Space {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	updatedSpace, err := client.Spaces.Update(space)
	require.NoError(t, err)
	require.NotNil(t, updatedSpace)

	return updatedSpace
}

func TestSpaceSetAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := CreateTestSpace(t, client)
	require.NotNil(t, space)
	defer DeleteTestSpace(t, client, space)

	spaceToCompare, err := client.Spaces.GetByID(space.GetID())
	require.NoError(t, err)
	require.NotNil(t, spaceToCompare)
	IsEqualSpaces(t, space, spaceToCompare)
}

func TestSpaceServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	spaces, err := client.Spaces.GetAll()
	require.NoError(t, err)
	require.NotNil(t, spaces)

	for _, space := range spaces {
		defer DeleteTestSpace(t, client, space)
	}
}

func TestSpaceServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// create 30 test spaces (to be deleted)
	for i := 0; i < 30; i++ {
		space := CreateTestSpace(t, client)
		require.NotNil(t, space)
		defer DeleteTestSpace(t, client, space)
	}

	allSpaces, err := client.Spaces.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allSpaces)
	require.True(t, len(allSpaces) >= 30)
}

func TestSpaceServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := getRandomName()
	resource, err := client.Spaces.GetByID(id)
	require.Equal(t, createResourceNotFoundError(octopusdeploy.ServiceSpaceService, "ID", id), err)
	require.Nil(t, resource)

	resources, err := client.Spaces.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		resourceToCompare, err := client.Spaces.GetByID(resource.GetID())
		require.NoError(t, err)
		IsEqualSpaces(t, resource, resourceToCompare)
	}
}

func TestSpaceGetByPartialName(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	spaces, err := client.Spaces.GetAll()
	require.NoError(t, err)
	require.NotNil(t, spaces)

	for _, space := range spaces {
		namedSpaces, err := client.Spaces.GetByPartialName(space.Name)
		require.NoError(t, err)
		require.NotNil(t, namedSpaces)
	}
}

func TestSpaceServiceUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	expected := CreateTestSpace(t, client)
	actual := UpdateTestSpace(t, client, expected)
	IsEqualSpaces(t, expected, actual)
	defer DeleteTestSpace(t, client, expected)
}
