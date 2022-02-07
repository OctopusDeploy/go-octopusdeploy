package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAzureServicePrincipalAccount(t *testing.T) {
	applicationID := uuid.New()
	applicationPassword := NewSensitiveValue(getRandomName())
	authenticationEndpoint := "https://login.microsoftonline.com/"
	azureEnvironment := "AzureCloud"
	invalidURI := "***"
	name := getRandomName()
	resourceManagerEndpoint := "https://management.azure.com/"
	spaceID := "space-id"
	subscriptionID := uuid.New()
	tenantedDeploymentMode := TenantedDeploymentMode("Untenanted")
	tenantID := uuid.New()

	testCases := []struct {
		TestName                string
		IsError                 bool
		ApplicationID           *uuid.UUID
		ApplicationPassword     *SensitiveValue
		AuthenticationEndpoint  string
		AzureEnvironment        string
		Name                    string
		ResourceManagerEndpoint string
		SpaceID                 string
		SubscriptionID          *uuid.UUID
		TenantedDeploymentMode  TenantedDeploymentMode
		TenantID                *uuid.UUID
	}{
		{"Valid", false, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"EmptyName", true, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, services.emptyString, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"WhitespaceName", true, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, services.whitespaceString, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"EmptySpaceID", false, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, services.emptyString, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"WhitespaceSpaceID", true, &applicationID, applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, services.whitespaceString, &subscriptionID, tenantedDeploymentMode, &tenantID},
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
			azureServicePrincipalAccount.AccountType = accounts.AccountType("AzureServicePrincipal")
			azureServicePrincipalAccount.Name = tc.Name
			azureServicePrincipalAccount.SpaceID = tc.SpaceID
			azureServicePrincipalAccount.TenantedDeploymentMode = tc.TenantedDeploymentMode
			if tc.IsError {
				require.Error(t, azureServicePrincipalAccount.Validate())
			} else {
				require.NoError(t, azureServicePrincipalAccount.Validate())

				require.Equal(t, accounts.AccountType("AzureServicePrincipal"), azureServicePrincipalAccount.GetAccountType())
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
	applicationPassword := NewSensitiveValue(getRandomName())
	name := getRandomName()
	subscriptionID := uuid.New()
	tenantID := uuid.New()

	account, err := NewAzureServicePrincipalAccount(name, subscriptionID, tenantID, applicationID, applicationPassword)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())
}
