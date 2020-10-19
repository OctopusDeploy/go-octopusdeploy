package client

import (
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/jinzhu/copier"
)

// accountService handles communication with account-related methods of the
// Octopus API.
type accountService struct {
	service
}

// newAccountService returns an account service with a preconfigured client.
func newAccountService(sling *sling.Sling, uriTemplate string) *accountService {
	accountService := &accountService{}
	accountService.service = newService(serviceAccountService, sling, uriTemplate, new(model.Account))

	return accountService
}

func toAccountResource(account model.IAccount) model.IAccount {
	var accountResource model.IAccount
	switch account.GetAccountType() {
	case "AmazonWebServicesAccount":
		accountResource = new(model.AmazonWebServicesAccount)
	case "AzureServicePrincipal":
		accountResource = new(model.AzureServicePrincipalAccount)
	case "AzureSubscription":
		accountResource = new(model.AzureSubscriptionAccount)
	case "SshKeyPair":
		accountResource = new(model.SSHKeyAccount)
	case "Token":
		accountResource = new(model.TokenAccount)
	case "UsernamePassword":
		accountResource = new(model.UsernamePasswordAccount)
	}

	copier.Copy(accountResource, account)
	return accountResource
}

func toAccountArray(accounts []*model.Account) []model.IAccount {
	items := []model.IAccount{}
	for _, account := range accounts {
		items = append(items, toAccountResource(account))
	}
	return items
}

func (s *accountService) getPagedResponse(path string) ([]model.IAccount, error) {
	resources := []*model.Account{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Accounts), path)
		if err != nil {
			return toAccountArray(resources), err
		}

		responseList := resp.(*model.Accounts)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return toAccountArray(resources), nil
}

// Add creates a new account.
func (s *accountService) Add(resource model.IAccount) (model.IAccount, error) {
	if resource == nil {
		return nil, createInvalidParameterError(operationAdd, "resource")
	}

	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	var account interface{}
	switch resource.GetAccountType() {
	case "AmazonWebServicesAccount":
		account = new(model.AmazonWebServicesAccount)
	case "AzureServicePrincipal":
		account = new(model.AzureServicePrincipalAccount)
	case "AzureSubscription":
		account = new(model.AzureSubscriptionAccount)
	case "SshKeyPair":
		account = new(model.SSHKeyAccount)
	case "Token":
		account = new(model.TokenAccount)
	case "UsernamePassword":
		account = new(model.UsernamePasswordAccount)
	}

	resp, err := apiAdd(s.getClient(), resource, account, path)
	if err != nil {
		return nil, err
	}

	return resp.(model.IAccount), nil
}

// GetAll returns all accounts. If none can be found or an error occurs, it
// returns an empty collection.
func (s *accountService) GetAll() ([]model.IAccount, error) {
	items := []*model.Account{}
	path, err := getAllPath(s)
	if err != nil {
		return toAccountArray(items), err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return toAccountArray(items), err
}

// GetByID returns the account that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s accountService) GetByID(id string) (model.IAccount, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(model.IAccount), nil
}

// GetByIDs returns the accounts that match the input IDs.
func (s *accountService) GetByIDs(ids []string) ([]model.IAccount, error) {
	if len(ids) == 0 {
		return []model.IAccount{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []model.IAccount{}, err
	}

	return s.getPagedResponse(path)
}

// GetByAccountType performs a lookup and returns the accounts with a matching
// account type.
func (s *accountService) GetByAccountType(accountType string) ([]model.IAccount, error) {
	path, err := getByAccountTypePath(s, accountType)
	if err != nil {
		return []model.IAccount{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName performs a lookup and returns a single instance of an account with a matching name.
func (s *accountService) GetByName(name string) (model.IAccount, error) {
	accounts, err := s.GetByPartialName(name)
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		if account.GetName() == name {
			return toAccountResource(account), nil
		}
	}

	return nil, nil
}

// GetByPartialName performs a lookup and returns instances of an account with
// a matching partial name.
func (s *accountService) GetByPartialName(name string) ([]model.IAccount, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []model.IAccount{}, err
	}

	return s.getPagedResponse(path)
}

// GetUsages lists the projects and deployments which are using an account.
func (s *accountService) GetUsages(account model.IAccount) (*model.AccountUsage, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := account.GetLinks()[linkUsages]

	resp, err := apiGet(s.getClient(), new(model.AccountUsage), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.AccountUsage), nil
}

// Update modifies an account based on the one provided as input.
func (s *accountService) Update(account model.IAccount) (model.IAccount, error) {
	if account == nil {
		return nil, createInvalidParameterError(operationUpdate, parameterAccount)
	}

	path, err := getUpdatePath(s, account)
	if err != nil {
		return nil, err
	}

	var resourceAccount interface{}
	switch account.GetAccountType() {
	case "AmazonWebServicesAccount":
		resourceAccount = new(model.AmazonWebServicesAccount)
	case "AzureServicePrincipal":
		resourceAccount = new(model.AzureServicePrincipalAccount)
	case "AzureSubscription":
		resourceAccount = new(model.AzureSubscriptionAccount)
	case "SshKeyPair":
		resourceAccount = new(model.SSHKeyAccount)
	case "Token":
		resourceAccount = new(model.TokenAccount)
	case "UsernamePassword":
		resourceAccount = new(model.UsernamePasswordAccount)
	}

	resp, err := apiUpdate(s.getClient(), account, resourceAccount, path)
	if err != nil {
		return nil, err
	}

	return resp.(model.IAccount), nil
}
