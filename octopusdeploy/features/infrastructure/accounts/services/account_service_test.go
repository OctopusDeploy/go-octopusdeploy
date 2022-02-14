package services

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createAccountService(t *testing.T) *accountService {
	service := NewAccountService(nil, octopusdeploy.TestURIAccounts)
	octopusdeploy.testNewService(t, service, octopusdeploy.TestURIAccounts, services.ServiceAccountService)
	return service
}

func TestAccountServiceAdd(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, octopusdeploy.createInvalidParameterError(services.OperationAdd, services.ParameterAccount))
	require.Nil(t, resource)

	resource, err = service.Add(&accounts.AccountResource{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestAccountServiceGetByID(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(octopusdeploy.emptyString)
	require.Equal(t, octopusdeploy.createInvalidParameterError(services.OperationGetByID, services.ParameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(octopusdeploy.whitespaceString)
	require.Equal(t, octopusdeploy.createInvalidParameterError(services.OperationGetByID, services.ParameterID), err)
	require.Nil(t, resource)
}

func TestAccountServiceNew(t *testing.T) {
	ServiceFunction := newAccountService
	client := &sling.Sling{}
	uriTemplate := octopusdeploy.emptyString
	ServiceName := services.ServiceAccountService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *accountService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, octopusdeploy.emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, octopusdeploy.whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			octopusdeploy.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestAccountServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", octopusdeploy.emptyString},
		{"Whitespace", octopusdeploy.whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createAccountService(t)
			require.NotNil(t, service)

			resource, err := service.GetByID(tc.parameter)
			require.Equal(t, octopusdeploy.createInvalidParameterError(services.OperationGetByID, services.ParameterID), err)
			require.Nil(t, resource)

			err = service.DeleteByID(tc.parameter)
			require.Error(t, err)
			require.Equal(t, octopusdeploy.createInvalidParameterError(services.OperationDeleteByID, services.ParameterID), err)
		})
	}
}

func TestAccountServiceUpdateWithEmptyAccount(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	account, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&accounts.AccountResource{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&accounts.AmazonWebServicesAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&services.AzureServicePrincipalAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&accounts.AzureSubscriptionAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&accounts.TokenAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&accounts.UsernamePasswordAccount{})
	require.Error(t, err)
	require.Nil(t, account)
}
