package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceWithEmptyProperties(t *testing.T) {
	userService := &UserService{}
	assert.Error(t, userService.validateInternalState())
}

func TestUserServiceWithEmptyPathProperty(t *testing.T) {
	userService := &UserService{sling: &sling.Sling{}}
	assert.Error(t, userService.validateInternalState())
}

func TestUserServiceWithEmptySlingProperty(t *testing.T) {
	userService := &UserService{path: "fake-path"}
	assert.Error(t, userService.validateInternalState())
}

func TestUserServiceWithValidProperties(t *testing.T) {
	userService := createTestUserService()
	assert.NoError(t, userService.validateInternalState())
}

func TestNewUserServiceWithNil(t *testing.T) {
	userService := NewUserService(nil)
	if !assert.Nil(t, userService) {
		assert.Error(t, userService.validateInternalState())
	}
}

func TestUserServiceWithEmptyClient(t *testing.T) {
	userService := NewUserService(&sling.Sling{})

	if !assert.NotNil(t, userService) {
		assert.NotNil(t, userService.sling)
		assert.Error(t, userService.validateInternalState())
	}
}

func TestUserServiceGetWithEmptyID(t *testing.T) {
	userService := createTestUserService()
	user, err := userService.Get("")

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUserServiceGetWithBlankID(t *testing.T) {
	userService := createTestUserService()
	user, err := userService.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, user)
}

func createTestUserService() *UserService {
	return &UserService{sling: &sling.Sling{}, path: "fake-path"}
}
