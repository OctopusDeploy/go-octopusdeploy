package octopusdeploy

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptyAccount(t *testing.T) {
	account := &Account{}

	require.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestAccountWithName(t *testing.T) {
	name := getRandomName()
	account := &Account{
		Name: name,
	}

	require.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestAccountWithAccountType(t *testing.T) {
	account := &Account{AccountType: accountTypeUsernamePassword}

	require.NotNil(t, account)
	assert.Error(t, account.Validate())
}

func TestAccountWithNameAndUsernamePasswordAccountType(t *testing.T) {
	name := getRandomName()
	account := &Account{
		AccountType:            accountTypeUsernamePassword,
		Name:                   name,
		TenantedDeploymentMode: "Untenanted",
	}

	require.NotNil(t, account)
	assert.NoError(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountWithInvalidParameters(t *testing.T) {
	name := getRandomName()
	account := NewAccount(name, accountTypeAzureServicePrincipal)
	require.NoError(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountOnlyWithValidSubscriptionNumber(t *testing.T) {
	name := getRandomName()
	account := NewAccount(name, accountTypeAzureServicePrincipal)
	subscriptionID := uuid.New()
	account.SubscriptionID = &subscriptionID
	require.NoError(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountWithInvalidClientID(t *testing.T) {
	name := getRandomName()
	account := NewAccount(name, accountTypeAzureServicePrincipal)
	subscriptionID := uuid.New()
	account.SubscriptionID = &subscriptionID
	assert.NoError(t, account.Validate())
}

func TestNewAccountForAzureServicePrincipalAccountOnlyWithValidProperties(t *testing.T) {
	name := getRandomName()
	account := NewAccount(name, accountTypeAzureServicePrincipal)
	subscriptionID := uuid.New()
	account.SubscriptionID = &subscriptionID
	applicationID := uuid.New()
	account.ApplicationID = &applicationID
	tenantID := uuid.New()
	account.TenantID = &tenantID
	assert.NoError(t, account.Validate())
}
