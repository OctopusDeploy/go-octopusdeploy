package model

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var accountName = "Name"

func TestEmptyAccount(t *testing.T) {
	account := &Account{}

	require.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestAccountWithName(t *testing.T) {
	account := &Account{Name: accountName}

	require.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestAccountWithAccountType(t *testing.T) {
	account := &Account{AccountType: accountTypeUsernamePassword}

	require.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestAccountWithNameAndUsernamePasswordAccountType(t *testing.T) {
	account := &Account{
		AccountType:            accountTypeUsernamePassword,
		Name:                   accountName,
		TenantedDeploymentMode: "Untenanted",
	}

	require.NotNil(t, account)
	assert.NoError(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountWithInvalidParameters(t *testing.T) {
	account := NewAccount(accountName, accountTypeAzureServicePrincipal)
	require.NoError(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountOnlyWithValidSubscriptionNumber(t *testing.T) {
	account := NewAccount(accountName, accountTypeAzureServicePrincipal)
	subscriptionID := uuid.New()
	account.SubscriptionID = &subscriptionID
	require.NoError(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountWithInvalidClientID(t *testing.T) {
	account := NewAccount(accountName, accountTypeAzureServicePrincipal)
	subscriptionID := uuid.New()
	account.SubscriptionID = &subscriptionID
	assert.NoError(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountOnlyWithValidProperties(t *testing.T) {
	account := NewAccount(accountName, accountTypeAzureServicePrincipal)
	subscriptionID := uuid.New()
	account.SubscriptionID = &subscriptionID
	applicationID := uuid.New()
	account.ApplicationID = &applicationID
	tenantID := uuid.New()
	account.TenantID = &tenantID
	assert.NoError(t, account.Validate())
}
