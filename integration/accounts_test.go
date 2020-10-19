package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountServiceAdd(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	account, err := octopusClient.Accounts.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestAccountServiceGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	accounts, err := octopusClient.Accounts.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, accounts)

	if len(accounts) == 0 {
		return
	}

	for _, account := range accounts {
		assert.NoError(t, err)
		assert.NotEmpty(t, account)
	}
}

func TestAccountServiceGetByName(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	accounts, err := octopusClient.Accounts.GetByPartialName("go")

	assert.NoError(t, err)
	assert.NotNil(t, accounts)

	if len(accounts) == 0 {
		return
	}

	for _, account := range accounts {
		assert.NoError(t, err)
		assert.NotEmpty(t, account)
	}
}

func TestAccountServiceGetByType(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	accounts, err := octopusClient.Accounts.GetByAccountType("UsernamePassword")

	assert.NoError(t, err)
	assert.NotNil(t, accounts)

	if len(accounts) == 0 {
		return
	}

	for _, account := range accounts {
		assert.NoError(t, err)
		assert.NotEmpty(t, account)
	}
}
