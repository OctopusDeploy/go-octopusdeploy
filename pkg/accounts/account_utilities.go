package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
)

func ToAccount(accountResource *AccountResource) (IAccount, error) {
	if accountResource == nil {
		return nil, internal.CreateInvalidParameterError("ToAccount", constants.ParameterAccountResource)
	}

	if err := accountResource.Validate(); err != nil {
		return nil, err
	}

	var account IAccount

	switch accountResource.GetAccountType() {
	case AccountTypeAmazonWebServicesAccount:
		amazonWebServicesAccount, err := NewAmazonWebServicesAccount(accountResource.GetName(), accountResource.AccessKey, accountResource.SecretKey)
		if err != nil {
			return nil, err
		}
		account = amazonWebServicesAccount
	case AccountTypeAzureServicePrincipal:
		azureServicePrincipalAccount, err := NewAzureServicePrincipalAccount(accountResource.GetName(), *accountResource.SubscriptionID, *accountResource.TenantID, *accountResource.ApplicationID, accountResource.ApplicationPassword)
		if err != nil {
			return nil, err
		}
		azureServicePrincipalAccount.AuthenticationEndpoint = accountResource.AuthenticationEndpoint
		azureServicePrincipalAccount.AzureEnvironment = accountResource.AzureEnvironment
		azureServicePrincipalAccount.ResourceManagerEndpoint = accountResource.ResourceManagerEndpoint
		account = azureServicePrincipalAccount
	case AccountTypeAzureSubscription:
		azureSubscriptionAccount, err := NewAzureSubscriptionAccount(accountResource.GetName(), *accountResource.SubscriptionID)
		if err != nil {
			return nil, err
		}
		azureSubscriptionAccount.AzureEnvironment = accountResource.AzureEnvironment
		azureSubscriptionAccount.CertificateBytes = accountResource.CertificateBytes
		azureSubscriptionAccount.CertificateThumbprint = accountResource.CertificateThumbprint
		azureSubscriptionAccount.ManagementEndpoint = accountResource.ManagementEndpoint
		azureSubscriptionAccount.StorageEndpointSuffix = accountResource.StorageEndpointSuffix
		account = azureSubscriptionAccount
	case AccountTypeGoogleCloudPlatformAccount:
		googleCloudPlatformAccount, err := NewGoogleCloudPlatformAccount(accountResource.GetName(), accountResource.JsonKey)
		if err != nil {
			return nil, err
		}
		account = googleCloudPlatformAccount
	case AccountTypeSSHKeyPair:
		sshKeyAccount, err := NewSSHKeyAccount(accountResource.GetName(), accountResource.Username, accountResource.PrivateKeyFile)
		if err != nil {
			return nil, err
		}
		sshKeyAccount.SetPrivateKeyPassphrase(accountResource.PrivateKeyPassphrase)
		account = sshKeyAccount
	case AccountTypeToken:
		tokenAccount, err := NewTokenAccount(accountResource.GetName(), accountResource.Token)
		if err != nil {
			return nil, err
		}
		account = tokenAccount
	case AccountTypeUsernamePassword:
		usernamePasswordAccount, err := NewUsernamePasswordAccount(accountResource.GetName())
		if err != nil {
			return nil, err
		}
		usernamePasswordAccount.SetPassword(accountResource.ApplicationPassword)
		usernamePasswordAccount.SetUsername(accountResource.Username)
		account = usernamePasswordAccount
	}

	account.SetDescription(accountResource.GetDescription())
	account.SetEnvironmentIDs(accountResource.GetEnvironmentIDs())
	account.SetLinks(accountResource.GetLinks())
	account.SetModifiedBy(accountResource.GetModifiedBy())
	account.SetModifiedOn(accountResource.GetModifiedOn())
	account.SetID(accountResource.GetID())
	account.SetSpaceID(accountResource.GetSpaceID())
	account.SetTenantedDeploymentMode(accountResource.GetTenantedDeploymentMode())
	account.SetTenantIDs(accountResource.GetTenantIDs())
	account.SetTenantTags(accountResource.GetTenantTags())

	return account, nil
}

func ToAccounts(accountResources *AccountResources) *Accounts {
	return &Accounts{
		Items:        ToAccountArray(accountResources.Items),
		PagedResults: accountResources.PagedResults,
	}
}

