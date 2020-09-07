package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

// AccountService handles communication with Account-related methods of the
// Octopus API.
type AccountService struct {
	sling *sling.Sling
	path  string
}

// NewAccountService returns an AccountService with a preconfigured client.
func NewAccountService(sling *sling.Sling) *AccountService {
	if sling == nil {
		fmt.Println(fmt.Errorf("AccountService: input parameter (sling) is nil"))
		return nil
	}

	return &AccountService{
		sling: sling,
		path:  "accounts",
	}
}

// Get returns an Account that matches the input ID.
func (s *AccountService) Get(id string) (*model.Account, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(id, " ")) == 0 {
		return nil, errors.New("AccountService: invalid parameter, ID")
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.Account), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}

// GetAll returns all of the Accounts for a Space.
func (s *AccountService) GetAll() (*[]model.Account, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.Account), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.Account), nil
}

// GetByName returns an Account that matches the input name.
func (s *AccountService) GetByName(name string) (*model.Account, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if len(strings.Trim(name, " ")) == 0 {
		return nil, errors.New("AccountService: invalid parameter, name")
	}

	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range *collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, errors.New("AccountService: item not found")
}

// Add creates a new Account.
func (s *AccountService) Add(resource *model.Account) (*model.Account, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if resource == nil {
		return nil, errors.New("AccountService: invalid parameter, resource")
	}

	err = resource.Validate()
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, resource, new(model.Account), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}

// Delete removes the Account that matches the input ID.
func (s *AccountService) Delete(id string) error {
	err := s.validateInternalState()
	if err != nil {
		return nil
	}

	if len(strings.Trim(id, " ")) == 0 {
		return errors.New("AccountService: invalid parameter, ID")
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update modifies an Account based on the one provided as input.
func (s *AccountService) Update(resource *model.Account) (*model.Account, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	if resource == nil {
		return nil, errors.New("AccountService: invalid parameter, resource")
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Account), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}

func (s *AccountService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("AccountService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("AccountService: the internal path is not set")
	}

	return nil
}
