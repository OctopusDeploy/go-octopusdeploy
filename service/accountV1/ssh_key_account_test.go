package accountV1

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"

	"github.com/stretchr/testify/require"
)

func TestSSHKeyAccountNew(t *testing.T) {
	accountType := AccountTypeSSHKeyPair
	description := ""
	environmentIDs := []string{}
	name := internal.getRandomName()
	privateKeyFile := resources.NewSensitiveValue(internal.getRandomName())
	var privateKeyPassphrase *resources.SensitiveValue
	spaceID := ""
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Untenanted")
	username := internal.getRandomName()

	account, err := NewSSHKeyAccount(name, username, privateKeyFile)

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

	// SSHKeyAccount
	require.Equal(t, privateKeyFile, account.PrivateKeyFile)
	require.Equal(t, privateKeyPassphrase, account.PrivateKeyPassphrase)
	require.Equal(t, username, account.Username)
}

func TestSSHKeyAccountNewWithConfigs(t *testing.T) {
	environmentIDs := []string{internal.getRandomName(), internal.getRandomName()}
	invalidID := internal.getRandomName()
	invalidName := internal.getRandomName()
	name := internal.getRandomName()
	description := "Description for " + name + " (OK to Delete)"
	privateKeyFile := resources.NewSensitiveValue(internal.getRandomName())
	privateKeyPassphrase := resources.NewSensitiveValue(internal.getRandomName())
	spaceID := internal.getRandomName()
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Tenanted")
	username := internal.getRandomName()

	options := func(a *SSHKeyAccount) {
		a.Description = description
		a.EnvironmentIDs = environmentIDs
		a.ID = invalidID
		a.Name = invalidName
		a.PrivateKeyFile = privateKeyFile
		a.PrivateKeyPassphrase = privateKeyPassphrase
		a.SpaceID = spaceID
		a.TenantedDeploymentMode = tenantedDeploymentMode
		a.Username = username
	}

	account, err := NewSSHKeyAccount(name, username, privateKeyFile, options)
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
	require.Equal(t, AccountTypeSSHKeyPair, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// SSHKeyAccount
	require.Equal(t, privateKeyFile, account.PrivateKeyFile)
	require.Equal(t, privateKeyPassphrase, account.PrivateKeyPassphrase)
	require.Equal(t, username, account.Username)
}
