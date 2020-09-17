package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type UserService struct {
	name  string       `validate:"required"`
	path  string       `validate:"required"`
	sling *sling.Sling `validate:"required"`
}

func NewUserService(sling *sling.Sling, uriTemplate string) *UserService {
	if sling == nil {
		return nil
	}

	path := strings.Split(uriTemplate, "{")[0]

	return &UserService{
		name:  "UserService",
		path:  path,
		sling: sling,
	}
}

func (s *UserService) Get(id string) (*model.User, error) {
	if isEmpty(id) {
		return nil, createInvalidParameterError("Get", "id")
	}

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

// GetAll returns all instances of a User.
func (s *UserService) GetAll() ([]model.User, error) {
	err := s.validateInternalState()

	items := new([]model.User)

	if err != nil {
		return *items, err
	}

	_, err = apiGet(s.sling, items, s.path+"/all")

	return *items, err
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

func (s *UserService) GetAuthenticationForUser(user *model.User) (*model.UserAuthentication, error) {
	if user == nil {
		return nil, createInvalidParameterError("GetAuthenticationForUser", "user")
	}

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

func (s *UserService) GetSpaces(user *model.User) (*[]model.Spaces, error) {
	if user == nil {
		return nil, createInvalidParameterError("GetSpaces", "user")
	}

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

// GetByName performs a lookup and returns the User with a matching name.
func (s *UserService) GetByName(name string) (*model.User, error) {
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
		if item.Username == name {
			return &item, nil
		}
	}

	return nil, createItemNotFoundError(s.name, "GetByName", name)
}

// Add creates a new User.
func (s *UserService) Add(user *model.User) (*model.User, error) {
	if user == nil {
		return nil, createInvalidParameterError("Add", "user")
	}

	err := user.Validate()

	if err != nil {
		return nil, err
	}

	err = s.validateInternalState()

	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.sling, user, new(model.User), s.path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

func (s *UserService) Delete(id string) error {
	if isEmpty(id) {
		return createInvalidParameterError("Delete", "id")
	}

	err := s.validateInternalState()

	if err != nil {
		return err
	}

	return apiDelete(s.sling, fmt.Sprintf(s.path+"/%s", id))
}

func (s *UserService) Update(user *model.User) (*model.User, error) {
	if user == nil {
		return nil, createInvalidParameterError("Update", "user")
	}

	err := s.validateInternalState()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf(s.path+"/%s", user.ID)
	resp, err := apiUpdate(s.sling, user, new(model.User), path)

	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

func (s *UserService) validateInternalState() error {
	if s.sling == nil {
		return createInvalidClientStateError(s.name)
	}

	if isEmpty(s.path) {
		return createInvalidPathError(s.name)
	}

	return nil
}

var _ ServiceInterface = &UserService{}
