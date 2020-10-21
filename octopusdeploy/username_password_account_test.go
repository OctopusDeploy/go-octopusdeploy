package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUsernamePasswordAccountNew(t *testing.T) {
	name := getRandomName()
	usernamePasswordAccount := NewUsernamePasswordAccount(name)

	require.NotNil(t, usernamePasswordAccount)
	require.NoError(t, usernamePasswordAccount.Validate())
	require.Equal(t, emptyString, usernamePasswordAccount.GetDescription())
	require.Equal(t, emptyString, usernamePasswordAccount.GetID())
	require.Equal(t, name, usernamePasswordAccount.GetName())
	require.Equal(t, accountTypeUsernamePassword, usernamePasswordAccount.GetAccountType())
	require.NotNil(t, usernamePasswordAccount.GetLinks())
}

func TestUsernamePasswordAccountSetDescription(t *testing.T) {
	description := getRandomName()
	name := getRandomName()

	usernamePasswordAccount := NewUsernamePasswordAccount(name)
	usernamePasswordAccount.Description = description

	require.NoError(t, usernamePasswordAccount.Validate())
	require.Equal(t, description, usernamePasswordAccount.Description)
	require.Equal(t, description, usernamePasswordAccount.GetDescription())
}

func TestUsernamePasswordAccountSetName(t *testing.T) {
	name := getRandomName()

	usernamePasswordAccount := NewUsernamePasswordAccount(name)

	require.NoError(t, usernamePasswordAccount.Validate())
	require.Equal(t, name, usernamePasswordAccount.Name)
	require.Equal(t, name, usernamePasswordAccount.GetName())
}

func TestUsernamePasswordAccountTypes(t *testing.T) {
	name := getRandomName()

	account := NewUsernamePasswordAccount(name)
	require.NoError(t, account.Validate())

	account.AccountType = "None"
	assert.Error(t, account.Validate())

	account.AccountType = "none"
	assert.Error(t, account.Validate())

	account.AccountType = accountTypeAzureSubscription
	assert.Error(t, account.Validate())

	account.AccountType = accountTypeAzureServicePrincipal
	assert.Error(t, account.Validate())

	account.AccountType = accountTypeAmazonWebServicesAccount
	assert.Error(t, account.Validate())

	account.AccountType = accountTypeSshKeyPair
	assert.Error(t, account.Validate())

	account.AccountType = accountTypeToken
	assert.Error(t, account.Validate())

	account.AccountType = accountTypeUsernamePassword
	assert.NoError(t, account.Validate())

	account.AccountType = "usernamePassword"
	assert.Error(t, account.Validate())
}
