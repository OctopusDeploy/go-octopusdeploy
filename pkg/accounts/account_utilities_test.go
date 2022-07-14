package accounts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/stretchr/testify/require"
)

func TestToAccount(t *testing.T) {
	var accountResource *AccountResource
	account, err := ToAccount(accountResource)
	require.Nil(t, account)
	require.Error(t, err)

	accountResource = NewAccountResource("", "")
	account, err = ToAccount(accountResource)
	require.Nil(t, account)
	require.Error(t, err)

	accountResource = NewAccountResource(internal.GetRandomName(), "")
	account, err = ToAccount(accountResource)
	require.Nil(t, account)
	require.Error(t, err)

	accountResource = NewAccountResource(internal.GetRandomName(), AccountTypeUsernamePassword)
	account, err = ToAccount(accountResource)
	require.NotNil(t, account)
	require.NoError(t, err)
	require.EqualValues(t, account.GetAccountType(), "UsernamePassword")

	account, err = ToAccountResource(accountResource)
	require.NotNil(t, account)
	require.NoError(t, err)
	require.EqualValues(t, account.GetAccountType(), "UsernamePassword")
}

func TestToAccountResource(t *testing.T) {
	var account IAccount
	accountResource, err := ToAccountResource(account)
	require.Nil(t, accountResource)
	require.Error(t, err)

	var amazonWebServicesAccount *AmazonWebServicesAccount
	accountResource, err = ToAccountResource(amazonWebServicesAccount)
	require.Nil(t, accountResource)
	require.Error(t, err)

	var azureServicePrincipalAccount *AzureServicePrincipalAccount
	accountResource, err = ToAccountResource(azureServicePrincipalAccount)
	require.Nil(t, accountResource)
	require.Error(t, err)

	var azureSubscriptionAccount *AzureSubscriptionAccount
	accountResource, err = ToAccountResource(azureSubscriptionAccount)
	require.Nil(t, accountResource)
	require.Error(t, err)

	var googleCloudPlatformAccount *GoogleCloudPlatformAccount
	accountResource, err = ToAccountResource(googleCloudPlatformAccount)
	require.Nil(t, accountResource)
	require.Error(t, err)

	var sshKeyAccount *SSHKeyAccount
	accountResource, err = ToAccountResource(sshKeyAccount)
	require.Nil(t, accountResource)
	require.Error(t, err)

	var tokenAccount *TokenAccount
	accountResource, err = ToAccountResource(tokenAccount)
	require.Nil(t, accountResource)
	require.Error(t, err)

	var usernamePasswordAccount *UsernamePasswordAccount
	accountResource, err = ToAccountResource(usernamePasswordAccount)
	require.Nil(t, accountResource)
	require.Error(t, err)
}
