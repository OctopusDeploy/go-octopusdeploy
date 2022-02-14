package accounts

func ToAccount(accountResource *AccountResource) (IAccount, error) {
	// TODO: IMPL
	return nil, nil
	// if octopusdeploy.isNil(accountResource) {
	// 	return nil, resources.CreateInvalidParameterError("ToAccount", octopusdeploy.ParameterAccountResource)
	// }

	// var account resources.IAccount
	// var err error
	// switch accountResource.GetAccountType() {
	// case AccountTypeAmazonWebresourcesAccount:
	// 	account, err = NewAmazonWebresourcesAccount(accountResource.GetName(), accountResource.AccessKey, accountResource.SecretKey)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeAzureServicePrincipal:
	// 	account, err = resources.NewAzureServicePrincipalAccount(accountResource.GetName(), *accountResource.SubscriptionID, *accountResource.TenantID, *accountResource.ApplicationID, accountResource.ApplicationPassword)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeAzureSubscription:
	// 	account, err = NewAzureSubscriptionAccount(accountResource.GetName(), *accountResource.SubscriptionID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeGoogleCloudPlatformAccount:
	// 	account, err = resources.NewGoogleCloudPlatformAccount(accountResource.GetName(), accountResource.JsonKey)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeSSHKeyPair:
	// 	account, err = NewSSHKeyAccount(accountResource.GetName(), accountResource.Username, accountResource.PrivateKeyFile)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeToken:
	// 	account, err = NewTokenAccount(accountResource.GetName(), accountResource.Token)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeUsernamePassword:
	// 	account, err = NewUsernamePasswordAccount(accountResource.GetName())
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// if err := copier.Copy(account, accountResource); err != nil {
	// 	return nil, err
	// }

	// return account, nil
}

// func ToAccountArray(accountResources []*AccountResource) []resources.IAccount {
// 	items := []resources.IAccount{}
// 	for _, accountResource := range accountResources {
// 		account, err := ToAccount(accountResource)
// 		if err != nil {
// 			return nil
// 		}
// 		items = append(items, account)
// 	}
// 	return items
// }
