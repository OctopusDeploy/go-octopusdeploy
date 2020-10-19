package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertEqualUsers(t *testing.T, expected *model.User, actual *model.User) {
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
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Identities, actual.Identities)
	assert.Equal(t, expected.IsActive, actual.IsActive)
	assert.Equal(t, expected.IsRequestor, actual.IsRequestor)
	assert.Equal(t, expected.IsService, actual.IsService)
	assert.True(t, IsEqualLinks(expected.Links, actual.Links))
	assert.Equal(t, expected.Password, actual.Password)
	assert.Equal(t, expected.Username, actual.Username)
}

func TestUsersAdd(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user := model.NewUser(getRandomName(), getRandomName())
	user.Password = getRandomName()

	actual, err := octopusClient.Users.Add(user)
	require.NoError(t, err)
	require.NotNil(t, actual)
	assert.NotEmpty(t, actual.GetLastModifiedBy())
	assert.NotEmpty(t, actual.GetLastModifiedOn())

	expected, err := octopusClient.Users.GetByID(actual.GetID())
	require.NoError(t, err)
	require.NotNil(t, actual)

	assertEqualUsers(t, expected, actual)

	err = octopusClient.Users.DeleteByID(actual.GetID())
	require.NoError(t, err)
}

func TestUsersGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	users, err := octopusClient.Users.GetAll()
	require.NoError(t, err)
	require.NotEmpty(t, users)
}

func TestUsersGetAuthentication(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	authentication, err := octopusClient.Users.GetAuthentication()
	require.NoError(t, err)
	require.NotNil(t, authentication)
}

func TestUsersGetAuthenticationForUser(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user, err := octopusClient.Users.GetMe()
	assert.NoError(t, err)
	require.NotEmpty(t, user)

	authentication, err := octopusClient.Users.GetAuthenticationForUser(user)
	require.NoError(t, err)
	require.NotNil(t, authentication)
}

func TestUsersGetByID(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	expected, err := octopusClient.Users.GetMe()
	require.NoError(t, err)
	require.NotEmpty(t, expected)

	actual, err := octopusClient.Users.GetByID(expected.GetID())
	require.NoError(t, err)
	require.NotEmpty(t, actual)

	assertEqualUsers(t, expected, actual)
}

func TestUsersGetMe(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user, err := octopusClient.Users.GetMe()
	require.NoError(t, err)
	require.NotNil(t, user)

	assert.True(t, user.IsActive)
	assert.False(t, user.IsService)
	assert.NotEmpty(t, user.EmailAddress)
}

func TestUsersGetSpaces(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user, err := octopusClient.Users.GetMe()
	require.NoError(t, err)
	require.NotNil(t, user)

	spaces, err := octopusClient.Users.GetSpaces(user)
	require.NoError(t, err)
	require.NotNil(t, spaces)
	require.GreaterOrEqual(t, len(*spaces), 1)
}

func TestUsersUpdate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user := model.NewUser(getRandomName(), getRandomName())
	user.Password = getRandomName()

	actual, err := octopusClient.Users.Add(user)
	require.NoError(t, err)
	require.NotNil(t, actual)
	assert.NotEmpty(t, actual.GetLastModifiedBy())
	assert.NotEmpty(t, actual.GetLastModifiedOn())

	expected, err := octopusClient.Users.GetByID(actual.GetID())
	require.NoError(t, err)
	require.NotNil(t, actual)

	assertEqualUsers(t, expected, actual)

	err = octopusClient.Users.DeleteByID(actual.GetID())
	require.NoError(t, err)
}
