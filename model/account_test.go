package model

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/enum"
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var accountName = "Name"

func TestEmptyAccount(t *testing.T) {
	account := &Account{}

	assert.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestAccountWithName(t *testing.T) {
	account := &Account{Name: accountName}

	assert.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestAccountWithAccountType(t *testing.T) {
	account := &Account{AccountType: enum.UsernamePassword}

	assert.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestAccountWithNameAndUsernamePasswordAccountType(t *testing.T) {
	account := &Account{
		AccountType: enum.UsernamePassword,
		Name:        accountName,
	}

	assert.NotNil(t, account)
	assert.NoError(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountWithInvalidParameters(t *testing.T) {
	account, err := NewAccount(accountName, enum.AzureServicePrincipal)

	assert.NoError(t, err)
	assert.NotNil(t, account)

	if err != nil {
		return
	}

	assert.Error(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountOnlyWithValidSubscriptionNumber(t *testing.T) {
	account, err := NewAccount(accountName, enum.AzureServicePrincipal)
	subscriptionID := uuid.New()
	account.SubscriptionID = &subscriptionID

	assert.NoError(t, err)
	assert.NotNil(t, account)

	if err != nil {
		return
	}

	assert.Error(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountWithInvalidClientID(t *testing.T) {
	account, err := NewAccount(accountName, enum.AzureServicePrincipal)
	subscriptionID := uuid.New()
	account.SubscriptionID = &subscriptionID
	applicationID := uuid.Nil
	account.ApplicationID = &applicationID

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountOnlyWithValidProperties(t *testing.T) {
	account, err := NewAccount(accountName, enum.AzureServicePrincipal)
	subscriptionID := uuid.New()
	account.SubscriptionID = &subscriptionID
	applicationID := uuid.New()
	account.ApplicationID = &applicationID
	tenantID := uuid.New()
	account.TenantID = &tenantID

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.NoError(t, account.Validate())
}

func TestNewAccountWithEmptyName(t *testing.T) {
	account, err := NewAccount(whitespaceString, enum.UsernamePassword)

	assert.Error(t, err)
	assert.Nil(t, account)
}
