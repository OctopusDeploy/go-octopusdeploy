package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// AccountService handles communication with Account-related methods of the
// Octopus API.
type AccountService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

// NewAccountService returns an AccountService with a preconfigured client.
func NewAccountService(sling *sling.Sling, uriTemplate string) *AccountService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &AccountService{
		name:  "AccountService",
		path:  path,
		sling: sling,
	}
}

// Get returns an Account that matches the input ID.
func (s *AccountService) Get(id string) (*model.Account, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Account), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}

// GetAll returns all instances of an Account.
func (s *AccountService) GetAll() ([]model.Account, error) {
	err := s.validateInternalState()

	items := new([]model.Account)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
}

// GetByName performs a lookup and returns the Account with a matching name.
func (s *AccountService) GetByName(name string) (*model.Account, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("GetByName", "name")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, createItemNotFoundError(s.name, "GetByName", name)
}

// GetUsage returns all projects and deployments which are using an Account.
func (s *AccountService) GetUsage(account model.Account) (*model.AccountUsage, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := account.Links["Usages"]
	resp, err := apiGet(s.sling, new(model.AccountUsage), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.AccountUsage), nil
}

// Add creates a new Account.
func (s *AccountService) Add(account *model.Account) (*model.Account, error) {
	if account == nil {
		return nil, createInvalidParameterError("Add", "account")
	}

	err := account.Validate()

	if err != nil {
		return nil, createValidationFailureError("Add", err)
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, account, new(model.Account), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}

// Delete removes the Account that matches the input ID.
func (s *AccountService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update modifies an Account based on the one provided as input.
func (s *AccountService) Update(account model.Account) (*model.Account, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = account.Validate()

	if err != nil {
		return nil, createValidationFailureError("Update", err)
	}

	path := fmt.Sprintf(s.path+"/%s", account.ID)
	resp, err := apiUpdate(s.sling, account, new(model.Account), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}

func (s *AccountService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &AccountService{}
