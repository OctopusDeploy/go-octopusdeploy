package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

const (
	TestAuthenticationServiceURITemplate = "authentication-service"
)

func TestNewAuthenticationService(t *testing.T) {
	service := NewAuthenticationService(nil, "")
	assert.Nil(t, service)
	createAuthenticationService(t)
}

func createAuthenticationService(t *testing.T) *AuthenticationService {
	service := NewAuthenticationService(&sling.Sling{}, TestAuthenticationServiceURITemplate)

	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
	assert.Equal(t, service.path, TestAuthenticationServiceURITemplate)
	assert.Equal(t, service.name, "AuthenticationService")

	return service
}
