package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIKeys(t *testing.T) {
	t.Run("Create", TestAPIKeysCreate)
	t.Run("GetByID", TestAPIKeysGetByID)
	t.Run("GetByUserID", TestAPIKeysGetByUserID)
}

func TestAPIKeysCreate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user := model.NewUser(getRandomName(), getRandomName())
	user.IsService = true

	user, err := octopusClient.Users.Add(user)
	require.NoError(t, err)
	require.NotNil(t, user)

	apiKey, err := model.NewAPIKey(getRandomName(), user.ID)
	require.NoError(t, err)
	require.NotNil(t, apiKey)

	createdAPIKey, err := octopusClient.APIKeys.Create(apiKey)
	require.NoError(t, err)
	require.NotNil(t, createdAPIKey)

	err = octopusClient.Users.DeleteByID(user.ID)
	assert.NoError(t, err)
}

func TestAPIKeysGetByID(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user := model.NewUser(getRandomName(), getRandomName())
	user.IsService = true

	user, err := octopusClient.Users.Add(user)
	require.NoError(t, err)
	require.NotNil(t, user)

	apiKeys, err := octopusClient.APIKeys.GetByUserID(user.ID)
	require.NoError(t, err)
	require.NotNil(t, apiKeys)

	for _, apiKey := range *apiKeys {
		key, _ := octopusClient.APIKeys.GetByID(user.ID, apiKey.ID)
		assert.NotNil(t, key)
	}

	err = octopusClient.Users.DeleteByID(user.ID)
	assert.NoError(t, err)
}

func TestAPIKeysGetByUserID(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	user := model.NewUser(getRandomName(), getRandomName())
	user.IsService = true
	assert.NotNil(t, user)

	user, err := octopusClient.Users.Add(user)
	require.NoError(t, err)
	require.NotNil(t, user)

	apiKey, err := model.NewAPIKey(getRandomName(), user.ID)
	require.NoError(t, err)
	require.NotNil(t, apiKey)

	createdAPIKey, err := octopusClient.APIKeys.Create(apiKey)
	require.NoError(t, err)
	require.NotNil(t, createdAPIKey)

	apiKeys, err := octopusClient.APIKeys.GetByUserID(user.ID)
	require.NoError(t, err)
	require.NotNil(t, apiKeys)

	for _, apiKey := range *apiKeys {
		assert.NotNil(t, apiKey)
		assert.NotNil(t, apiKey.ID)
	}

	err = octopusClient.Users.DeleteByID(user.ID)
	assert.NoError(t, err)
}
