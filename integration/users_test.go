package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsers(t *testing.T) {
	t.Run("Add", TestUsersAdd)
	t.Run("GetAll", TestUsersGetAll)
	t.Run("GetAuthentication", TestUsersGetAuthentication)
	t.Run("GetAuthenticationForUser", TestUsersGetAuthenticationForUser)
	t.Run("GetByID", TestUsersGetByID)
	t.Run("GetMe", TestUsersGetMe)
	t.Run("GetSpaces", TestUsersGetSpaces)
}

func TestUsersAdd(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user := model.NewUser(getRandomName(), getRandomName())
	user.Password = getRandomName()

	user, err := octopusClient.Users.Add(user)
	assert.NoError(t, err)
	require.NotNil(t, user)

	assert.True(t, user.IsActive)
	assert.False(t, user.IsService)
	assert.Empty(t, user.EmailAddress)
}

func TestUsersGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	users, err := octopusClient.Users.GetAll()
	assert.NoError(t, err)
	require.NotEmpty(t, users)
}

func TestUsersGetAuthentication(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	authentication, err := octopusClient.Users.GetAuthentication()
	assert.NoError(t, err)
	require.NotEmpty(t, authentication)

	// TODO: add more asserts here
}

func TestUsersGetAuthenticationForUser(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user, err := octopusClient.Users.GetMe()
	assert.NoError(t, err)
	require.NotEmpty(t, user)

	authentication, err := octopusClient.Users.GetAuthenticationForUser(user)
	assert.NoError(t, err)
	require.NotEmpty(t, authentication)

	// TODO: add more asserts here
}

func TestUsersGetByID(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user, err := octopusClient.Users.GetMe()
	assert.NoError(t, err)
	require.NotEmpty(t, user)

	userToVerify, err := octopusClient.Users.GetByID(user.ID)

	assert.NoError(t, err)
	require.NotEmpty(t, userToVerify)

	// TODO: add more asserts here
}

func TestUsersGetMe(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user, err := octopusClient.Users.GetMe()
	assert.NoError(t, err)
	require.NotNil(t, user)

	assert.True(t, user.IsActive)
	assert.False(t, user.IsService)
	assert.NotEmpty(t, user.EmailAddress)
}

func TestUsersGetSpaces(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user, err := octopusClient.Users.GetMe()
	assert.NoError(t, err)
	require.NotEmpty(t, user)

	spaces, err := octopusClient.Users.GetSpaces(user)
	assert.NoError(t, err)
	require.NotNil(t, spaces)

	// TODO: add more asserts here
}
