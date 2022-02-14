package accounts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"

	"github.com/stretchr/testify/require"
)

func TestSSHKeyAccountNew(t *testing.T) {
	accountType := AccountTypeSSHKeyPair
	description := ""
	environmentIDs := []string{}
	name := getRandomName()
	privateKeyFile := resources.NewSensitiveValue(getRandomName())
	var privateKeyPassphrase *resources.SensitiveValue
	spaceID := ""
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Untenanted")
	username := getRandomName()

	account, err := NewSSHKeyAccount(name, username, privateKeyFile)

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

	// SSHKeyAccount
	require.Equal(t, privateKeyFile, account.PrivateKeyFile)
	require.Equal(t, privateKeyPassphrase, account.PrivateKeyPassphrase)
	require.Equal(t, username, account.Username)
}

func TestSSHKeyAccountNewWithConfigs(t *testing.T) {
	environmentIDs := []string{getRandomName(), getRandomName()}
	invalidID := getRandomName()
	invalidName := getRandomName()
	name := getRandomName()
	description := "Description for " + name + " (OK to Delete)"
	privateKeyFile := resources.NewSensitiveValue(getRandomName())
	privateKeyPassphrase := resources.NewSensitiveValue(getRandomName())
	spaceID := getRandomName()
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Tenanted")
	username := getRandomName()

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

	// account
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
