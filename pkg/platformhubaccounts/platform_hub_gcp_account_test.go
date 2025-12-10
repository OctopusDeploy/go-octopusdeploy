package platformhubaccounts

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestPlatformHubGcpAccountNew(t *testing.T) {
	jsonKey := core.NewSensitiveValue(internal.GetRandomName())
	accountType := AccountTypePlatformHubGcpAccount
	description := ""
	name := internal.GetRandomName()

	account, err := NewPlatformHubGcpAccount(name, jsonKey)

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

	// PlatformHubGcpAccount
	require.Equal(t, jsonKey, account.JsonKey)
}

func TestPlatformHubGcpAccountMarshalJSON(t *testing.T) {
	jsonKey := core.NewSensitiveValue(internal.GetRandomName())
	name := internal.GetRandomName()

	jsonKeyAsJSON, err := json.Marshal(jsonKey)
	require.NoError(t, err)
	require.NotNil(t, jsonKeyAsJSON)

	expectedJson := fmt.Sprintf(`{
		"JsonKey": %s,
		"AccountType": "GoogleCloudAccount",
		"Name": "%s"
	}`, jsonKeyAsJSON, name)

	account, err := NewPlatformHubGcpAccount(name, jsonKey)
	require.NoError(t, err)
	require.NotNil(t, account)

	accountAsJSON, err := json.Marshal(account)
	require.NoError(t, err)
	require.NotNil(t, accountAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(accountAsJSON))
}

func TestPlatformHubGcpAccountNewWithConfigs(t *testing.T) {
	jsonKey := core.NewSensitiveValue(internal.GetRandomName())
	accountType := AccountTypePlatformHubGcpAccount
	id := internal.GetRandomName()
	modifiedBy := internal.GetRandomName()
	modifiedOn := time.Now()
	name := internal.GetRandomName()
	description := "Description for " + name + " (OK to Delete)"

	account, err := NewPlatformHubGcpAccount(name, jsonKey)
	require.NoError(t, err)
	require.NotNil(t, account)
	require.NoError(t, account.Validate())

	account.Description = description
	account.ID = id
	account.ModifiedBy = modifiedBy
	account.ModifiedOn = &modifiedOn

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

	// PlatformHubGcpAccount
	require.Equal(t, jsonKey, account.JsonKey)
}
