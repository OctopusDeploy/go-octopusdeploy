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
	resources := []IAccount{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(Accounts), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*Accounts)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
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

	resp, err := apiGet(s.getClient(), new(Accounts), path)
	if err != nil {
		return &Accounts{}, err
	}

	return resp.(*Accounts), nil
}

// GetAll returns all accounts. If none can be found or an error occurs, it
// returns an empty collection.
func (s *accountService) GetAll() ([]IAccount, error) {
	items := []*Account{}
	path := s.BasePath + "/all"

	_, err := apiGet(s.getClient(), &items, path)
	return toAccountArray(items), err
}

// GetByID returns the account that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s accountService) GetByID(id string) (IAccount, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError(operationGetByID, parameterID)
	}

	path := s.BasePath + "/" + id
	resp, err := apiGet(s.getClient(), new(Account), path)
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
