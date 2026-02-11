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

func TestPlatformHubUsernamePasswordAccountNew(t *testing.T) {
	username := internal.GetRandomName()
	password := core.NewSensitiveValue(internal.GetRandomName())
	accountType := AccountTypePlatformHubUsernamePasswordAccount
	description := ""
	name := internal.GetRandomName()

	account, err := NewPlatformHubUsernamePasswordAccount(name, username, password)

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

	require.Equal(t, username, account.Username)
	require.Equal(t, password, account.Password)
}

func TestPlatformHubUsernamePasswordAccountMarshalJSON(t *testing.T) {
	username := internal.GetRandomName()
	password := core.NewSensitiveValue(internal.GetRandomName())
	name := internal.GetRandomName()

	passwordAsJSON, err := json.Marshal(password)
	require.NoError(t, err)
	require.NotNil(t, passwordAsJSON)

	expectedJson := fmt.Sprintf(`{
		"Username": "%s",
		"Password": %s,
		"AccountType": "UsernamePassword",
		"Name": "%s"
	}`, username, passwordAsJSON, name)

	account, err := NewPlatformHubUsernamePasswordAccount(name, username, password)
	require.NoError(t, err)
	require.NotNil(t, account)

	accountAsJSON, err := json.Marshal(account)
	require.NoError(t, err)
	require.NotNil(t, accountAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(accountAsJSON))
}

func TestPlatformHubUsernamePasswordAccountNewWithConfigs(t *testing.T) {
	username := internal.GetRandomName()
	password := core.NewSensitiveValue(internal.GetRandomName())
	accountType := AccountTypePlatformHubUsernamePasswordAccount
	id := internal.GetRandomName()
	modifiedBy := internal.GetRandomName()
	modifiedOn := time.Now()
	name := internal.GetRandomName()
	description := "Description for " + name + " (OK to Delete)"

	account, err := NewPlatformHubUsernamePasswordAccount(name, username, password)
	require.NoError(t, err)
	require.NotNil(t, account)
	require.NoError(t, account.Validate())

	account.Description = description
	account.ID = id
	account.ModifiedBy = modifiedBy
	account.ModifiedOn = &modifiedOn

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

	require.Equal(t, username, account.Username)
	require.Equal(t, password, account.Password)
}
