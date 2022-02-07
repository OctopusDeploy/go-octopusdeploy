package octopusdeploy

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/services"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewAPIKeyService(t *testing.T) {
	ServiceFunction := newAPIKeyService
	client := &sling.Sling{}
	uriTemplate := services.emptyString
	ServiceName := ServiceAPIKeyService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *apiKeyService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, services.emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, services.whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			services.testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestAPIKeyServiceGetWithEmptyID(t *testing.T) {
	service := createAPIKeyService(t)
	resource, err := service.GetByUserID(services.emptyString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByUserID, ParameterUserID))
	assert.Nil(t, resource)

	resource, err = service.GetByUserID(services.whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByUserID, ParameterUserID))
	assert.Nil(t, resource)
}

func TestAPIKeyServiceDeleteByID(t *testing.T) {
	service := createAPIKeyService(t)
	user := createServiceAccountUser(t)
	resource := NewAPIKey(getRandomName(), user.GetID())
	assert.NotNil(t, resource)

	resource, err := service.Create(resource)
	assert.NoError(t, err)
	assert.NotNil(t, resource)
}

func createServiceAccountUser(t *testing.T) *User {
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
		TestURIUserIdentityMetadata,
	)
	services.testNewService(t, service, TestURIUsers, ServiceUserService)

	user := NewUser(getRandomName(), getRandomName())
	user.IsService = true
	assert.NotNil(t, user)

	resource, err := service.Add(user)
	assert.NoError(t, err)
	assert.NotNil(t, resource)

	return resource
}

func createAPIKeyService(t *testing.T) *apiKeyService {
	service := newAPIKeyService(nil, TestURIAPIKeys)
	services.testNewService(t, service, TestURIAPIKeys, ServiceAPIKeyService)
	return service
}
