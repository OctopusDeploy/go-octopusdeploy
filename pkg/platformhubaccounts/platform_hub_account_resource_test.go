package platformhubaccounts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/stretchr/testify/require"
)

func TestToPlatformHubAccount(t *testing.T) {
	// Test with nil resource
	var accountResource *PlatformHubAccountResource
	account, err := accountResource.ToPlatformHubAccount()
	require.Nil(t, account)
	require.Error(t, err)

	// Test with empty name and account type
	accountResource = NewPlatformHubAccountResource("", "")
	account, err = accountResource.ToPlatformHubAccount()
	require.Nil(t, account)
	require.Error(t, err)

	// Test with name but empty account type
	accountResource = NewPlatformHubAccountResource(internal.GetRandomName(), "")
	account, err = accountResource.ToPlatformHubAccount()
	require.Nil(t, account)
	require.Error(t, err)

	// Test with unknown account type (default case)
	accountResource = NewPlatformHubAccountResource(internal.GetRandomName(), PlatformHubAccountType("UnknownAccountType"))
	account, err = accountResource.ToPlatformHubAccount()
	require.Nil(t, account)
	require.Error(t, err)

	// Test with valid AWS account
	accessKey := internal.GetRandomName()
	name := internal.GetRandomName()
	secretKey := core.NewSensitiveValue(internal.GetRandomName())

	accountResource = NewPlatformHubAccountResource(name, AccountTypePlatformHubAwsAccount)
	accountResource.AccessKey = accessKey
	accountResource.SecretKey = secretKey
	accountResource.Description = "Test description"
	accountResource.ID = "test-id"

	account, err = accountResource.ToPlatformHubAccount()
	require.NotNil(t, account)
	require.NoError(t, err)
	require.Equal(t, AccountTypePlatformHubAwsAccount, account.GetAccountType())
	require.Equal(t, name, account.GetName())
	require.Equal(t, "Test description", account.GetDescription())
	require.Equal(t, "test-id", account.GetID())

	// Verify it's an AWS account
	awsAccount, ok := account.(*PlatformHubAwsAccount)
	require.True(t, ok)
	require.Equal(t, accessKey, awsAccount.AccessKey)
	require.Equal(t, secretKey, awsAccount.SecretKey)
}

func TestToPlatformHubAccountResource(t *testing.T) {
	// Test with nil account
	accountResource, err := ToPlatformHubAccountResource(nil)
	require.Nil(t, accountResource)
	require.Error(t, err)

	// Test with valid AWS account
	accessKey := internal.GetRandomName()
	name := internal.GetRandomName()
	secretKey := core.NewSensitiveValue(internal.GetRandomName())

	awsAccount, err := NewPlatformHubAwsAccount(name, accessKey, secretKey)
	require.NoError(t, err)
	require.NotNil(t, awsAccount)

	awsAccount.Description = "Test description"
	awsAccount.ID = "test-id"

	accountResource, err = ToPlatformHubAccountResource(awsAccount)
	require.NotNil(t, accountResource)
	require.NoError(t, err)
	require.Equal(t, AccountTypePlatformHubAwsAccount, accountResource.GetAccountType())
	require.Equal(t, name, accountResource.GetName())
	require.Equal(t, "Test description", accountResource.GetDescription())
	require.Equal(t, "test-id", accountResource.GetID())
	require.Equal(t, accessKey, accountResource.AccessKey)
	require.Equal(t, secretKey, accountResource.SecretKey)

	// Test that converting a resource returns it as-is
	accountResource2, err := ToPlatformHubAccountResource(accountResource)
	require.NotNil(t, accountResource2)
	require.NoError(t, err)
	require.Equal(t, accountResource, accountResource2)
}

func TestToPlatformHubAccountResourceUnknownType(t *testing.T) {
	// Create a resource with an unknown account type
	name := internal.GetRandomName()
	accountResource := NewPlatformHubAccountResource(name, PlatformHubAccountType("UnknownAccountType"))

	// This should fail when trying to convert back to account
	account, err := accountResource.ToPlatformHubAccount()
	require.Nil(t, account)
	require.Error(t, err)
}
