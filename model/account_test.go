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
	assert.Error(t, account.Validate())
}

func TestAccountCreationWithInvalidUUIDs(t *testing.T) {
	account, err := NewAccount(accountName, enum.AzureServicePrincipal)
	subscriptionNumber := uuid.New()
	account.SubscriptionNumber = &subscriptionNumber
	clientID := uuid.New()
	account.ClientID = &clientID
	tenantID := uuid.New()
	account.TenantID = &tenantID

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.NoError(t, account.Validate())

	account.ClientID = &uuid.Nil

	assert.Error(t, account.Validate())
}

func TestNewAccountForUsernamePasswordWithValidParameters(t *testing.T) {
	account, err := NewAccount(accountName, enum.UsernamePassword)

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountWithInvalidParameters(t *testing.T) {
	account, err := NewAccount(accountName, enum.AzureServicePrincipal)

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountOnlyWithValidSubscriptionNumber(t *testing.T) {
	account, err := NewAccount(accountName, enum.AzureServicePrincipal)
	subscriptionNumber := uuid.New()
	account.SubscriptionNumber = &subscriptionNumber

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountWithInvalidClientID(t *testing.T) {
	account, err := NewAccount(accountName, enum.AzureServicePrincipal)
	subscriptionNumber := uuid.New()
	account.SubscriptionNumber = &subscriptionNumber
	clientID := uuid.Nil
	account.ClientID = &clientID

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountOnlyWithValidProperties(t *testing.T) {
	account, err := NewAccount(accountName, enum.AzureServicePrincipal)
	subscriptionNumber := uuid.New()
	account.SubscriptionNumber = &subscriptionNumber
	clientID := uuid.New()
	account.ClientID = &clientID
	tenantID := uuid.New()
	account.TenantID = &tenantID

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.NoError(t, account.Validate())
}

func TestNewAccountWithEmptyName(t *testing.T) {
	account, err := NewAccount(" ", enum.UsernamePassword)

	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestNewAccountWithLongEmptyName(t *testing.T) {
	account, err := NewAccount("       ", enum.UsernamePassword)

	assert.Error(t, err)
	assert.Nil(t, account)
}
