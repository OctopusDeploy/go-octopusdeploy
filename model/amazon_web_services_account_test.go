package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAmazonWebServicesAccountNew(t *testing.T) {
	accessKey := getRandomName()
	name := getRandomName()
	secretKey := NewSensitiveValue(getRandomName())

	amazonWebServicesAccount := NewAmazonWebServicesAccount(name, accessKey, secretKey)

	require.NotNil(t, amazonWebServicesAccount)
	require.NoError(t, amazonWebServicesAccount.Validate())
	require.Equal(t, emptyString, amazonWebServicesAccount.Description)
	require.Equal(t, emptyString, amazonWebServicesAccount.GetDescription())
	require.Equal(t, emptyString, amazonWebServicesAccount.ID)
	require.Equal(t, emptyString, amazonWebServicesAccount.GetID())
	require.Equal(t, name, amazonWebServicesAccount.Name)
	require.Equal(t, name, amazonWebServicesAccount.GetName())
	require.Equal(t, accountTypeAmazonWebServicesAccount, amazonWebServicesAccount.AccountType)
	require.Equal(t, accountTypeAmazonWebServicesAccount, amazonWebServicesAccount.GetAccountType())
	require.NotNil(t, amazonWebServicesAccount.Links)
	require.NotNil(t, amazonWebServicesAccount.GetLinks())
}

func TestAmazonWebServicesAccountSetDescription(t *testing.T) {
	accessKey := getRandomName()
	description := getRandomName()
	name := getRandomName()
	secretKey := NewSensitiveValue(getRandomName())

	amazonWebServicesAccount := NewAmazonWebServicesAccount(name, accessKey, secretKey)
	amazonWebServicesAccount.Description = description

	require.NoError(t, amazonWebServicesAccount.Validate())
	require.Equal(t, description, amazonWebServicesAccount.Description)
	require.Equal(t, description, amazonWebServicesAccount.GetDescription())
}

func TestAmazonWebServicesAccountSetName(t *testing.T) {
	accessKey := getRandomName()
	name := getRandomName()
	secretKey := NewSensitiveValue(getRandomName())

	amazonWebServicesAccount := NewAmazonWebServicesAccount(name, accessKey, secretKey)

	require.NoError(t, amazonWebServicesAccount.Validate())
	require.Equal(t, name, amazonWebServicesAccount.Name)
	require.Equal(t, name, amazonWebServicesAccount.GetName())
}

func TestAmazonWebServicesAccountTypes(t *testing.T) {
	accessKey := getRandomName()
	name := getRandomName()
	secretKey := NewSensitiveValue(getRandomName())

	amazonWebServicesAccount := NewAmazonWebServicesAccount(name, accessKey, secretKey)
	require.NoError(t, amazonWebServicesAccount.Validate())

	amazonWebServicesAccount.AccountType = "None"
	assert.Error(t, amazonWebServicesAccount.Validate())

	amazonWebServicesAccount.AccountType = "none"
	assert.Error(t, amazonWebServicesAccount.Validate())

	amazonWebServicesAccount.AccountType = accountTypeAzureSubscription
	assert.Error(t, amazonWebServicesAccount.Validate())

	amazonWebServicesAccount.AccountType = accountTypeAzureServicePrincipal
	assert.Error(t, amazonWebServicesAccount.Validate())

	amazonWebServicesAccount.AccountType = accountTypeUsernamePassword
	assert.Error(t, amazonWebServicesAccount.Validate())

	amazonWebServicesAccount.AccountType = accountTypeSshKeyPair
	assert.Error(t, amazonWebServicesAccount.Validate())

	amazonWebServicesAccount.AccountType = accountTypeToken
	assert.Error(t, amazonWebServicesAccount.Validate())

	amazonWebServicesAccount.AccountType = accountTypeAmazonWebServicesAccount
	assert.NoError(t, amazonWebServicesAccount.Validate())

	amazonWebServicesAccount.AccountType = "amazonWebServicesAccount"
	assert.Error(t, amazonWebServicesAccount.Validate())
}