func ToAccountResource(account IAccount) (*AccountResource, error) {
	if IsNil(account) {
		return nil, internal.CreateInvalidParameterError("ToAccountResource", constants.ParameterAccount)
	}

	// conversion unnecessary if input account is *AccountResource
	if v, ok := account.(*AccountResource); ok {
		return v, nil
	}

	accountResource := NewAccountResource(account.GetName(), account.GetAccountType())

	switch accountResource.GetAccountType() {
	case AccountTypeAmazonWebServicesAccount:
		amazonWebServicesAccount := account.(*AmazonWebServicesAccount)
		accountResource.AccessKey = amazonWebServicesAccount.AccessKey
		accountResource.SecretKey = amazonWebServicesAccount.SecretKey
	case AccountTypeAzureServicePrincipal:
		azureServicePrincipalAccount := account.(*AzureServicePrincipalAccount)
		accountResource.ApplicationID = azureServicePrincipalAccount.ApplicationID
		accountResource.ApplicationPassword = azureServicePrincipalAccount.ApplicationPassword
		accountResource.AuthenticationEndpoint = azureServicePrincipalAccount.AuthenticationEndpoint
		accountResource.AzureEnvironment = azureServicePrincipalAccount.AzureEnvironment
		accountResource.ResourceManagerEndpoint = azureServicePrincipalAccount.ResourceManagerEndpoint
		accountResource.SubscriptionID = azureServicePrincipalAccount.SubscriptionID
		accountResource.TenantID = azureServicePrincipalAccount.TenantID
	case AccountTypeAzureSubscription:
		azureSubscriptionAccount := account.(*AzureSubscriptionAccount)
		accountResource.AzureEnvironment = azureSubscriptionAccount.AzureEnvironment
		accountResource.CertificateBytes = azureSubscriptionAccount.CertificateBytes
		accountResource.CertificateThumbprint = azureSubscriptionAccount.CertificateThumbprint
		accountResource.ManagementEndpoint = azureSubscriptionAccount.ManagementEndpoint
		accountResource.StorageEndpointSuffix = azureSubscriptionAccount.StorageEndpointSuffix
		accountResource.SubscriptionID = azureSubscriptionAccount.SubscriptionID
	case AccountTypeGoogleCloudPlatformAccount:
		googleCloudPlatformAccount := account.(*GoogleCloudPlatformAccount)
		accountResource.JsonKey = googleCloudPlatformAccount.JsonKey
	case AccountTypeSSHKeyPair:
		sshKeyAccount := account.(*SSHKeyAccount)
		accountResource.PrivateKeyFile = sshKeyAccount.PrivateKeyFile
		accountResource.PrivateKeyPassphrase = sshKeyAccount.PrivateKeyPassphrase
		accountResource.Username = sshKeyAccount.Username
	case AccountTypeToken:
		tokenAccount := account.(*TokenAccount)
		accountResource.Token = tokenAccount.Token
	case AccountTypeUsernamePassword:
		usernamePasswordAccount := account.(*UsernamePasswordAccount)
		accountResource.Username = usernamePasswordAccount.Username
		accountResource.ApplicationPassword = usernamePasswordAccount.Password
	}

	accountResource.SetDescription(account.GetDescription())
	accountResource.SetEnvironmentIDs(account.GetEnvironmentIDs())
	accountResource.SetLinks(account.GetLinks())
	accountResource.SetModifiedBy(account.GetModifiedBy())
	accountResource.SetModifiedOn(account.GetModifiedOn())
	accountResource.SetID(account.GetID())
	accountResource.SetSpaceID(account.GetSpaceID())
	accountResource.SetTenantedDeploymentMode(account.GetTenantedDeploymentMode())
	accountResource.SetTenantIDs(account.GetTenantIDs())
	accountResource.SetTenantTags(account.GetTenantTags())

	return accountResource, nil
}

func ToAccountArray(accountResources []*AccountResource) []IAccount {
	items := []IAccount{}
	for _, accountResource := range accountResources {
		account, err := ToAccount(accountResource)
		if err != nil {
			return nil
		}
		items = append(items, account)
	}
	return items
}
