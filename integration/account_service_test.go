package integration

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestAmazonWebServicesAccount(t *testing.T, client *octopusdeploy.client) services.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	accessKey := getRandomName()
	name := getRandomName()
	secretKey := services.NewSensitiveValue(getRandomName())

	account, err := accounts.NewAmazonWebServicesAccount(name, accessKey, secretKey)
	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	createdAccount, err := client.Accounts.Add(account)
	require.NoError(t, err)
	require.NotNil(t, createdAccount)
	require.NotEmpty(t, createdAccount.GetID())
	require.Equal(t, accounts.AccountTypeAmazonWebServicesAccount, createdAccount.GetAccountType())
	require.Equal(t, name, createdAccount.GetName())

	return createdAccount
}

func CreateTestAzureServicePrincipalAccount(t *testing.T, client *octopusdeploy.client) services.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	applicationID := uuid.New()
	applicationPassword := services.NewSensitiveValue(getRandomName())
	azureEnvironment := getRandomAzureEnvironment()
	name := getRandomName()
	subscriptionID := uuid.New()
	tenantID := uuid.New()

	account, err := services.NewAzureServicePrincipalAccount(name, subscriptionID, tenantID, applicationID, applicationPassword)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// set Azure environment fields
	if !IsEmpty(azureEnvironment.Name) {
		account.AzureEnvironment = azureEnvironment.Name
		account.AuthenticationEndpoint = azureEnvironment.AuthenticationEndpoint
		account.ResourceManagerEndpoint = azureEnvironment.ResourceManagerEndpoint
	}

	require.NoError(t, account.Validate())

	createdAccount, err := client.Accounts.Add(account)
	require.NoError(t, err)
	require.NotNil(t, createdAccount)
	require.NotEmpty(t, createdAccount.GetID())
	require.Equal(t, accounts.AccountTypeAzureServicePrincipal, createdAccount.GetAccountType())
	require.Equal(t, name, createdAccount.GetName())

	return createdAccount
}

func CreateTestAzureSubscriptionAccount(t *testing.T, client *octopusdeploy.client) services.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	azureEnvironment := getRandomAzureEnvironment()
	name := getRandomName()
	subscriptionID := uuid.New()

	account, err := accounts.NewAzureSubscriptionAccount(name, subscriptionID)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// set Azure environment fields
	if !IsEmpty(azureEnvironment.Name) {
		account.AzureEnvironment = azureEnvironment.Name
		account.ManagementEndpoint = azureEnvironment.ManagementEndpoint
		account.StorageEndpointSuffix = azureEnvironment.StorageEndpointSuffix
	}

	require.NoError(t, account.Validate())

	resource, err := client.Accounts.Add(account)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func CreateTestSSHKeyAccount(t *testing.T, client *octopusdeploy.client) services.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()
	username := getRandomName()
	privateKeyFile := services.NewSensitiveValue(getRandomName())

	account, err := accounts.NewSSHKeyAccount(name, username, privateKeyFile)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	resource, err := client.Accounts.Add(account)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func CreateTestTokenAccount(t *testing.T, client *octopusdeploy.client) services.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()
	token := services.NewSensitiveValue(getRandomName())

	account, err := accounts.NewTokenAccount(name, token)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	resource, err := client.Accounts.Add(account)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func CreateTestUsernamePasswordAccount(t *testing.T, client *octopusdeploy.client) services.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()

	account, err := accounts.NewUsernamePasswordAccount(name)
	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	resource, err := client.Accounts.Add(account)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func DeleteTestAccount(t *testing.T, client *octopusdeploy.client, account services.IAccount) {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Accounts.DeleteByID(account.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedAccount, err := client.Accounts.GetByID(account.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedAccount)
}

func IsEqualAccounts(t *testing.T, expected services.IAccount, actual services.IAccount) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	// type check
	assert.IsType(t, expected, actual)

	// IResource
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.True(t, IsEqualLinks(expected.GetLinks(), actual.GetLinks()))

	// IAccount
	assert.Equal(t, expected.GetAccountType(), actual.GetAccountType())
	assert.Equal(t, expected.GetName(), actual.GetName())
}

func UpdateAccount(t *testing.T, client *octopusdeploy.client, account services.IAccount) services.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	updatedAccount, err := client.Accounts.Update(account)
	require.NoError(t, err)
	require.NotNil(t, updatedAccount)

	return updatedAccount
}

func ValidateAccount(t *testing.T, account services.IAccount) {
	require.NoError(t, account.Validate())
	require.NotEmpty(t, account.GetLinks())
	require.NotEmpty(t, account.GetID())
	require.NotEmpty(t, account.GetModifiedBy())
	require.NotEmpty(t, account.GetModifiedOn())
	require.NotEmpty(t, account.GetSpaceID())

	// TODO: validate other fields/methods
}

func TestAccountServiceDeleteAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	accounts, err := client.Accounts.GetAll()
	require.NoError(t, err)
	require.NotNil(t, accounts)

	for _, account := range accounts {
		defer DeleteTestAccount(t, client, account)
	}
}

func TestAccountServiceAddGetDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	amazonWebServicesAccount := CreateTestAmazonWebServicesAccount(t, client)
	ValidateAccount(t, amazonWebServicesAccount)
	defer DeleteTestAccount(t, client, amazonWebServicesAccount)

	azureServicePrincipalAccount := CreateTestAzureServicePrincipalAccount(t, client)
	ValidateAccount(t, azureServicePrincipalAccount)
	defer DeleteTestAccount(t, client, azureServicePrincipalAccount)

	azureSubscriptionAccount := CreateTestAzureSubscriptionAccount(t, client)
	ValidateAccount(t, azureSubscriptionAccount)
	defer DeleteTestAccount(t, client, azureSubscriptionAccount)

	sshKeyAccount := CreateTestSSHKeyAccount(t, client)
	ValidateAccount(t, sshKeyAccount)
	defer DeleteTestAccount(t, client, sshKeyAccount)

	tokenAccount := CreateTestTokenAccount(t, client)
	ValidateAccount(t, tokenAccount)
	defer DeleteTestAccount(t, client, tokenAccount)

	usernamePasswordAccount := CreateTestUsernamePasswordAccount(t, client)
	ValidateAccount(t, usernamePasswordAccount)
	defer DeleteTestAccount(t, client, usernamePasswordAccount)

	accounts, err := client.Accounts.GetAll()
	require.NoError(t, err)
	require.NotNil(t, accounts)

	for _, account := range accounts {
		name := account.GetName()
		query := services.AccountsQuery{
			PartialName: name,
			Take:        1,
		}
		namedAccounts, err := client.Accounts.Get(query)
		require.NoError(t, err)
		require.NotNil(t, namedAccounts)
		IsEqualAccounts(t, account, namedAccounts.Items[0])

		accountToCompare, err := client.Accounts.GetByID(account.GetID())
		require.NoError(t, err)
		require.NotNil(t, accountToCompare)
		IsEqualAccounts(t, account, accountToCompare)

		accountUsages, err := client.Accounts.GetUsages(accounts[0])
		require.NoError(t, err)
		require.NotNil(t, accountUsages)
	}

	accountTypes := []accounts.AccountType{
		accounts.AccountTypeNone,
		accounts.AccountTypeUsernamePassword,
		accounts.AccountTypeSSHKeyPair,
		accounts.AccountTypeAzureSubscription,
		accounts.AccountTypeAzureServicePrincipal,
		accounts.AccountTypeAmazonWebServicesAccount,
		accounts.AccountTypeToken,
	}

	for _, accountType := range accountTypes {
		query := services.AccountsQuery{AccountType: accountType}
		accounts, err := client.Accounts.Get(query)
		require.NoError(t, err)

		for _, account := range accounts.Items {
			accountToCompare, err := client.Accounts.GetByID(account.GetID())
			require.NoError(t, err)
			IsEqualAccounts(t, account, accountToCompare)
		}
	}
}

func TestAccountServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := getRandomName()
	resource, err := client.Accounts.GetByID(id)
	require.Error(t, err)
	require.Nil(t, resource)

	apiError := err.(*services.APIError)
	assert.Equal(t, 404, apiError.StatusCode)

	accounts, err := client.Accounts.GetAll()
	require.NoError(t, err)
	require.NotNil(t, accounts)

	for _, account := range accounts {
		accountToCompare, err := client.Accounts.GetByID(account.GetID())
		require.NoError(t, err)
		IsEqualAccounts(t, account, accountToCompare)
	}
}

func TestAccountServiceGetByIDs(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	accounts, err := client.Accounts.GetAll()
	require.NoError(t, err)
	require.NotNil(t, accounts)

	ids := []string{}
	for _, account := range accounts {
		ids = append(ids, account.GetID())
	}

	query := services.AccountsQuery{IDs: ids}

	accountsByIDs, err := client.Accounts.Get(query)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(accounts), len(accountsByIDs.Items))
}

func TestAccountServiceTokenAccounts(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	for i := 0; i < 10; i++ {
		tokenAccount := CreateTestTokenAccount(t, client)
		ValidateAccount(t, tokenAccount)
		defer DeleteTestAccount(t, client, tokenAccount)
	}

	accounts, err := client.Accounts.GetAll()
	require.NoError(t, err)
	require.NotNil(t, accounts)

	for _, account := range accounts {
		accountToCompare, err := client.Accounts.GetByID(account.GetID())
		require.NoError(t, err)
		IsEqualAccounts(t, account, accountToCompare)
	}
}

func TestAccountServiceUpdate(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	expected := CreateTestAzureServicePrincipalAccount(t, client)
	actual := UpdateAccount(t, client, expected)
	IsEqualAccounts(t, expected, actual)
	ValidateAccount(t, actual)
	defer DeleteTestAccount(t, client, expected)

	expected = CreateTestAzureSubscriptionAccount(t, client)
	expected.SetName(getRandomName())
	actual = UpdateAccount(t, client, expected)
	IsEqualAccounts(t, expected, actual)
	ValidateAccount(t, actual)
	defer DeleteTestAccount(t, client, expected)

	expected = CreateTestSSHKeyAccount(t, client)
	expected.SetName(getRandomName())
	actual = UpdateAccount(t, client, expected)
	IsEqualAccounts(t, expected, actual)
	ValidateAccount(t, actual)
	defer DeleteTestAccount(t, client, expected)

	expected = CreateTestTokenAccount(t, client)
	expected.SetName(getRandomName())
	actual = UpdateAccount(t, client, expected)
	IsEqualAccounts(t, expected, actual)
	ValidateAccount(t, actual)
	defer DeleteTestAccount(t, client, expected)

	expected = CreateTestUsernamePasswordAccount(t, client)
	expected.SetName(getRandomName())
	actual = UpdateAccount(t, client, expected)
	IsEqualAccounts(t, expected, actual)
	ValidateAccount(t, actual)
	defer DeleteTestAccount(t, client, expected)
}
