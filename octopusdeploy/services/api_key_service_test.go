package services

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewAPIKeyService(t *testing.T) {
	ServiceFunction := newAPIKeyService
	client := &sling.Sling{}
	uriTemplate := emptyString
	ServiceName := ServiceAPIKeyService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *apiKeyService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", ServiceFunction, nil, uriTemplate},
		{"EmptyURITemplate", ServiceFunction, client, emptyString},
		{"URITemplateWithWhitespace", ServiceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, ServiceName)
		})
	}
}

func TestAPIKeyServiceGetWithEmptyID(t *testing.T) {
	service := createAPIKeyService(t)
	resource, err := service.GetByUserID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByUserID, ParameterUserID))
	assert.Nil(t, resource)

	resource, err = service.GetByUserID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(OperationGetByUserID, ParameterUserID))
	assert.Nil(t, resource)
}

func TestAPIKeyServiceDeleteByID(t *testing.T) {
	service := createAPIKeyService(t)
	user := createServiceAccountUser(t)
	resource := NewAPIKey(octopusdeploy.getRandomName(), user.GetID())
	assert.NotNil(t, resource)

	resource, err := service.Create(resource)
	assert.NoError(t, err)
	assert.NotNil(t, resource)
}

func createServiceAccountUser(t *testing.T) *User {
	service := octopusdeploy.newUserService(nil,
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
	testNewService(t, service, TestURIUsers, ServiceUserService)

	user := NewUser(octopusdeploy.getRandomName(), octopusdeploy.getRandomName())
	user.IsService = true
	assert.NotNil(t, user)

	resource, err := service.Add(user)
	assert.NoError(t, err)
	assert.NotNil(t, resource)

	return resource
}

func createAPIKeyService(t *testing.T) *apiKeyService {
	service := newAPIKeyService(nil, TestURIAPIKeys)
	testNewService(t, service, TestURIAPIKeys, ServiceAPIKeyService)
	return service
}
