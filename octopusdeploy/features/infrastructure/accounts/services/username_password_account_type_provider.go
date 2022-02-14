package services

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/features/infrastructure/accounts/resources"
)

type usernamePasswordAccountProvider struct {
}

func (p usernamePasswordAccountProvider) AccountType() resources.AccountType {
	return resources.AccountType(resources.UsernamePasswordAccountType)
}

func (p usernamePasswordAccountProvider) Factory() resources.IAccount {
	return new(resources.UsernamePasswordAccount)
}

func doUsernamePasswordAccountTypeRegistration() *usernamePasswordAccountProvider {
	provider := &usernamePasswordAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
	return provider
}

var _ = doUsernamePasswordAccountTypeRegistration()
var _ IAccountTypeProvider = new(usernamePasswordAccountProvider)
