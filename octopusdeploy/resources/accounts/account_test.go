package accounts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	a := &account{}
	name := octopusdeploy.getRandomName()

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &account{
		Name: name,
	}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &account{AccountType: AccountTypeUsernamePassword}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &account{
		AccountType:            AccountTypeUsernamePassword,
		Name:                   name,
		TenantedDeploymentMode: resources.TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.NoError(t, a.Validate())

	a = &account{
		AccountType:            AccountTypeUsernamePassword,
		Name:                   "All",
		TenantedDeploymentMode: resources.TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &account{
		AccountType:            AccountTypeUsernamePassword,
		Name:                   "all",
		TenantedDeploymentMode: resources.TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = newAccount(name, AccountTypeAzureServicePrincipal)
	require.NoError(t, a.Validate())
}
