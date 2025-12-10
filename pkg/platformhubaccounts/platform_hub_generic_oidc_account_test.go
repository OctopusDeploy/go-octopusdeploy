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

func TestPlatformHubGenericOidcAccountNew(t *testing.T) {
	accountType := AccountTypePlatformHubGenericOidcAccount
	description := ""
	name := internal.GetRandomName()

	account, err := NewPlatformHubGenericOidcAccount(name)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// resource
	require.Equal(t, "", account.ID)
	require.Equal(t, "", account.ModifiedBy)
	require.Nil(t, account.ModifiedOn)
	require.NotNil(t, account.Links)

	// IResource
	require.Equal(t, "", account.GetID())
	require.Equal(t, "", account.GetModifiedBy())
	require.Nil(t, account.GetModifiedOn())
	require.NotNil(t, account.GetLinks())

	// IPlatformHubAccount
	require.Equal(t, accountType, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// PlatformHubGenericOidcAccount
	require.Nil(t, account.ExecutionSubjectKeys)
	require.Equal(t, "", account.Audience)
}

func TestPlatformHubGenericOidcAccountMarshalJSON(t *testing.T) {
	name := internal.GetRandomName()
	executionSubjectKeys := []string{"space", "project"}
	audience := "api://default"

	expectedJson := fmt.Sprintf(`{
		"ExecutionSubjectKeys": ["space", "project"],
		"Audience": "%s",
		"AccountType": "GenericOidcAccount",
		"Name": "%s"
	}`, audience, name)

	account, err := NewPlatformHubGenericOidcAccount(name)
	require.NoError(t, err)
	require.NotNil(t, account)

	account.ExecutionSubjectKeys = executionSubjectKeys
	account.Audience = audience

	accountAsJSON, err := json.Marshal(account)
	require.NoError(t, err)
	require.NotNil(t, accountAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(accountAsJSON))
}

func TestPlatformHubGenericOidcAccountNewWithConfigs(t *testing.T) {
	accountType := AccountTypePlatformHubGenericOidcAccount
	id := internal.GetRandomName()
	modifiedBy := internal.GetRandomName()
	modifiedOn := time.Now()
	name := internal.GetRandomName()
	description := "Description for " + name + " (OK to Delete)"
	executionSubjectKeys := []string{"space", "environment", "project"}
	audience := "api://default"

	account, err := NewPlatformHubGenericOidcAccount(name)
	require.NoError(t, err)
	require.NotNil(t, account)
	require.NoError(t, account.Validate())

	account.Description = description
	account.ID = id
	account.ModifiedBy = modifiedBy
	account.ModifiedOn = &modifiedOn
	account.ExecutionSubjectKeys = executionSubjectKeys
	account.Audience = audience

	// resource
	require.Equal(t, id, account.ID)
	require.Equal(t, modifiedBy, account.ModifiedBy)
	require.Equal(t, &modifiedOn, account.ModifiedOn)
	require.NotNil(t, account.Links)

	// IResource
	require.Equal(t, id, account.GetID())
	require.Equal(t, modifiedBy, account.GetModifiedBy())
	require.Equal(t, &modifiedOn, account.GetModifiedOn())
	require.NotNil(t, account.GetLinks())

	// IPlatformHubAccount
	require.Equal(t, accountType, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// PlatformHubGenericOidcAccount
	require.Equal(t, executionSubjectKeys, account.ExecutionSubjectKeys)
	require.Equal(t, audience, account.Audience)
}
