package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services/api"
	"github.com/dghubble/sling"
)

// AccountService handles communication with the account endpoint.
type AccountService struct {
	services.CanDeleteService
}

// NewAccountService returns the service with a preconfigured client.
func NewAccountService(sling *sling.Sling, uriTemplate string) *AccountService {
	return &AccountService{
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceAccountService, sling, uriTemplate),
		},
	}
}

// Add creates a new account.
func (s *AccountService) Add(account IAccount) (IAccount, error) {
	if IsNil(account) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterAccount)
	}

	response, err := services.ApiAdd(s.GetClient(), account, account, s.GetBasePath())
	if err != nil {
		return nil, err
	}

	return response.(IAccount), nil
}

// Get returns a collection of accounts based on the criteria defined by its
// input query parameter. If an error occurs, an empty collection is returned
// along with the associated error.
func (s *AccountService) Get(accountsQuery ...AccountsQuery) (*Accounts, error) {
	values := make(map[string]interface{})
	path, err := s.GetURITemplate().Expand(values)
	if err != nil {
		return &Accounts{}, err
	}

	if accountsQuery != nil {
		path, err = s.GetURITemplate().Expand(accountsQuery[0])
		if err != nil {
			return &Accounts{}, err
		}
	}

	response, err := api.ApiGet(s.GetClient(), new(resources.Resources[*AccountResource]), path)
	if err != nil {
		return &Accounts{}, err
	}

	return ToAccounts(response.(*resources.Resources[*AccountResource])), nil
}

// GetAll returns all accounts. If none are found or an error occurs, it
// returns an empty collection.
func (s *AccountService) GetAll() ([]IAccount, error) {
	items := []*AccountResource{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return ToAccountArray(items), err
	}

	_, err = api.ApiGet(s.GetClient(), &items, path)
	return ToAccountArray(items), err
}

// GetByID returns the account that matches the input ID. If one is not found,
// it returns nil and an error.
func (s *AccountService) GetByID(id string) (IAccount, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := api.ApiGet(s.GetClient(), new(AccountResource), path)
	if err != nil {
		return nil, err
	}

	return ToAccount(resp.(*AccountResource))
}

// GetUsages lists the projects and deployments which are using an account.
func (s *AccountService) GetUsages(account IAccount) (*AccountUsage, error) {
	path := account.GetLinks()[constants.LinkUsages]
	resp, err := api.ApiGet(s.GetClient(), new(AccountUsage), path)
	if err != nil {
		return nil, err
	}

	return resp.(*AccountUsage), nil
}

// Update modifies an account based on the one provided as input.
func (s *AccountService) Update(account IAccount) (IAccount, error) {
	if account == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationUpdate, constants.ParameterAccount)
	}

	path, err := services.GetUpdatePath(s, account)
	if err != nil {
		return nil, err
	}

	accountResource, err := ToAccountResource(account)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), accountResource, new(AccountResource), path)
	if err != nil {
		return nil, err
	}

	return ToAccount(resp.(*AccountResource))
}
