package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
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

func TestAccountServiceGetWithEmptyID(t *testing.T) {
	service := NewAccountService(&sling.Sling{})

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "accounts")

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestAccountServiceGetWithEmptyName(t *testing.T) {
	service := NewAccountService(&sling.Sling{})

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "accounts")

	resource, err := service.GetByName("")

	assert.Error(t, err)
	assert.Nil(t, resource)

	resource, err = service.GetByName(" ")

	assert.Error(t, err)
	assert.Nil(t, resource)
}

func TestAccountServiceAddWithNilAccount(t *testing.T) {
	service := NewAccountService(&sling.Sling{})

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "accounts")

	account, err := service.Add(nil)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestAccountServiceAddWithInvalidAccount(t *testing.T) {
	service := NewAccountService(&sling.Sling{})

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "accounts")

	account, err := service.Add(&model.Account{})

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestAccountServiceDeleteWithEmptyID(t *testing.T) {
	service := NewAccountService(&sling.Sling{})

	assert.NotNil(t, service)
	assert.Equal(t, service.path, "accounts")

	err := service.Delete("")

	assert.Error(t, err)

	err = service.Delete(" ")

	assert.Error(t, err)
}
