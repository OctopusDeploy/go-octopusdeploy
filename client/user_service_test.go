package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createUserService(t *testing.T) *userService {
	service := newUserService(nil,
		TestURIUsers,
		TestURIAPIKeys,
		TestURIAuthenticateOctopusID,
		TestURICurrentUser,
		TestURIExternalUserSearch,
		TestURIRegister,
		TestURISignIn,
		TestURISignOut,
		TestURIUserAuthentication,
		TestURIUserIdentityMetadata)
	testNewService(t, service, TestURIUsers, serviceUserService)
	return service
}

func TestUserService(t *testing.T) {
	serviceFunction := newUserService
	client := &sling.Sling{}
	uriTemplate := emptyString
	apiKeysPath := emptyString
	authenticateOctopusIDPath := emptyString
	currentUserPath := emptyString
	externalUserSearchPath := emptyString
	registerPath := emptyString
	signInPath := emptyString
	signOutPath := emptyString
	userAuthenticationPath := emptyString
	userIdentityMetadataPath := emptyString
	serviceName := serviceUserService

	testCases := []struct {
		name                      string
		f                         func(*sling.Sling, string, string, string, string, string, string, string, string, string, string) *userService
		client                    *sling.Sling
		uriTemplate               string
		apiKeysPath               string
		authenticateOctopusIDPath string
		currentUserPath           string
		externalUserSearchPath    string
		registerPath              string
		signInPath                string
		signOutPath               string
		userAuthenticationPath    string
		userIdentityMetadataPath  string
	}{
		{"NilClient",
			serviceFunction,
			nil,
			uriTemplate,
			apiKeysPath,
			authenticateOctopusIDPath,
			currentUserPath,
			externalUserSearchPath,
			registerPath,
			signInPath,
			signOutPath,
			userAuthenticationPath,
			userIdentityMetadataPath},
		{"EmptyURITemplate",
			serviceFunction,
			client,
			emptyString,
			apiKeysPath,
			authenticateOctopusIDPath,
			currentUserPath,
			externalUserSearchPath,
			registerPath,
			signInPath,
			signOutPath,
			userAuthenticationPath,
			userIdentityMetadataPath},
		{"URITemplateWithWhitespace",
			serviceFunction,
			client,
			whitespaceString,
			apiKeysPath,
			authenticateOctopusIDPath,
			currentUserPath,
			externalUserSearchPath,
			registerPath,
			signInPath,
			signOutPath,
			userAuthenticationPath,
			userIdentityMetadataPath},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client,
				tc.apiKeysPath,
				tc.uriTemplate,
				tc.authenticateOctopusIDPath,
				tc.currentUserPath,
				tc.externalUserSearchPath,
				tc.registerPath,
				tc.signInPath,
				tc.signOutPath,
				tc.userAuthenticationPath,
				tc.userIdentityMetadataPath,
			)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func CreateTestUser(t *testing.T, service *userService) *model.User {
	if service == nil {
		service = createUserService(t)
	}
	require.NotNil(t, service)

	username := getRandomName()
	displayName := getRandomName()
	password := getRandomName()

	user := model.NewUser(username, displayName)
	user.Password = password
	require.NoError(t, user.Validate())

	createdUser, err := service.Add(user)
	require.NoError(t, err)
	require.NotNil(t, createdUser)
	require.NotEmpty(t, createdUser.GetID())
	require.Equal(t, username, createdUser.Username)
	require.Equal(t, displayName, createdUser.DisplayName)

	return createdUser
}

func DeleteTestUser(t *testing.T, service *userService, user model.User) error {
	require.NotNil(t, user)

	if service == nil {
		service = createUserService(t)
	}
	require.NotNil(t, service)

	err := service.DeleteByID(user.GetID())
	assert.NoError(t, err)

	return err
}

func TestUserServiceGetByID(t *testing.T) {
	service := createUserService(t)
	require.NotNil(t, service)

	user, err := service.GetByID(emptyString)
	require.Error(t, err)
	require.Nil(t, user)

	user, err = service.GetByID(whitespaceString)
	require.Error(t, err)
	require.Nil(t, user)

	users, err := service.GetAll()
	for _, user := range users {
		user, err := service.GetByID(user.GetID())
		require.NoError(t, err)
		require.NotNil(t, user)
	}
}

func TestUserServiceGetMe(t *testing.T) {
	service := createUserService(t)
	require.NotNil(t, service)

	user, err := service.GetMe()
	require.NoError(t, err)
	require.NotNil(t, user)
}
