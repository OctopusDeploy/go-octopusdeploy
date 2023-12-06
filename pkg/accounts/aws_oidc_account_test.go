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
	audience := "some default audience"
	roleArn := "role arn::with:two:colons:for:some:reason"
	sessionDuration := "3600"
	deploymentSubjectKeys := []string{"space", "project", "tenant", "environment"}
	healthCheckSubjectKeys := []string{"space", "target"}
	accountTestSubjectKeys := []string{"space", "account"}
	invalidDeploymentSubjectKeys := []string{"space", "target"}
	invalidHealthCheckSubjectKeys := []string{"space", "project"}
	invalidAccountTestSubjectKeys := []string{"space", "project"}

	testCases := []struct {
		TestName               string
		IsError                bool
		Name                   string
		SpaceID                string
		TenantedDeploymentMode core.TenantedDeploymentMode
		Audience               string
		RoleArn                string
		SessionDuration        string
		DeploymentSubjectKeys  []string
		HealthCheckSubjectKeys []string
		AccountTestSubjectKeys []string
	}{
		{"Valid", false, name, spaceID, tenantedDeploymentMode, audience, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"EmptyName", true, "", spaceID, tenantedDeploymentMode, audience, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"WhitespaceName", true, " ", spaceID, tenantedDeploymentMode, audience, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"EmptySpaceID", false, name, "", tenantedDeploymentMode, audience, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"WhitespaceSpaceID", false, name, " ", tenantedDeploymentMode, audience, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"NilSubjectKeys", false, name, spaceID, tenantedDeploymentMode, audience, roleArn, sessionDuration, nil, nil, nil},
		{"InvalidDeploymentSubjectKeys", true, name, spaceID, tenantedDeploymentMode, audience, roleArn, sessionDuration, invalidDeploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"InvalidHealthCheckSubjectKeys", true, name, spaceID, tenantedDeploymentMode, audience, roleArn, sessionDuration, deploymentSubjectKeys, invalidHealthCheckSubjectKeys, invalidAccountTestSubjectKeys},
		{"InvalidAccountTestSubjectKeys", true, name, spaceID, tenantedDeploymentMode, audience, roleArn, sessionDuration, deploymentSubjectKeys, healthCheckSubjectKeys, invalidAccountTestSubjectKeys},
	}
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			awsOIDCAccount := &AwsOIDCAccount{
				RoleArn:                tc.RoleArn,
				SessionDuration:        tc.SessionDuration,
				Audience:               tc.Audience,
				DeploymentSubjectKeys:  tc.DeploymentSubjectKeys,
				HealthCheckSubjectKeys: tc.HealthCheckSubjectKeys,
				AccountTestSubjectKeys: tc.AccountTestSubjectKeys,
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
