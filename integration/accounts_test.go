package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"

	"github.com/stretchr/testify/assert"
)

func TestAccountServiceAdd(t *testing.T) {
	octopusClient := getOctopusClient()

	account, err := octopusClient.Accounts.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestAccountServiceGetAll(t *testing.T) {
	octopusClient := getOctopusClient()

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

	accounts, err := octopusClient.Accounts.GetByAccountType(enum.UsernamePassword)

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
