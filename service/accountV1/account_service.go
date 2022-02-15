package accountV1

import (
	"fmt"
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	service2 "github.com/OctopusDeploy/go-octopusdeploy/service"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources/infrastructure/accounts"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/service"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
)

const accountsV1BasePath = "accounts"

// accountService handles communication with accountV1-related methods of the
// Octopus API.
type accountService struct {
	client *service2.SpaceScopedClient
	service2.SpaceScopedService
	service2.GetsByIDer[IAccount]
	service2.ResourceQueryer[IAccount]
	service2.CanAddService[IAccount]
	service2.CanUpdateService[IAccount]
	service2.CanDeleteService[IAccount]
}

// NewAccountService returns an accountV1 service with a preconfigured client.
func NewAccountService(client *service2.SpaceScopedClient) *accountService {
	accountService := &accountService{
		SpaceScopedService: service2.NewSpaceScopedService(service2.ServiceAccountService, accountsV1BasePath, client),
	}

	return accountService
}

// Add creates a new accountV1.
func (s *accountService) Add(account IAccount) (IAccount, error) {
	if account == nil {
		return nil, internal.CreateInvalidParameterError(service2.OperationAdd, octopusdeploy.ParameterAccount)
	}

	accountResource, err := accounts.ToAccountResource(s.GetClient(), account)
	if err != nil {
		return nil, err
	}

	response, err := service2.ApiAdd(s.GetClient(), AccountResource)(accountResource, new(AccountResource))
	if err != nil {
		return nil, err
	}

	return ToAccount(response.(*AccountResource))
}

// Query returns a collection of accounts based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s accountService) Query(accountsQuery ...service2.AccountsQuery) (*accounts.Accounts, error) {
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

// GetByID returns the accountV1 that matches the input ID. If one is not found,
// it returns nil and an error.
func (s accountService) GetByID(id string) (service.IAccount, error) {
	resp, err := s.client.apiGetByID(new(AccountResource), id)
	if err != nil {
		return nil, err
	}

	return ToAccount(resp.(*AccountResource))
}

// GetUsages lists the projects and deployments which are using an accountV1.
func (s *accountService) GetUsages(account service.IAccount) (*AccountUsage, error) {
	path := account.GetLinks()[linkUsages]
	resp, err := s.client.apiGet(new(AccountUsage), path)
	if err != nil {
		return nil, err
	}

	return resp.(*AccountUsage), nil
}

// Update modifies an accountV1 based on the one provided as input.
func (s *accountService) Update(account service.IAccount) (service.IAccount, error) {
	if account == nil {
		return nil, octopusdeploy.createInvalidParameterError(service2.OperationUpdate, service.ParameterAccount)
	}

	accountResource, err := accounts.ToAccountResource(s.client, account)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.apiUpdate(accountResource, new(AccountResource))
	if err != nil {
		return nil, err
	}

	return ToAccount(resp.(*AccountResource))
}
