package accounts

import (
	"encoding/json"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountResource(t *testing.T) {
	name := getRandomName()
	spaceID := getRandomName()
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Untenanted")

	testCases := []struct {
		TestName               string
		IsError                bool
		Name                   string
		SpaceID                string
		TenantedDeploymentMode resources.TenantedDeploymentMode
	}{
		{"Valid", false, name, spaceID, tenantedDeploymentMode},
		{"EmptyName", true, emptyString, spaceID, tenantedDeploymentMode},
		{"WhitespaceName", true, whitespaceString, spaceID, tenantedDeploymentMode},
		{"EmptySpaceID", false, name, emptyString, tenantedDeploymentMode},
		{"WhitespaceSpaceID", true, name, whitespaceString, tenantedDeploymentMode},
	}
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			accountInline := &Account{
				AccountType: AccountType("None"),
				Name:        tc.Name,
			}
			accountWithNew := NewAccount(tc.Name, AccountType("None"))
			accountInline.SpaceID = tc.SpaceID
			accountWithNew.SpaceID = tc.SpaceID
			accountInline.TenantedDeploymentMode = tc.TenantedDeploymentMode
			accountWithNew.TenantedDeploymentMode = tc.TenantedDeploymentMode
			if tc.IsError {
				require.Error(t, accountInline.Validate())
				require.Error(t, accountWithNew.Validate())
			} else {
				require.NoError(t, accountInline.Validate())
				require.NoError(t, accountWithNew.Validate())
			}
		})
	}
}

func TestAccountResourceAsJSON(t *testing.T) {
	accountType := AccountType("None")
	description := "description"
	id := "id-value"
	spaceID := "space-id"
	tenantedDeploymentMode := resources.TenantedDeploymentMode("Untenanted")
	tenantIDs := []string{
		"tenant-id-1",
		"tenant-id-2",
	}
	tenantTags := []string{
		"tenant-tag-1",
		"tenant-tag-2",
	}

	var account AccountResource
	err := json.Unmarshal([]byte(exampleAsJSON), &account)
	require.NoError(t, err)
	require.NotNil(t, account)
	require.NoError(t, account.Validate())

	assert.Equal(t, accountType, account.AccountType)
	assert.Equal(t, accountType, account.GetAccountType())
	assert.Equal(t, description, account.Description)
	assert.Equal(t, id, account.GetID())
	assert.Equal(t, spaceID, account.SpaceID)
	assert.Equal(t, tenantedDeploymentMode, account.TenantedDeploymentMode)
	assert.Equal(t, tenantIDs, account.TenantIDs)
	assert.Equal(t, tenantTags, account.TenantTags)

	accountAsJSON, err := json.Marshal(account)
	require.NoError(t, err)
	require.NotNil(t, accountAsJSON)

	jsonassert.New(t).Assertf(exampleAsJSON, string(accountAsJSON))
}

const exampleAsJSON string = `{
	"AccountType": "None",
	"Description": "description",
	"EnvironmentIds": [
		"environment-id-1",
		"environment-id-2"
	],
	"Id": "id-value",
	"Name": "name",
	"SpaceId": "space-id",
	"TenantedDeploymentParticipation": "Untenanted",
	"TenantIds": [
		"tenant-id-1",
		"tenant-id-2"
	],
	"TenantTags": [
		"tenant-tag-1",
		"tenant-tag-2"
	]
  }`
