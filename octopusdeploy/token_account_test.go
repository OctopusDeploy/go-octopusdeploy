package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTokenAccountNew(t *testing.T) {
	name := getRandomName()
	token := NewSensitiveValue("new-value")

	tokenAccount := NewTokenAccount(name, token)

	require.NotNil(t, tokenAccount)
	require.NoError(t, tokenAccount.Validate())
	require.Equal(t, emptyString, tokenAccount.Description)
	require.Equal(t, emptyString, tokenAccount.GetDescription())
	require.Equal(t, emptyString, tokenAccount.ID)
	require.Equal(t, emptyString, tokenAccount.GetID())
	require.Equal(t, name, tokenAccount.Name)
	require.Equal(t, name, tokenAccount.GetName())
	require.Equal(t, accountTypeToken, tokenAccount.AccountType)
	require.Equal(t, accountTypeToken, tokenAccount.GetAccountType())
	require.NotNil(t, tokenAccount.Links)
	require.NotNil(t, tokenAccount.GetLinks())
}

func TestTokenAccountSetDescription(t *testing.T) {
	description := getRandomName()
	name := getRandomName()
	token := NewSensitiveValue("new-value")

	tokenAccount := NewTokenAccount(name, token)
	tokenAccount.Description = description

	require.NoError(t, tokenAccount.Validate())
	require.Equal(t, description, tokenAccount.Description)
	require.Equal(t, description, tokenAccount.GetDescription())
}

func TestTokenAccountSetName(t *testing.T) {
	name := getRandomName()
	token := NewSensitiveValue("new-value")

	tokenAccount := NewTokenAccount(name, token)

	require.NoError(t, tokenAccount.Validate())
	require.Equal(t, name, tokenAccount.Name)
	require.Equal(t, name, tokenAccount.GetName())
}

func TestTokenAccountTypes(t *testing.T) {
	name := getRandomName()
	token := NewSensitiveValue("new-value")

	account := NewTokenAccount(name, token)
	require.NoError(t, account.Validate())

	account.AccountType = "None"
	assert.Error(t, account.Validate())

	account.AccountType = "none"
	assert.Error(t, account.Validate())

	account.AccountType = accountTypeUsernamePassword
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
	assert.NoError(t, account.Validate())

	account.AccountType = "token"
	assert.Error(t, account.Validate())
}
