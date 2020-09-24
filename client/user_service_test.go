package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	serviceFunction := newUserService
	client := &sling.Sling{}
	uriTemplate := emptyString
	serviceName := serviceUserService

	testCases := []struct {
		name        string
		f           func(*sling.Sling, string) *userService
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

func TestUserServiceGetWithEmptyID(t *testing.T) {
	userService := createTestUserService()
	user, err := userService.GetByID(emptyString)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUserServiceGetWithBlankID(t *testing.T) {
	userService := createTestUserService()
	user, err := userService.GetByID(whitespaceString)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func createTestUserService() *userService {
	return &userService{sling: &sling.Sling{}, path: "fake-path"}
}
