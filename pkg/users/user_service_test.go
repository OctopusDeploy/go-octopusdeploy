package users

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/require"
)

func createUserService(t *testing.T) *UserService {
	service := NewUserService(nil,
		constants.TestURIUsers,
		constants.TestURIAPIKeys,
		constants.TestURIAuthenticateOctopusID,
		constants.TestURICurrentUser,
		constants.TestURIExternalUserSearch,
		constants.TestURIRegister,
		constants.TestURISignIn,
		constants.TestURISignOut,
		constants.TestURIUserAuthentication,
		constants.TestURIUserIdentityMetadata)
	services.NewServiceTests(t, service, constants.TestURIUsers, constants.ServiceUserService)
	return service
}

func TestUserService(t *testing.T) {
	ServiceFunction := NewUserService
	client := &sling.Sling{}
	uriTemplate := ""
	apiKeysPath := ""
	authenticateOctopusIDPath := ""
	currentUserPath := ""
	externalUserSearchPath := ""
	registerPath := ""
	signInPath := ""
	signOutPath := ""
	userAuthenticationPath := ""
	userIdentityMetadataPath := ""
	ServiceName := constants.ServiceUserService

	testCases := []struct {
		name                      string
		f                         func(*sling.Sling, string, string, string, string, string, string, string, string, string, string) *UserService
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
			"",
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
			" ",
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
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestUserServiceGetByID(t *testing.T) {
	service := createUserService(t)
	require.NotNil(t, service)

	user, err := service.GetByID("")
	require.Error(t, err)
	require.Nil(t, user)

	user, err = service.GetByID(" ")
	require.Error(t, err)
	require.Nil(t, user)
}

func TestUserServiceNilInputs(t *testing.T) {
	service := createUserService(t)
	require.NotNil(t, service)

	user, err := service.Add(nil)
	require.Nil(t, user)
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationAdd, constants.ParameterUser), err)

	apiKey, err := service.GetAPIKeyByID(nil, "")
	require.Nil(t, apiKey)
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetAPIKeyByID, constants.ParameterUser), err)

	apiKeys, err := service.GetAPIKeys(nil)
	require.Nil(t, apiKeys)
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGetAPIKeys, constants.ParameterUser), err)
}
