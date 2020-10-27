package octopusdeploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	a := &account{}
	name := getRandomName()

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &account{
		Name: name,
	}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &account{accountType: AccountTypeUsernamePassword}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &account{
		accountType:            AccountTypeUsernamePassword,
		Name:                   name,
		TenantedDeploymentMode: TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.NoError(t, a.Validate())

	a = &account{
		accountType:            AccountTypeUsernamePassword,
		Name:                   "All",
		TenantedDeploymentMode: TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &account{
		accountType:            AccountTypeUsernamePassword,
		Name:                   "all",
		TenantedDeploymentMode: TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = newAccount(name, AccountTypeAzureServicePrincipal)
	require.NoError(t, a.Validate())
}
