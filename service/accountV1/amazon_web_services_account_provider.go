package accountV1

type amazonWebServicesAccountProvider struct {
}

func (p amazonWebServicesAccountProvider) AccountType() AccountType {
	return AccountType(AmazonWebServicesAccountType)
}

func (p amazonWebServicesAccountProvider) Factory() IAccount {
	return new(AmazonWebServicesAccount)
}

func init() {
	provider := &amazonWebServicesAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
}
