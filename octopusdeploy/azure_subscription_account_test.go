package octopusdeploy

import (
	"testing"

	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAzureSubscriptionAccountNew(t *testing.T) {
	name := getRandomName()
	subscriptionID := uuid.New()

	account, err := NewAzureSubscriptionAccount(name, subscriptionID)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	require.Equal(t, emptyString, account.Description)
	require.Equal(t, emptyString, account.GetDescription())
	require.Equal(t, emptyString, account.GetID())
	require.Equal(t, name, account.Name)
	require.Equal(t, name, account.GetName())
	require.Equal(t, AccountTypeAzureSubscription, account.GetAccountType())
	require.NotNil(t, account.Links)
	require.NotNil(t, account.GetLinks())
}

func TestAzureSubscriptionAccountSetDescription(t *testing.T) {
	description := getRandomName()
	name := getRandomName()
	subscriptionID := uuid.New()

	account, err := NewAzureSubscriptionAccount(name, subscriptionID)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	account.Description = description

	require.Equal(t, description, account.Description)
	require.Equal(t, description, account.GetDescription())
}

func TestAzureSubscriptionAccountSetName(t *testing.T) {
	name := getRandomName()
	subscriptionID := uuid.New()

	account, err := NewAzureSubscriptionAccount(name, subscriptionID)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	require.Equal(t, name, account.Name)
	require.Equal(t, name, account.GetName())
}
