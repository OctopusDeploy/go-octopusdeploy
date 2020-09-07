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

	user := model.NewUser(getRandomName(), getRandomName())
	user.IsService = true
	user, err := octopusClient.Users.Add(user)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	apiKeys, err := octopusClient.APIKeys.Get(user.ID)

	assert.NoError(t, err)
	assert.NotNil(t, apiKeys)

	for _, apiKey := range *apiKeys {
		assert.NotNil(t, apiKey)
		assert.NotNil(t, apiKey.ID)
	}

	err = octopusClient.Users.Delete(user.ID)

	assert.NoError(t, err)
}

func TestGetAPIKeyByID(t *testing.T) {

	user := model.NewUser(getRandomName(), getRandomName())
	user.IsService = true
	user, err := octopusClient.Users.Add(user)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	apiKeys, err := octopusClient.APIKeys.Get(user.ID)

	assert.NoError(t, err)
	assert.NotNil(t, apiKeys)

	for _, apiKey := range *apiKeys {
		key, _ := octopusClient.APIKeys.GetByID(user.ID, apiKey.ID)
		assert.NotNil(t, key)
	}

	err = octopusClient.Users.Delete(user.ID)

	assert.NoError(t, err)
}

func TestCreateAPIKey(t *testing.T) {

	apiKey, err := model.NewAPIKey(getRandomName())

	assert.NoError(t, err)
	assert.NotNil(t, apiKey)

	user := model.NewUser(getRandomName(), getRandomName())
	user.IsService = true
	user, err = octopusClient.Users.Add(user)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	apiKey.UserID = &user.ID
	createdAPIKey, err := octopusClient.APIKeys.Create(apiKey)

	assert.NoError(t, err)
	assert.NotNil(t, createdAPIKey)

	err = octopusClient.Users.Delete(user.ID)

	assert.NoError(t, err)
}
