package client

import (
	"errors"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/go-playground/validator"
)

// AccountService handles communication with Account-related methods of the
// Octopus API.
type AccountService struct {
	sling *sling.Sling `validate:"required"`
	path  string       `validate:"required"`
}

// NewAccountService returns an AccountService with a preconfigured client.
func NewAccountService(sling *sling.Sling) *AccountService {
	if sling == nil {
		return nil
	}

	return &AccountService{
		sling: sling,
		path:  "accounts",
	}
}

// Get returns an Account that matches the input ID.
func (s *AccountService) Get(id string) (*model.Account, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("AccountService", "id")
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

	accounts := new([]model.Account)

	if err != nil {
		return *accounts, err
	}

	_, err = apiGet(s.sling, accounts, s.path+"/all")

	return *accounts, err
}

// GetByName performs a lookup and returns the Account with a matching name.
func (s *AccountService) GetByName(name string) (*model.Account, error) {
	if isEmpty(name) {
		return nil, createInvalidParameterError("AccountService", "accountName")
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

	return nil, errors.New("AccountService: item not found")
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

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = account.Validate()

	if err != nil {
		return nil, createValidationFailureError("Add", err)
	}

	resp, err := apiAdd(s.sling, account, new(model.Account), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}

// Delete removes the Account that matches the input ID.
func (s *AccountService) Delete(accountID string) error {
	if isEmpty(accountID) {
		return createInvalidParameterError("Delete", "accountID")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", accountID))
}

// Update modifies an Account based on the one provided as input.
func (s *AccountService) Update(account model.Account) (*model.Account, error) {
	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	err = account.Validate()

	if err != nil {
		return nil, createValidationFailureError("AccountService", err)
	}

	path := fmt.Sprintf(s.path+"/%s", account.ID)
	resp, err := apiUpdate(s.sling, account, new(model.Account), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}

func (s *AccountService) validateInternalState() error {
	validate := validator.New()
	err := validate.Struct(s)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		return err
	}

	return nil
}

var _ ServiceInterface = &AccountService{}
