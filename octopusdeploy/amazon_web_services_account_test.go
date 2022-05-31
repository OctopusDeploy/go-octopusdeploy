package octopusdeploy

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestAmazonWebServicesAccountNew(t *testing.T) {
	accessKey := getRandomName()
	accountType := AccountTypeAmazonWebServicesAccount
	description := emptyString
	environmentIDs := []string{}
	name := getRandomName()
	secretKey := NewSensitiveValue(getRandomName())
	spaceID := emptyString
	tenantedDeploymentMode := TenantedDeploymentMode("Untenanted")

	account, err := NewAmazonWebServicesAccount(name, accessKey, secretKey)

	require.NotNil(t, account)
	require.NoError(t, err)
	require.NoError(t, account.Validate())

	// resource
	require.Equal(t, emptyString, account.ID)
	require.Equal(t, emptyString, account.ModifiedBy)
	require.Nil(t, account.ModifiedOn)
	require.NotNil(t, account.Links)

	// IResource
	require.Equal(t, emptyString, account.GetID())
	require.Equal(t, emptyString, account.GetModifiedBy())
	require.Nil(t, account.GetModifiedOn())
	require.NotNil(t, account.GetLinks())

	// account
	require.Equal(t, description, account.Description)
	require.Equal(t, environmentIDs, account.EnvironmentIDs)
	require.Equal(t, name, account.Name)
	require.Equal(t, spaceID, account.SpaceID)
	require.Equal(t, tenantedDeploymentMode, account.TenantedDeploymentMode)

	// IAccount
	require.Equal(t, accountType, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// AmazonWebServicesAccount
	require.Equal(t, accessKey, account.AccessKey)
	require.Equal(t, secretKey, account.SecretKey)
}

func TestAmazonWebServicesAccountMarshalJSON(t *testing.T) {
	accessKey := getRandomName()
	name := getRandomName()
	secretKey := NewSensitiveValue(getRandomName())

	secretKeyAsJSON, err := json.Marshal(secretKey)
	require.NoError(t, err)
	require.NotNil(t, secretKeyAsJSON)

	expectedJson := fmt.Sprintf(`{
		"AccessKey": "%s",
		"AccountType": "AmazonWebServicesAccount",
		"Name": "%s",
		"SecretKey": %s,
		"TenantedDeploymentParticipation": "Untenanted"
	}`, accessKey, name, secretKeyAsJSON)

	account, err := NewAmazonWebServicesAccount(name, accessKey, secretKey)
	require.NoError(t, err)
	require.NotNil(t, account)

	accountAsJSON, err := json.Marshal(account)
	require.NoError(t, err)
	require.NotNil(t, accountAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(accountAsJSON))
}

func TestAmazonWebServicesAccountNewWithConfigs(t *testing.T) {
	accessKey := getRandomName()
	accountType := AccountTypeAmazonWebServicesAccount
	environmentIDs := []string{"environment-id-1", "environment-id-2"}
	id := getRandomName()
	modifiedBy := getRandomName()
	modifiedOn := time.Now()
	name := getRandomName()
	description := "Description for " + name + " (OK to Delete)"
	secretKey := NewSensitiveValue(getRandomName())
	spaceID := getRandomName()
	tenantedDeploymentMode := TenantedDeploymentMode("Tenanted")

	account, err := NewAmazonWebServicesAccount(name, accessKey, secretKey)
	require.NoError(t, err)
	require.NotNil(t, account)
	require.NoError(t, account.Validate())

	account.Description = description
	account.EnvironmentIDs = environmentIDs
	account.ID = id
	account.ModifiedBy = modifiedBy
	account.ModifiedOn = &modifiedOn
	account.SpaceID = spaceID
	account.TenantedDeploymentMode = tenantedDeploymentMode

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

	// account
	require.Equal(t, description, account.Description)
	require.Equal(t, environmentIDs, account.EnvironmentIDs)
	require.Equal(t, name, account.Name)
	require.Equal(t, spaceID, account.SpaceID)
	require.Equal(t, tenantedDeploymentMode, account.TenantedDeploymentMode)

	// IAccount
	require.Equal(t, accountType, account.GetAccountType())
	require.Equal(t, description, account.GetDescription())
	require.Equal(t, name, account.GetName())

	// AmazonWebServicesAccount
	require.Equal(t, accessKey, account.AccessKey)
	require.Equal(t, secretKey, account.SecretKey)
}
