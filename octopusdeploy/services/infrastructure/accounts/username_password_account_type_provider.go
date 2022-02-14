package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/infrastructure/accounts"
)

type usernamePasswordAccountProvider struct {
}

func (p usernamePasswordAccountProvider) AccountType() accounts.AccountType {
	return accounts.AccountType(accounts.UsernamePasswordAccountType)
}

func (p usernamePasswordAccountProvider) Factory() accounts.IAccount {
	return new(accounts.UsernamePasswordAccount)
}

func doUsernamePasswordAccountTypeRegistration() *usernamePasswordAccountProvider {
	provider := &usernamePasswordAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
	return provider
}

var _ = doUsernamePasswordAccountTypeRegistration()
var _ IAccountTypeProvider = new(usernamePasswordAccountProvider)
