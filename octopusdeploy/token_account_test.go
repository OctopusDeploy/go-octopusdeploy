package octopusdeploy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTokenAccountNew(t *testing.T) {
	accountType := AccountTypeToken
	description := emptyString
	environmentIDs := []string{}
	name := getRandomName()
	spaceID := emptyString
	tenantedDeploymentMode := TenantedDeploymentMode("Untenanted")
	token := NewSensitiveValue(getRandomName())

	account, err := NewTokenAccount(name, token)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// resource
	require.Equal(t, emptyString, account.ID)
	require.Equal(t, emptyString, account.ModifiedBy)
	require.Nil(t, account.ModifiedOn)
	require.NotNil(t, account.Links)

	// IResource
	require.Equal(t, emptyString, account.GetID())
	require.Equal(t, emptyString, account.GetModifiedBy())
	require.Nil(t, account.GetModifiedOn())
	require.NotNil(t, account.GetLinks())

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
	invalidModifiedBy := getRandomName()
	invalidModifiedOn := time.Now()
	invalidName := getRandomName()
	name := getRandomName()
	description := "Description for " + name + " (OK to Delete)"
	spaceID := getRandomName()
	tenantedDeploymentMode := TenantedDeploymentMode("Tenanted")
	token := NewSensitiveValue(getRandomName())

	options := func(a *TokenAccount) {
		a.Description = description
		a.EnvironmentIDs = environmentIDs
		a.ID = invalidID
		a.ModifiedBy = invalidModifiedBy
		a.ModifiedOn = &invalidModifiedOn
		a.Name = invalidName
		a.SpaceID = spaceID
		a.TenantedDeploymentMode = tenantedDeploymentMode
		a.Token = token
	}

	account, err := NewTokenAccount(name, token, options)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// resource
	require.Equal(t, emptyString, account.ID)
	require.Equal(t, emptyString, account.ModifiedBy)
	require.Nil(t, account.ModifiedOn)
	require.NotNil(t, account.Links)

	// IResource
	require.Equal(t, emptyString, account.GetID())
	require.Equal(t, emptyString, account.GetModifiedBy())
	require.Nil(t, account.GetModifiedOn())
	require.NotNil(t, account.GetLinks())

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
	token := NewSensitiveValue(getRandomName())

	account, err := NewTokenAccount(name, token)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	require.Equal(t, name, account.Name)
	require.Equal(t, name, account.GetName())
}
