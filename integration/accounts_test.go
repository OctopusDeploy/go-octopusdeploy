package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	if octopusClient == nil {
		octopusClient = initTest()
	}
}

func TestAddAccount(t *testing.T) {
	account, err := octopusClient.Accounts.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestGetAllAccounts(t *testing.T) {
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

func TestGetUsage(t *testing.T) {
	accounts, err := octopusClient.Accounts.GetAll()

	assert.NoError(t, err)

	for _, account := range accounts {
		accountUsages, err := octopusClient.Accounts.GetUsage(account)

		assert.NoError(t, err)
		assert.NotNil(t, accountUsages)
	}
}
