package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccountResource(t *testing.T) {
	invalidTenantedDeploymentMode := "***"
	name := getRandomName()
	spaceID := "space-id"
	tenantedDeploymentMode := "Untenanted"

	testCases := []struct {
		TestName               string
		IsError                bool
		Name                   string
		SpaceID                string
		TenantedDeploymentMode string
	}{
		{"Valid", false, name, spaceID, tenantedDeploymentMode},
		{"EmptyName", true, emptyString, spaceID, tenantedDeploymentMode},
		{"WhitespaceName", true, whitespaceString, spaceID, tenantedDeploymentMode},
		{"EmptySpaceID", false, name, emptyString, tenantedDeploymentMode},
		{"WhitespaceSpaceID", true, name, whitespaceString, tenantedDeploymentMode},
		{"EmptyTenantedDeploymentMode", true, name, spaceID, emptyString},
		{"WhitespaceTenantedDeploymentMode", true, name, spaceID, whitespaceString},
		{"InvalidTenantedDeploymentMode", true, name, spaceID, invalidTenantedDeploymentMode},
	}
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			accountResourceInline := &AccountResource{
				AccountType: "None",
				Name:        tc.Name,
			}
			accountResourceWithNew := newAccountResource(tc.Name, "None")
			accountResourceInline.SpaceID = tc.SpaceID
			accountResourceWithNew.SpaceID = tc.SpaceID
			accountResourceInline.TenantedDeploymentMode = tc.TenantedDeploymentMode
			accountResourceWithNew.TenantedDeploymentMode = tc.TenantedDeploymentMode
			if tc.IsError {
				require.Error(t, accountResourceInline.Validate())
				require.Error(t, accountResourceWithNew.Validate())
			} else {
				require.NoError(t, accountResourceInline.Validate())
				require.NoError(t, accountResourceWithNew.Validate())
			}
		})
	}
}

func TestAccountResourceAsJSON(t *testing.T) {
	description := "description"
	id := "id-value"
	lastModifiedBy := "john.smith@example.com"
	lastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")
	links := map[string]string{
		"Self": "/test",
		"Foo":  "/test-2",
	}
	spaceID := "space-id"
	tenantedDeploymentMode := "Untenanted"
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

	assert.Equal(t, description, account.Description)
	assert.Equal(t, id, account.ID)
	assert.Equal(t, id, account.GetID())
	assert.Equal(t, lastModifiedBy, account.LastModifiedBy)
	assert.Equal(t, lastModifiedBy, account.GetLastModifiedBy())
	assert.Equal(t, lastModifiedOn, *account.LastModifiedOn)
	assert.Equal(t, lastModifiedOn, *account.GetLastModifiedOn())
	assert.Equal(t, links, account.Links)
	assert.Equal(t, spaceID, account.SpaceID)
	assert.Equal(t, tenantedDeploymentMode, account.TenantedDeploymentMode)
	assert.Equal(t, tenantIDs, account.TenantIDs)
	assert.Equal(t, tenantTags, account.TenantTags)

	newID := "new-id-value"
	newLastModifiedBy := "alice.smith@example.com"
	newLastModifiedOn, _ := time.Parse(time.RFC3339, "2020-10-02T00:44:11.284Z")

	account.SetID(newID)
	account.SetLastModifiedBy(newLastModifiedBy)
	account.SetLastModifiedOn(&newLastModifiedOn)

	assert.Equal(t, description, account.Description)
	assert.Equal(t, newID, account.ID)
	assert.Equal(t, newID, account.GetID())
	assert.Equal(t, newLastModifiedBy, account.LastModifiedBy)
	assert.Equal(t, newLastModifiedBy, account.GetLastModifiedBy())
	assert.Equal(t, newLastModifiedOn, *account.LastModifiedOn)
	assert.Equal(t, newLastModifiedOn, *account.GetLastModifiedOn())
	assert.Equal(t, links, account.Links)
	assert.Equal(t, spaceID, account.SpaceID)
	assert.Equal(t, tenantedDeploymentMode, account.TenantedDeploymentMode)
	assert.Equal(t, tenantIDs, account.TenantIDs)
	assert.Equal(t, tenantTags, account.TenantTags)

	account.SetID(id)
	account.SetLastModifiedBy(lastModifiedBy)
	account.SetLastModifiedOn(&lastModifiedOn)

	accountAsJSON, err := json.Marshal(account)
	require.NoError(t, err)
	require.NotNil(t, accountAsJSON)

	jsonassert.New(t).Assertf(exampleAsJSON, string(accountAsJSON))
}

const exampleAsJSON string = `{
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
