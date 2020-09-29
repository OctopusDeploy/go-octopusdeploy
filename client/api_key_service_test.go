package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"

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
	assert := assert.New(t)

	resource, err := service.GetByUserID(emptyString)

	assert.Equal(err, createInvalidParameterError(operationGetByUserID, parameterUserID))
	assert.Nil(resource)

	resource, err = service.GetByUserID(whitespaceString)

	assert.Equal(err, createInvalidParameterError(operationGetByUserID, parameterUserID))
	assert.Nil(resource)
}

func TestAPIKeyServiceDeleteByID(t *testing.T) {
	service := createAPIKeyService(t)
	assert := assert.New(t)

	user := createServiceAccountUser(t)
	resource, err := model.NewAPIKey(getRandomName(), user.ID)

	assert.NoError(err)
	assert.NotNil(resource)

	resource, err = service.Create(resource)

	assert.NoError(err)
	assert.NotNil(resource)
}

func createServiceAccountUser(t *testing.T) *model.User {
	service := newUserService(nil, TestURIUsers)
	assert := assert.New(t)

	testNewService(t, service, TestURIUsers, serviceUserService)

	user := model.NewUser(getRandomName(), getRandomName())
	user.IsService = true
	assert.NotNil(t, user)

	resource, err := service.Add(user)
	assert.NoError(err)
	assert.NotNil(resource)

	return resource
}

func createAPIKeyService(t *testing.T) *apiKeyService {
	service := newAPIKeyService(nil, TestURIAPIKeys)
	testNewService(t, service, TestURIAPIKeys, serviceAPIKeyService)
	return service
}
