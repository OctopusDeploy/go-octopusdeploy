package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenericOIDCAccount(t *testing.T) {
	name := internal.GetRandomName()
	audience := "api://default"
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
		Audience               string
		DeploymentSubjectKeys  []string
		HealthCheckSubjectKeys []string
		AccountTestSubjectKeys []string
	}{
		{"Valid", false, name, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"EmptyName", true, "", audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"NilSubjectKeys", false, name, "", nil, nil, nil},
		{"InvalidDeploymentSubjectKeys", true, name, "", invalidDeploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"InvalidHealthCheckSubjectKeys", true, name, "", deploymentSubjectKeys, invalidHealthCheckSubjectKeys, invalidAccountTestSubjectKeys},
		{"InvalidAccountTestSubjectKeys", true, name, "", deploymentSubjectKeys, healthCheckSubjectKeys, invalidAccountTestSubjectKeys},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			genericOIDCAccount := &GenericOIDCAccount{
				Audience:               tc.Audience,
				DeploymentSubjectKeys:  tc.DeploymentSubjectKeys,
				HealthCheckSubjectKeys: tc.HealthCheckSubjectKeys,
				AccountTestSubjectKeys: tc.AccountTestSubjectKeys,
			}
			genericOIDCAccount.AccountType = AccountTypeGenericOIDCAccount
			genericOIDCAccount.Name = tc.Name
			if tc.IsError {
				require.Error(t, genericOIDCAccount.Validate())
			} else {
				require.NoError(t, genericOIDCAccount.Validate())

				require.Equal(t, AccountTypeGenericOIDCAccount, genericOIDCAccount.GetAccountType())
				require.Equal(t, tc.Name, genericOIDCAccount.GetName())
			}
			genericOIDCAccount.SetName(tc.Name)
			if tc.IsError {
				require.Error(t, genericOIDCAccount.Validate())
			} else {
				require.NoError(t, genericOIDCAccount.Validate())
				require.Equal(t, tc.Name, genericOIDCAccount.GetName())
			}
		})
	}
}

func TestGenericOIDCAccountNew(t *testing.T) {
	name := internal.GetRandomName()

	account, err := NewGenericOIDCAccount(name)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())
}
