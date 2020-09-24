package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var azureSubscriptionName string = getRandomName()

func TestCreateInvalidAzureSubscriptionAccount(t *testing.T) {
	octopusClient := getOctopusClient()

	account, err := model.NewAzureSubscriptionAccount(azureSubscriptionName, uuid.Nil)

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Error(t, account.Validate())

	account, err = octopusClient.Accounts.Add(account)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestCreateAndDeleteAndGetMinimalAzureSubscriptionAccount(t *testing.T) {
	octopusClient := getOctopusClient()

	account, err := model.NewAzureSubscriptionAccount(azureSubscriptionName, uuid.New())

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

	assert.NotEmpty(t, createdAccount.ID)
	assert.NotEmpty(t, createdAccount.LastModifiedBy)
	assert.NotEmpty(t, createdAccount.LastModifiedOn)
	assert.NotEmpty(t, createdAccount.Links)
	assert.NotEmpty(t, createdAccount.SpaceID)

	assert.Equal(t, account.ApplicationID, createdAccount.ApplicationID)
	assert.Equal(t, account.SubscriptionID, createdAccount.SubscriptionID)
	assert.Equal(t, account.TenantedDeploymentParticipation, createdAccount.TenantedDeploymentParticipation)

	err = octopusClient.Accounts.DeleteByID(createdAccount.ID)

	assert.NoError(t, err)

	deletedAccount, err := octopusClient.Accounts.GetByID(createdAccount.ID)

	assert.Error(t, err)
	assert.Nil(t, deletedAccount)
}
