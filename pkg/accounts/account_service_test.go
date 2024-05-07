package accounts

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/uritemplates"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createAccountService(t *testing.T) *AccountService {
	name := constants.ServiceAccountService
	uriTemplate := constants.TestURIAccounts

	service := NewAccountService(nil, uriTemplate)
	require.NotNil(t, service)
	require.NotNil(t, service.GetClient())

	template, err := uritemplates.Parse(uriTemplate)
	require.NoError(t, err)
	require.Equal(t, service.GetURITemplate(), template)
	require.Equal(t, service.GetName(), name)

	return service
}

func TestAccountServiceAdd(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterAccount), err)
	require.Nil(t, resource)

	resource, err = service.Add(&AccountResource{})
	require.Error(t, err)
	require.Nil(t, resource)
}

func TestAccountServiceGetByID(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	resource, err := service.GetByID("")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID), err)
	require.Nil(t, resource)

	resource, err = service.GetByID(" ")
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID), err)
	require.Nil(t, resource)
}

func TestAccountServiceNew(t *testing.T) {
	client := &sling.Sling{}
	serviceName := constants.ServiceAccountService
	serviceFunction := NewAccountService
	uriTemplate := ""

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *AccountService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, ""},
		{"URITemplateWithWhitespace", serviceFunction, client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.NewServiceTests(t, service, uriTemplate, serviceName)
		})
	}
}

func TestAccountServiceParameters(t *testing.T) {
	testCases := []struct {
		name      string
		parameter string
	}{
		{"Empty", ""},
		{"Whitespace", " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := createAccountService(t)
			require.NotNil(t, service)

			resource, err := service.GetByID(tc.parameter)
			require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetByID, "id"), err)
			require.Nil(t, resource)

			err = service.DeleteByID(tc.parameter)
			require.Error(t, err)
			require.Equal(t, internal.CreateInvalidParameterError(constants.OperationDeleteByID, constants.ParameterID), err)
		})
	}
}

func TestAccountServiceUpdateWithEmptyAccount(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	account, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&AccountResource{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&AmazonWebServicesAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&AzureServicePrincipalAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&AzureOIDCAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&AwsOIDCAccount{})
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
