package accounts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAzureOIDCAccount(t *testing.T) {
	applicationID := uuid.New()
	authenticationEndpoint := "https://login.microsoftonline.com/"
	azureEnvironment := "AzureCloud"
	invalidURI := "***"
	name := internal.GetRandomName()
	resourceManagerEndpoint := "https://management.azure.com/"
	spaceID := "space-id"
	subscriptionID := uuid.New()
	tenantedDeploymentMode := core.TenantedDeploymentMode("Untenanted")
	tenantID := uuid.New()
	audience := "api://AzureADTokenExchange"
	deploymentSubjectKeys := []string{"space", "project", "tenant", "environment"}
	healthCheckSubjectKeys := []string{"space", "target"}
	accountTestSubjectKeys := []string{"space", "account"}
	invalidAccountTestSubjectKeys := []string{"space", "account", "project"}

	testCases := []struct {
		TestName                string
		IsError                 bool
		ApplicationID           *uuid.UUID
		AuthenticationEndpoint  string
		AzureEnvironment        string
		Name                    string
		ResourceManagerEndpoint string
		SpaceID                 string
		SubscriptionID          *uuid.UUID
		TenantedDeploymentMode  core.TenantedDeploymentMode
		TenantID                *uuid.UUID
		Audience                string
		DeploymentSubjectKeys   []string
		HealthCheckSubjectKeys  []string
		AccountTestSubjectKeys  []string
	}{
		{"Valid", false, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"EmptyName", true, &applicationID, authenticationEndpoint, azureEnvironment, "", resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"WhitespaceName", true, &applicationID, authenticationEndpoint, azureEnvironment, " ", resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"EmptySpaceID", false, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, "", &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"WhitespaceSpaceID", false, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, " ", &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"NilApplicationID", true, nil, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"NilSubscriptionID", true, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, nil, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"NilTenantID", true, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, nil, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"InvalidAuthenticationEndpoint", true, &applicationID, invalidURI, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"InvalidResourceManagerEndpoint", true, &applicationID, authenticationEndpoint, azureEnvironment, name, invalidURI, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys},
		{"NilSubjectKeys", false, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, "", nil, nil, nil},
		{"InvalidSubjectKeys", true, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, "", deploymentSubjectKeys, healthCheckSubjectKeys, invalidAccountTestSubjectKeys},
	}
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			azureOIDCAccount := &AzureOIDCAccount{
				ApplicationID:           tc.ApplicationID,
				AuthenticationEndpoint:  tc.AuthenticationEndpoint,
				AzureEnvironment:        tc.AzureEnvironment,
				ResourceManagerEndpoint: tc.ResourceManagerEndpoint,
				SubscriptionID:          tc.SubscriptionID,
				TenantID:                tc.TenantID,
				Audience:                tc.Audience,
				DeploymentSubjectKeys:   tc.DeploymentSubjectKeys,
				HealthCheckSubjectKeys:  tc.HealthCheckSubjectKeys,
				AccountTestSubjectKeys:  tc.AccountTestSubjectKeys,
			}
			azureOIDCAccount.AccountType = AccountType("AzureOIDC")
			azureOIDCAccount.Name = tc.Name
			azureOIDCAccount.SpaceID = tc.SpaceID
			azureOIDCAccount.TenantedDeploymentMode = tc.TenantedDeploymentMode
			if tc.IsError {
				require.Error(t, azureOIDCAccount.Validate())
			} else {
				require.NoError(t, azureOIDCAccount.Validate())

				require.Equal(t, AccountType("AzureOIDC"), azureOIDCAccount.GetAccountType())
				require.Equal(t, tc.Name, azureOIDCAccount.GetName())
			}
			azureOIDCAccount.SetName(tc.Name)
			if tc.IsError {
				require.Error(t, azureOIDCAccount.Validate())
			} else {
				require.NoError(t, azureOIDCAccount.Validate())
				require.Equal(t, tc.Name, azureOIDCAccount.GetName())
			}
		})
	}
}

func TestAzureOIDCAccountNew(t *testing.T) {
	applicationID := uuid.New()
	name := internal.GetRandomName()
	subscriptionID := uuid.New()
	tenantID := uuid.New()

	account, err := NewAzureOIDCAccount(name, subscriptionID, tenantID, applicationID)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())
}
