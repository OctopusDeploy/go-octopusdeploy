package accountV1

type azureServicePrincipalAccountProvider struct {
}

func (p azureServicePrincipalAccountProvider) AccountType() AccountType {
	return AccountType(AzureServicePrincipalAccountType)
}

func (p azureServicePrincipalAccountProvider) Factory() IAccount {
	return new(AzureServicePrincipalAccount)
}

func init() {
	provider := &azureServicePrincipalAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
}
