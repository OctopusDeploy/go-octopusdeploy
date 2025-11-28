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

func TestPlatformHubAwsAccountNew(t *testing.T) {
	accessKey := internal.GetRandomName()
	accountType := AccountTypePlatformHubAwsAccount
	description := ""
	name := internal.GetRandomName()
	secretKey := core.NewSensitiveValue(internal.GetRandomName())

	account, err := NewPlatformHubAwsAccount(name, accessKey, secretKey)

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

	// PlatformHubAwsAccount
	require.Equal(t, accessKey, account.AccessKey)
	require.Equal(t, secretKey, account.SecretKey)
}

func TestPlatformHubAwsAccountMarshalJSON(t *testing.T) {
	accessKey := internal.GetRandomName()
	name := internal.GetRandomName()
	secretKey := core.NewSensitiveValue(internal.GetRandomName())

	secretKeyAsJSON, err := json.Marshal(secretKey)
	require.NoError(t, err)
	require.NotNil(t, secretKeyAsJSON)

	expectedJson := fmt.Sprintf(`{
		"AccessKey": "%s",
		"AccountType": "AmazonWebServicesAccount",
		"Name": "%s",
		"SecretKey": %s
	}`, accessKey, name, secretKeyAsJSON)

	account, err := NewPlatformHubAwsAccount(name, accessKey, secretKey)
	require.NoError(t, err)
	require.NotNil(t, account)

	accountAsJSON, err := json.Marshal(account)
	require.NoError(t, err)
	require.NotNil(t, accountAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(accountAsJSON))
}

func TestPlatformHubAwsAccountNewWithConfigs(t *testing.T) {
	accessKey := internal.GetRandomName()
	accountType := AccountTypePlatformHubAwsAccount
	id := internal.GetRandomName()
	modifiedBy := internal.GetRandomName()
	modifiedOn := time.Now()
	name := internal.GetRandomName()
	description := "Description for " + name + " (OK to Delete)"
	secretKey := core.NewSensitiveValue(internal.GetRandomName())

	account, err := NewPlatformHubAwsAccount(name, accessKey, secretKey)
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

	// PlatformHubAwsAccount
	require.Equal(t, accessKey, account.AccessKey)
	require.Equal(t, secretKey, account.SecretKey)
}
