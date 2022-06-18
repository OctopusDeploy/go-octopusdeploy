package e2e

// import (
// 	"testing"

// // 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// func TestAPIKeysCreate(t *testing.T) {
// 	octopusClient := getOctopusClient()
// 	require.NotNil(t, octopusClient)

// 	user := NewUser(internal.GetRandomName(), internal.GetRandomName())
// 	user.IsService = true

// 	user, err := octopusClient.Users.Add(user)
// 	require.NoError(t, err)
// 	require.NotNil(t, user)

// 	apiKey := NewAPIKey(internal.GetRandomName(), user.id)
// 	require.NotNil(t, apiKey)

// 	createdAPIKey, err := octopusClient.APIKeys.Create(apiKey)
// 	require.NoError(t, err)
// 	require.NotNil(t, createdAPIKey)

// 	err = octopusClient.Users.DeleteByID(user.id)
// 	assert.NoError(t, err)
// }

// func TestAPIKeysGetByID(t *testing.T) {
// 	octopusClient := getOctopusClient()
// 	require.NotNil(t, octopusClient)

// 	user := NewUser(internal.GetRandomName(), internal.GetRandomName())
// 	user.IsService = true

// 	user, err := octopusClient.Users.Add(user)
// 	require.NoError(t, err)
// 	require.NotNil(t, user)

// 	apiKeys, err := octopusClient.APIKeys.GetByUserID(user.id)
// 	require.NoError(t, err)
// 	require.NotNil(t, apiKeys)

// 	for _, apiKey := range *apiKeys {
// 		key, _ := octopusClient.APIKeys.GetByID(user.id, apiKey.id)
// 		assert.NotNil(t, key)
// 	}

// 	err = octopusClient.Users.DeleteByID(user.id)
// 	assert.NoError(t, err)
// }

// func TestAPIKeysGetByUserID(t *testing.T) {
// 	octopusClient := getOctopusClient()
// 	require.NotNil(t, octopusClient)

// 	user := NewUser(internal.GetRandomName(), internal.GetRandomName())
// 	user.IsService = true
// 	assert.NotNil(t, user)

// 	user, err := octopusClient.Users.Add(user)
// 	require.NoError(t, err)
// 	require.NotNil(t, user)

// 	apiKey := NewAPIKey(internal.GetRandomName(), user.id)
// 	require.NotNil(t, apiKey)

// 	createdAPIKey, err := octopusClient.APIKeys.Create(apiKey)
// 	require.NoError(t, err)
// 	require.NotNil(t, createdAPIKey)

// 	apiKeys, err := octopusClient.APIKeys.GetByUserID(user.id)
// 	require.NoError(t, err)
// 	require.NotNil(t, apiKeys)

// 	for _, apiKey := range *apiKeys {
// 		assert.NotNil(t, apiKey)
// 		assert.NotNil(t, apiKey.id)
// 	}

// 	err = octopusClient.Users.DeleteByID(user.id)
// 	assert.NoError(t, err)
// }
