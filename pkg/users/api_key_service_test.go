package users

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/services"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewAPIKeyService(t *testing.T) {
	ServiceFunction := NewAPIKeyService
	client := &sling.Sling{}
	uriTemplate := ""
	ServiceName := constants.ServiceAPIKeyService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *ApiKeyService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, ""},
		{"URITemplateWithWhitespace", ServiceFunction, client, " "},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.NewServiceTests(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestAPIKeyServiceGetWithEmptyID(t *testing.T) {
	service := createAPIKeyService(t)
	resource, err := service.GetByUserID("")

	assert.Equal(t, err, internal.CreateInvalidParameterError("GetByUserID", "userID"))
	assert.Nil(t, resource)

	resource, err = service.GetByUserID(" ")

	assert.Equal(t, err, internal.CreateInvalidParameterError("GetByUserID", "userID"))
	assert.Nil(t, resource)
}

func TestAPIKeyServiceDeleteByID(t *testing.T) {
	service := createAPIKeyService(t)
	user := createServiceAccountUser(t)
	resource := NewAPIKey(internal.GetRandomName(), user.GetID())
	assert.NotNil(t, resource)

	resource, err := service.Create(resource)
	assert.NoError(t, err)
	assert.NotNil(t, resource)
}

func createServiceAccountUser(t *testing.T) *User {
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
		constants.TestURIUserIdentityMetadata,
	)
	services.NewServiceTests(t, service, constants.TestURIUsers, constants.ServiceUserService)

	user := NewUser(internal.GetRandomName(), internal.GetRandomName())
	user.IsService = true
	assert.NotNil(t, user)

	resource, err := service.Add(user)
	assert.NoError(t, err)
	assert.NotNil(t, resource)

	return resource
}

func createAPIKeyService(t *testing.T) *ApiKeyService {
	service := NewAPIKeyService(nil, constants.TestURIAPIKeys)
	services.NewServiceTests(t, service, constants.TestURIAPIKeys, constants.ServiceAPIKeyService)
	return service
}
