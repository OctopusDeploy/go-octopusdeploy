package octopusdeploy

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

const accountsV1BasePath = "accounts"

// accountService handles communication with account-related methods of the
// Octopus API.
type accountService struct {
	spaceScopedService
	canDeleteService
}

// NewAccountService returns an account service with a preconfigured client.
func NewAccountService(client SpaceScopedClient) *accountService {
	accountService := &accountService{
		spaceScopedService: newSpaceScopedService(ServiceAccountService, client),
	}

	return accountService
}

// Add creates a new account.
func (s *accountService) Add(account IAccount) (IAccount, error) {
	if account == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterAccount)
	}

	accountResource, err := ToAccountResource(s.getClient(), account)
	if err != nil {
		return nil, err
	}

	response, err := s.client.apiAdd(accountResource, new(AccountResource))
	if err != nil {
		return nil, err
	}

	return ToAccount(response.(*AccountResource))
}

// Query returns a collection of accounts based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s accountService) Query(accountsQuery ...AccountsQuery) (*Accounts, error) {
	template := uritemplates.Parse(fmt.Sprintf("%s{/id}{?skip,take,ids,partialName,accountType}", s.BasePath))

	values := make(map[string]interface{})
	path, err := s.uriTemplate.Expand(values)
	if err != nil {
		return &Accounts{}, err
	}

	if accountsQuery != nil {
		path, err = s.uriTemplate.Expand(accountsQuery[0])
		if err != nil {
			return &Accounts{}, err
		}
	}

	response, err := s.client.apiQuery(new(AccountResources), path)
	if err != nil {
		return &Accounts{}, err
	}

	return ToAccounts(response.(*AccountResources)), nil
}

// GetByID returns the account that matches the input ID. If one is not found,
// it returns nil and an error.
func (s accountService) GetByID(id string) (IAccount, error) {
	resp, err := s.client.apiGetByID(new(AccountResource), id)
	if err != nil {
		return nil, err
	}

	return ToAccount(resp.(*AccountResource))
}

// GetUsages lists the projects and deployments which are using an account.
func (s *accountService) GetUsages(account IAccount) (*AccountUsage, error) {
	path := account.GetLinks()[linkUsages]
	resp, err := s.client.apiGet(new(AccountUsage), path)
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

	accountResource, err := ToAccountResource(s.client, account)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.apiUpdate(accountResource, new(AccountResource))
	if err != nil {
		return nil, err
	}

	return ToAccount(resp.(*AccountResource))
}
