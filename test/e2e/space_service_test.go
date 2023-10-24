package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/spaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestSpace(t *testing.T, client *client.Client) *spaces.Space {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	user := CreateTestUser(t, client)

	name := getShortRandomName()

	space := spaces.NewSpace(name)
	require.NoError(t, space.Validate())

	space.SpaceManagersTeamMembers = []string{user.GetID()}

	createdSpace, err := client.Spaces.Add(space)
	require.NoError(t, err)
	require.NotNil(t, createdSpace)
	require.NotEmpty(t, createdSpace.GetID())
	require.Equal(t, name, createdSpace.Name)

	return createdSpace
}

func DeleteTestSpace(t *testing.T, client *client.Client, space *spaces.Space) {
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

func IsEqualSpaces(t *testing.T, expected *spaces.Space, actual *spaces.Space) {
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

	// space
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.IsDefault, actual.IsDefault)
	assert.Equal(t, expected.SpaceManagersTeamMembers, actual.SpaceManagersTeamMembers)
	assert.Equal(t, expected.SpaceManagersTeams, actual.SpaceManagersTeams)
	assert.Equal(t, expected.TaskQueueStopped, actual.TaskQueueStopped)
}

func GetDefaultSpace(t *testing.T, client *client.Client) *spaces.Space {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	spaces, err := client.Spaces.GetAll()
	require.NotNil(t, spaces)
	require.NotEmpty(t, spaces)
	require.NoError(t, err)

	for _, space := range spaces {
		if space.IsDefault {
			return space
		}
	}

	return nil
}

func UpdateTestSpace(t *testing.T, client *client.Client, space *spaces.Space) *spaces.Space {
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

// TODO: fix test
// func TestSpaceServiceGetAll(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	// create 30 test spaces (to be deleted)
// 	for i := 0; i < 30; i++ {
// 		space := CreateTestSpace(t, client)
// 		require.NotNil(t, space)
// 		defer DeleteTestSpace(t, client, space)
// 	}

// 	allSpaces, err := client.Spaces.GetAll()
// 	require.NoError(t, err)
// 	require.NotNil(t, allSpaces)
// 	require.True(t, len(allSpaces) >= 30)
// }

func TestSpaceServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := internal.GetRandomName()
	resource, err := client.Spaces.GetByID(id)
	require.Error(t, err)
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

func TestSpaceServiceGetByIDOrName(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	idOrName := internal.GetRandomName()
	resource, err := client.Spaces.GetByIDOrName(idOrName)
	require.Error(t, err)
	require.Nil(t, resource)

	resources, err := client.Spaces.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		resourceToCompare, err := client.Spaces.GetByIDOrName(resource.GetID())
		require.NoError(t, err)
		IsEqualSpaces(t, resource, resourceToCompare)
	}
}

func TestSpaceGetByName(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	name := internal.GetRandomName()
	resource, err := client.Spaces.GetByName(name)
	require.Error(t, err)
	require.Nil(t, resource)

	spaces, err := client.Spaces.GetAll()
	require.NoError(t, err)
	require.NotNil(t, spaces)

	for _, space := range spaces {
		namedSpaces, err := client.Spaces.GetByName(space.Name)
		require.NoError(t, err)
		require.NotNil(t, namedSpaces)
	}
}

// TODO: fix test
// func TestSpaceServiceUpdate(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	expected := CreateTestSpace(t, client)
// 	actual := UpdateTestSpace(t, client, expected)
// 	IsEqualSpaces(t, expected, actual)
// 	defer DeleteTestSpace(t, client, expected)
// }

// === NEW ===

func TestSpaceSetAddGetDelete_NewClient(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := CreateTestSpace_NewClient(t, client)
	require.NotNil(t, space)
	defer DeleteTestSpace_NewClient(t, client, space)

	spaceToCompare, err := spaces.GetByID(client, space.GetID())
	require.NoError(t, err)
	require.NotNil(t, spaceToCompare)
	IsEqualSpaces(t, space, spaceToCompare)
}

func CreateTestSpace_NewClient(t *testing.T, client *client.Client) *spaces.Space {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	user := CreateTestUser_NewClient(t, client)

	name := getShortRandomName()

	space := spaces.NewSpace(name)
	require.NoError(t, space.Validate())

	space.SpaceManagersTeamMembers = []string{user.GetID()}

	// TODO: newclient space Add function
	createdSpace, err := client.Spaces.Add(space)
	require.NoError(t, err)
	require.NotNil(t, createdSpace)
	require.NotEmpty(t, createdSpace.GetID())
	require.Equal(t, name, createdSpace.Name)

	return createdSpace
}

func DeleteTestSpace_NewClient(t *testing.T, client *client.Client, space *spaces.Space) {
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
		updatedSpace, err := spaces.Update(client, space)
		assert.NoError(t, err)

		spaceID = updatedSpace.GetID()
		assert.NotEmpty(t, spaceID)
	}

	err := client.Spaces.DeleteByID(spaceID)
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedAccount, err := spaces.GetByID(client, spaceID)
	assert.Error(t, err)
	assert.Nil(t, deletedAccount)
}
