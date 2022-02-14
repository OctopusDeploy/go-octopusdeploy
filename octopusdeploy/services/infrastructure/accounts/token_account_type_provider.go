package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/infrastructure/accounts"
)

type tokenAccountProvider struct {
	IAccountTypeProvider
}

func (p tokenAccountProvider) AccountType() accounts.AccountType {
	return accounts.AccountType(accounts.TokenAccountType)
}

func (p tokenAccountProvider) Factory() accounts.IAccount {
	return new(accounts.TokenAccount)
}

func doTokenAccountTypeRegistration() *tokenAccountProvider {
	provider := tokenAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
	return &provider
}

var _ = doTokenAccountTypeRegistration()
var _ IAccountTypeProvider = new(tokenAccountProvider)
