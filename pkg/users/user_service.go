package users

import (
	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/permissions"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/spaces"
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
)

type UserService struct {
	apiKeysPath               string
	authenticateOctopusIDPath string
	currentUserPath           string
	externalUserSearchPath    string
	registerPath              string
	signInPath                string
	signOutPath               string
	userAuthenticationPath    string
	userIdentityMetadataPath  string

	services.CanDeleteService
}

func NewUserService(
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
	userIdentityMetadataPath string) *UserService {

	return &UserService{
		apiKeysPath:               apiKeysPath,
		authenticateOctopusIDPath: authenticateOctopusIDPath,
		currentUserPath:           currentUserPath,
		externalUserSearchPath:    externalUserSearchPath,
		registerPath:              registerPath,
		signInPath:                signInPath,
		signOutPath:               signOutPath,
		userAuthenticationPath:    userAuthenticationPath,
		userIdentityMetadataPath:  userIdentityMetadataPath,
		CanDeleteService: services.CanDeleteService{
			Service: services.NewService(constants.ServiceUserService, sling, uriTemplate),
		},
	}
}

// Add creates a new user.
func (s *UserService) Add(user *User) (*User, error) {
	if IsNil(user) {
		return nil, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterUser)
	}

	path, err := services.GetAddPath(s, user)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiAdd(s.GetClient(), user, new(User), path)
	if err != nil {
		return nil, err
	}

	return resp.(*User), nil
}

// GetAll returns all users. If none can be found or an error occurs, it
// returns an empty collection.
func (s *UserService) GetAll() ([]*User, error) {
	items := []*User{}
	path, err := services.GetAllPath(s)
	if err != nil {
		return items, err
	}

	_, err = services.ApiGet(s.GetClient(), &items, path)
	return items, err
}

func (s *UserService) GetAPIKeyByID(user *User, apiKeyID string) (*APIKey, error) {
	if user == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetAPIKeyByID, constants.ParameterUser)
	}

	// TODO: validate apiKeyID

	path := internal.TrimTemplate(user.Links[constants.LinkAPIKeys]) + "/" + apiKeyID

	response, err := services.ApiGet(s.GetClient(), new(APIKey), path)
	if err != nil {
		return nil, err
	}

	return response.(*APIKey), nil
}

func (s *UserService) GetAPIKeys(user *User, apiQuery ...APIQuery) (*APIKeys, error) {
	if user == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetAPIKeys, constants.ParameterUser)
	}

	// URI template: /api/users/[user-id]/apikeys{/id}{?skip,take}
	//
	// The URI template associated with this service endpoint permits a get-by
	// criteria (i.e. {/id}) or get-all criteria with support for constraining
	// the results (i.e. {?skip,take}). This function assumes that if the get-by
	// criteria is specified then the query parameters will be ignored.
	// Otherwise, this function assumes the get-all criteria and will apply
	// constraints to the results through query parameters.

	// Permissions required: TODO

	path := internal.TrimTemplate(user.Links[constants.LinkAPIKeys])

	v, _ := query.Values(apiQuery)
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := services.ApiGet(s.GetClient(), new(APIKeys), path)
	if err != nil {
		return nil, err
	}

	return response.(*APIKeys), nil
}

func (s *UserService) GetAuthentication() (*UserAuthentication, error) {
	path := internal.TrimTemplate(s.userAuthenticationPath)
	resp, err := services.ApiGet(s.GetClient(), new(UserAuthentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*UserAuthentication), nil
}

func (s *UserService) GetAuthenticationByUser(user *User) (*UserAuthentication, error) {
	if user == nil {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetAuthenticationByUser, constants.ParameterUser)
	}

	path := internal.TrimTemplate(s.userAuthenticationPath) + "/" + user.GetID()

	resp, err := services.ApiGet(s.GetClient(), new(UserAuthentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*UserAuthentication), nil
}

// Get returns a collection of users based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
func (s *UserService) Get(usersQuery UsersQuery) (*Users, error) {
	path, err := s.GetURITemplate().Expand(usersQuery)
	if err != nil {
		return &Users{}, err
	}

	response, err := services.ApiGet(s.GetClient(), new(Users), path)
	if err != nil {
		return &Users{}, err
	}

	return response.(*Users), nil
}

// GetByID returns the user that matches the input ID. If one cannot be found,
// it returns nil and an error.
func (s *UserService) GetByID(id string) (*User, error) {
	if internal.IsEmpty(id) {
		return nil, internal.CreateInvalidParameterError(constants.OperationGetByID, constants.ParameterID)
	}

	path, err := services.GetByIDPath(s, id)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiGet(s.GetClient(), new(User), path)
	if err != nil {
		return nil, internal.CreateResourceNotFoundError("user", "ID", id)
	}

	return resp.(*User), nil
}

// GetMe returns the user associated with the key used to invoke this API.
func (s *UserService) GetMe() (*User, error) {
	if err := services.ValidateInternalState(s); err != nil {
		return nil, err
	}

	path := internal.TrimTemplate(s.GetPath())
	path = path + "/me"

	resp, err := services.ApiGet(s.GetClient(), new(User), path)
	if err != nil {
		return nil, err
	}

	return resp.(*User), nil
}

func (s *UserService) GetPermissions(user *User, userQuery ...UserQuery) (*permissions.UserPermissionSet, error) {
	if user == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterUser)
	}

	v, _ := query.Values(userQuery)
	path := internal.TrimTemplate(user.Links[constants.LinkPermissions])
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := services.ApiGet(s.GetClient(), new(permissions.UserPermissionSet), path)
	return response.(*permissions.UserPermissionSet), err
}

func (s *UserService) GetPermissionsConfiguration(user *User, userQuery ...UserQuery) (*permissions.UserPermissionSet, error) {
	if user == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterUser)
	}

	v, _ := query.Values(userQuery)
	path := internal.TrimTemplate(user.Links[constants.LinkPermissionsConfiguration])
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := services.ApiGet(s.GetClient(), new(permissions.UserPermissionSet), path)
	return response.(*permissions.UserPermissionSet), err
}

func (s *UserService) GetSpaces(user *User) ([]*spaces.Space, error) {
	if user == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterUser)
	}

	// TODO: check permissions

	path := internal.TrimTemplate(user.Links[constants.LinkSpaces])
	items := []*spaces.Space{}
	_, err := services.ApiGet(s.GetClient(), &items, path)
	return items, err
}

func (s *UserService) GetTeams(user *User, userQuery ...UserQuery) (*[]permissions.ProjectedTeamReferenceDataItem, error) {
	if user == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterUser)
	}

	v, _ := query.Values(userQuery)
	path := internal.TrimTemplate(user.Links[constants.LinkTeams])
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := services.ApiGet(s.GetClient(), new([]permissions.ProjectedTeamReferenceDataItem), path)
	if err != nil {
		return nil, err
	}

	return response.(*[]permissions.ProjectedTeamReferenceDataItem), nil
}

// Update modifies a user based on the one provided as input.
func (s *UserService) Update(user *User) (*User, error) {
	if user == nil {
		return nil, internal.CreateRequiredParameterIsEmptyOrNilError(constants.ParameterUser)
	}

	path, err := services.GetUpdatePath(s, user)
	if err != nil {
		return nil, err
	}

	resp, err := services.ApiUpdate(s.GetClient(), user, new(User), path)
	if err != nil {
		return nil, err
	}

	return resp.(*User), nil
}
