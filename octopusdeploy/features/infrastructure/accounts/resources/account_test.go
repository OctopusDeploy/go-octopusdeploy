package resources

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	a := &Account{}
	name := getRandomName()

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &Account{
		Name: name,
	}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &Account{AccountType: AccountTypeUsernamePassword}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &Account{
		AccountType:            AccountTypeUsernamePassword,
		Name:                   name,
		TenantedDeploymentMode: resources.TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.NoError(t, a.Validate())

	a = &Account{
		AccountType:            AccountTypeUsernamePassword,
		Name:                   "All",
		TenantedDeploymentMode: resources.TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &Account{
		AccountType:            AccountTypeUsernamePassword,
		Name:                   "all",
		TenantedDeploymentMode: resources.TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = NewAccount(name, AccountTypeAzureServicePrincipal)
	require.NoError(t, a.Validate())
}
