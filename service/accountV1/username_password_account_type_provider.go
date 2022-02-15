package accountV1

type usernamePasswordAccountProvider struct {
}

func (p usernamePasswordAccountProvider) AccountType() AccountType {
	return AccountType(UsernamePasswordAccountType)
}

func (p usernamePasswordAccountProvider) Factory() IAccount {
	return new(UsernamePasswordAccount)
}

func doUsernamePasswordAccountTypeRegistration() *usernamePasswordAccountProvider {
	provider := &usernamePasswordAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
	return provider
}

var _ = doUsernamePasswordAccountTypeRegistration()
var _ IAccountTypeProvider = new(usernamePasswordAccountProvider)
