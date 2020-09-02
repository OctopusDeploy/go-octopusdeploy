package client

import (
	"testing"

	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

func TestNewAccountServiceWithNil(t *testing.T) {
	service := NewAccountService(nil)
	assert.Nil(t, service)
}

func TestNewAccountServiceWithEmptyClient(t *testing.T) {
	service := NewAccountService(&sling.Sling{})
	assert.NotNil(t, service)
	assert.Equal(t, service.path, "accounts")
	assert.NotNil(t, service.sling)
}
