package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
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
//
// Deprecated: Use accounts.Add
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
//
// Deprecated: Use accounts.Get
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
//
// Deprecated: Use accounts.Get
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

// ----- new -----

const (
	template = "/api/{spaceId}/accounts{/id}{?skip,take,ids,partialName,accountType}"
)

// Get returns a collection of accounts based on the criteria defined by its
// input query parameter.
func Get(client newclient.Client, spaceID string, accountsQuery *AccountsQuery) (*Accounts, error) {
	res, err := newclient.GetByQuery[AccountResource](client, template, spaceID, accountsQuery)
	if err != nil {
		return nil, err
	}
	return ToAccounts(res), nil
}

// Add creates a new account.
func Add(client newclient.Client, account IAccount) (IAccount, error) {
	res, err := newclient.Add[AccountResource](client, template, account.GetSpaceID(), account)
	if err != nil {
		return nil, err
	}
	return ToAccount(res)
}

// GetByID returns the account that matches the input ID.
func GetByID(client newclient.Client, spaceID string, ID string) (IAccount, error) {

	res, err := newclient.GetByID[AccountResource](client, template, spaceID, ID)
	if err != nil {
		return nil, err
	}

	return ToAccount(res)
}

// Update modifies an account based on the one provided as input.
func Update(client newclient.Client, account IAccount) (IAccount, error) {
	accountResource, err := ToAccountResource(account)
	if err != nil {
		return nil, err
	}

	res, err := newclient.Add[AccountResource](client, template, account.GetSpaceID(), accountResource)
	if err != nil {
		return nil, err
	}

	return ToAccount(res)
}

// DeleteByID will delete a account with the provided id.
func DeleteByID(client newclient.Client, spaceID string, id string) error {
	return newclient.DeleteByID(client, template, spaceID, id)
}
