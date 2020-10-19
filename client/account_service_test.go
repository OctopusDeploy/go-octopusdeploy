package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestAmazonWebServicesAccount(t *testing.T, service *accountService) model.IAccount {
	if service == nil {
		service = createAccountService(t)
	}
	require.NotNil(t, service)

	accessKey := getRandomName()
	name := getRandomName()
	secretKey := model.NewSensitiveValue(getRandomName())

	account := model.NewAmazonWebServicesAccount(name, accessKey, secretKey)
	require.NoError(t, account.Validate())

	createdAccount, err := service.Add(account)
	require.NoError(t, err)
	require.NotNil(t, createdAccount)
	require.NotEmpty(t, createdAccount.GetID())
	require.Equal(t, "AmazonWebServicesAccount", createdAccount.GetAccountType())
	require.Equal(t, name, createdAccount.GetName())

	return createdAccount
}

func CreateTestAzureServicePrincipalAccount(t *testing.T, service *accountService) model.IAccount {
	if service == nil {
		service = createAccountService(t)
	}
	require.NotNil(t, service)

	applicationID := uuid.New()
	applicationPassword := model.NewSensitiveValue(getRandomName())
	azureEnvironment := getRandomAzureEnvironment()
	name := getRandomName()
	subscriptionID := uuid.New()
	tenantID := uuid.New()

	account := model.NewAzureServicePrincipalAccount(name, subscriptionID, tenantID, applicationID, applicationPassword)

	// set Azure environment fields
	if !isEmpty(azureEnvironment.Name) {
		account.AzureEnvironment = azureEnvironment.Name
		account.AuthenticationEndpoint = azureEnvironment.AuthenticationEndpoint
		account.ResourceManagerEndpoint = azureEnvironment.ResourceManagerEndpoint
	}

	require.NoError(t, account.Validate())

	createdAccount, err := service.Add(account)
	require.NoError(t, err)
	require.NotNil(t, createdAccount)
	require.NotEmpty(t, createdAccount.GetID())
	require.Equal(t, "AzureServicePrincipal", createdAccount.GetAccountType())
	require.Equal(t, name, createdAccount.GetName())

	return createdAccount
}

