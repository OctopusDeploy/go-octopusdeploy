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

func TestCreateInvalidTokenAccount(t *testing.T) {
	account, err := model.NewTokenAccount("", model.SensitiveValue{})

	assert.Error(t, err)
	assert.Nil(t, account)

	account, err = model.NewTokenAccount(getRandomName(), generateSensitiveValue())

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.NoError(t, account.Validate())

	createdAccount, err := octopusClient.Accounts.Add(account)

	assert.NoError(t, err)
	assert.NotNil(t, createdAccount)
}

func TestCreateAndDeleteAndGetTokenAccount(t *testing.T) {
	sensitiveValue, err := model.NewSensitiveValue(getRandomName())

	assert.NoError(t, err)
	assert.NotNil(t, sensitiveValue)

	account, err := model.NewTokenAccount(getRandomName(), *sensitiveValue)

	assert.NoError(t, err)
	assert.NotNil(t, account)

	if err != nil {
		return
	}

	createdAccount, err := octopusClient.Accounts.Add(account)

	assert.NoError(t, err)
	assert.NotNil(t, createdAccount)
	assert.NoError(t, createdAccount.Validate())

	if err != nil {
		return
	}

	assert.NotEmpty(t, createdAccount.ID)
	assert.NotEmpty(t, createdAccount.LastModifiedBy)
	assert.NotEmpty(t, createdAccount.LastModifiedOn)
	assert.NotEmpty(t, createdAccount.Links)
	assert.NotEmpty(t, createdAccount.SpaceID)

	assert.Equal(t, account.ApplicationID, createdAccount.ApplicationID)
	assert.Equal(t, account.SubscriptionID, createdAccount.SubscriptionID)
	assert.Equal(t, account.TenantedDeploymentParticipation, createdAccount.TenantedDeploymentParticipation)

	err = octopusClient.Accounts.Delete(createdAccount.ID)

	assert.NoError(t, err)

	deletedAccount, err := octopusClient.Accounts.Get(createdAccount.ID)

	assert.Error(t, err)
	assert.Nil(t, deletedAccount)
}
