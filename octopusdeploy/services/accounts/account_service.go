package accounts

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/access_management"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

const accountsV1BasePath = "accounts"

// accountService handles communication with account-related methods of the
// Octopus API.
type accountService struct {
	client *octopusdeploy.SpaceScopedClient
	services.SpaceScopedService
	services.GetsByIDer[octopusdeploy.IAccount]
	services.ResourceQueryer[accounts.Accounts]
	services.CanAddService[octopusdeploy.IAccount]
	services.CanUpdateService[octopusdeploy.IAccount]
	services.CanDeleteService[octopusdeploy.IAccount]
}

// NewAccountService returns an account service with a preconfigured client.
func NewAccountService(client *octopusdeploy.SpaceScopedClient) *accountService {
	accountService := &accountService{
		SpaceScopedService: services.NewSpaceScopedService(octopusdeploy.ServiceAccountService, accountsV1BasePath, client),
	}

	return accountService
}

// Add creates a new account.
func (s *accountService) Add(account octopusdeploy.IAccount) (octopusdeploy.IAccount, error) {
	if account == nil {
		return nil, octopusdeploy.CreateInvalidParameterError(octopusdeploy.OperationAdd, octopusdeploy.ParameterAccount)
	}

	accountResource, err := accounts.ToAccountResource(s.GetClient(), account)
	if err != nil {
		return nil, err
	}

	response, err := octopusdeploy.ApiAdd(s.GetClient(), accounts.AccountResource)(accountResource, new(accounts.AccountResource))
	if err != nil {
		return nil, err
	}

	return accounts.ToAccount(response.(*accounts.AccountResource))
}

// Query returns a collection of accounts based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s accountService) Query(accountsQuery ...octopusdeploy.AccountsQuery) (*accounts.Accounts, error) {
	template := uritemplates.Parse(fmt.Sprintf("%s{/id}{?skip,take,ids,partialName,accountType}", s.BasePath))

	values := make(map[string]interface{})
	path, err := s.uriTemplate.Expand(values)
	if err != nil {
		return &accounts.Accounts{}, err
	}

	if accountsQuery != nil {
		path, err = s.uriTemplate.Expand(accountsQuery[0])
		if err != nil {
			return &accounts.Accounts{}, err
		}
	}

	response, err := s.client.apiQuery(new(accounts.AccountResources), path)
	if err != nil {
		return &accounts.Accounts{}, err
	}

	return accounts.ToAccounts(response.(*accounts.AccountResources)), nil
}

// GetByID returns the account that matches the input ID. If one is not found,
// it returns nil and an error.
func (s accountService) GetByID(id string) (octopusdeploy.IAccount, error) {
	resp, err := s.client.apiGetByID(new(accounts.AccountResource), id)
	if err != nil {
		return nil, err
	}

	return accounts.ToAccount(resp.(*accounts.AccountResource))
}

// GetUsages lists the projects and deployments which are using an account.
func (s *accountService) GetUsages(account octopusdeploy.IAccount) (*accounts.AccountUsage, error) {
	path := account.GetLinks()[linkUsages]
	resp, err := s.client.apiGet(new(accounts.AccountUsage), path)
	if err != nil {
		return nil, err
	}

	return resp.(*accounts.AccountUsage), nil
}

// Update modifies an account based on the one provided as input.
func (s *accountService) Update(account octopusdeploy.IAccount) (octopusdeploy.IAccount, error) {
	if account == nil {
		return nil, octopusdeploy.createInvalidParameterError(octopusdeploy.OperationUpdate, octopusdeploy.ParameterAccount)
	}

	accountResource, err := accounts.ToAccountResource(s.client, account)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.apiUpdate(accountResource, new(accounts.AccountResource))
	if err != nil {
		return nil, err
	}

	return accounts.ToAccount(resp.(*accounts.AccountResource))
}
