package accounts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccount(t *testing.T) {
	a := &account{}
	name := internal.GetRandomName()

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
		TenantedDeploymentMode: core.TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.NoError(t, a.Validate())

	a = &account{
		AccountType:            AccountTypeUsernamePassword,
		Name:                   "All",
		TenantedDeploymentMode: core.TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = &account{
		AccountType:            AccountTypeUsernamePassword,
		Name:                   "all",
		TenantedDeploymentMode: core.TenantedDeploymentMode("Untenanted"),
	}

	require.NotNil(t, a)
	assert.Error(t, a.Validate())

	a = newAccount(name, AccountTypeAzureServicePrincipal)
	require.NoError(t, a.Validate())
}
