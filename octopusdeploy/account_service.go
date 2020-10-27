package octopusdeploy

import (
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
	"github.com/jinzhu/copier"
)

// accountService handles communication with account-related methods of the
// Octopus API.
type accountService struct {
	canDeleteService
}

// newAccountService returns an account service with a preconfigured client.
func newAccountService(sling *sling.Sling, uriTemplate string) *accountService {
	accountService := &accountService{}
	accountService.service = newService(ServiceAccountService, sling, uriTemplate)

	return accountService
}

func toAccount(accountResource *AccountResource) (IAccount, error) {
	if isNil(accountResource) {
		return nil, createInvalidParameterError("toAccount", ParameterAccountResource)
	}

	var account IAccount
	var err error
	switch accountResource.GetAccountType() {
	case AccountTypeAmazonWebServicesAccount:
		account, err = NewAmazonWebServicesAccount(accountResource.GetName(), accountResource.AccessKey, accountResource.SecretKey)
		if err != nil {
			return nil, err
		}
	case AccountTypeAzureServicePrincipal:
		account, err = NewAzureServicePrincipalAccount(accountResource.GetName(), *accountResource.SubscriptionID, *accountResource.TenantID, *accountResource.ApplicationID, accountResource.ApplicationPassword)
		if err != nil {
			return nil, err
		}
	case AccountTypeAzureSubscription:
		account, err = NewAzureSubscriptionAccount(accountResource.GetName(), *accountResource.SubscriptionID)
		if err != nil {
			return nil, err
		}
	case AccountTypeSSHKeyPair:
		account, err = NewSSHKeyAccount(accountResource.GetName(), accountResource.Username, accountResource.PrivateKeyFile)
		if err != nil {
			return nil, err
		}
	case AccountTypeToken:
		account, err = NewTokenAccount(accountResource.GetName(), accountResource.Token)
		if err != nil {
			return nil, err
		}
	case AccountTypeUsernamePassword:
		account, err = NewUsernamePasswordAccount(accountResource.GetName())
		if err != nil {
			return nil, err
		}
	}

	err = copier.Copy(account, accountResource)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func toAccounts(accountResources *AccountResources) *Accounts {
	return &Accounts{
		Items:        toAccountArray(accountResources.Items),
		PagedResults: accountResources.PagedResults,
	}
}

func toAccountResource(account IAccount) (*AccountResource, error) {
	if isNil(account) {
		return nil, createInvalidParameterError("toAccountResource", ParameterAccount)
	}

	accountResource := NewAccountResource(account.GetName(), account.GetAccountType())

	err := copier.Copy(&accountResource, account)
	if err != nil {
		return nil, err
	}

	return accountResource, nil
}

func toAccountArray(accountResources []*AccountResource) []IAccount {
	items := []IAccount{}
	for _, accountResource := range accountResources {
		account, err := toAccount(accountResource)
		if err != nil {
			return nil
		}
		items = append(items, account)
	}
	return items
}

// Add creates a new account.
func (s *accountService) Add(account IAccount) (IAccount, error) {
	if account == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterAccount)
	}

	accountResource, err := toAccountResource(account)
	if err != nil {
		return nil, err
	}

	response, err := apiAdd(s.getClient(), accountResource, new(AccountResource), s.BasePath)
	if err != nil {
		return nil, err
	}

	return toAccount(response.(*AccountResource))
}

// Get returns a collection of accounts based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s accountService) Get(accountsQuery AccountsQuery) (*Accounts, error) {
	v, _ := query.Values(accountsQuery)
	path := s.BasePath
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := apiGet(s.getClient(), new(AccountResources), path)
	if err != nil {
		return &Accounts{}, err
	}

	return toAccounts(response.(*AccountResources)), nil
}

// GetAll returns all accounts. If none can be found or an error occurs, it
// returns an empty collection.
func (s *accountService) GetAll() ([]IAccount, error) {
	items := []*AccountResource{}
	path := s.BasePath + "/all"

	_, err := apiGet(s.getClient(), &items, path)
	return toAccountArray(items), err
}

// GetByID returns the account that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s accountService) GetByID(id string) (IAccount, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError(OperationGetByID, ParameterID)
	}

	path := s.BasePath + "/" + id
	resp, err := apiGet(s.getClient(), new(AccountResource), path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(IAccount), nil
}

// GetUsages lists the projects and deployments which are using an account.
func (s *accountService) GetUsages(account IAccount) (*AccountUsage, error) {
	path := account.GetLinks()[linkUsages]
	resp, err := apiGet(s.getClient(), new(AccountUsage), path)
	if err != nil {
		return nil, err
	}

	return resp.(*AccountUsage), nil
}

// Update modifies an account based on the one provided as input.
func (s *accountService) Update(account IAccount) (IAccount, error) {
	if account == nil {
		return nil, createInvalidParameterError(OperationUpdate, ParameterAccount)
	}

	path, err := getUpdatePath(s, account)
	if err != nil {
		return nil, err
	}

	accountResource, err := toAccountResource(account)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), accountResource, new(AccountResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(IAccount), nil
}
