package accountV1

type tokenAccountProvider struct {
	IAccountTypeProvider
}

func (p tokenAccountProvider) AccountType() AccountType {
	return AccountType(TokenAccountType)
}

func (p tokenAccountProvider) Factory() IAccount {
	return new(TokenAccount)
}

func doTokenAccountTypeRegistration() *tokenAccountProvider {
	provider := tokenAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
	return &provider
}

var _ = doTokenAccountTypeRegistration()
var _ IAccountTypeProvider = new(tokenAccountProvider)
