package octopusdeploy

import (
	"github.com/dghubble/sling"
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
	accountService.service = newService(serviceAccountService, sling, uriTemplate, new(Account))

	return accountService
}

func toAccountResource(account IAccount) IAccount {
	var accountResource IAccount
	switch account.GetAccountType() {
	case accountTypeAmazonWebServicesAccount:
		accountResource = new(AmazonWebServicesAccount)
	case accountTypeAzureServicePrincipal:
		accountResource = new(AzureServicePrincipalAccount)
	case accountTypeAzureSubscription:
		accountResource = new(AzureSubscriptionAccount)
	case accountTypeSshKeyPair:
		accountResource = new(SSHKeyAccount)
	case accountTypeToken:
		accountResource = new(TokenAccount)
	case accountTypeUsernamePassword:
		accountResource = new(UsernamePasswordAccount)
	}

	copier.Copy(accountResource, account)
	return accountResource
}

func toAccountArray(accounts []*Account) []IAccount {
	items := []IAccount{}
	for _, account := range accounts {
		items = append(items, toAccountResource(account))
	}
	return items
}

func (s *accountService) getPagedResponse(path string) ([]IAccount, error) {
	resources := []*Account{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Accounts), path)
		if err != nil {
			return toAccountArray(resources), err
		}

		responseList := resp.(*Accounts)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return toAccountArray(resources), nil
}

// Add creates a new account.
func (s *accountService) Add(resource IAccount) (IAccount, error) {
	if resource == nil {
		return nil, createInvalidParameterError(operationAdd, parameterResource)
	}

	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	var account interface{}
	switch resource.GetAccountType() {
	case accountTypeAmazonWebServicesAccount:
		account = new(AmazonWebServicesAccount)
	case accountTypeAzureServicePrincipal:
		account = new(AzureServicePrincipalAccount)
	case accountTypeAzureSubscription:
		account = new(AzureSubscriptionAccount)
	case accountTypeSshKeyPair:
		account = new(SSHKeyAccount)
	case accountTypeToken:
		account = new(TokenAccount)
	case accountTypeUsernamePassword:
		account = new(UsernamePasswordAccount)
	}

	resp, err := apiAdd(s.getClient(), resource, account, path)
	if err != nil {
		return nil, err
	}

	return resp.(IAccount), nil
}

// GetAll returns all accounts. If none can be found or an error occurs, it
// returns an empty collection.
func (s *accountService) GetAll() ([]IAccount, error) {
	items := []*Account{}
	path, err := getAllPath(s)
	if err != nil {
		return toAccountArray(items), err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return toAccountArray(items), err
}

// GetByID returns the account that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s accountService) GetByID(id string) (IAccount, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, createResourceNotFoundError(s.getName(), "ID", id)
	}

	return resp.(IAccount), nil
}

// GetByIDs returns the accounts that match the input IDs.
func (s *accountService) GetByIDs(ids []string) ([]IAccount, error) {
	if len(ids) == 0 {
		return []IAccount{}, nil
	}

	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []IAccount{}, err
	}

	return s.getPagedResponse(path)
}

// GetByAccountType performs a lookup and returns the accounts with a matching
// account type.
func (s *accountService) GetByAccountType(accountType string) ([]IAccount, error) {
	path, err := getByAccountTypePath(s, accountType)
	if err != nil {
		return []IAccount{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName performs a lookup and returns a single instance of an account with a matching name.
func (s *accountService) GetByName(name string) (IAccount, error) {
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
func (s *accountService) GetByPartialName(name string) ([]IAccount, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []IAccount{}, err
	}

	return s.getPagedResponse(path)
}

// GetUsages lists the projects and deployments which are using an account.
func (s *accountService) GetUsages(account IAccount) (*AccountUsage, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

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
		return nil, createInvalidParameterError(operationUpdate, parameterAccount)
	}

	path, err := getUpdatePath(s, account)
	if err != nil {
		return nil, err
	}

	var resourceAccount interface{}
	switch account.GetAccountType() {
	case accountTypeAmazonWebServicesAccount:
		resourceAccount = new(AmazonWebServicesAccount)
	case accountTypeAzureServicePrincipal:
		resourceAccount = new(AzureServicePrincipalAccount)
	case accountTypeAzureSubscription:
		resourceAccount = new(AzureSubscriptionAccount)
	case accountTypeSshKeyPair:
		resourceAccount = new(SSHKeyAccount)
	case accountTypeToken:
		resourceAccount = new(TokenAccount)
	case accountTypeUsernamePassword:
		resourceAccount = new(UsernamePasswordAccount)
	}

	resp, err := apiUpdate(s.getClient(), account, resourceAccount, path)
	if err != nil {
		return nil, err
	}

	return resp.(IAccount), nil
}
