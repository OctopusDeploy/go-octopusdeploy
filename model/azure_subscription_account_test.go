package model

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAzureSubscriptionAccountNew(t *testing.T) {
	name := getRandomName()
	subscriptionID := uuid.New()

	azureSubscriptionAccount := NewAzureSubscriptionAccount(name, subscriptionID)

	require.NotNil(t, azureSubscriptionAccount)
	require.NoError(t, azureSubscriptionAccount.Validate())
	require.Equal(t, emptyString, azureSubscriptionAccount.Description)
	require.Equal(t, emptyString, azureSubscriptionAccount.GetDescription())
	require.Equal(t, emptyString, azureSubscriptionAccount.ID)
	require.Equal(t, emptyString, azureSubscriptionAccount.GetID())
	require.Equal(t, name, azureSubscriptionAccount.Name)
	require.Equal(t, name, azureSubscriptionAccount.GetName())
	require.Equal(t, accountTypeAzureSubscription, azureSubscriptionAccount.AccountType)
	require.Equal(t, accountTypeAzureSubscription, azureSubscriptionAccount.GetAccountType())
	require.NotNil(t, azureSubscriptionAccount.Links)
	require.NotNil(t, azureSubscriptionAccount.GetLinks())
}

func TestAzureSubscriptionAccountSetDescription(t *testing.T) {
	description := getRandomName()
	name := getRandomName()
	subscriptionID := uuid.New()

	azureSubscriptionAccount := NewAzureSubscriptionAccount(name, subscriptionID)
	azureSubscriptionAccount.Description = description

	require.NoError(t, azureSubscriptionAccount.Validate())
	require.Equal(t, description, azureSubscriptionAccount.Description)
	require.Equal(t, description, azureSubscriptionAccount.GetDescription())
}

func TestAzureSubscriptionAccountSetName(t *testing.T) {
	name := getRandomName()
	subscriptionID := uuid.New()

	azureSubscriptionAccount := NewAzureSubscriptionAccount(name, subscriptionID)

	require.NoError(t, azureSubscriptionAccount.Validate())
	require.Equal(t, name, azureSubscriptionAccount.Name)
	require.Equal(t, name, azureSubscriptionAccount.GetName())
}

func TestAzureSubscriptionAccountTypes(t *testing.T) {
	name := getRandomName()
	subscriptionID := uuid.New()

	azureSubscriptionAccount := NewAzureSubscriptionAccount(name, subscriptionID)
	require.NoError(t, azureSubscriptionAccount.Validate())

	azureSubscriptionAccount.AccountType = "None"
	assert.Error(t, azureSubscriptionAccount.Validate())

	azureSubscriptionAccount.AccountType = "none"
	assert.Error(t, azureSubscriptionAccount.Validate())

	azureSubscriptionAccount.AccountType = accountTypeAmazonWebServicesAccount
	assert.Error(t, azureSubscriptionAccount.Validate())

	azureSubscriptionAccount.AccountType = accountTypeAzureServicePrincipal
	assert.Error(t, azureSubscriptionAccount.Validate())

	azureSubscriptionAccount.AccountType = accountTypeUsernamePassword
	assert.Error(t, azureSubscriptionAccount.Validate())

	azureSubscriptionAccount.AccountType = accountTypeSshKeyPair
	assert.Error(t, azureSubscriptionAccount.Validate())

	azureSubscriptionAccount.AccountType = accountTypeToken
	assert.Error(t, azureSubscriptionAccount.Validate())

	azureSubscriptionAccount.AccountType = accountTypeAzureSubscription
	assert.NoError(t, azureSubscriptionAccount.Validate())

	azureSubscriptionAccount.AccountType = "azureSubscriptionAccount"
	assert.Error(t, azureSubscriptionAccount.Validate())
}
