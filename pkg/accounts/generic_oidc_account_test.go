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
	customClaims := map[string]string{
		"claim1": "value1",
		"claim2": "value2",
	}

	testCases := []struct {
		TestName              string
		IsError               bool
		Name                  string
		Audience              string
		DeploymentSubjectKeys []string
		CustomClaims          map[string]string
	}{
		{"Valid", false, name, audience, deploymentSubjectKeys, nil},
		{"ValidWithCustomClaims", false, name, audience, deploymentSubjectKeys, customClaims},
		{"EmptyName", true, "", audience, deploymentSubjectKeys, nil},
		{"NilSubjectKeys", false, name, "", nil, nil},
		{"InvalidDeploymentSubjectKeys", true, name, "", invalidDeploymentSubjectKeys, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			genericOIDCAccount := &GenericOIDCAccount{
				Audience:              tc.Audience,
				DeploymentSubjectKeys: tc.DeploymentSubjectKeys,
				CustomClaims:          tc.CustomClaims,
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
