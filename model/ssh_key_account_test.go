package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSSHKeyAccountTypes(t *testing.T) {
	account := NewSSHKeyAccount("ssh-key-account-name", "ssh-key-account-username", NewSensitiveValue("new-value"))
	require.NoError(t, account.Validate())

	account.AccountType = "None"
	assert.Error(t, account.Validate())

	account.AccountType = "none"
	assert.Error(t, account.Validate())

	account.AccountType = "UsernamePassword"
	assert.Error(t, account.Validate())

	account.AccountType = "AzureSubscription"
	assert.Error(t, account.Validate())

	account.AccountType = "AzureServicePrincipal"
	assert.Error(t, account.Validate())

	account.AccountType = "AmazonWebServicesAccount"
	assert.Error(t, account.Validate())

	account.AccountType = "AmazonWebServicesRoleAccount"
	assert.Error(t, account.Validate())

	account.AccountType = "Token"
	assert.Error(t, account.Validate())

	account.AccountType = "SshKeyPair"
	assert.NoError(t, account.Validate())

	account.AccountType = "sshKeyPair"
	assert.Error(t, account.Validate())
}

func TestSSHKeyAccountMarshalJSON(t *testing.T) {
	account := NewSSHKeyAccount("account-name", "account-username", NewSensitiveValue("new-value"))

	environmentIDs := []string{
		"environment-id-1",
		"environment-id-2",
	}
	lastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")
	links := map[string]string{
		"Self": "/api/foo/bar/quux",
		"test": "/api/xyzzy",
	}

	privateKeyPassphrase := NewSensitiveValue(emptyString)

	tenantIDs := []string{
		"tenant-id-1",
		"tenant-id-2",
	}
	tenantTags := []string{
		"tenant-tag-1",
		"tenant-tag-2",
	}

	account.Description = "account-description"
	account.EnvironmentIDs = environmentIDs
	account.ID = "account-id"
	account.LastModifiedBy = "john.smith@example.com"
	account.LastModifiedOn = &lastModifiedOn
	account.Links = links
	account.PrivateKeyPassphrase = &privateKeyPassphrase
	account.SpaceID = "space-id"
	account.TenantIDs = tenantIDs
	account.TenantTags = tenantTags

	require.NoError(t, account.Validate())

	jsonEncoding, err := json.Marshal(account)
	require.NoError(t, err)
	require.NotNil(t, jsonEncoding)

	actual := string(jsonEncoding)

	jsonassert.New(t).Assertf(actual, sshKeyAccountAsJSON)
}

func TestSSHKeyAccountUnmarshalJSON(t *testing.T) {
	var account SSHKeyAccount
	err := json.Unmarshal([]byte(sshKeyAccountAsJSON), &account)

	require.NoError(t, err)
	require.NotNil(t, account)

	environmentIDs := []string{
		"environment-id-1",
		"environment-id-2",
	}
	lastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")
	links := map[string]string{
		"Self": "/api/foo/bar/quux",
		"test": "/api/xyzzy",
	}

	privateKeyFile := NewSensitiveValue("new-value")
	privateKeyPassphrase := NewSensitiveValue(emptyString)

	tenantIDs := []string{
		"tenant-id-1",
		"tenant-id-2",
	}
	tenantTags := []string{
		"tenant-tag-1",
		"tenant-tag-2",
	}

	assert.Equal(t, "account-id", account.ID)
	assert.Equal(t, environmentIDs, account.EnvironmentIDs)
	assert.Equal(t, lastModifiedOn, *account.LastModifiedOn)
	assert.Equal(t, "john.smith@example.com", account.LastModifiedBy)
	assert.Equal(t, links, account.Links)
	assert.Equal(t, privateKeyFile, *account.PrivateKeyFile)
	assert.Equal(t, privateKeyPassphrase, *account.PrivateKeyPassphrase)
	assert.Equal(t, "space-id", account.SpaceID)
	assert.Equal(t, "Untenanted", account.TenantedDeploymentMode)
	assert.Equal(t, tenantIDs, account.TenantIDs)
	assert.Equal(t, tenantTags, account.TenantTags)
	assert.Equal(t, "account-username", account.Username)
}

const sshKeyAccountAsJSON string = `{
	"AccountType": "SshKeyPair",
	"Description": "account-description",
	"EnvironmentIds": [
		"environment-id-1",
		"environment-id-2"
	],
	"Id": "account-id",
	"LastModifiedOn": "2020-10-02T00:44:11.284Z",
	"LastModifiedBy": "john.smith@example.com",
	"Links": {
		"Self": "/api/foo/bar/quux",
		"test": "/api/xyzzy"
	},
	"Name": "account-name",
	"PrivateKeyFile": {
		"HasValue": true,
		"NewValue": "new-value"
	},
	"PrivateKeyPassphrase": {
		"HasValue": false,
		"NewValue": null
	},
	"SpaceId": "space-id",
	"TenantedDeploymentParticipation": "Untenanted",
	"TenantIds": [
		"tenant-id-1",
		"tenant-id-2"
	],
	"TenantTags": [
		"tenant-tag-1",
		"tenant-tag-2"
	],
	"Username": "account-username"
}`
