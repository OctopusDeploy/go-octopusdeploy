package services

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/features/infrastructure/accounts/resources"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/infrastructure/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

const accountsV1BasePath = "accounts"

// accountService handles communication with account-related methods of the
// Octopus API.
type accountService struct {
	client *services.SpaceScopedClient
	services.SpaceScopedService
	services.GetsByIDer[resources.IAccount]
	services.ResourceQueryer[resources.IAccount]
	services.CanAddService[resources.IAccount]
	services.CanUpdateService[resources.IAccount]
	services.CanDeleteService[resources.IAccount]
}

// NewAccountService returns an account service with a preconfigured client.
func NewAccountService(client *services.SpaceScopedClient) *accountService {
	accountService := &accountService{
		SpaceScopedService: services.NewSpaceScopedService(services.ServiceAccountService, accountsV1BasePath, client),
	}

	return accountService
}

// Add creates a new account.
func (s *accountService) Add(account resources.IAccount) (resources.IAccount, error) {
	if account == nil {
		return nil, octopusdeploy.CreateInvalidParameterError(services.OperationAdd, octopusdeploy.ParameterAccount)
	}

	accountResource, err := accounts.ToAccountResource(s.GetClient(), account)
	if err != nil {
		return nil, err
	}

	response, err := services.ApiAdd(s.GetClient(), resources.AccountResource)(accountResource, new(resources.AccountResource))
	if err != nil {
		return nil, err
	}

	return resources.ToAccount(response.(*resources.AccountResource))
}

// Query returns a collection of accounts based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s accountService) Query(accountsQuery ...services.AccountsQuery) (*accounts.Accounts, error) {
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
func (s accountService) GetByID(id string) (services.IAccount, error) {
	resp, err := s.client.apiGetByID(new(resources.AccountResource), id)
	if err != nil {
		return nil, err
	}

	return resources.ToAccount(resp.(*resources.AccountResource))
}

// GetUsages lists the projects and deployments which are using an account.
func (s *accountService) GetUsages(account services.IAccount) (*resources.AccountUsage, error) {
	path := account.GetLinks()[linkUsages]
	resp, err := s.client.apiGet(new(resources.AccountUsage), path)
	if err != nil {
		return nil, err
	}

	return resp.(*resources.AccountUsage), nil
}

// Update modifies an account based on the one provided as input.
func (s *accountService) Update(account services.IAccount) (services.IAccount, error) {
	if account == nil {
		return nil, octopusdeploy.createInvalidParameterError(services.OperationUpdate, services.ParameterAccount)
	}

	accountResource, err := accounts.ToAccountResource(s.client, account)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.apiUpdate(accountResource, new(resources.AccountResource))
	if err != nil {
		return nil, err
	}

	return resources.ToAccount(resp.(*resources.AccountResource))
}
