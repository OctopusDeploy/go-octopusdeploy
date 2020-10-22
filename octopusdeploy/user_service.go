package octopusdeploy

import (
	"fmt"

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

	canDeleteService
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
		new(User))

	return userService
}

// Add creates a new user.
func (s userService) Add(user *User) (*User, error) {
	if user == nil {
		return nil, createInvalidParameterError(operationAdd, parameterUser)
	}

	path, err := getAddPath(s, user)
	if err != nil {
		return nil, err
	}

	resp, err := apiAdd(s.getClient(), user, new(User), path)
	if err != nil {
		return nil, err
	}

	return resp.(*User), nil
}

// GetAll returns all users. If none can be found or an error occurs, it
// returns an empty collection.
func (s userService) GetAll() ([]*User, error) {
	items := []*User{}
	path, err := getAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

func (s userService) GetAuthentication() (*UserAuthentication, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.getPath())
	path = path + "/authentication"

	resp, err := apiGet(s.getClient(), new(UserAuthentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*UserAuthentication), nil
}

func (s userService) GetAuthenticationForUser(user *User) (*UserAuthentication, error) {
	if user == nil {
		return nil, createInvalidParameterError(operationGetAuthenticationForUser, parameterUser)
	}

	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/authentication/%s", user.GetID())

	resp, err := apiGet(s.getClient(), new(UserAuthentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*UserAuthentication), nil
}

// GetByID returns the user that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s userService) GetByID(id string) (*User, error) {
	path, err := getByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := apiGet(s.getClient(), new(User), path)
	if err != nil {
		return nil, createResourceNotFoundError("user", "ID", id)
	}

	return resp.(*User), nil
}

func (s userService) GetMe() (*User, error) {
	err := validateInternalState(s)
	if err != nil {
		return nil, err
	}

	path := trimTemplate(s.getPath())
	path = path + "/me"

	resp, err := apiGet(s.getClient(), new(User), path)
	if err != nil {
		return nil, err
	}

	return resp.(*User), nil
}

func (s userService) GetSpaces(user *User) ([]*Spaces, error) {
	if user == nil {
		return nil, createInvalidParameterError("GetSpaces", "user")
	}

	items := []*Spaces{}
	err := validateInternalState(s)
	if err != nil {
		return items, err
	}

	path := trimTemplate(s.getPath())
	path = fmt.Sprintf(path+"/%s/spaces", user.GetID())

	_, err = apiGet(s.getClient(), &items, path)
	return items, err
}

// Update modifies a user based on the one provided as input.
func (s userService) Update(resource User) (*User, error) {
	path, err := getUpdatePath(s, &resource)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), resource, new(User), path)
	if err != nil {
		return nil, err
	}

	return resp.(*User), nil
}
