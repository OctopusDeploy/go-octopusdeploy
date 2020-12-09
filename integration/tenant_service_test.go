package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualTenants(t *testing.T, expected *octopusdeploy.Tenant, actual *octopusdeploy.Tenant) {
	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	assert.Equal(t, expected.ClonedFromTenantID, actual.ClonedFromTenantID)
	assert.Equal(t, expected.Description, actual.Description)
	assert.Equal(t, expected.GetID(), actual.GetID())
	assert.Equal(t, expected.Links, actual.Links)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.ProjectEnvironments, actual.ProjectEnvironments)
	assert.Equal(t, expected.SpaceID, actual.SpaceID)
	assert.Equal(t, expected.TenantTags, actual.TenantTags)
}

func CreateTestTenant(t *testing.T, octopusClient *octopusdeploy.Client) *octopusdeploy.Tenant {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	name := getRandomName()

	tenant := octopusdeploy.NewTenant(name)

	createdTenant, err := octopusClient.Tenants.Add(tenant)
	require.NoError(t, err)
	require.NotNil(t, createdTenant)
	require.NotEmpty(t, createdTenant.GetID())

	return createdTenant
}

func DeleteTestTenant(t *testing.T, client *octopusdeploy.Client, tenant *octopusdeploy.Tenant) {
	require.NotNil(t, tenant)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := client.Tenants.DeleteByID(tenant.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedTenant, err := client.Tenants.GetByID(tenant.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedTenant)
}

func TestTenantAddGetAndDelete(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	expected := CreateTestTenant(t, client)
	defer DeleteTestTenant(t, client, expected)

	actual, err := client.Tenants.GetByID(expected.GetID())
	assert.NoError(t, err)
	AssertEqualTenants(t, expected, actual)
}

func TestTenantServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// create 10 test tenants (to be deleted)
	for i := 0; i < 10; i++ {
		tenant := CreateTestTenant(t, client)
		require.NotNil(t, tenant)
		defer DeleteTestTenant(t, client, tenant)
	}

	allTenants, err := client.Tenants.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allTenants)
	require.True(t, len(allTenants) >= 10)
}

func TestTenantGetByPartialName(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	expected := CreateTestTenant(t, client)
	defer DeleteTestTenant(t, client, expected)

	resources, err := client.Tenants.GetByPartialName(expected.Name)
	assert.NoError(t, err)
	assert.NotNil(t, resources)

	for _, actual := range resources {
		AssertEqualTenants(t, expected, actual)
	}
}

func TestTenantUpdate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	expected := CreateTestTenant(t, octopusClient)
	defer DeleteTestTenant(t, octopusClient, expected)

	expected.Name = getRandomName()
	expected.Description = getRandomName()

	actual, err := octopusClient.Tenants.Update(expected)
	assert.NoError(t, err)
	assert.NotNil(t, actual)

	AssertEqualTenants(t, expected, actual)
}
