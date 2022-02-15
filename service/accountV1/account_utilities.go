package accountV1

func ToAccount(accountResource *AccountResource) (IAccount, error) {
	// TODO: IMPL
	return nil, nil
	// if octopusdeploy.isNil(accountResource) {
	// 	return nil, resources.CreateInvalidParameterError("ToAccount", octopusdeploy.ParameterAccountResource)
	// }

	// var accountV1 resources.IAccount
	// var err error
	// switch accountResource.GetAccountType() {
	// case AccountTypeAmazonWebresourcesAccount:
	// 	accountV1, err = NewAmazonWebresourcesAccount(accountResource.GetName(), accountResource.AccessKey, accountResource.SecretKey)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeAzureServicePrincipal:
	// 	accountV1, err = resources.NewAzureServicePrincipalAccount(accountResource.GetName(), *accountResource.SubscriptionID, *accountResource.TenantID, *accountResource.ApplicationID, accountResource.ApplicationPassword)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeAzureSubscription:
	// 	accountV1, err = NewAzureSubscriptionAccount(accountResource.GetName(), *accountResource.SubscriptionID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeGoogleCloudPlatformAccount:
	// 	accountV1, err = resources.NewGoogleCloudPlatformAccount(accountResource.GetName(), accountResource.JsonKey)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeSSHKeyPair:
	// 	accountV1, err = NewSSHKeyAccount(accountResource.GetName(), accountResource.Username, accountResource.PrivateKeyFile)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeToken:
	// 	accountV1, err = NewTokenAccount(accountResource.GetName(), accountResource.Token)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// case AccountTypeUsernamePassword:
	// 	accountV1, err = NewUsernamePasswordAccount(accountResource.GetName())
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// if err := copier.Copy(accountV1, accountResource); err != nil {
	// 	return nil, err
	// }

	// return accountV1, nil
}

// func ToAccountArray(accountResources []*AccountResource) []resources.IAccount {
// 	items := []resources.IAccount{}
// 	for _, accountResource := range accountResources {
// 		accountV1, err := ToAccount(accountResource)
// 		if err != nil {
// 			return nil
// 		}
// 		items = append(items, accountV1)
// 	}
// 	return items
// }
