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
	return &AccountService{
		sling: sling,
		path:  "accounts",
	}
}

// Get returns an Account that matches the input ID.
func (s *AccountService) Get(id string) (*model.Account, error) {
	if len(strings.Trim(id, " ")) == 0 {
		return nil, errors.New("client: invalid ID")
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
	var p []model.Account
	path := s.path
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(model.Accounts), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*model.Accounts)
		p = append(p, r.Items...)
		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

// GetByName returns an Account that matches the input name.
func (s *AccountService) GetByName(name string) (*model.Account, error) {
	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range *collection {
		if item.Name == name {
			return &item, nil
		}
	}

	return nil, errors.New("client: item not found")
}

// Add creates a new Account.
func (s *AccountService) Add(resource *model.Account) (*model.Account, error) {
	resp, err := apiAdd(s.sling, resource, new(model.Account), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}

// Delete removes the Account that matches the input ID.
func (s *AccountService) Delete(id string) error {
	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

// Update modifies an Account based on the one provided as input.
func (s *AccountService) Update(resource *model.Account) (*model.Account, error) {
	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.Account), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.Account), nil
}
