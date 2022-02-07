package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSSHKeyAccountNew(t *testing.T) {
	accountType := AccountTypeSSHKeyPair
	description := services.emptyString
	environmentIDs := []string{}
	name := octopusdeploy.getRandomName()
	privateKeyFile := octopusdeploy.NewSensitiveValue(octopusdeploy.getRandomName())
	var privateKeyPassphrase *octopusdeploy.SensitiveValue
	spaceID := services.emptyString
	tenantedDeploymentMode := octopusdeploy.TenantedDeploymentMode("Untenanted")
	username := octopusdeploy.getRandomName()

	account, err := NewSSHKeyAccount(name, username, privateKeyFile)

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

	// SSHKeyAccount
	require.Equal(t, privateKeyFile, account.PrivateKeyFile)
	require.Equal(t, privateKeyPassphrase, account.PrivateKeyPassphrase)
	require.Equal(t, username, account.Username)
}

func TestSSHKeyAccountNewWithConfigs(t *testing.T) {
	environmentIDs := []string{octopusdeploy.getRandomName(), octopusdeploy.getRandomName()}
	invalidID := octopusdeploy.getRandomName()
	invalidModifiedBy := octopusdeploy.getRandomName()
	invalidModifiedOn := time.Now()
	invalidName := octopusdeploy.getRandomName()
	name := octopusdeploy.getRandomName()
	description := "Description for " + name + " (OK to Delete)"
	privateKeyFile := octopusdeploy.NewSensitiveValue(octopusdeploy.getRandomName())
	privateKeyPassphrase := octopusdeploy.NewSensitiveValue(octopusdeploy.getRandomName())
	spaceID := octopusdeploy.getRandomName()
	tenantedDeploymentMode := octopusdeploy.TenantedDeploymentMode("Tenanted")
	username := octopusdeploy.getRandomName()

	options := func(a *SSHKeyAccount) {
		a.Description = description
		a.EnvironmentIDs = environmentIDs
		a.ID = invalidID
		a.ModifiedBy = invalidModifiedBy
		a.ModifiedOn = &invalidModifiedOn
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
	require.Equal(t, AccountTypeSSHKeyPair, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// SSHKeyAccount
	require.Equal(t, privateKeyFile, account.PrivateKeyFile)
	require.Equal(t, privateKeyPassphrase, account.PrivateKeyPassphrase)
	require.Equal(t, username, account.Username)
}
