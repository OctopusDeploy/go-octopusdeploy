package accountV1

type azureScriptionAccountProvider struct {
}

func (p azureScriptionAccountProvider) AccountType() AccountType {
	return AccountType(AzureSubscriptionAccountType)
}

func (p azureScriptionAccountProvider) Factory() IAccount {
	return new(AzureSubscriptionAccount)
}

func init() {
	provider := &azureScriptionAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
}
