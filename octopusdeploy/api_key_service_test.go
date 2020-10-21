package octopusdeploy

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewAPIKeyService(t *testing.T) {
	serviceFunction := newAPIKeyService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceAPIKeyService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *apiKeyService
		client      *sling.Sling
		uriTemplate string
	}{
		{"NilClient", serviceFunction, nil, uriTemplate},
		{"EmptyURITemplate", serviceFunction, client, emptyString},
		{"URITemplateWithWhitespace", serviceFunction, client, whitespaceString},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			service := tc.f(tc.client, tc.uriTemplate)
			testNewService(t, service, uriTemplate, serviceName)
		})
	}
}

func TestAPIKeyServiceGetWithEmptyID(t *testing.T) {
	service := createAPIKeyService(t)
	resource, err := service.GetByUserID(emptyString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByUserID, parameterUserID))
	assert.Nil(t, resource)

	resource, err = service.GetByUserID(whitespaceString)

	assert.Equal(t, err, createInvalidParameterError(operationGetByUserID, parameterUserID))
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
	testNewService(t, service, TestURIUsers, serviceUserService)

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
	testNewService(t, service, TestURIAPIKeys, serviceAPIKeyService)
	return service
}
