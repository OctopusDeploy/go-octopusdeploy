package accounts

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountResource(t *testing.T) {
	name := internal.GetRandomName()
	spaceID := internal.GetRandomName()
	tenantedDeploymentMode := core.TenantedDeploymentMode("Untenanted")

	testCases := []struct {
		TestName               string
		IsError                bool
		Name                   string
		SpaceID                string
		TenantedDeploymentMode core.TenantedDeploymentMode
	}{
		{"Valid", false, name, spaceID, tenantedDeploymentMode},
		{"EmptyName", true, "", spaceID, tenantedDeploymentMode},
		{"WhitespaceName", true, " ", spaceID, tenantedDeploymentMode},
		{"EmptySpaceID", false, name, "", tenantedDeploymentMode},
		{"WhitespaceSpaceID", false, name, " ", tenantedDeploymentMode},
	}
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			accountInline := &account{
				AccountType: AccountType("None"),
				Name:        tc.Name,
			}
			accountWithNew := newAccount(tc.Name, AccountType("None"))
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
	lastModifiedBy := "john.smith@example.com"
	lastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")
	links := map[string]string{
		"Self": "/test",
		"Foo":  "/test-2",
	}
	spaceID := "space-id"
	tenantedDeploymentMode := core.TenantedDeploymentMode("Untenanted")
	tenantIDs := []string{
		"tenant-id-1",
		"tenant-id-2",
	}
	tenantTags := []string{
		"tenant-tag-1",
		"tenant-tag-2",
	}

	exampleAsJSON := `{
		"AccountType": "None",
		"Description": "description",
		"EnvironmentIds": [
			"environment-id-1",
			"environment-id-2"
		],
		"Id": "id-value",
		"LastModifiedOn": "2020-10-02T00:44:11.284Z",
		"LastModifiedBy": "john.smith@example.com",
		"Links": {
		"Self": "/test",
		"Foo": "/test-2"
		},
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

	var account AccountResource
	err := json.Unmarshal([]byte(exampleAsJSON), &account)
	require.NoError(t, err)
	require.NotNil(t, account)
	require.NoError(t, account.Validate())

	assert.Equal(t, accountType, account.AccountType)
	assert.Equal(t, accountType, account.GetAccountType())
	assert.Equal(t, description, account.Description)
	assert.Equal(t, id, account.GetID())
	assert.Equal(t, lastModifiedBy, account.GetModifiedBy())
	assert.Equal(t, lastModifiedOn, *account.GetModifiedOn())
	assert.Equal(t, links, account.Links)
	assert.Equal(t, spaceID, account.SpaceID)
	assert.Equal(t, tenantedDeploymentMode, account.TenantedDeploymentMode)
	assert.Equal(t, tenantIDs, account.TenantIDs)
	assert.Equal(t, tenantTags, account.TenantTags)

	accountAsJSON, err := json.Marshal(account)
	require.NoError(t, err)
	require.NotNil(t, accountAsJSON)

	jsonassert.New(t).Assertf(exampleAsJSON, string(accountAsJSON))
}
