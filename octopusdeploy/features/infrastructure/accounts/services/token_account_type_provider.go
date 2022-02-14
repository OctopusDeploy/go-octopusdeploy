package services

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/features/infrastructure/accounts/resources"
)

type tokenAccountProvider struct {
	IAccountTypeProvider
}

func (p tokenAccountProvider) AccountType() resources.AccountType {
	return resources.AccountType(resources.TokenAccountType)
}

func (p tokenAccountProvider) Factory() resources.IAccount {
	return new(resources.TokenAccount)
}

func doTokenAccountTypeRegistration() *tokenAccountProvider {
	provider := tokenAccountProvider{}
	GetAccountTypeCacheAdder().Add(provider)
	return &provider
}

var _ = doTokenAccountTypeRegistration()
var _ IAccountTypeProvider = new(tokenAccountProvider)
