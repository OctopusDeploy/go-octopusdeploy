package accountV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"

	"github.com/stretchr/testify/require"
)

func TestTokenAccountNew(t *testing.T) {
	accountType := AccountTypeToken
	description := ""
	environmentIDs := []string{}
	name := internal.getRandomName()
	spaceID := ""
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Untenanted")
	token := resources.NewSensitiveValue(internal.getRandomName())

	account, err := NewTokenAccount(name, token)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// Resource
	require.Equal(t, "", account.ID)

	// IResource
	require.Equal(t, "", account.GetID())

	// accountV1
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
	environmentIDs := []string{internal.getRandomName(), internal.getRandomName()}
	invalidID := internal.getRandomName()
	invalidName := internal.getRandomName()
	name := internal.getRandomName()
	description := "Description for " + name + " (OK to Delete)"
	spaceID := internal.getRandomName()
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Tenanted")
	token := resources.NewSensitiveValue(internal.getRandomName())

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

	// accountV1
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
	name := internal.getRandomName()
	token := resources.NewSensitiveValue(internal.getRandomName())

	account, err := NewTokenAccount(name, token)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	require.Equal(t, name, account.Name)
	require.Equal(t, name, account.GetName())
}
