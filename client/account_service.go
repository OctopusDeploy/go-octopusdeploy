package client

import (
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

// accountService handles communication with Account-related methods of the Octopus API.
type accountService struct {
	name        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

// newAccountService returns an accountService with a preconfigured client.
func newAccountService(sling *sling.Sling, uriTemplate string) *accountService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &accountService{
		name:        serviceAccountService,
		sling:       sling,
		uriTemplate: template,
	}
}

func (s accountService) getClient() *sling.Sling {
	return s.sling
}

func (s accountService) getName() string {
	return s.name
}

func (s accountService) getPagedResponse(path string) ([]model.Account, error) {
	resources := []model.Account{}
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.getClient(), new(model.Accounts), path)
		if err != nil {
			return resources, err
		}

		responseList := resp.(*model.Accounts)
		resources = append(resources, responseList.Items...)
		path, loadNextPage = LoadNextPage(responseList.PagedResults)
	}

	return resources, nil
}

func (s accountService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// Add creates a new account.
func (s accountService) Add(resource *model.Account) (*model.Account, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.Account), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}

// DeleteByID deletes the account that matches the input ID.
func (s accountService) DeleteByID(id string) error {
	return deleteByID(s, id)
}

// GetAll returns all accounts. If none can be found or an error occurs, it
// returns an empty collection.
func (s accountService) GetAll() ([]model.Account, error) {
	items := []model.Account{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// GetByID returns the account that matches the input ID. If one cannot be
// found, it returns nil and an error.
func (s accountService) GetByID(id string) (*model.Account, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.Account), path)
	if err != nil {
		return nil, createResourceNotFoundError("account", "ID", id)
	}

	return resp.(*model.Account), nil
}

// GetByIDs returns the accounts that match the input IDs.
func (s accountService) GetByIDs(ids []string) ([]model.Account, error) {
	path, err := getByIDsPath(s, ids)
	if err != nil {
		return []model.Account{}, err
	}

	return s.getPagedResponse(path)
}

// GetByAccountType performs a lookup and returns the Accounts with a matching AccountType.
func (s accountService) GetByAccountType(accountType enum.AccountType) ([]model.Account, error) {
	path, err := getByAccountTypePath(s, accountType)
	if err != nil {
		return []model.Account{}, err
	}

	return s.getPagedResponse(path)
}

// GetByName performs a lookup and returns a single instance of an account with a matching name.
func (s accountService) GetByName(name string) (*model.Account, error) {
	resourceList, err := s.GetByPartialName(name)
	if err != nil {
		return nil, err
	}

	for _, resource := range resourceList {
		if resource.Name == name {
			return &resource, nil
		}
	}

	return nil, nil
}

// GetByPartialName performs a lookup and returns instances of an Account with a matching partial name.
func (s accountService) GetByPartialName(name string) ([]model.Account, error) {
	path, err := getByPartialNamePath(s, name)
	if err != nil {
		return []model.Account{}, err
	}

	return s.getPagedResponse(path)
}

// GetUsages lists the projects and deployments which are using an account.
func (s accountService) GetUsages(account model.Account) (*model.AccountUsage, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := account.Links[linkUsages]

	resp, err := apiGet(s.getClient(), new(model.AccountUsage), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.AccountUsage), nil
}

// Update modifies an account based on the one provided as input.
func (s accountService) Update(resource model.Account) (*model.Account, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.Account), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}

var _ ServiceInterface = &accountService{}
