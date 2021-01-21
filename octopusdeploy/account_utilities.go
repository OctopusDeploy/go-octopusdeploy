package octopusdeploy

import "github.com/jinzhu/copier"

func ToAccount(accountResource *AccountResource) (IAccount, error) {
	if isNil(accountResource) {
		return nil, createInvalidParameterError("ToAccount", ParameterAccountResource)
	}

	var account IAccount
	var err error
	switch accountResource.GetAccountType() {
	case AccountTypeAmazonWebServicesAccount:
		account, err = NewAmazonWebServicesAccount(accountResource.GetName(), accountResource.AccessKey, accountResource.SecretKey)
		if err != nil {
			return nil, err
		}
	case AccountTypeAzureServicePrincipal:
		account, err = NewAzureServicePrincipalAccount(accountResource.GetName(), *accountResource.SubscriptionID, *accountResource.TenantID, *accountResource.ApplicationID, accountResource.ApplicationPassword)
		if err != nil {
			return nil, err
		}
	case AccountTypeAzureSubscription:
		account, err = NewAzureSubscriptionAccount(accountResource.GetName(), *accountResource.SubscriptionID)
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

	err = copier.Copy(account, accountResource)
	if err != nil {
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

func ToAccountResource(account IAccount) (*AccountResource, error) {
	if isNil(account) {
		return nil, createInvalidParameterError("ToAccountResource", ParameterAccount)
	}

	accountResource := NewAccountResource(account.GetName(), account.GetAccountType())

	if err := copier.Copy(&accountResource, account); err != nil {
		return nil, err
	}

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
