package client

import (
	"fmt"
	"strings"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/OctopusDeploy/go-octopusdeploy/uritemplates"
	"github.com/dghubble/sling"
)

type userService struct {
	name        string                    `validate:"required"`
	path        string                    `validate:"required"`
	sling       *sling.Sling              `validate:"required"`
	uriTemplate *uritemplates.UriTemplate `validate:"required"`
}

func newUserService(sling *sling.Sling, uriTemplate string) *userService {
	if sling == nil {
		sling = getDefaultClient()
	}

	template, err := uritemplates.Parse(strings.TrimSpace(uriTemplate))
	if err != nil {
		return nil
	}

	return &userService{
		name:        serviceUserService,
		path:        strings.TrimSpace(uriTemplate),
		sling:       sling,
		uriTemplate: template,
	}
}

func (s userService) getClient() *sling.Sling {
	return s.sling
}

func (s userService) getName() string {
	return s.name
}

func (s userService) getURITemplate() *uritemplates.UriTemplate {
	return s.uriTemplate
}

// Add creates a new user.
func (s userService) Add(resource *model.User) (*model.User, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, new(model.User), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

// DeleteByID deletes the user that matches the input ID.
func (s userService) DeleteByID(id string) error {
	err := deleteByID(s, id)
	if err == ErrItemNotFound {
		return createResourceNotFoundError("user", "ID", id)
	}

	return err
}

// GetByID returns the user that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s userService) GetByID(id string) (*model.User, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(model.User), path)
	if err != nil {
		return nil, createResourceNotFoundError("user", "ID", id)
	}

	return resp.(*model.User), nil
}

func (s userService) GetMe() (*model.User, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)
	path = path + "/me"

	resp, err := apiGet(s.getClient(), new(model.User), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

// GetAll returns all tenants. If none can be found or an error occurs, it
// returns an empty collection.
func (s userService) GetAll() ([]model.User, error) {
	items := []model.User{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

func (s userService) GetAuthentication() (*model.UserAuthentication, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)
	path = path + "/authentication"

	resp, err := apiGet(s.getClient(), new(model.UserAuthentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.UserAuthentication), nil
}

func (s userService) GetAuthenticationForUser(user *model.User) (*model.UserAuthentication, error) {
	if user == nil {
		return nil, createInvalidParameterError("GetAuthenticationForUser", "user")
	}

	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)
	path = fmt.Sprintf(path+"/authentication/%s", user.ID)

	resp, err := apiGet(s.getClient(), new(model.UserAuthentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.UserAuthentication), nil
}

func (s userService) GetSpaces(user *model.User) (*[]model.Spaces, error) {
	if user == nil {
		return nil, createInvalidParameterError("GetSpaces", "user")
	}

	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.path)
	path = fmt.Sprintf(path+"/%s/spaces", user.ID)

	resp, err := apiGet(s.getClient(), new([]model.Spaces), path)
	if err != nil {
		return nil, err
	}

	return resp.(*[]model.Spaces), nil
}

// Update modifies a user based on the one provided as input.
func (s userService) Update(resource model.User) (*model.User, error) {
	path, err := getUpdatePath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(model.User), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

var _ ServiceInterface = &userService{}
