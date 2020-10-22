package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateTestAmazonWebServicesAccount(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	accessKey := getRandomName()
	name := getRandomName()
	secretKey := octopusdeploy.NewSensitiveValue(getRandomName())

	account := octopusdeploy.NewAmazonWebServicesAccount(name, accessKey, secretKey)
	require.NoError(t, account.Validate())

	createdAccount, err := client.Accounts.Add(account)
	require.NoError(t, err)
	require.NotNil(t, createdAccount)
	require.NotEmpty(t, createdAccount.GetID())
	require.Equal(t, accountTypeAmazonWebServicesAccount, createdAccount.GetAccountType())
	require.Equal(t, name, createdAccount.GetName())

	return createdAccount
}

func CreateTestAzureServicePrincipalAccount(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	applicationID := uuid.New()
	applicationPassword := octopusdeploy.NewSensitiveValue(getRandomName())
	azureEnvironment := getRandomAzureEnvironment()
	name := getRandomName()
	subscriptionID := uuid.New()
	tenantID := uuid.New()

	account := octopusdeploy.NewAzureServicePrincipalAccount(name, subscriptionID, tenantID, applicationID, applicationPassword)

	// set Azure environment fields
	if !isEmpty(azureEnvironment.Name) {
		account.AzureEnvironment = azureEnvironment.Name
		account.AuthenticationEndpoint = azureEnvironment.AuthenticationEndpoint
		account.ResourceManagerEndpoint = azureEnvironment.ResourceManagerEndpoint
	}

	require.NoError(t, account.Validate())

	createdAccount, err := client.Accounts.Add(account)
	require.NoError(t, err)
	require.NotNil(t, createdAccount)
	require.NotEmpty(t, createdAccount.GetID())
	require.Equal(t, accountTypeAzureServicePrincipal, createdAccount.GetAccountType())
	require.Equal(t, name, createdAccount.GetName())

	return createdAccount
}

