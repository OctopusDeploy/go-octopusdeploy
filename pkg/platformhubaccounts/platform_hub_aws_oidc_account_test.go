package platformhubaccounts

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestPlatformHubAwsOIDCAccount(t *testing.T) {
	name := internal.GetRandomName()
	roleArn := "arn:aws:iam::123456789012:role/MyRole"
	sessionDuration := "3600"
	deploymentSubjectKeys := []string{"space", "environment", "project"}
	healthCheckSubjectKeys := []string{"space", "target"}
	accountTestSubjectKeys := []string{"space", "account"}
	invalidDeploymentSubjectKeys := []string{"space", "invalid"}
	invalidHealthCheckSubjectKeys := []string{"space", "project"}
	invalidAccountTestSubjectKeys := []string{"space", "environment"}

	testCases := []struct {
		TestName               string
		IsError                bool
		Name                   string
		RoleArn                string
		SessionDuration        string
		DeploymentSubjectKeys  []string
		HealthCheckSubjectKeys []string
		AccountTestSubjectKeys []string
	}{
		{"Valid", false, name, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"EmptyName", true, "", roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"WhitespaceName", true, " ", roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"EmptyRoleArn", true, name, "", sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"NilSubjectKeys", false, name, roleArn, sessionDuration, nil, nil, nil},
		{"EmptySessionDuration", false, name, roleArn, "", deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"InvalidDeploymentSubjectKeys", true, name, roleArn, sessionDuration, invalidDeploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"InvalidHealthCheckSubjectKeys", true, name, roleArn, sessionDuration, deploymentSubjectKeys, invalidHealthCheckSubjectKeys, accountTestSubjectKeys},
		{"InvalidAccountTestSubjectKeys", true, name, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, invalidAccountTestSubjectKeys},
	}
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			account := &PlatformHubAwsOIDCAccount{
				RoleArn:                tc.RoleArn,
				SessionDuration:        tc.SessionDuration,
				DeploymentSubjectKeys:  tc.DeploymentSubjectKeys,
				HealthCheckSubjectKeys: tc.HealthCheckSubjectKeys,
				AccountTestSubjectKeys: tc.AccountTestSubjectKeys,
			}
			account.AccountType = AccountTypePlatformHubAwsOIDCAccount
			account.Name = tc.Name

			if tc.IsError {
				require.Error(t, account.Validate())
			} else {
				require.NoError(t, account.Validate())
				require.Equal(t, AccountTypePlatformHubAwsOIDCAccount, account.GetAccountType())
				require.Equal(t, tc.Name, account.GetName())
			}

			account.SetName(tc.Name)
			if tc.IsError {
				require.Error(t, account.Validate())
			} else {
				require.NoError(t, account.Validate())
				require.Equal(t, tc.Name, account.GetName())
			}
		})
	}
}

func TestPlatformHubAwsOIDCAccountNew(t *testing.T) {
	name := internal.GetRandomName()
	roleArn := "arn:aws:iam::123456789012:role/MyRole"
	accountType := AccountTypePlatformHubAwsOIDCAccount
	description := ""

	account, err := NewPlatformHubAwsOIDCAccount(name, roleArn)

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

	// IPlatformHubAccount
	require.Equal(t, accountType, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// PlatformHubAwsOIDCAccount
	require.Equal(t, roleArn, account.RoleArn)
}

func TestPlatformHubAwsOIDCAccountMarshalJSON(t *testing.T) {
	name := internal.GetRandomName()
	roleArn := "arn:aws:iam::123456789012:role/MyRole"
	sessionDuration := "3600"
	deploymentSubjectKeys := []string{"space", "environment"}
	healthCheckSubjectKeys := []string{"space", "target"}
	accountTestSubjectKeys := []string{"space", "account"}

	deploymentSubjectKeysJSON, err := json.Marshal(deploymentSubjectKeys)
	require.NoError(t, err)
	require.NotNil(t, deploymentSubjectKeysJSON)

	healthCheckSubjectKeysJSON, err := json.Marshal(healthCheckSubjectKeys)
	require.NoError(t, err)
	require.NotNil(t, healthCheckSubjectKeysJSON)

	accountTestSubjectKeysJSON, err := json.Marshal(accountTestSubjectKeys)
	require.NoError(t, err)
	require.NotNil(t, accountTestSubjectKeysJSON)

	expectedJson := fmt.Sprintf(`{
		"AccountType": "AmazonWebServicesOidcAccount",
		"Name": "%s",
		"RoleArn": "%s",
		"SessionDuration": "%s",
		"DeploymentSubjectKeys": %s,
		"HealthCheckSubjectKeys": %s,
		"AccountTestSubjectKeys": %s
	}`, name, roleArn, sessionDuration, deploymentSubjectKeysJSON, healthCheckSubjectKeysJSON, accountTestSubjectKeysJSON)

	account, err := NewPlatformHubAwsOIDCAccount(name, roleArn)
	require.NoError(t, err)
	require.NotNil(t, account)

	account.SessionDuration = sessionDuration
	account.DeploymentSubjectKeys = deploymentSubjectKeys
	account.HealthCheckSubjectKeys = healthCheckSubjectKeys
	account.AccountTestSubjectKeys = accountTestSubjectKeys

	accountAsJSON, err := json.Marshal(account)
	require.NoError(t, err)
	require.NotNil(t, accountAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(accountAsJSON))
}

func TestPlatformHubAwsOIDCAccountNewWithConfigs(t *testing.T) {
	name := internal.GetRandomName()
	roleArn := "arn:aws:iam::123456789012:role/MyRole"
	sessionDuration := "3600"
	deploymentSubjectKeys := []string{"space", "environment", "project"}
	healthCheckSubjectKeys := []string{"space", "target"}
	accountTestSubjectKeys := []string{"space", "account"}
	accountType := AccountTypePlatformHubAwsOIDCAccount
	id := internal.GetRandomName()
	modifiedBy := internal.GetRandomName()
	modifiedOn := time.Now()
	description := "Description for " + name + " (OK to Delete)"

	account, err := NewPlatformHubAwsOIDCAccount(name, roleArn)
	require.NoError(t, err)
	require.NotNil(t, account)
	require.NoError(t, account.Validate())

	account.Description = description
	account.SessionDuration = sessionDuration
	account.DeploymentSubjectKeys = deploymentSubjectKeys
	account.HealthCheckSubjectKeys = healthCheckSubjectKeys
	account.AccountTestSubjectKeys = accountTestSubjectKeys
	account.ID = id
	account.ModifiedBy = modifiedBy
	account.ModifiedOn = &modifiedOn

	// resource
	require.Equal(t, id, account.ID)
	require.Equal(t, modifiedBy, account.ModifiedBy)
	require.Equal(t, &modifiedOn, account.ModifiedOn)
	require.NotNil(t, account.Links)

	// IResource
	require.Equal(t, id, account.GetID())
	require.Equal(t, modifiedBy, account.GetModifiedBy())
	require.Equal(t, &modifiedOn, account.GetModifiedOn())
	require.NotNil(t, account.GetLinks())

	// IPlatformHubAccount
	require.Equal(t, accountType, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// PlatformHubAwsOIDCAccount
	require.Equal(t, roleArn, account.RoleArn)
	require.Equal(t, sessionDuration, account.SessionDuration)
	require.Equal(t, deploymentSubjectKeys, account.DeploymentSubjectKeys)
	require.Equal(t, healthCheckSubjectKeys, account.HealthCheckSubjectKeys)
	require.Equal(t, accountTestSubjectKeys, account.AccountTestSubjectKeys)
}
