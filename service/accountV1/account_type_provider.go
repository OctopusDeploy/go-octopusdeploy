package accountV1

import (
	"encoding/json"
)

type AccountDeserializationContainer struct {
	Account IAccount
}

var accountTypeProviderCacher *accountTypeProviderCache

type IAccountTypeProvider interface {
	AccountType() AccountType
	Factory() IAccount
}

type AccountTypeProviderCacheAdder interface {
	Add(provider IAccountTypeProvider)
}

type accountTypeProviderCache struct {
	providers map[AccountType]IAccountTypeProvider
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

	provider := accountTypeProviderCacher.getProviderForDiscriminator(AccountType(v.AccountType))

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
			providers: make(map[AccountType]IAccountTypeProvider),
		}
	}

	return accountTypeProviderCacher
}

func (c accountTypeProviderCache) getProviderForDiscriminator(discriminator AccountType) IAccountTypeProvider {
	return c.providers[discriminator]
}

func (c accountTypeProviderCache) Add(provider IAccountTypeProvider) {
	c.providers[provider.AccountType()] = provider
}
