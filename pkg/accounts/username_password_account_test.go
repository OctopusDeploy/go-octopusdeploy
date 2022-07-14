package accounts

import (
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/stretchr/testify/require"
)

func TestUsernamePasswordAccountNew(t *testing.T) {
	accountType := AccountTypeUsernamePassword
	description := ""
	environmentIDs := []string{}
	name := internal.GetRandomName()
	var password *core.SensitiveValue
	spaceID := ""
	tenantedDeploymentMode := core.TenantedDeploymentMode("Untenanted")
	username := ""

	account, err := NewUsernamePasswordAccount(name)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// resource
	require.Equal(t, "", account.ID)
	require.Equal(t, "", account.ModifiedBy)
	require.Nil(t, account.ModifiedOn)
	require.NotNil(t, account.Links)

	// IResource
	require.Equal(t, "", account.GetID())
	require.Equal(t, "", account.GetModifiedBy())
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
	invalidID := internal.GetRandomName()
	invalidModifiedBy := internal.GetRandomName()
	invalidModifiedOn := time.Now()
	invalidName := internal.GetRandomName()
	name := internal.GetRandomName()
	description := "Description for " + name + " (OK to Delete)"
	password := core.NewSensitiveValue("password")
	spaceID := internal.GetRandomName()
	tenantedDeploymentMode := core.TenantedDeploymentMode("Tenanted")
	username := internal.GetRandomName()

	options := func(a *UsernamePasswordAccount) {
		a.Description = description
		a.EnvironmentIDs = environmentIDs
		a.ID = invalidID
		a.ModifiedBy = invalidModifiedBy
		a.ModifiedOn = &invalidModifiedOn
		a.Name = invalidName
		a.Password = core.NewSensitiveValue("password")
		a.SpaceID = spaceID
		a.TenantedDeploymentMode = tenantedDeploymentMode
		a.Username = username
	}

	account, err := NewUsernamePasswordAccount(name, options)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// resource
	require.Equal(t, "", account.ID)
	require.Equal(t, "", account.ModifiedBy)
	require.Nil(t, account.ModifiedOn)
	require.NotNil(t, account.Links)

	// IResource
	require.Equal(t, "", account.GetID())
	require.Equal(t, "", account.GetModifiedBy())
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
	name := internal.GetRandomName()

	account, err := NewUsernamePasswordAccount(name)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	require.Equal(t, name, account.Name)
	require.Equal(t, name, account.GetName())
}
