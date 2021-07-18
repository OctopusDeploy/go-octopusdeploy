package octopusdeploy

import (
	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
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
	userService.service = newService(ServiceUserService, sling, uriTemplate)

	return userService
}

// Add creates a new user.
func (s userService) Add(user *User) (*User, error) {
	if user == nil {
		return nil, createInvalidParameterError(OperationAdd, ParameterUser)
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

func (s userService) GetAPIKeyByID(user *User, apiKeyID string) (*APIKey, error) {
	if user == nil {
		return nil, createInvalidParameterError(OperationGetAPIKeyByID, ParameterUser)
	}

	// TODO: validate apiKeyID

	path := trimTemplate(user.Links[linkAPIKeys]) + "/" + apiKeyID

	response, err := apiGet(s.getClient(), new(APIKey), path)
	if err != nil {
		return nil, err
	}

	return response.(*APIKey), nil
}

func (s userService) GetAPIKeys(user *User, apiQuery ...APIQuery) (*APIKeys, error) {
	if user == nil {
		return nil, createInvalidParameterError(OperationGetAPIKeys, ParameterUser)
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

	path := trimTemplate(user.Links[linkAPIKeys])

	v, _ := query.Values(apiQuery)
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := apiGet(s.getClient(), new(APIKeys), path)
	if err != nil {
		return nil, err
	}

	return response.(*APIKeys), nil
}

func (s userService) GetAuthentication() (*UserAuthentication, error) {
	path := trimTemplate(s.userAuthenticationPath)
	resp, err := apiGet(s.getClient(), new(UserAuthentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*UserAuthentication), nil
}

func (s userService) GetAuthenticationByUser(user *User) (*UserAuthentication, error) {
	if user == nil {
		return nil, createInvalidParameterError(OperationGetAuthenticationByUser, ParameterUser)
	}

	path := trimTemplate(s.userAuthenticationPath) + "/" + user.GetID()

	resp, err := apiGet(s.getClient(), new(UserAuthentication), path)
	if err != nil {
		return nil, err
	}

	return resp.(*UserAuthentication), nil
}

// Get returns a collection of users based on the criteria defined by its input
// query parameter. If an error occurs, an empty collection is returned along
// with the associated error.
func (s userService) Get(usersQuery UsersQuery) (*Users, error) {
	path, err := s.getURITemplate().Expand(usersQuery)
	if err != nil {
		return &Users{}, err
	}

	response, err := apiGet(s.getClient(), new(Users), path)
	if err != nil {
		return &Users{}, err
	}

	return response.(*Users), nil
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

// GetMe returns the user associated with the key used to invoke this API.
func (s userService) GetMe() (*User, error) {
	if err := validateInternalState(s); err != nil {
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

func (s userService) GetPermissions(user *User, userQuery ...UserQuery) (*UserPermissionSet, error) {
	if user == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterUser)
	}

	v, _ := query.Values(userQuery)
	path := trimTemplate(user.Links[linkPermissions])
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := apiGet(s.getClient(), new(UserPermissionSet), path)
	return response.(*UserPermissionSet), err
}

func (s userService) GetPermissionsConfiguration(user *User, userQuery ...UserQuery) (*UserPermissionSet, error) {
	if user == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterUser)
	}

	v, _ := query.Values(userQuery)
	path := trimTemplate(user.Links[linkPermissionsConfiguration])
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := apiGet(s.getClient(), new(UserPermissionSet), path)
	return response.(*UserPermissionSet), err
}

func (s userService) GetSpaces(user *User) ([]*Space, error) {
	if user == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterUser)
	}

	// TODO: check permissions

	path := trimTemplate(user.Links[linkSpaces])
	items := []*Space{}
	_, err := apiGet(s.getClient(), &items, path)
	return items, err
}

func (s userService) GetTeams(user *User, userQuery ...UserQuery) (*[]ProjectedTeamReferenceDataItem, error) {
	if user == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterUser)
	}

	v, _ := query.Values(userQuery)
	path := trimTemplate(user.Links[linkTeams])
	encodedQueryString := v.Encode()
	if len(encodedQueryString) > 0 {
		path += "?" + encodedQueryString
	}

	response, err := apiGet(s.getClient(), new([]ProjectedTeamReferenceDataItem), path)
	if err != nil {
		return nil, err
	}

	return response.(*[]ProjectedTeamReferenceDataItem), nil
}

// Update modifies a user based on the one provided as input.
func (s userService) Update(user *User) (*User, error) {
	if user == nil {
		return nil, createRequiredParameterIsEmptyOrNilError(ParameterUser)
	}

	path, err := getUpdatePath(s, user)
	if err != nil {
		return nil, err
	}

	resp, err := apiUpdate(s.getClient(), user, new(User), path)
	if err != nil {
		return nil, err
	}

	return resp.(*User), nil
}
