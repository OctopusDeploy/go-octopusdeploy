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
	invalidDeploymentSubjectKeys := []string{"space", "target"}

	testCases := []struct {
		TestName              string
		IsError               bool
		Name                  string
		Audience              string
		DeploymentSubjectKeys []string
	}{
		{"Valid", false, name, audience, deploymentSubjectKeys},
		{"EmptyName", true, "", audience, deploymentSubjectKeys},
		{"NilSubjectKeys", false, name, "", nil},
		{"InvalidDeploymentSubjectKeys", true, name, "", invalidDeploymentSubjectKeys},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			genericOIDCAccount := &GenericOIDCAccount{
				Audience:              tc.Audience,
				DeploymentSubjectKeys: tc.DeploymentSubjectKeys,
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
