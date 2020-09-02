package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

func init() {
	octopusClient = initTest()
}

func TestGetByID(t *testing.T) {
	apiKey, user := createAPIKey(t)
	assert.NotNil(t, apiKey)
	deleteUser(*user, t)
}

func createUser(t *testing.T) (*model.User, error) {
	user := model.NewUser(getRandomName(), getRandomName())
	user.Password = getRandomName()
	return octopusClient.Users.Add(user)
}

func deleteUser(user model.User, t *testing.T) error {
	return octopusClient.Users.Delete(user.ID)
}

func createAPIKey(t *testing.T) (*model.APIKey, *model.User) {
	user, _ := createUser(t)
	apiKey, _ := octopusClient.APIKeys.Get(user.ID)
	return apiKey, user
}
