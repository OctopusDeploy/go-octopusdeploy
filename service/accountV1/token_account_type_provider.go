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

func init() {
	provider := tokenAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
}
