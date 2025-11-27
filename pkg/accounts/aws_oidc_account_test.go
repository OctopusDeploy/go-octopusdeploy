package accounts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/stretchr/testify/require"
)

func TestAwsOIDCAccount(t *testing.T) {
	name := internal.GetRandomName()
	spaceID := "space-id"
	tenantedDeploymentMode := core.TenantedDeploymentMode("Untenanted")
	roleArn := "role arn::with:two:colons:for:some:reason"
	sessionDuration := "3600"
	deploymentSubjectKeys := []string{"space", "project", "tenant", "environment"}
	healthCheckSubjectKeys := []string{"space", "target"}
	accountTestSubjectKeys := []string{"space", "account"}
	invalidDeploymentSubjectKeys := []string{"space", "target"}
	invalidHealthCheckSubjectKeys := []string{"space", "project"}
	invalidAccountTestSubjectKeys := []string{"space", "project"}
	customClaims := map[string]string{
		"claim1": "value1",
		"claim2": "value2",
	}

	testCases := []struct {
		TestName               string
		IsError                bool
		Name                   string
		SpaceID                string
		TenantedDeploymentMode core.TenantedDeploymentMode
		RoleArn                string
		SessionDuration        string
		DeploymentSubjectKeys  []string
		HealthCheckSubjectKeys []string
		AccountTestSubjectKeys []string
		CustomClaims           map[string]string
	}{
		{"Valid", false, name, spaceID, tenantedDeploymentMode, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"ValidWithCustomClaims", false, name, spaceID, tenantedDeploymentMode, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, customClaims},
		{"EmptyName", true, "", spaceID, tenantedDeploymentMode, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"WhitespaceName", true, " ", spaceID, tenantedDeploymentMode, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"EmptySpaceID", false, name, "", tenantedDeploymentMode, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"WhitespaceSpaceID", false, name, " ", tenantedDeploymentMode, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"NilSubjectKeys", false, name, spaceID, tenantedDeploymentMode, roleArn, sessionDuration, nil, nil, nil, nil},
		{"InvalidDeploymentSubjectKeys", true, name, spaceID, tenantedDeploymentMode, roleArn, sessionDuration, invalidDeploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"InvalidHealthCheckSubjectKeys", true, name, spaceID, tenantedDeploymentMode, roleArn, sessionDuration, deploymentSubjectKeys, invalidHealthCheckSubjectKeys, invalidAccountTestSubjectKeys, nil},
		{"InvalidAccountTestSubjectKeys", true, name, spaceID, tenantedDeploymentMode, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, invalidAccountTestSubjectKeys, nil},
	}
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			awsOIDCAccount := &AwsOIDCAccount{
				RoleArn:                tc.RoleArn,
				SessionDuration:        tc.SessionDuration,
				DeploymentSubjectKeys:  tc.DeploymentSubjectKeys,
				HealthCheckSubjectKeys: tc.HealthCheckSubjectKeys,
				AccountTestSubjectKeys: tc.AccountTestSubjectKeys,
				CustomClaims:           tc.CustomClaims,
			}
			awsOIDCAccount.AccountType = AccountTypeAwsOIDC
			awsOIDCAccount.Name = tc.Name
			awsOIDCAccount.SpaceID = tc.SpaceID
			awsOIDCAccount.TenantedDeploymentMode = tc.TenantedDeploymentMode
			if tc.IsError {
				require.Error(t, awsOIDCAccount.Validate())
			} else {
				require.NoError(t, awsOIDCAccount.Validate())

				require.Equal(t, AccountTypeAwsOIDC, awsOIDCAccount.GetAccountType())
				require.Equal(t, tc.Name, awsOIDCAccount.GetName())
			}
			awsOIDCAccount.SetName(tc.Name)
			if tc.IsError {
				require.Error(t, awsOIDCAccount.Validate())
			} else {
				require.NoError(t, awsOIDCAccount.Validate())
				require.Equal(t, tc.Name, awsOIDCAccount.GetName())
			}
		})
	}
}

func TestAwsOIDCAccountNew(t *testing.T) {
	name := internal.GetRandomName()
	roleArn := internal.GetRandomName()

	account, err := NewAwsOIDCAccount(name, roleArn)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())
}
