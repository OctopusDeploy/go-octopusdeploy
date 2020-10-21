package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func cleanTenant(t *testing.T, octopusClient *octopusdeploy.Client, tenantID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	err := octopusClient.Tenants.DeleteByID(tenantID)
	assert.NoError(t, err)
}

func CreateTestTenant(t *testing.T, octopusClient *octopusdeploy.Client) *octopusdeploy.Tenant {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	name := getRandomName()
	lifecycleID := getRandomName()

	tenant := octopusdeploy.NewTenant(name, lifecycleID)

	createdTenant, err := octopusClient.Tenants.Add(tenant)
	require.NoError(t, err)
	require.NotNil(t, createdTenant)
	require.NotEmpty(t, createdTenant.GetID())

	return createdTenant
}

func DeleteTestTenant(t *testing.T, octopusClient *octopusdeploy.Client, tenant *octopusdeploy.Tenant) error {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	return octopusClient.Tenants.DeleteByID(tenant.GetID())
}

func IsEqualTenants(t *testing.T, expected *octopusdeploy.Tenant, actual *octopusdeploy.Tenant) {
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

func TestTenantAddGetAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	expected := CreateTestTenant(t, octopusClient)
	defer cleanTenant(t, octopusClient, expected.GetID())

	actual, err := octopusClient.Tenants.GetByID(expected.GetID())
	assert.NoError(t, err)
	IsEqualTenants(t, expected, actual)
}

func TestTenantServiceGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	tenants := []octopusdeploy.Tenant{}

	// create 30 test tenants (to be deleted)
	for i := 0; i < 30; i++ {
		tenant := CreateTestTenant(t, octopusClient)
		require.NotNil(t, tenant)
		tenants = append(tenants, *tenant)
	}

	allTenants, err := octopusClient.Tenants.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allTenants)
	require.True(t, len(allTenants) >= 30)

	for _, tenant := range tenants {
		require.NotNil(t, tenant)
		require.NotEmpty(t, tenant.GetID())
		err = DeleteTestTenant(t, octopusClient, &tenant)
		require.NoError(t, err)
	}
}

func TestTenantGetByPartialName(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	expected := CreateTestTenant(t, octopusClient)
	defer cleanTenant(t, octopusClient, expected.GetID())

	resources, err := octopusClient.Tenants.GetByPartialName(expected.Name)
	assert.NoError(t, err)
	assert.NotNil(t, resources)

	for _, actual := range resources {
		IsEqualTenants(t, expected, actual)
	}
}

func TestTenantUpdate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	expected := CreateTestTenant(t, octopusClient)
	defer cleanTenant(t, octopusClient, expected.GetID())

	expected.Name = getRandomName()
	expected.Description = getRandomName()

	actual, err := octopusClient.Tenants.Update(expected)
	assert.NoError(t, err)
	assert.NotNil(t, actual)

	IsEqualTenants(t, expected, actual)
}
