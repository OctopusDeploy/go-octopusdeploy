package e2e

import (
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateSimpleTestUser(t *testing.T, client *client.Client) *users.User {
	user := users.NewUser(internal.GetRandomName(), internal.GetRandomName())
	user.IsService = true

	user, err := users.Add(client, user)
	require.NoError(t, err)
	require.NotNil(t, user)

	return user
}

func CreateTestAPIKey(t *testing.T, client *client.Client, user *users.User) *users.CreateAPIKey {
	apiKey := users.NewAPIKey(internal.GetRandomName(), user.ID)
	require.NotNil(t, apiKey)

	expiry := time.Now().Add(time.Hour).Round(time.Millisecond)
	apiKey.Expires = &expiry

	createdAPIKey, err := client.APIKeys.Create(apiKey)
	require.NoError(t, err)
	require.NotNil(t, createdAPIKey)

	return createdAPIKey
}

func CleanupTestUser(t *testing.T, client *client.Client, user *users.User) {
	err := users.DeleteByID(client, user.ID)
	assert.NoError(t, err)
}

func TestAPIKeysCreate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	user := CreateSimpleTestUser(t, client)
	CreateTestAPIKey(t, client, user)

	apiKeys, err := client.APIKeys.GetByUserID(user.ID)
	require.NotNil(t, apiKeys)
	require.NoError(t, err)

	CleanupTestUser(t, client, user)
}

func TestAPIKeysGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	user := CreateSimpleTestUser(t, client)
	CreateTestAPIKey(t, client, user)
	CreateTestAPIKey(t, client, user)

	apiKeys, err := client.APIKeys.GetByUserID(user.ID)
	require.NoError(t, err)
	require.NotNil(t, apiKeys)

	for _, apiKey := range apiKeys {
		key, _ := client.APIKeys.GetByID(user.ID, apiKey.ID)
		assert.NotNil(t, key)
	}

	CleanupTestUser(t, client, user)
}

func TestAPIKeysGetByUserID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	user := CreateSimpleTestUser(t, client)
	CreateTestAPIKey(t, client, user)
	CreateTestAPIKey(t, client, user)

	apiKeys, err := client.APIKeys.GetByUserID(user.ID)
	require.NoError(t, err)
	require.NotNil(t, apiKeys)

	for _, apiKey := range apiKeys {
		assert.NotNil(t, apiKey)
		assert.NotNil(t, apiKey.ID)
	}

	CleanupTestUser(t, client, user)
}
