package client

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
)

const (
	TestAccountServiceURITemplate = "accounts-service"
)

func TestNewAccountService(t *testing.T) {
	service := NewAccountService(nil, "")
	assert.Nil(t, service)
	createAccountService(t)
}

func TestAccountServiceGetWithEmptyID(t *testing.T) {
	service := createAccountService(t)

	resource, err := service.Get("")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "id"))
	assert.Nil(t, resource)

	resource, err = service.Get(" ")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Get", "id"))
	assert.Nil(t, resource)
}

func TestAccountServiceGetWithEmptyName(t *testing.T) {
	service := createAccountService(t)

	resource, err := service.GetByName("")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("GetByName", "name"))
	assert.Nil(t, resource)

	resource, err = service.GetByName(" ")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("GetByName", "name"))
	assert.Nil(t, resource)
}

func TestAccountServiceAddWithNilAccount(t *testing.T) {
	service := createAccountService(t)

	account, err := service.Add(nil)

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Add", "account"))
	assert.Nil(t, account)
}

func TestAccountServiceAddWithInvalidAccount(t *testing.T) {
	service := createAccountService(t)

	account, err := service.Add(&model.Account{})

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestAccountServiceDeleteWithEmptyID(t *testing.T) {
	service := createAccountService(t)

	err := service.Delete("")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Delete", "id"))

	err = service.Delete(" ")

	assert.Error(t, err)
	assert.Equal(t, err, createInvalidParameterError("Delete", "id"))
}

func TestAccountServiceUpdateWithEmptyAccount(t *testing.T) {
	service := createAccountService(t)

	account, err := service.Update(model.Account{})

	assert.Error(t, err)
	assert.Nil(t, account)
}

func createAccountService(t *testing.T) *AccountService {
	service := NewAccountService(&sling.Sling{}, TestAccountServiceURITemplate)

	assert.NotNil(t, service)
	assert.NotNil(t, service.sling)
	assert.Equal(t, service.path, TestAccountServiceURITemplate)
	assert.Equal(t, service.name, "AccountService")

	return service
}