func CreateTestAzureSubscriptionAccount(t *testing.T, service *accountService) model.IAccount {
	if service == nil {
		service = createAccountService(t)
	}
	require.NotNil(t, service)

	azureEnvironment := getRandomAzureEnvironment()
	name := getRandomName()
	subscriptionID := uuid.New()

	account := model.NewAzureSubscriptionAccount(name, subscriptionID)

	// set Azure environment fields
	if !isEmpty(azureEnvironment.Name) {
		account.AzureEnvironment = azureEnvironment.Name
		account.ManagementEndpoint = azureEnvironment.ManagementEndpoint
		account.StorageEndpointSuffix = azureEnvironment.StorageEndpointSuffix
	}

	require.NoError(t, account.Validate())

	resource, err := service.Add(account)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func CreateTestSSHKeyAccount(t *testing.T, service *accountService) model.IAccount {
	if service == nil {
		service = createAccountService(t)
	}
	require.NotNil(t, service)

	name := getRandomName()
	username := getRandomName()
	privateKeyFile := model.NewSensitiveValue(getRandomName())

	account := model.NewSSHKeyAccount(name, username, privateKeyFile)

	require.NoError(t, account.Validate())

	resource, err := service.Add(account)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func CreateTestTokenAccount(t *testing.T, service *accountService) model.IAccount {
	if service == nil {
		service = createAccountService(t)
	}
	require.NotNil(t, service)

	name := getRandomName()
	token := model.NewSensitiveValue(getRandomName())

	account := model.NewTokenAccount(name, token)

	require.NoError(t, account.Validate())

	resource, err := service.Add(account)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func CreateTestUsernamePasswordAccount(t *testing.T, service *accountService) model.IAccount {
	if service == nil {
		service = createAccountService(t)
	}
	require.NotNil(t, service)

	name := getRandomName()

	account := model.NewUsernamePasswordAccount(name)

	require.NoError(t, account.Validate())

	resource, err := service.Add(account)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func createAccountService(t *testing.T) *accountService {
	service := newAccountService(nil, TestURIAccounts)
	testNewService(t, service, TestURIAccounts, serviceAccountService)
	return service
}

func DeleteTestAccount(t *testing.T, service *accountService, account model.IAccount) error {
	if service == nil {
		service = createAccountService(t)
	}
	require.NotNil(t, service)

	return service.DeleteByID(account.GetID())
}

func IsEqualAccounts(t *testing.T, expected model.IAccount, actual model.IAccount) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	// IResource
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, IsEqualLinks(expected.GetLinks(), actual.GetLinks()))

	// IAccount
	assert.Equal(t, expected.GetAccountType(), actual.GetAccountType())
	assert.Equal(t, expected.GetName(), actual.GetName())
}

func UpdateAccount(t *testing.T, service *accountService, account model.IAccount) model.IAccount {
	if service == nil {
		service = createAccountService(t)
	}
	require.NotNil(t, service)

	updatedAccount, err := service.Update(account)
	assert.NoError(t, err)
	require.NotNil(t, updatedAccount)

	return updatedAccount.(model.IAccount)
}

func ValidateAccount(t *testing.T, account model.IAccount) {
	require.NoError(t, account.Validate())
	require.NotEmpty(t, account.GetID())
	require.NotEmpty(t, account.GetLastModifiedBy())
	require.NotEmpty(t, account.GetLastModifiedOn())
	require.NotEmpty(t, account.GetLinks())

	baseAccount, ok := account.(*model.Account)
	if ok {
		require.NotEmpty(t, baseAccount.SpaceID)
	} else {
		switch account.GetAccountType() {
		case "AmazonWebServicesAccount":
			ValidateAmazonWebServicesAccount(t, account.(*model.AmazonWebServicesAccount))
		case "AzureServicePrincipal":
			ValidateAzureServicePrincipalAccount(t, account.(*model.AzureServicePrincipalAccount))
		case "AzureSubscription":
			ValidateAzureSubscriptionAccount(t, account.(*model.AzureSubscriptionAccount))
		case "SshKeyPair":
			ValidateSSHKeyAccount(t, account.(*model.SSHKeyAccount))
		case "Token":
			ValidateTokenAccount(t, account.(*model.TokenAccount))
		case "UsernamePassword":
			ValidateUsernamePasswordAccount(t, account.(*model.UsernamePasswordAccount))
		}
	}
}

func ValidateAmazonWebServicesAccount(t *testing.T, account *model.AmazonWebServicesAccount) {
	require.NotEmpty(t, account.SpaceID)
}

func ValidateAzureServicePrincipalAccount(t *testing.T, account *model.AzureServicePrincipalAccount) {
	require.NotEmpty(t, account.SpaceID)
}

func ValidateAzureSubscriptionAccount(t *testing.T, account *model.AzureSubscriptionAccount) {
	require.NotEmpty(t, account.SpaceID)
}

func ValidateSSHKeyAccount(t *testing.T, account *model.SSHKeyAccount) {
	require.NotEmpty(t, account.SpaceID)
}

func ValidateTokenAccount(t *testing.T, account *model.TokenAccount) {
	require.NotEmpty(t, account.SpaceID)
}

func ValidateUsernamePasswordAccount(t *testing.T, account *model.UsernamePasswordAccount) {
	require.NotEmpty(t, account.SpaceID)
}

func TestAccountServiceAdd(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	resource, err := service.Add(nil)
	require.Equal(t, err, createInvalidParameterError(operationAdd, parameterResource))
	require.Nil(t, resource)

	resource, err = service.Add(&model.Account{})
	require.Error(t, err)
	require.Nil(t, resource)

	amazonWebServicesAccount := CreateTestAmazonWebServicesAccount(t, service)
	ValidateAccount(t, amazonWebServicesAccount)
	defer DeleteTestAccount(t, service, amazonWebServicesAccount)

	azureServicePrincipalAccount := CreateTestAzureServicePrincipalAccount(t, service)
	ValidateAccount(t, azureServicePrincipalAccount)
	defer DeleteTestAccount(t, service, azureServicePrincipalAccount)

	azureSubscriptionAccount := CreateTestAzureSubscriptionAccount(t, service)
	ValidateAccount(t, azureSubscriptionAccount)
	defer DeleteTestAccount(t, service, azureSubscriptionAccount)

	sshKeyAccount := CreateTestSSHKeyAccount(t, service)
	ValidateAccount(t, sshKeyAccount)
	defer DeleteTestAccount(t, service, sshKeyAccount)

	tokenAccount := CreateTestTokenAccount(t, service)
	ValidateAccount(t, tokenAccount)
	defer DeleteTestAccount(t, service, tokenAccount)

	usernamePasswordAccount := CreateTestUsernamePasswordAccount(t, service)
	ValidateAccount(t, usernamePasswordAccount)
	defer DeleteTestAccount(t, service, usernamePasswordAccount)
}

func TestAccountServiceDeleteAll(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	accounts, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, accounts)

	for _, account := range accounts {
		defer DeleteTestAccount(t, service, account)
	}
}

func TestAccountServiceGetAll(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		require.NotNil(t, resource)
		assert.NotEmpty(t, resource.GetID())
	}
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

	id := getRandomName()
	resource, err = service.GetByID(id)
	require.Equal(t, createResourceNotFoundError(service.getName(), "ID", id), err)
	require.Nil(t, resource)

	resources, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, resources)

	for _, resource := range resources {
		resourceToCompare, err := service.GetByID(resource.GetID())
		require.NoError(t, err)
		IsEqualAccounts(t, resource, resourceToCompare)
	}
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

