package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndDeleteAndGetUsernamePasswordAccount(t *testing.T) {
	name := getRandomName()
	description := getRandomName()
	username := getRandomName()
	password := getRandomName()

	account, err := model.NewUsernamePasswordAccount(name)
	account.Description = description
	account.Username = username
	account.Password = &model.SensitiveValue{NewValue: &password}

	assert.NoError(t, err)
	assert.NotNil(t, account)

	if err != nil {
		return
	}

	assert.NoError(t, err)
	assert.NotNil(t, account)

	if err != nil {
		return
	}

	assert.NoError(t, account.Validate())

	createdAccount, err := octopusClient.Accounts.Add(account)

	assert.NoError(t, err)
	assert.NotNil(t, createdAccount)
	assert.NoError(t, createdAccount.Validate())

	if err != nil {
		return
	}

	assert.Equal(t, account.Name, createdAccount.Name)
	assert.Equal(t, account.TenantedDeploymentParticipation, createdAccount.TenantedDeploymentParticipation)
	assert.Equal(t, account.Description, createdAccount.Description)
	assert.Equal(t, account.AccountType, createdAccount.AccountType)
	assert.Equal(t, account.Username, createdAccount.Username)
	assert.NotEmpty(t, createdAccount.Password)
	assert.True(t, createdAccount.Password.HasValue)
	assert.Nil(t, createdAccount.Password.NewValue)
	assert.NotEmpty(t, createdAccount.Links)

	err = octopusClient.Accounts.Delete(createdAccount.ID)

	assert.NoError(t, err)

	deletedAccount, err := octopusClient.Accounts.Get(createdAccount.ID)

	assert.Error(t, err)
	assert.Nil(t, deletedAccount)
}
