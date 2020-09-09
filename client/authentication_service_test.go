package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthenticationServiceWithNil(t *testing.T) {
	service := NewAuthenticationService(nil)
	assert.Nil(t, service)
}

func TestAuthenticationServiceWithEmptyClient(t *testing.T) {
	service := NewAuthenticationService(&sling.Sling{})
	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
}
