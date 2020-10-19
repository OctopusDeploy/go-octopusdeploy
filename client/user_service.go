package client

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
)

type userService struct {
	apiKeysPath               string
	authenticateOctopusIDPath string
	currentUserPath           string
	externalUserSearchPath    string
	registerPath              string
	signInPath                string
	signOutPath               string
	userAuthenticationPath    string
	userIdentityMetadataPath  string

	service
}

func newUserService(
	sling *sling.Sling,
	uriTemplate string,
	apiKeysPath string,
	authenticateOctopusIDPath string,
	currentUserPath string,
	externalUserSearchPath string,
	registerPath string,
	signInPath string,
	signOutPath string,
	userAuthenticationPath string,
	userIdentityMetadataPath string) *userService {
	userService := &userService{
		apiKeysPath:               apiKeysPath,
		authenticateOctopusIDPath: authenticateOctopusIDPath,
		currentUserPath:           currentUserPath,
		externalUserSearchPath:    externalUserSearchPath,
		registerPath:              registerPath,
		signInPath:                signInPath,
		signOutPath:               signOutPath,
		userAuthenticationPath:    userAuthenticationPath,
		userIdentityMetadataPath:  userIdentityMetadataPath,
	}
	userService.service = newService(serviceUserService,
		sling,
		uriTemplate,
		new(model.User))

	return userService
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

	path := trimTemplate(s.getPath())
	path = path + "/me"

	resp, err := apiGet(s.getClient(), s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

// Add creates a new user.
func (s userService) Add(resource *model.User) (*model.User, error) {
	path, err := getAddPath(s, resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}

// GetAll returns all tenants. If none can be found or an error occurs, it
// returns an empty collection.
func (s userService) GetAll() ([]*model.User, error) {
	items := []*model.User{}
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

	path := trimTemplate(s.getPath())
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

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/authentication/%s", user.ID)

	resp, err := apiGet(s.getClient(), new(model.UserAuthentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.UserAuthentication), nil
}

func (s userService) GetSpaces(user *model.User) ([]*model.Spaces, error) {
	if user == nil {
		return nil, createInvalidParameterError("GetSpaces", "user")
	}

	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/%s/spaces", user.ID)

	resp, err := apiGet(s.getClient(), new([]model.Spaces), path)
	if err != nil {
		return nil, err
	}

	return resp.([]*model.Spaces), nil
}

// Update modifies a user based on the one provided as input.
func (s userService) Update(resource model.User) (*model.User, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, s.itemType, path)
	if err != nil {
		return nil, err
	}

	return resp.(*model.User), nil
}
