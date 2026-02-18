package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
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
		amazonWebServicesAccount.Region = accountResource.Region
		account = amazonWebServicesAccount
	case AccountTypeAwsOIDC:
		awsOIDCAccount, err := NewAwsOIDCAccount(accountResource.GetName(), accountResource.RoleArn)
		if err != nil {
			return nil, err
		}
		awsOIDCAccount.RoleArn = accountResource.RoleArn
		awsOIDCAccount.SessionDuration = accountResource.SessionDuration
		awsOIDCAccount.Region = accountResource.Region
		awsOIDCAccount.Audience = accountResource.Audience
		awsOIDCAccount.DeploymentSubjectKeys = accountResource.DeploymentSubjectKeys
		awsOIDCAccount.AccountTestSubjectKeys = accountResource.AccountTestSubjectKeys
		awsOIDCAccount.HealthCheckSubjectKeys = accountResource.HealthCheckSubjectKeys
		awsOIDCAccount.CustomClaims = accountResource.CustomClaims
		account = awsOIDCAccount
	case AccountTypeAzureServicePrincipal:
		azureServicePrincipalAccount, err := NewAzureServicePrincipalAccount(accountResource.GetName(), *accountResource.SubscriptionID, *accountResource.TenantID, *accountResource.ApplicationID, accountResource.ApplicationPassword)
		if err != nil {
			return nil, err
		}
		azureServicePrincipalAccount.AuthenticationEndpoint = accountResource.AuthenticationEndpoint
		azureServicePrincipalAccount.AzureEnvironment = accountResource.AzureEnvironment
		azureServicePrincipalAccount.ResourceManagerEndpoint = accountResource.ResourceManagerEndpoint
		account = azureServicePrincipalAccount
	case AccountTypeAzureOIDC:
		azureOIDCAccount, err := NewAzureOIDCAccount(accountResource.GetName(), *accountResource.SubscriptionID, *accountResource.TenantID, *accountResource.ApplicationID)
		if err != nil {
			return nil, err
		}
		azureOIDCAccount.AuthenticationEndpoint = accountResource.AuthenticationEndpoint
		azureOIDCAccount.AzureEnvironment = accountResource.AzureEnvironment
		azureOIDCAccount.ResourceManagerEndpoint = accountResource.ResourceManagerEndpoint
		azureOIDCAccount.Audience = accountResource.Audience
		azureOIDCAccount.DeploymentSubjectKeys = accountResource.DeploymentSubjectKeys
		azureOIDCAccount.AccountTestSubjectKeys = accountResource.AccountTestSubjectKeys
		azureOIDCAccount.HealthCheckSubjectKeys = accountResource.HealthCheckSubjectKeys
		azureOIDCAccount.CustomClaims = accountResource.CustomClaims
		account = azureOIDCAccount
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
	case AccountTypeGenericOIDCAccount:
		genericOIDCAccount, err := NewGenericOIDCAccount(accountResource.GetName())
		if err != nil {
			return nil, err
		}
		genericOIDCAccount.Audience = accountResource.Audience
		genericOIDCAccount.DeploymentSubjectKeys = accountResource.DeploymentSubjectKeys
		genericOIDCAccount.CustomClaims = accountResource.CustomClaims
		account = genericOIDCAccount
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
	account.SetSlug(accountResource.GetSlug())
	account.SetSpaceID(accountResource.GetSpaceID())
	account.SetTenantedDeploymentMode(accountResource.GetTenantedDeploymentMode())
	account.SetTenantIDs(accountResource.GetTenantIDs())
	account.SetTenantTags(accountResource.GetTenantTags())
	account.SetSlug(accountResource.GetSlug())

	return account, nil
}

func ToAccounts(accountResources *resources.Resources[*AccountResource]) *Accounts {
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
		accountResource.Region = amazonWebServicesAccount.Region
	case AccountTypeAwsOIDC:
		awsOIDCAccount := account.(*AwsOIDCAccount)
		accountResource.RoleArn = awsOIDCAccount.RoleArn
		accountResource.SessionDuration = awsOIDCAccount.SessionDuration
		accountResource.Region = awsOIDCAccount.Region
		accountResource.Audience = awsOIDCAccount.Audience
		accountResource.DeploymentSubjectKeys = awsOIDCAccount.DeploymentSubjectKeys
		accountResource.AccountTestSubjectKeys = awsOIDCAccount.AccountTestSubjectKeys
		accountResource.HealthCheckSubjectKeys = awsOIDCAccount.HealthCheckSubjectKeys
		accountResource.CustomClaims = awsOIDCAccount.CustomClaims
	case AccountTypeAzureServicePrincipal:
		azureServicePrincipalAccount := account.(*AzureServicePrincipalAccount)
		accountResource.ApplicationID = azureServicePrincipalAccount.ApplicationID
		accountResource.ApplicationPassword = azureServicePrincipalAccount.ApplicationPassword
		accountResource.AuthenticationEndpoint = azureServicePrincipalAccount.AuthenticationEndpoint
		accountResource.AzureEnvironment = azureServicePrincipalAccount.AzureEnvironment
		accountResource.ResourceManagerEndpoint = azureServicePrincipalAccount.ResourceManagerEndpoint
		accountResource.SubscriptionID = azureServicePrincipalAccount.SubscriptionID
		accountResource.TenantID = azureServicePrincipalAccount.TenantID
	case AccountTypeAzureOIDC:
		azureOIDCAccount := account.(*AzureOIDCAccount)
		accountResource.ApplicationID = azureOIDCAccount.ApplicationID
		accountResource.AuthenticationEndpoint = azureOIDCAccount.AuthenticationEndpoint
		accountResource.AzureEnvironment = azureOIDCAccount.AzureEnvironment
		accountResource.ResourceManagerEndpoint = azureOIDCAccount.ResourceManagerEndpoint
		accountResource.SubscriptionID = azureOIDCAccount.SubscriptionID
		accountResource.TenantID = azureOIDCAccount.TenantID
		accountResource.Audience = azureOIDCAccount.Audience
		accountResource.DeploymentSubjectKeys = azureOIDCAccount.DeploymentSubjectKeys
		accountResource.AccountTestSubjectKeys = azureOIDCAccount.AccountTestSubjectKeys
		accountResource.HealthCheckSubjectKeys = azureOIDCAccount.HealthCheckSubjectKeys
		accountResource.CustomClaims = azureOIDCAccount.CustomClaims
	case AccountTypeAzureSubscription:
		azureSubscriptionAccount := account.(*AzureSubscriptionAccount)
		accountResource.AzureEnvironment = azureSubscriptionAccount.AzureEnvironment
		accountResource.CertificateBytes = azureSubscriptionAccount.CertificateBytes
		accountResource.CertificateThumbprint = azureSubscriptionAccount.CertificateThumbprint
		accountResource.ManagementEndpoint = azureSubscriptionAccount.ManagementEndpoint
		accountResource.StorageEndpointSuffix = azureSubscriptionAccount.StorageEndpointSuffix
		accountResource.SubscriptionID = azureSubscriptionAccount.SubscriptionID
	case AccountTypeGenericOIDCAccount:
		genericOidcAccount := account.(*GenericOIDCAccount)
		accountResource.DeploymentSubjectKeys = genericOidcAccount.DeploymentSubjectKeys
		accountResource.Audience = genericOidcAccount.Audience
		accountResource.CustomClaims = genericOidcAccount.CustomClaims
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
	accountResource.SetSlug(account.GetSlug())
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
