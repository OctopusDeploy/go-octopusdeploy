package model

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAzureServicePrincipalAccount(t *testing.T) {
	accountType := "AzureServicePrincipal"
	applicationID := uuid.New()
	applicationPassword := NewSensitiveValue(getRandomName())
	authenticationEndpoint := "https://login.microsoftonline.com/"
	azureEnvironment := "AzureCloud"
	invalidAccountType := "***"
	invalidURI := "***"
	invalidTenantedDeploymentMode := "***"
	name := getRandomName()
	resourceManagerEndpoint := "https://management.azure.com/"
	spaceID := "space-id"
	subscriptionID := uuid.New()
	tenantedDeploymentMode := "Untenanted"
	tenantID := uuid.New()

	testCases := []struct {
		TestName                string
		IsError                 bool
		AccountType             string
		ApplicationID           *uuid.UUID
		ApplicationPassword     *SensitiveValue
		AuthenticationEndpoint  string
		AzureEnvironment        string
		Name                    string
		ResourceManagerEndpoint string
		SpaceID                 string
		SubscriptionID          *uuid.UUID
		TenantedDeploymentMode  string
		TenantID                *uuid.UUID
	}{
		{"Valid", false, accountType, &applicationID, &applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"EmptyName", true, accountType, &applicationID, &applicationPassword, authenticationEndpoint, azureEnvironment, emptyString, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"WhitespaceName", true, accountType, &applicationID, &applicationPassword, authenticationEndpoint, azureEnvironment, whitespaceString, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"EmptySpaceID", false, accountType, &applicationID, &applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, emptyString, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"WhitespaceSpaceID", true, accountType, &applicationID, &applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, whitespaceString, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"InvalidAccountType", true, invalidAccountType, &applicationID, &applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"NilApplicationPassword", true, accountType, &applicationID, nil, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"NilApplicationID", true, accountType, nil, &applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"NilSubscriptionID", true, accountType, &applicationID, &applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, nil, tenantedDeploymentMode, &tenantID},
		{"NilTenantID", true, accountType, &applicationID, &applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, nil},
		{"InvalidAuthenticationEndpoint", true, accountType, &applicationID, &applicationPassword, invalidURI, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"InvalidResourceManagerEndpoint", true, accountType, &applicationID, &applicationPassword, authenticationEndpoint, azureEnvironment, name, invalidURI, spaceID, &subscriptionID, tenantedDeploymentMode, &tenantID},
		{"InvalidTenantedDeploymentMode", true, accountType, &applicationID, &applicationPassword, authenticationEndpoint, azureEnvironment, name, resourceManagerEndpoint, spaceID, &subscriptionID, invalidTenantedDeploymentMode, &tenantID},
	}
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			azureServicePrincipalAccount := &AzureServicePrincipalAccount{
				AccountType:             tc.AccountType,
				ApplicationID:           tc.ApplicationID,
				ApplicationPassword:     tc.ApplicationPassword,
				AuthenticationEndpoint:  tc.AuthenticationEndpoint,
				AzureEnvironment:        tc.AzureEnvironment,
				ResourceManagerEndpoint: tc.ResourceManagerEndpoint,
				SubscriptionID:          tc.SubscriptionID,
				TenantID:                tc.TenantID,
			}
			azureServicePrincipalAccount.Name = tc.Name
			azureServicePrincipalAccount.SpaceID = tc.SpaceID
			azureServicePrincipalAccount.TenantedDeploymentMode = tc.TenantedDeploymentMode
			if tc.IsError {
				require.Error(t, azureServicePrincipalAccount.Validate())
			} else {
				require.NoError(t, azureServicePrincipalAccount.Validate())

				require.Equal(t, tc.AccountType, azureServicePrincipalAccount.GetAccountType())
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

	azureServicePrincipalAccount := NewAzureServicePrincipalAccount(name, subscriptionID, tenantID, applicationID, applicationPassword)
	require.NotNil(t, azureServicePrincipalAccount)
	require.NoError(t, azureServicePrincipalAccount.Validate())
}
