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

	// Test with valid GCP account
	jsonKey := core.NewSensitiveValue(internal.GetRandomName())
	name = internal.GetRandomName()

	accountResource = NewPlatformHubAccountResource(name, AccountTypePlatformHubGcpAccount)
	accountResource.JsonKey = jsonKey
	accountResource.Description = "Test GCP description"
	accountResource.ID = "test-gcp-id"

	account, err = accountResource.ToPlatformHubAccount()
	require.NotNil(t, account)
	require.NoError(t, err)
	require.Equal(t, AccountTypePlatformHubGcpAccount, account.GetAccountType())
	require.Equal(t, name, account.GetName())
	require.Equal(t, "Test GCP description", account.GetDescription())
	require.Equal(t, "test-gcp-id", account.GetID())

	// Verify it's a GCP account
	gcpAccount, ok := account.(*PlatformHubGcpAccount)
	require.True(t, ok)
	require.Equal(t, jsonKey, gcpAccount.JsonKey)

	// Test with valid Generic OIDC account
	name = internal.GetRandomName()
	executionSubjectKeys := []string{"space", "project"}
	audience := "api://default"

	accountResource = NewPlatformHubAccountResource(name, AccountTypePlatformHubGenericOidcAccount)
	accountResource.ExecutionSubjectKeys = executionSubjectKeys
	accountResource.Audience = audience
	accountResource.Description = "Test Generic OIDC description"
	accountResource.ID = "test-oidc-id"

	account, err = accountResource.ToPlatformHubAccount()
	require.NotNil(t, account)
	require.NoError(t, err)
	require.Equal(t, AccountTypePlatformHubGenericOidcAccount, account.GetAccountType())
	require.Equal(t, name, account.GetName())
	require.Equal(t, "Test Generic OIDC description", account.GetDescription())
	require.Equal(t, "test-oidc-id", account.GetID())

	// Verify it's a Generic OIDC account
	oidcAccount, ok := account.(*PlatformHubGenericOidcAccount)
	require.True(t, ok)
	require.Equal(t, executionSubjectKeys, oidcAccount.ExecutionSubjectKeys)
	require.Equal(t, audience, oidcAccount.Audience)

	// Test with valid UsernamePassword account
	name = internal.GetRandomName()
	username := internal.GetRandomName()
	password := core.NewSensitiveValue(internal.GetRandomName())

	accountResource = NewPlatformHubAccountResource(name, AccountTypePlatformHubUsernamePasswordAccount)
	accountResource.Username = username
	accountResource.Password = password
	accountResource.Description = "Test UsernamePassword description"
	accountResource.ID = "test-userpass-id"

	account, err = accountResource.ToPlatformHubAccount()
	require.NotNil(t, account)
	require.NoError(t, err)
	require.Equal(t, AccountTypePlatformHubUsernamePasswordAccount, account.GetAccountType())
	require.Equal(t, name, account.GetName())
	require.Equal(t, "Test UsernamePassword description", account.GetDescription())
	require.Equal(t, "test-userpass-id", account.GetID())

	usernamePasswordAccount, ok := account.(*PlatformHubUsernamePasswordAccount)
	require.True(t, ok)
	require.Equal(t, username, usernamePasswordAccount.Username)
	require.Equal(t, password, usernamePasswordAccount.Password)
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

	// Test with valid GCP account
	jsonKey := core.NewSensitiveValue(internal.GetRandomName())
	name = internal.GetRandomName()

	gcpAccount, err := NewPlatformHubGcpAccount(name, jsonKey)
	require.NoError(t, err)
	require.NotNil(t, gcpAccount)

	gcpAccount.Description = "Test GCP description"
	gcpAccount.ID = "test-gcp-id"

	accountResource, err = ToPlatformHubAccountResource(gcpAccount)
	require.NotNil(t, accountResource)
	require.NoError(t, err)
	require.Equal(t, AccountTypePlatformHubGcpAccount, accountResource.GetAccountType())
	require.Equal(t, name, accountResource.GetName())
	require.Equal(t, "Test GCP description", accountResource.GetDescription())
	require.Equal(t, "test-gcp-id", accountResource.GetID())
	require.Equal(t, jsonKey, accountResource.JsonKey)

	// Test with valid Generic OIDC account
	name = internal.GetRandomName()
	executionSubjectKeys := []string{"space", "project", "environment"}
	audience := "api://default"

	oidcAccount, err := NewPlatformHubGenericOidcAccount(name)
	require.NoError(t, err)
	require.NotNil(t, oidcAccount)

	oidcAccount.Description = "Test Generic OIDC description"
	oidcAccount.ID = "test-oidc-id"
	oidcAccount.ExecutionSubjectKeys = executionSubjectKeys
	oidcAccount.Audience = audience

	accountResource, err = ToPlatformHubAccountResource(oidcAccount)
	require.NotNil(t, accountResource)
	require.NoError(t, err)
	require.Equal(t, AccountTypePlatformHubGenericOidcAccount, accountResource.GetAccountType())
	require.Equal(t, name, accountResource.GetName())
	require.Equal(t, "Test Generic OIDC description", accountResource.GetDescription())
	require.Equal(t, "test-oidc-id", accountResource.GetID())
	require.Equal(t, executionSubjectKeys, accountResource.ExecutionSubjectKeys)
	require.Equal(t, audience, accountResource.Audience)

	// Test with valid UsernamePassword account
	name = internal.GetRandomName()
	username := internal.GetRandomName()
	password := core.NewSensitiveValue(internal.GetRandomName())

	usernamePasswordAccount, err := NewPlatformHubUsernamePasswordAccount(name, username, password)
	require.NoError(t, err)
	require.NotNil(t, usernamePasswordAccount)

	usernamePasswordAccount.Description = "Test UsernamePassword description"
	usernamePasswordAccount.ID = "test-userpass-id"

	accountResource, err = ToPlatformHubAccountResource(usernamePasswordAccount)
	require.NotNil(t, accountResource)
	require.NoError(t, err)
	require.Equal(t, AccountTypePlatformHubUsernamePasswordAccount, accountResource.GetAccountType())
	require.Equal(t, name, accountResource.GetName())
	require.Equal(t, "Test UsernamePassword description", accountResource.GetDescription())
	require.Equal(t, "test-userpass-id", accountResource.GetID())
	require.Equal(t, username, accountResource.Username)
	require.Equal(t, password, accountResource.Password)
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
