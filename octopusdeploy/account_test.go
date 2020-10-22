package octopusdeploy

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	account := &Account{}
	applicationID := uuid.New()
	name := getRandomName()
	subscriptionID := uuid.New()
	tenantID := uuid.New()

	require.NotNil(t, account)
	assert.Error(t, account.Validate())

	account = &Account{
		Name: name,
	}

	require.NotNil(t, account)
	assert.Error(t, account.Validate())

	account = &Account{AccountType: accountTypeUsernamePassword}

	require.NotNil(t, account)
	assert.Error(t, account.Validate())

	account = &Account{
		AccountType:            accountTypeUsernamePassword,
		Name:                   name,
		TenantedDeploymentMode: "Untenanted",
	}

	require.NotNil(t, account)
	assert.NoError(t, account.Validate())

	account = &Account{
		AccountType:            accountTypeUsernamePassword,
		Name:                   "All",
		TenantedDeploymentMode: "Untenanted",
	}

	require.NotNil(t, account)
	assert.Error(t, account.Validate())

	account = &Account{
		AccountType:            accountTypeUsernamePassword,
		Name:                   "all",
		TenantedDeploymentMode: "Untenanted",
	}

	require.NotNil(t, account)
	assert.Error(t, account.Validate())

	account = NewAccount(name, accountTypeAzureServicePrincipal)
	require.NoError(t, account.Validate())

	account = NewAccount(name, accountTypeAzureServicePrincipal)
	account.SubscriptionID = &subscriptionID
	require.NoError(t, account.Validate())

	account = NewAccount(name, accountTypeAzureServicePrincipal)
	account.SubscriptionID = &subscriptionID
	assert.NoError(t, account.Validate())

	account = NewAccount(name, accountTypeAzureServicePrincipal)
	account.SubscriptionID = &subscriptionID
	account.ApplicationID = &applicationID
	account.TenantID = &tenantID
	assert.NoError(t, account.Validate())
}
