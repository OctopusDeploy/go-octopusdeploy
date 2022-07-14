package accounts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/stretchr/testify/require"
)

func TestSSHKeyAccountNew(t *testing.T) {
	accountType := AccountTypeSSHKeyPair
	description := ""
	environmentIDs := []string{}
	name := internal.GetRandomName()
	privateKeyFile := core.NewSensitiveValue(internal.GetRandomName())
	var privateKeyPassphrase *core.SensitiveValue
	spaceID := ""
	tenantedDeploymentMode := core.TenantedDeploymentMode("Untenanted")
	username := internal.GetRandomName()

	account, err := NewSSHKeyAccount(name, username, privateKeyFile)

	require.NoError(t, err)
	require.NotNil(t, account)
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

	// SSHKeyAccount
	require.Equal(t, privateKeyFile, account.PrivateKeyFile)
	require.Equal(t, privateKeyPassphrase, account.PrivateKeyPassphrase)
	require.Equal(t, username, account.Username)
}

func TestSSHKeyAccountNewWithConfigs(t *testing.T) {
	name := internal.GetRandomName()
	privateKeyFile := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	account, err := NewSSHKeyAccount(name, username, privateKeyFile)
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

	description := internal.GetRandomName()
	account.Description = description

	// IAccount
	require.Equal(t, AccountTypeSSHKeyPair, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())
}