func TestAccountServiceGetByAccountType(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	accountTypes := []string{"None", "UsernamePassword", "SshKeyPair", "AzureSubscription", "AzureServicePrincipal", "AmazonWebServicesAccount", "AmazonWebServicesRoleAccount", "Token"}

	for _, accountType := range accountTypes {
		accounts, err := service.GetByAccountType(accountType)
		require.NoError(t, err)
		require.NotNil(t, accounts)

		for _, account := range accounts {
			accountToCompare, err := service.GetByID(account.GetID())
			require.NoError(t, err)
			IsEqualAccounts(t, account, accountToCompare)
		}
	}
}

func TestAccountServiceGetByName(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	amazonWebServicesAccount := CreateTestAmazonWebServicesAccount(t, service)
	ValidateAccount(t, amazonWebServicesAccount)
	defer DeleteTestAccount(t, service, amazonWebServicesAccount)

	azureServicePrincipalAccount := CreateTestAzureServicePrincipalAccount(t, service)
	ValidateAccount(t, azureServicePrincipalAccount)
	defer DeleteTestAccount(t, service, azureServicePrincipalAccount)

	azureSubscriptionAccount := CreateTestAzureSubscriptionAccount(t, service)
	ValidateAccount(t, azureSubscriptionAccount)
	defer DeleteTestAccount(t, service, azureSubscriptionAccount)

	sshKeyAccount := CreateTestSSHKeyAccount(t, service)
	ValidateAccount(t, sshKeyAccount)
	defer DeleteTestAccount(t, service, sshKeyAccount)

	tokenAccount := CreateTestTokenAccount(t, service)
	ValidateAccount(t, tokenAccount)
	defer DeleteTestAccount(t, service, tokenAccount)

	usernamePasswordAccount := CreateTestUsernamePasswordAccount(t, service)
	ValidateAccount(t, usernamePasswordAccount)
	defer DeleteTestAccount(t, service, usernamePasswordAccount)

	accounts, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, accounts)

	for _, account := range accounts {
		accountToCompare, err := service.GetByName(account.GetName())
		require.NoError(t, err)
		IsEqualAccounts(t, account, accountToCompare)
	}
}

