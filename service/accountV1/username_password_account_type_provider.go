package accountV1

type usernamePasswordAccountProvider struct {
}

func (p usernamePasswordAccountProvider) AccountType() AccountType {
	return AccountType(UsernamePasswordAccountType)
}

func (p usernamePasswordAccountProvider) Factory() IAccount {
	return new(UsernamePasswordAccount)
}

func init() {
	provider := &usernamePasswordAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
}
