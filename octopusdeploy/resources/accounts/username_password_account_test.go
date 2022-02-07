package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUsernamePasswordAccountNew(t *testing.T) {
	accountType := AccountTypeUsernamePassword
	description := services.emptyString
	environmentIDs := []string{}
	name := octopusdeploy.getRandomName()
	var password *octopusdeploy.SensitiveValue
	spaceID := services.emptyString
	tenantedDeploymentMode := octopusdeploy.TenantedDeploymentMode("Untenanted")
	username := services.emptyString

	account, err := NewUsernamePasswordAccount(name)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// Resource
	require.Equal(t, services.emptyString, account.ID)
	require.Equal(t, services.emptyString, account.ModifiedBy)
	require.Nil(t, account.ModifiedOn)
	require.NotNil(t, account.Links)

	// IResource
	require.Equal(t, services.emptyString, account.GetID())
	require.Equal(t, services.emptyString, account.GetModifiedBy())
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

	// UsernamePasswordAccount
	require.Equal(t, password, account.Password)
	require.Equal(t, username, account.Username)
}

func TestUsernamePasswordAccountNewWithConfigs(t *testing.T) {
	environmentIDs := []string{"environment-id-1", "environment-id-2"}
	invalidID := octopusdeploy.getRandomName()
	invalidModifiedBy := octopusdeploy.getRandomName()
	invalidModifiedOn := time.Now()
	invalidName := octopusdeploy.getRandomName()
	name := octopusdeploy.getRandomName()
	description := "Description for " + name + " (OK to Delete)"
	password := octopusdeploy.NewSensitiveValue("password")
	spaceID := octopusdeploy.getRandomName()
	tenantedDeploymentMode := octopusdeploy.TenantedDeploymentMode("Tenanted")
	username := octopusdeploy.getRandomName()

	options := func(a *UsernamePasswordAccount) {
		a.Description = description
		a.EnvironmentIDs = environmentIDs
		a.ID = invalidID
		a.ModifiedBy = invalidModifiedBy
		a.ModifiedOn = &invalidModifiedOn
		a.Name = invalidName
		a.Password = octopusdeploy.NewSensitiveValue("password")
		a.SpaceID = spaceID
		a.TenantedDeploymentMode = tenantedDeploymentMode
		a.Username = username
	}

	account, err := NewUsernamePasswordAccount(name, options)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// Resource
	require.Equal(t, services.emptyString, account.ID)
	require.Equal(t, services.emptyString, account.ModifiedBy)
	require.Nil(t, account.ModifiedOn)
	require.NotNil(t, account.Links)

	// IResource
	require.Equal(t, services.emptyString, account.GetID())
	require.Equal(t, services.emptyString, account.GetModifiedBy())
	require.Nil(t, account.GetModifiedOn())
	require.NotNil(t, account.GetLinks())

	// account
	require.Equal(t, description, account.Description)
	require.Equal(t, environmentIDs, account.EnvironmentIDs)
	require.Equal(t, name, account.Name)
	require.Equal(t, spaceID, account.SpaceID)
	require.Equal(t, tenantedDeploymentMode, account.TenantedDeploymentMode)

	// IAccount
	require.Equal(t, AccountTypeUsernamePassword, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// UsernamePasswordAccount
	require.Equal(t, password, account.Password)
	require.Equal(t, username, account.Username)
}

func TestUsernamePasswordAccountSetName(t *testing.T) {
	name := octopusdeploy.getRandomName()

	account, err := NewUsernamePasswordAccount(name)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	require.Equal(t, name, account.Name)
	require.Equal(t, name, account.GetName())
}
