package accounts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAzureServicePrincipalAccount(t *testing.T) {
	applicationID := uuid.New()
	applicationPassword := core.NewSensitiveValue(internal.GetRandomName())
	authenticationEndpoint := "https://login.microsoftonline.com/"
	azureEnvironment := "AzureCloud"
	invalidURI := "***"
	name := internal.GetRandomName()
	resourceManagerEndpoint := "https://management.azure.com/"
	spaceID := "space-id"
	subscriptionID := uuid.New()
	tenantedDeploymentMode := core.TenantedDeploymentMode("Untenanted")
	tenantID := uuid.New()

	testCases := []struct {
		TestName                string
		IsError                 bool
		ApplicationID           *uuid.UUID
		ApplicationPassword     *core.SensitiveValue
		AuthenticationEndpoint  string
		AzureEnvironment        string
		Name                    string
		ResourceManagerEndpoint string
		SpaceID                 string
		SubscriptionID          *uuid.UUID
		TenantedDeploymentMode  core.TenantedDeploymentMode
		TenantID                *uuid.UUID
	}{
		{"Valid", false, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"EmptyName", true, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, "", resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"WhitespaceName", true, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, " ", resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"EmptySpaceID", false, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, "", &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"WhitespaceSpaceID", false, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, " ", &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"NilApplicationPassword", true, &applicationID, nil, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"NilApplicationID", true, nil, applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"NilSubscriptionID", true, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, nil, tenantedDeploymentMode, &tenantID},
		{"NilTenantID", true, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, nil},
		{"InvalidAuthenticationEndpoint", true, &applicationID, applicationPassword, invalidURI, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"InvalidResourceManagerEndpoint", true, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, name, invalidURI, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
	}
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			azureServicePrincipalAccount := &AzureServicePrincipalAccount{
				ApplicationID:           tc.ApplicationID,
				ApplicationPassword:     tc.ApplicationPassword,
				AuthenticationEndpoint:  tc.AuthenticationEndpoint,
				AzureEnvironment:        tc.AzureEnvironment,
				ResourceManagerEndpoint: tc.ResourceManagerEndpoint,
				SubscriptionID:          tc.SubscriptionID,
				TenantID:                tc.TenantID,
			}
			azureServicePrincipalAccount.AccountType = AccountType("AzureServicePrincipal")
			azureServicePrincipalAccount.Name = tc.Name
			azureServicePrincipalAccount.SpaceID = tc.SpaceID
			azureServicePrincipalAccount.TenantedDeploymentMode = tc.TenantedDeploymentMode
			if tc.IsError {
				require.Error(t, azureServicePrincipalAccount.Validate())
			} else {
				require.NoError(t, azureServicePrincipalAccount.Validate())

				require.Equal(t, AccountType("AzureServicePrincipal"), azureServicePrincipalAccount.GetAccountType())
				require.Equal(t, tc.Name, azureServicePrincipalAccount.GetName())
			}
			azureServicePrincipalAccount.SetName(tc.Name)
			if tc.IsError {
				require.Error(t, azureServicePrincipalAccount.Validate())
			} else {
				require.NoError(t, azureServicePrincipalAccount.Validate())
				require.Equal(t, tc.Name, azureServicePrincipalAccount.GetName())
			}
		})
	}
}

func TestAzureServicePrincipalAccountNew(t *testing.T) {
	applicationID := uuid.New()
	applicationPassword := core.NewSensitiveValue(internal.GetRandomName())
	name := internal.GetRandomName()
	subscriptionID := uuid.New()
	tenantID := uuid.New()

	account, err := NewAzureServicePrincipalAccount(name, subscriptionID, tenantID, applicationID, applicationPassword)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())
}
