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
	invalidDeploymentSubjectKeys := []string{"space", "target"}
	invalidHealthCheckSubjectKeys := []string{"space", "project"}
	invalidAccountTestSubjectKeys := []string{"space", "project"}
	customClaims := map[string]string{
		"claim1": "value1",
		"claim2": "value2",
	}

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
		CustomClaims            map[string]string
	}{
		{"Valid", false, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"ValidWithCustomClaims", false, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, customClaims},
		{"EmptyName", true, &applicationID, authenticationEndpoint, azureEnvironment, "", resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"WhitespaceName", true, &applicationID, authenticationEndpoint, azureEnvironment, " ", resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"EmptySpaceID", false, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, "", &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"WhitespaceSpaceID", false, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, " ", &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"NilApplicationID", true, nil, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"NilSubscriptionID", true, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, nil, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"NilTenantID", true, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, nil, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"InvalidAuthenticationEndpoint", true, &applicationID, invalidURI, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"InvalidResourceManagerEndpoint", true, &applicationID, authenticationEndpoint, azureEnvironment, name, invalidURI, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, audience, deploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"NilSubjectKeys", false, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, "", nil, nil, nil, nil},
		{"InvalidDeploymentSubjectKeys", true, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, "", invalidDeploymentSubjectKeys, healthCheckSubjectKeys, accountTestSubjectKeys, nil},
		{"InvalidHealthCheckSubjectKeys", true, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, "", deploymentSubjectKeys, invalidHealthCheckSubjectKeys, invalidAccountTestSubjectKeys, nil},
		{"InvalidAccountTestSubjectKeys", true, &applicationID, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID, "", deploymentSubjectKeys, healthCheckSubjectKeys, invalidAccountTestSubjectKeys, nil},
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
				CustomClaims:            tc.CustomClaims,
			}
			azureOIDCAccount.AccountType = AccountTypeAzureOIDC
			azureOIDCAccount.Name = tc.Name
			azureOIDCAccount.SpaceID = tc.SpaceID
			azureOIDCAccount.TenantedDeploymentMode = tc.TenantedDeploymentMode
			if tc.IsError {
				require.Error(t, azureOIDCAccount.Validate())
			} else {
				require.NoError(t, azureOIDCAccount.Validate())

				require.Equal(t, AccountTypeAzureOIDC, azureOIDCAccount.GetAccountType())
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
