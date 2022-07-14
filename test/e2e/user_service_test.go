package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestUser(t *testing.T, client *client.Client) *users.User {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	username := internal.GetRandomName()
	displayName := internal.GetRandomName()
	password := internal.GetRandomName()

	user := users.NewUser(username, displayName)
	user.Password = password
	require.NoError(t, user.Validate())

	createdUser, err := client.Users.Add(user)
	require.NotNil(t, createdUser)
	require.NoError(t, err)

	userToCompare, err := client.Users.GetByID(createdUser.GetID())
	require.NotNil(t, userToCompare)
	require.NoError(t, err)

	AssertEqualUsers(t, createdUser, userToCompare)

	return createdUser
}

func DeleteTestUser(t *testing.T, client *client.Client, user *users.User) {
	require.NotNil(t, user)

	// you cannot delete your own account
	if user.IsRequestor {
		return
	}

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Users.DeleteByID(user.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedUser, err := client.Users.GetByID(user.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedUser)
}

func AssertEqualUsers(t *testing.T, expected *users.User, actual *users.User) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	assert.Equal(t, expected.CanPasswordBeEdited, actual.CanPasswordBeEdited)
	assert.Equal(t, expected.DisplayName, actual.DisplayName)
	assert.Equal(t, expected.EmailAddress, actual.EmailAddress)
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.Equal(t, expected.Identities, actual.Identities)
	assert.Equal(t, expected.IsActive, actual.IsActive)
	assert.Equal(t, expected.IsRequestor, actual.IsRequestor)
	assert.Equal(t, expected.IsService, actual.IsService)
	assert.True(t, internal.IsLinksEqual(expected.Links, actual.Links))
	assert.Equal(t, expected.Password, actual.Password)
	assert.Equal(t, expected.Username, actual.Username)
}

func TestUserServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	user := CreateTestUser(t, client)
	require.NotNil(t, user)
	defer DeleteTestUser(t, client, user)

	usersToTest, err := client.Users.GetAll()
	require.NoError(t, err)
	require.NotNil(t, usersToTest)

	for _, user := range usersToTest {
		query := users.UsersQuery{
			IDs: []string{user.GetID()},
		}

		usersToCompare, err := client.Users.Get(query)
		require.NoError(t, err)
		require.NotNil(t, usersToCompare)

		for _, userToCompare := range usersToCompare.Items {
			if user.GetID() == userToCompare.GetID() {
				AssertEqualUsers(t, user, userToCompare)
			}
		}

		userToCompare, err := client.Users.GetByID(user.GetID())
		require.NoError(t, err)
		require.NotNil(t, userToCompare)
		AssertEqualUsers(t, user, userToCompare)
	}
}

func TestUserServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	users, err := client.Users.GetAll()
	require.NoError(t, err)
	require.NotNil(t, users)

	for _, user := range users {
		defer DeleteTestUser(t, client, user)
	}
}

// TODO: fix test
// func TestUserServiceGetAPIKeys(t *testing.T) {
// 	client := getOctopusClient()
// 	require.NotNil(t, client)

// 	user, err := client.Users.GetMe()
// 	require.NotNil(t, user)
// 	require.NoError(t, err)

// 	apiKeys, err := client.Users.GetAPIKeys(user)
// 	require.NotNil(t, apiKeys)
// 	require.NoError(t, err)

// 	for _, apiKey := range apiKeys.Items {
// 		apiKeyToConfirm, err := client.Users.GetAPIKeyByID(user, apiKey.GetID())
// 		require.NotNil(t, apiKeyToConfirm)
// 		require.NoError(t, err)

// 		t.Log(apiKeyToConfirm.GetID())
// 	}
// }

func TestUserServiceGetAuthenticationByUser(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	user, err := client.Users.GetMe()
	require.NoError(t, err)
	require.NotEmpty(t, user)

	userAuthentication, err := client.Users.GetAuthenticationByUser(user)
	require.NotNil(t, userAuthentication)
	require.NoError(t, err)
}

func TestUserServiceGetAuthentication(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	userAuthentication, err := client.Users.GetAuthentication()
	require.NotNil(t, userAuthentication)
	require.NoError(t, err)
}

func TestUserServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	users, err := client.Users.GetAll()
	require.NotNil(t, users)
	require.NoError(t, err)

	for _, user := range users {
		user, err := client.Users.GetByID(user.GetID())
		require.NoError(t, err)
		require.NotNil(t, user)
	}
}

func TestUserServiceGetMe(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	user, err := client.Users.GetMe()
	require.NotNil(t, user)
	require.NoError(t, err)
}

func TestUserServiceGetPermissions(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	user, err := client.Users.GetMe()
	require.NotNil(t, user)
	require.NoError(t, err)

	userPermissionSet, err := client.Users.GetPermissions(user)
	require.NotNil(t, userPermissionSet)
	require.NoError(t, err)
}

func TestUserServiceGetPermissionsConfiguration(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	user, err := client.Users.GetMe()
	require.NotNil(t, user)
	require.NoError(t, err)

	userPermissionSet, err := client.Users.GetPermissionsConfiguration(user)
	require.NotNil(t, userPermissionSet)
	require.NoError(t, err)
}

func TestUserServiceGetSpaces(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	user, err := client.Users.GetMe()
	require.NotNil(t, user)
	require.NoError(t, err)

	spaces, err := client.Users.GetSpaces(nil)
	require.Nil(t, spaces)
	require.Equal(t, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterUser), err)

	spaces, err = client.Users.GetSpaces(user)
	require.NotNil(t, spaces)
	require.NoError(t, err)

	for _, space := range spaces {
		spaceToConfirm, err := client.Spaces.GetByID(space.GetID())
		require.NoError(t, err)
		require.NotNil(t, spaceToConfirm)
		IsEqualSpaces(t, space, spaceToConfirm)
	}
}

func TestUserServiceGetTeams(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	user, err := client.Users.GetMe()
	require.NotNil(t, user)
	require.NoError(t, err)

	teams, err := client.Users.GetTeams(user)
	require.NotNil(t, teams)
	require.NoError(t, err)
}
