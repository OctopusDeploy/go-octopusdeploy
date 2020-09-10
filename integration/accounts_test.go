package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var name string = getRandomName()
var description string = getRandomName()
var username string = getRandomName()
var password string = getRandomName()

func init() {
	if octopusClient == nil {
		octopusClient = initTest()
	}
}

func TestEmptyAccount(t *testing.T) {
	account, err := octopusClient.Accounts.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestCreateAndDeleteAndGetUsernamePasswordAccount(t *testing.T) {
	account, err := model.NewAccount(name, enum.UsernamePassword)
	account.Description = description
	account.Username = username
	account.Password = &model.SensitiveValue{NewValue: &password}

	assert.NoError(t, err)
	assert.NotNil(t, account)

	account, err = octopusClient.Accounts.Add(account)

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.NoError(t, account.Validate())

	verificationAccount, err := octopusClient.Accounts.Get(account.ID)

	assert.NoError(t, err)
	assert.NotNil(t, verificationAccount)
	assert.NoError(t, verificationAccount.Validate())

	assert.Equal(t, account.Name, verificationAccount.Name)
	assert.Equal(t, account.TenantedDeploymentParticipation, verificationAccount.TenantedDeploymentParticipation)
	assert.Equal(t, account.Description, verificationAccount.Description)
	assert.Equal(t, account.AccountType, verificationAccount.AccountType)
	assert.Equal(t, account.Username, verificationAccount.Username)
	assert.NotNil(t, verificationAccount.Password)
	assert.True(t, verificationAccount.Password.HasValue)
	assert.Nil(t, verificationAccount.Password.NewValue)
	assert.NotEmpty(t, verificationAccount.Links)
	assert.Equal(t, account.Links["Self"], verificationAccount.Links["Self"])

	err = octopusClient.Accounts.Delete(account.ID)

	assert.NoError(t, err)

	account, err = octopusClient.Accounts.Get(account.ID)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestCreateInvalidAzureServicePrincipalAccount(t *testing.T) {
	account, err := model.NewAccount(name, enum.AzureServicePrincipal)

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Error(t, account.Validate())

	account, err = octopusClient.Accounts.Add(account)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestCreateInvalidAzureSubscriptionAccount(t *testing.T) {
	account, err := model.NewAccount(name, enum.AzureSubscription)

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Error(t, account.Validate())

	account, err = octopusClient.Accounts.Add(account)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestCreateAndDeleteAndGetMinimalAzureSubscriptionAccount(t *testing.T) {
	account, err := model.NewAccount(name, enum.AzureSubscription)
	subscriptionNumber := uuid.New()
	account.SubscriptionNumber = &subscriptionNumber
	clientID := uuid.New()
	account.ClientID = &clientID

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.NoError(t, account.Validate())

	account, err = octopusClient.Accounts.Add(account)

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.NoError(t, account.Validate())
	assert.NotEmpty(t, account.ID)
	assert.Len(t, account.EnvironmentIDs, 0)
	assert.Len(t, account.TenantIDs, 0)
	assert.Len(t, account.TenantTags, 0)
	assert.Empty(t, account.AzureEnvironment)
	assert.Empty(t, account.Description)
	assert.NotEmpty(t, account.SpaceID)
	assert.NotEmpty(t, account.LastModifiedBy)
	assert.NotEmpty(t, account.LastModifiedOn)

	err = octopusClient.Accounts.Delete(account.ID)

	assert.NoError(t, err)

	account, err = octopusClient.Accounts.Get(account.ID)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestGetAll(t *testing.T) {
	accounts, err := octopusClient.Accounts.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, accounts)
}
