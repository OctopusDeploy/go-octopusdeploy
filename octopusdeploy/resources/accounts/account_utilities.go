package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/jinzhu/copier"
)

func ToAccount(accountResource *AccountResource) (octopusdeploy.IAccount, error) {
	if octopusdeploy.isNil(accountResource) {
		return nil, octopusdeploy.CreateInvalidParameterError("ToAccount", octopusdeploy.ParameterAccountResource)
	}

	var account octopusdeploy.IAccount
	var err error
	switch accountResource.GetAccountType() {
	case AccountTypeAmazonWebServicesAccount:
		account, err = NewAmazonWebServicesAccount(accountResource.GetName(), accountResource.AccessKey, accountResource.SecretKey)
		if err != nil {
			return nil, err
		}
	case AccountTypeAzureServicePrincipal:
		account, err = octopusdeploy.NewAzureServicePrincipalAccount(accountResource.GetName(), *accountResource.SubscriptionID, *accountResource.TenantID, *accountResource.ApplicationID, accountResource.ApplicationPassword)
		if err != nil {
			return nil, err
		}
	case AccountTypeAzureSubscription:
		account, err = NewAzureSubscriptionAccount(accountResource.GetName(), *accountResource.SubscriptionID)
		if err != nil {
			return nil, err
		}
	case AccountTypeGoogleCloudPlatformAccount:
		account, err = octopusdeploy.NewGoogleCloudPlatformAccount(accountResource.GetName(), accountResource.JsonKey)
		if err != nil {
			return nil, err
		}
	case AccountTypeSSHKeyPair:
		account, err = NewSSHKeyAccount(accountResource.GetName(), accountResource.Username, accountResource.PrivateKeyFile)
		if err != nil {
			return nil, err
		}
	case AccountTypeToken:
		account, err = NewTokenAccount(accountResource.GetName(), accountResource.Token)
		if err != nil {
			return nil, err
		}
	case AccountTypeUsernamePassword:
		account, err = NewUsernamePasswordAccount(accountResource.GetName())
		if err != nil {
			return nil, err
		}
	}

	if err := copier.Copy(account, accountResource); err != nil {
		return nil, err
	}

	return account, nil
}

func ToAccounts(accountResources *AccountResources) *Accounts {
	return &Accounts{
		Items:        ToAccountArray(accountResources.Items),
		PagedResults: accountResources.PagedResults,
	}
}

func ToAccountResource(client octopusdeploy.spaceScopedClient, account octopusdeploy.IAccount) (*AccountResource, error) {
	if octopusdeploy.isNil(account) {
		return nil, octopusdeploy.CreateInvalidParameterError("ToAccountResource", octopusdeploy.ParameterAccount)
	}

	spaceID, err := octopusdeploy.GetSpaceIDForResource(account, client)
	if err != nil {
		return nil, err
	}
	accountResource := NewAccountResource(spaceID, account.GetName(), account.GetAccountType())

	if err := copier.Copy(&accountResource, account); err != nil {
		return nil, err
	}

	return accountResource, nil
}

func ToAccountArray(accountResources []*AccountResource) []octopusdeploy.IAccount {
	items := []octopusdeploy.IAccount{}
	for _, accountResource := range accountResources {
		account, err := ToAccount(accountResource)
		if err != nil {
			return nil
		}
		items = append(items, account)
	}
	return items
}
