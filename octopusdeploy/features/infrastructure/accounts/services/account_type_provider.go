package services

import (
	"encoding/json"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/features/infrastructure/accounts/resources"
)

type AccountDeserializationContainer struct {
	Account resources.IAccount
}

var accountTypeProviderCacher *accountTypeProviderCache

type IAccountTypeProvider interface {
	AccountType() resources.AccountType
	Factory() resources.IAccount
}

type AccountTypeProviderCacheAdder interface {
	Add(provider IAccountTypeProvider)
}

type accountTypeProviderCache struct {
	providers map[resources.AccountType]IAccountTypeProvider
	AccountTypeProviderCacheAdder
}

func (c *AccountDeserializationContainer) UnmarshalJSON(js []byte) error {
	v := struct {
		AccountType string
	}{}

	err := json.Unmarshal(js, &v)
	if err != nil {
		return err
	}

	provider := accountTypeProviderCacher.getProviderForDiscriminator(resources.AccountType(v.AccountType))

	a := provider.Factory()
	err = json.Unmarshal(js, a)
	if err != nil {
		return err
	}
	c.Account = a

	return nil
}

func GetAccountTypeCacheAdder() AccountTypeProviderCacheAdder {
	if accountTypeProviderCacher == nil {
		accountTypeProviderCacher = &accountTypeProviderCache{
			providers: make(map[resources.AccountType]IAccountTypeProvider),
		}
	}

	return accountTypeProviderCacher
}

func (c accountTypeProviderCache) getProviderForDiscriminator(discriminator resources.AccountType) IAccountTypeProvider {
	return c.providers[discriminator]
}

func (c accountTypeProviderCache) Add(provider IAccountTypeProvider) {
	c.providers[provider.AccountType()] = provider
}
