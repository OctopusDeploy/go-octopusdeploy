package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type UserService struct {
	sling *sling.Sling
	path  string
}

func NewUserService(sling *sling.Sling) *UserService {
	return &UserService{
		sling: sling,
		path:  "users",
	}
}

func (s *UserService) Get(id string) (*model.User, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", id)
	resp, err := apiGet(s.sling, new(model.User), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

func (s *UserService) GetMe() (*model.User, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new(model.User), s.path+"/me")

	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

func (s *UserService) GetAll() (*[]model.User, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new([]model.User), s.path+"/all")

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.User), nil
}

func (s *UserService) GetAuthentication() (*model.UserAuthentication, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.sling, new(model.UserAuthentication), s.path+"/authentication")

	if err != nil {
		return nil, err
	}

	return resp.(*model.UserAuthentication), nil
}

func (s *UserService) GetAuthenticationForUser(user model.User) (*model.UserAuthentication, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/authentication/%s", user.ID)
	resp, err := apiGet(s.sling, new(model.UserAuthentication), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.UserAuthentication), nil
}

func (s *UserService) GetSpaces(user model.User) (*[]model.Spaces, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s/spaces", user.ID)
	resp, err := apiGet(s.sling, new([]model.Spaces), path)

	if err != nil {
		return nil, err
	}

	return resp.(*[]model.Spaces), nil
}

func (s *UserService) GetByName(name string) (*model.User, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	collection, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, item := range *collection {
		if item.Username == name {
			return &item, nil
		}
	}

	return nil, errors.New("client: item not found")
}

func (s *UserService) Add(resource *model.User) (*model.User, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, resource, new(model.User), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

func (s *UserService) Delete(id string) error {
	err := s.validateInternalState()
	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *UserService) Update(resource *model.User) (*model.User, error) {
	err := s.validateInternalState()
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", resource.ID)
	resp, err := apiUpdate(s.sling, resource, new(model.User), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

func (s *UserService) validateInternalState() error {
	if s.sling == nil {
		return fmt.Errorf("UserService: the internal client is nil")
	}

	if len(strings.Trim(s.path, " ")) == 0 {
		return errors.New("UserService: the internal path is not set")
	}

	return nil
}
