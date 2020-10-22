package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createAccountService(t *testing.T) *accountService {
	service := newAccountService(nil, TestURIAccounts)
	testNewService(t, service, TestURIAccounts, serviceAccountService)
	return service
}

func TestAccountServiceAdd(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
	require.Nil(t, resource)

	resource, err = service.Add(&Account{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestAccountServiceGetByID(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID(emptyString)
	require.Equal(t, createInvalidParameterError(operationGetByID, parameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(whitespaceString)
	require.Equal(t, createInvalidParameterError(operationGetByID, parameterID), err)
	require.Nil(t, resource)
}

func TestAccountServiceNew(t *testing.T) {
	serviceFunction := newAccountService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceAccountService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *accountService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, emptyString},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestAccountServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", emptyString},
		{"Whitespace", whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createAccountService(t)
			require.NotNil(t, service)

			resource, err := service.GetByID(tc.parameter)
			require.Equal(t, createInvalidParameterError(operationGetByID, parameterID), err)
			require.Nil(t, resource)

			err = service.DeleteByID(tc.parameter)
			require.Error(t, err)
			require.Equal(t, createInvalidParameterError(operationDeleteByID, parameterID), err)
		})
	}
}

func TestAccountServiceUpdateWithEmptyAccount(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	account, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&Account{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&AmazonWebServicesAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&AzureServicePrincipalAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&AzureSubscriptionAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&TokenAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&UsernamePasswordAccount{})
	require.Error(t, err)
	require.Nil(t, account)
}
