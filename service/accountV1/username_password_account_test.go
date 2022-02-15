package accountV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"

	"github.com/stretchr/testify/require"
)

func TestUsernamePasswordAccountNew(t *testing.T) {
	accountType := AccountTypeUsernamePassword
	description := ""
	environmentIDs := []string{}
	name := internal.getRandomName()
	var password *resources.SensitiveValue
	spaceID := ""
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Untenanted")
	username := ""

	account, err := NewUsernamePasswordAccount(name)

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

	// UsernamePasswordAccount
	require.Equal(t, password, account.Password)
	require.Equal(t, username, account.Username)
}

func TestUsernamePasswordAccountNewWithConfigs(t *testing.T) {
	environmentIDs := []string{"environment-id-1", "environment-id-2"}
	invalidID := internal.getRandomName()
	invalidName := internal.getRandomName()
	name := internal.getRandomName()
	description := "Description for " + name + " (OK to Delete)"
	password := resources.NewSensitiveValue("password")
	spaceID := internal.getRandomName()
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Tenanted")
	username := internal.getRandomName()

	options := func(a *UsernamePasswordAccount) {
		a.Description = description
		a.EnvironmentIDs = environmentIDs
		a.ID = invalidID
		a.Name = invalidName
		a.Password = resources.NewSensitiveValue("password")
		a.SpaceID = spaceID
		a.TenantedDeploymentMode = tenantedDeploymentMode
		a.Username = username
	}

	account, err := NewUsernamePasswordAccount(name, options)

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
	require.Equal(t, AccountTypeUsernamePassword, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// UsernamePasswordAccount
	require.Equal(t, password, account.Password)
	require.Equal(t, username, account.Username)
}

func TestUsernamePasswordAccountSetName(t *testing.T) {
	name := internal.getRandomName()

	account, err := NewUsernamePasswordAccount(name)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	require.Equal(t, name, account.Name)
	require.Equal(t, name, account.GetName())
}
