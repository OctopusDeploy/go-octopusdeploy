package services

import (
	"testing"

	"github.com/dghubble/sling"
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
	testNewService(t, service, TestURIUsers, ServiceUserService)
	return service
}

func TestUserService(t *testing.T) {
	ServiceFunction := newUserService
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
	ServiceName := ServiceUserService

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
			ServiceFunction,
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
			ServiceFunction,
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
			ServiceFunction,
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
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
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
}

func TestUserServiceNilInputs(t *testing.T) {
	service := createUserService(t)
	require.NotNil(t, service)

	user, err := service.Add(nil)
	require.Nil(t, user)
	require.Equal(t, createInvalidParameterError(OperationAdd, ParameterUser), err)

	apiKey, err := service.GetAPIKeyByID(nil, emptyString)
	require.Nil(t, apiKey)
	require.Equal(t, createInvalidParameterError(OperationGetAPIKeyByID, ParameterUser), err)

	apiKeys, err := service.GetAPIKeys(nil)
	require.Nil(t, apiKeys)
	require.Equal(t, createInvalidParameterError(OperationGetAPIKeys, ParameterUser), err)
}
