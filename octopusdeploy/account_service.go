package octopusdeploy

import (
	"github.com/dghubble/sling"
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

// Add creates a new account.
func (s *accountService) Add(account IAccount) (IAccount, error) {
	if account == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterAccount)
	}

	accountResource, err := ToAccountResource(account)
	if err != nil {
		return nil, err
	}

	response, err := apiAdd(s.getClient(), accountResource, new(AccountResource), s.BasePath)
	if err != nil {
		return nil, err
	}

	return ToAccount(response.(*AccountResource))
}

// Get returns a collection of accounts based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s accountService) Get(accountsQuery ...AccountsQuery) (*Accounts, error) {
	values := make(map[string]interface{})
	path, err := s.URITemplate.Expand(values)
	if err != nil {
		return &Accounts{}, err
	}

	if accountsQuery != nil {
		path, err = s.URITemplate.Expand(accountsQuery[0])
		if err != nil {
			return &Accounts{}, err
		}
	}

	response, err := apiGet(s.getClient(), new(AccountResources), path)
	if err != nil {
		return &Accounts{}, err
	}

	return ToAccounts(response.(*AccountResources)), nil
}

// GetAll returns all accounts. If none can be found or an error occurs, it
// returns an empty collection.
func (s *accountService) GetAll() ([]IAccount, error) {
	items := []*AccountResource{}
	path := s.BasePath + "/all"

	_, err := apiGet(s.getClient(), &items, path)
	return ToAccountArray(items), err
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
		return nil, err
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

	accountResource, err := ToAccountResource(account)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), accountResource, new(AccountResource), path)
	if err != nil {
		return nil, err
	}

	return resp.(IAccount), nil
}