func TestAccountServiceGetByPartialName(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	accounts, err := service.GetByPartialName(emptyString)
	require.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	require.NotNil(t, accounts)
	require.Len(t, accounts, 0)

	accounts, err = service.GetByPartialName(whitespaceString)
	require.Equal(t, err, createInvalidParameterError(operationGetByPartialName, parameterName))
	require.NotNil(t, accounts)
	require.Len(t, accounts, 0)

	accounts, err = service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, accounts)

	for _, account := range accounts {
		namedAccounts, err := service.GetByPartialName(account.GetName())
		require.NoError(t, err)
		require.NotNil(t, namedAccounts)
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
			require.Equal(t, err, createInvalidParameterError(operationGetByID, parameterID))
			require.Nil(t, resource)

			resourceList, err := service.GetByPartialName(tc.parameter)
			require.Equal(t, createInvalidParameterError(operationGetByPartialName, parameterName), err)
			require.NotNil(t, resourceList)

			err = service.DeleteByID(tc.parameter)
			require.Error(t, err)
			require.Equal(t, err, createInvalidParameterError(operationDeleteByID, parameterID))
		})
	}
}

func TestAccountServiceGetUsages(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	accounts, err := service.GetAll()
	require.NoError(t, err)

	if len(accounts) > 0 {
		accountUsages, err := service.GetUsages(accounts[0])
		require.NoError(t, err)
		require.NotNil(t, accountUsages)
	}
}

func TestAccountServiceGetByIDs(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	accounts, err := service.GetAll()
	require.NoError(t, err)
	require.NotNil(t, accounts)

	ids := []string{}
	for _, account := range accounts {
		ids = append(ids, account.GetID())
	}

	accountsByIDs, err := service.GetByIDs(ids)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(accounts), len(accountsByIDs))
}

func TestAccountServiceUpdateWithEmptyAccount(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	account, err := service.Update(nil)
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&model.Account{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&model.AmazonWebServicesAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&model.AzureServicePrincipalAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&model.AzureSubscriptionAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&model.TokenAccount{})
	require.Error(t, err)
	require.Nil(t, account)

	account, err = service.Update(&model.UsernamePasswordAccount{})
	require.Error(t, err)
	require.Nil(t, account)
}

func TestAccountServiceUpdate(t *testing.T) {
	service := createAccountService(t)
	require.NotNil(t, service)

	expected := CreateTestAzureServicePrincipalAccount(t, service)
	actual := UpdateAccount(t, service, expected)
	IsEqualAccounts(t, expected, actual)
	ValidateAccount(t, actual)
	defer DeleteTestAccount(t, service, expected)

	expected = CreateTestAzureSubscriptionAccount(t, service)
	expected.SetName(getRandomName())
	actual = UpdateAccount(t, service, expected)
	IsEqualAccounts(t, expected, actual)
	ValidateAccount(t, actual)
	defer DeleteTestAccount(t, service, expected)

	expected = CreateTestSSHKeyAccount(t, service)
	expected.SetName(getRandomName())
	actual = UpdateAccount(t, service, expected)
	IsEqualAccounts(t, expected, actual)
	ValidateAccount(t, actual)
	defer DeleteTestAccount(t, service, expected)

	expected = CreateTestTokenAccount(t, service)
	expected.SetName(getRandomName())
	actual = UpdateAccount(t, service, expected)
	IsEqualAccounts(t, expected, actual)
	ValidateAccount(t, actual)
	defer DeleteTestAccount(t, service, expected)

	expected = CreateTestUsernamePasswordAccount(t, service)
	expected.SetName(getRandomName())
	actual = UpdateAccount(t, service, expected)
	IsEqualAccounts(t, expected, actual)
	ValidateAccount(t, actual)
	defer DeleteTestAccount(t, service, expected)
}
