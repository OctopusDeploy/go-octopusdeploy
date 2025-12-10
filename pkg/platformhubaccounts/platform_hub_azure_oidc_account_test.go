package platformhubaccounts

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestPlatformHubAzureOidcAccountNew(t *testing.T) {
	subscriptionID := "00000000-0000-0000-0000-000000000000"
	applicationID := "11111111-1111-1111-1111-111111111111"
	tenantID := "22222222-2222-2222-2222-222222222222"
	accountType := AccountTypePlatformHubAzureOidcAccount
	description := ""
	name := internal.GetRandomName()

	account, err := NewPlatformHubAzureOidcAccount(name, subscriptionID, applicationID, tenantID)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	require.Equal(t, "", account.ID)
	require.Equal(t, "", account.ModifiedBy)
	require.Nil(t, account.ModifiedOn)
	require.NotNil(t, account.Links)

	require.Equal(t, "", account.GetID())
	require.Equal(t, "", account.GetModifiedBy())
	require.Nil(t, account.GetModifiedOn())
	require.NotNil(t, account.GetLinks())

	require.Equal(t, accountType, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	require.Equal(t, subscriptionID, account.SubscriptionID)
	require.Equal(t, applicationID, account.ApplicationID)
	require.Equal(t, tenantID, account.TenantID)
}

func TestPlatformHubAzureOidcAccountMarshalJSON(t *testing.T) {
	subscriptionID := "00000000-0000-0000-0000-000000000000"
	applicationID := "11111111-1111-1111-1111-111111111111"
	tenantID := "22222222-2222-2222-2222-222222222222"
	name := internal.GetRandomName()

	expectedJson := fmt.Sprintf(`{
		"SubscriptionNumber": "%s",
		"ClientId": "%s",
		"TenantId": "%s",
		"AccountType": "AzureOidc",
		"Name": "%s"
	}`, subscriptionID, applicationID, tenantID, name)

	account, err := NewPlatformHubAzureOidcAccount(name, subscriptionID, applicationID, tenantID)
	require.NoError(t, err)
	require.NotNil(t, account)

	accountAsJSON, err := json.Marshal(account)
	require.NoError(t, err)
	require.NotNil(t, accountAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(accountAsJSON))
}

func TestPlatformHubAzureOidcAccountNewWithConfigs(t *testing.T) {
	subscriptionID := "00000000-0000-0000-0000-000000000000"
	applicationID := "11111111-1111-1111-1111-111111111111"
	tenantID := "22222222-2222-2222-2222-222222222222"
	accountType := AccountTypePlatformHubAzureOidcAccount
	id := internal.GetRandomName()
	modifiedBy := internal.GetRandomName()
	modifiedOn := time.Now()
	name := internal.GetRandomName()
	description := "Description for " + name + " (OK to Delete)"

	account, err := NewPlatformHubAzureOidcAccount(name, subscriptionID, applicationID, tenantID)
	require.NoError(t, err)
	require.NotNil(t, account)
	require.NoError(t, account.Validate())

	account.Description = description
	account.ID = id
	account.ModifiedBy = modifiedBy
	account.ModifiedOn = &modifiedOn
	account.ExecutionSubjectKeys = []string{"space", "project"}
	account.Audience = "api://default"

	require.Equal(t, id, account.ID)
	require.Equal(t, modifiedBy, account.ModifiedBy)
	require.Equal(t, &modifiedOn, account.ModifiedOn)
	require.NotNil(t, account.Links)

	require.Equal(t, id, account.GetID())
	require.Equal(t, modifiedBy, account.GetModifiedBy())
	require.Equal(t, &modifiedOn, account.GetModifiedOn())
	require.NotNil(t, account.GetLinks())

	require.Equal(t, accountType, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	require.Equal(t, subscriptionID, account.SubscriptionID)
	require.Equal(t, applicationID, account.ApplicationID)
	require.Equal(t, tenantID, account.TenantID)
}