func CreateTestAzureSubscriptionAccount(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	azureEnvironment := getRandomAzureEnvironment()
	name := getRandomName()
	subscriptionID := uuid.New()

	account := octopusdeploy.NewAzureSubscriptionAccount(name, subscriptionID)

	// set Azure environment fields
	if !isEmpty(azureEnvironment.Name) {
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

func CreateTestSSHKeyAccount(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()
	username := getRandomName()
	privateKeyFile := octopusdeploy.NewSensitiveValue(getRandomName())

	account := octopusdeploy.NewSSHKeyAccount(name, username, privateKeyFile)

	require.NoError(t, account.Validate())

	resource, err := client.Accounts.Add(account)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func CreateTestTokenAccount(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()
	token := octopusdeploy.NewSensitiveValue(getRandomName())

	account := octopusdeploy.NewTokenAccount(name, token)

	require.NoError(t, account.Validate())

	resource, err := client.Accounts.Add(account)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func CreateTestUsernamePasswordAccount(t *testing.T, client *octopusdeploy.Client) octopusdeploy.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := getRandomName()

	account := octopusdeploy.NewUsernamePasswordAccount(name)

	require.NoError(t, account.Validate())

	resource, err := client.Accounts.Add(account)
	require.NoError(t, err)
	require.NotNil(t, resource)

	return resource
}

func DeleteTestAccount(t *testing.T, client *octopusdeploy.Client, account octopusdeploy.IAccount) {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Accounts.DeleteByID(account.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	accounts, err := client.Accounts.GetByID(account.GetID())
	assert.Error(t, err)
	assert.Nil(t, accounts)
}

func IsEqualAccounts(t *testing.T, expected octopusdeploy.IAccount, actual octopusdeploy.IAccount) {
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

func UpdateAccount(t *testing.T, client *octopusdeploy.Client, account octopusdeploy.IAccount) octopusdeploy.IAccount {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	updatedAccount, err := client.Accounts.Update(account)
	assert.NoError(t, err)
	require.NotNil(t, updatedAccount)

	return updatedAccount.(octopusdeploy.IAccount)
}

func ValidateAccount(t *testing.T, account octopusdeploy.IAccount) {
	require.NoError(t, account.Validate())
	require.NotEmpty(t, account.GetID())
	require.NotEmpty(t, account.GetModifiedBy())
	require.NotEmpty(t, account.GetModifiedOn())
	require.NotEmpty(t, account.GetLinks())

	baseAccount, ok := account.(*octopusdeploy.Account)
	if ok {
		require.NotEmpty(t, baseAccount.SpaceID)
	} else {
		switch account.GetAccountType() {
		case accountTypeAmazonWebServicesAccount:
			ValidateAmazonWebServicesAccount(t, account.(*octopusdeploy.AmazonWebServicesAccount))
		case accountTypeAzureServicePrincipal:
			ValidateAzureServicePrincipalAccount(t, account.(*octopusdeploy.AzureServicePrincipalAccount))
		case accountTypeAzureSubscription:
			ValidateAzureSubscriptionAccount(t, account.(*octopusdeploy.AzureSubscriptionAccount))
		case accountTypeSshKeyPair:
			ValidateSSHKeyAccount(t, account.(*octopusdeploy.SSHKeyAccount))
		case accountTypeToken:
			ValidateTokenAccount(t, account.(*octopusdeploy.TokenAccount))
		case accountTypeUsernamePassword:
			ValidateUsernamePasswordAccount(t, account.(*octopusdeploy.UsernamePasswordAccount))
		}
	}
}

func ValidateAmazonWebServicesAccount(t *testing.T, account *octopusdeploy.AmazonWebServicesAccount) {
	require.NotEmpty(t, account.SpaceID)
}

func ValidateAzureServicePrincipalAccount(t *testing.T, account *octopusdeploy.AzureServicePrincipalAccount) {
	require.NotEmpty(t, account.SpaceID)
}

func ValidateAzureSubscriptionAccount(t *testing.T, account *octopusdeploy.AzureSubscriptionAccount) {
	require.NotEmpty(t, account.SpaceID)
}

func ValidateSSHKeyAccount(t *testing.T, account *octopusdeploy.SSHKeyAccount) {
	require.NotEmpty(t, account.SpaceID)
}

func ValidateTokenAccount(t *testing.T, account *octopusdeploy.TokenAccount) {
	require.NotEmpty(t, account.SpaceID)
}

func ValidateUsernamePasswordAccount(t *testing.T, account *octopusdeploy.UsernamePasswordAccount) {
	require.NotEmpty(t, account.SpaceID)
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

func TestAccountServiceGetByID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	id := getRandomName()
	resource, err := client.Accounts.GetByID(id)
	require.Equal(t, createResourceNotFoundError(serviceAccountService, "ID", id), err)
	require.Nil(t, resource)
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
		query := octopusdeploy.AccountsQuery{PartialName: account.GetName()}
		namedAccounts, err := client.Accounts.Get(query)
		require.NoError(t, err)
		require.NotNil(t, namedAccounts)
		IsEqualAccounts(t, account, namedAccounts.Items[0])

		query = octopusdeploy.AccountsQuery{IDs: []string{account.GetID()}}
		namedAccounts, err = client.Accounts.Get(query)
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

	accountTypes := []string{"None", accountTypeUsernamePassword, accountTypeSshKeyPair, accountTypeAzureSubscription, accountTypeAzureServicePrincipal, accountTypeAmazonWebServicesAccount, "AmazonWebServicesRoleAccount", accountTypeToken}

	for _, accountType := range accountTypes {
		query := octopusdeploy.AccountsQuery{AccountType: accountType}
		accounts, err := client.Accounts.Get(query)
		require.NoError(t, err)

		for _, account := range accounts.Items {
			accountToCompare, err := client.Accounts.GetByID(account.GetID())
			require.NoError(t, err)
			IsEqualAccounts(t, account, accountToCompare)
		}
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

	query := octopusdeploy.AccountsQuery{IDs: ids}

	accountsByIDs, err := client.Accounts.Get(query)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(accounts), len(accountsByIDs.Items))
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
