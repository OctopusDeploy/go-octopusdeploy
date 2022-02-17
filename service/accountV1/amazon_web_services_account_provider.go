package accountV1

type amazonWebServicesAccountProvider struct {
}

func (p amazonWebServicesAccountProvider) AccountType() AccountType {
	return AccountType(UsernamePasswordAccountType)
}

func (p amazonWebServicesAccountProvider) Factory() IAccount {
	return new(UsernamePasswordAccount)
}

func init() {
	provider := &amazonWebServicesAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
}
