package accounts

import (
	"encoding/json"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/infrastructure/accounts"
)

type AccountDeserializationContainer struct {
	Account accounts.IAccount
}

var accountTypeProviderCacher *accountTypeProviderCache

type IAccountTypeProvider interface {
	AccountType() accounts.AccountType
	Factory() accounts.IAccount
}

type AccountTypeProviderCacheAdder interface {
	Add(provider IAccountTypeProvider)
}

type accountTypeProviderCache struct {
	providers map[accounts.AccountType]IAccountTypeProvider
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

	provider := accountTypeProviderCacher.getProviderForDiscriminator(accounts.AccountType(v.AccountType))

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
			providers: make(map[accounts.AccountType]IAccountTypeProvider),
		}
	}

	return accountTypeProviderCacher
}

func (c accountTypeProviderCache) getProviderForDiscriminator(discriminator accounts.AccountType) IAccountTypeProvider {
	return c.providers[discriminator]
}

func (c accountTypeProviderCache) Add(provider IAccountTypeProvider) {
	c.providers[provider.AccountType()] = provider
}
