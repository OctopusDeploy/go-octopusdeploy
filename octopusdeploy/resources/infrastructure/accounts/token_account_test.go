package accounts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"

	"github.com/stretchr/testify/require"
)

func TestTokenAccountNew(t *testing.T) {
	accountType := AccountTypeToken
	description := ""
	environmentIDs := []string{}
	name := getRandomName()
	spaceID := ""
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Untenanted")
	token := resources.NewSensitiveValue(getRandomName())

	account, err := NewTokenAccount(name, token)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// Resource
	require.Equal(t, "", account.ID)

	// IResource
	require.Equal(t, "", account.GetID())

	// account
	require.Equal(t, description, account.Description)
	require.Equal(t, environmentIDs, account.EnvironmentIDs)
	require.Equal(t, name, account.Name)
	require.Equal(t, spaceID, account.SpaceID)
	require.Equal(t, tenantedDeploymentMode, account.TenantedDeploymentMode)

	// IAccount
	require.Equal(t, accountType, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// TokenAccount
	require.Equal(t, token, account.Token)
}

func TestTokenAccountNewWithConfigs(t *testing.T) {
	environmentIDs := []string{getRandomName(), getRandomName()}
	invalidID := getRandomName()
	invalidName := getRandomName()
	name := getRandomName()
	description := "Description for " + name + " (OK to Delete)"
	spaceID := getRandomName()
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Tenanted")
	token := resources.NewSensitiveValue(getRandomName())

	options := func(a *TokenAccount) {
		a.Description = description
		a.EnvironmentIDs = environmentIDs
		a.ID = invalidID
		a.Name = invalidName
		a.SpaceID = spaceID
		a.TenantedDeploymentMode = tenantedDeploymentMode
		a.Token = token
	}

	account, err := NewTokenAccount(name, token, options)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// Resource
	require.Equal(t, "", account.ID)

	// IResource
	require.Equal(t, "", account.GetID())

	// account
	require.Equal(t, description, account.Description)
	require.Equal(t, environmentIDs, account.EnvironmentIDs)
	require.Equal(t, name, account.Name)
	require.Equal(t, spaceID, account.SpaceID)
	require.Equal(t, tenantedDeploymentMode, account.TenantedDeploymentMode)

	// IAccount
	require.Equal(t, AccountTypeToken, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// TokenAccount
	require.Equal(t, token, account.Token)
}

func TestTokenAccountSetName(t *testing.T) {
	name := getRandomName()
	token := resources.NewSensitiveValue(getRandomName())

	account, err := NewTokenAccount(name, token)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	require.Equal(t, name, account.Name)
	require.Equal(t, name, account.GetName())
}
