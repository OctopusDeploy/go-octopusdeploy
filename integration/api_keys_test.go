package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

func init() {
	if octopusClient == nil {
		octopusClient = initTest()
	}
}

func TestGetAPIKeys(t *testing.T) {

	user, err := createUser(t)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	apiKeys, err := octopusClient.APIKeys.Get(user.ID)

	assert.NoError(t, err)
	assert.NotNil(t, apiKeys)

	for _, apiKey := range *apiKeys {
		assert.NotNil(t, apiKey)
		assert.NotNil(t, apiKey.ID)
	}

	deleteUser(t, *user)
}

func TestGetByID(t *testing.T) {

	user, err := createUser(t)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	apiKeys, err := octopusClient.APIKeys.Get(user.ID)

	assert.NoError(t, err)
	assert.NotNil(t, apiKeys)

	for _, apiKey := range *apiKeys {
		key, _ := octopusClient.APIKeys.GetByID(user.ID, apiKey.ID)
		assert.NotNil(t, key)
	}

	deleteUser(t, *user)
}

func TestCreateAPIKey(t *testing.T) {

	apiKey, err := model.NewAPIKey(getRandomName())

	assert.NoError(t, err)
	assert.NotNil(t, apiKey)

	user, err := createUser(t)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	apiKey.UserID = &user.ID
	createdAPIKey, err := octopusClient.APIKeys.Create(apiKey)

	assert.NoError(t, err)
	assert.NotNil(t, createdAPIKey)
}

func createUser(t *testing.T) (*model.User, error) {
	user := model.NewUser(getRandomName(), getRandomName())
	user.IsService = true
	return octopusClient.Users.Add(user)
}

func deleteUser(t *testing.T, user model.User) error {
	return octopusClient.Users.Delete(user.ID)
}
