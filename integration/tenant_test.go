package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/client"
	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func cleanTenant(t *testing.T, octopusClient *client.Client, tenantID string) {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	err := octopusClient.Tenants.DeleteByID(tenantID)
	assert.NoError(t, err)
}

func createTestTenant(t *testing.T, octopusClient *client.Client, name string) model.Tenant {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	p := createTenant(name)
	resource, err := octopusClient.Tenants.Add(&p)
	require.NoError(t, err)

	return *resource
}

func createTenant(tenantName string) model.Tenant {
	p := model.NewTenant(tenantName, "Lifecycles-1")
	return *p
}

func isEqualTenants(t *testing.T, expected model.Tenant, actual model.Tenant) {
	assert := assert.New(t)

	// equality cannot be determined through a direct comparison (below)
	// because APIs like GetByPartialName do not include the fields,
	// LastModifiedBy and LastModifiedOn
	//
	// assert.EqualValues(expected, actual)
	//
	// this statement (above) is expected to succeed, but it fails due to these
	// missing fields

	assert.Equal(expected.ClonedFromTenantID, actual.ClonedFromTenantID)
	assert.Equal(expected.Description, actual.Description)
	assert.Equal(expected.ID, actual.ID)
	assert.Equal(expected.Links, actual.Links)
	assert.Equal(expected.Name, actual.Name)
	assert.Equal(expected.ProjectEnvironments, actual.ProjectEnvironments)
	assert.Equal(expected.SpaceID, actual.SpaceID)
	assert.Equal(expected.TenantTags, actual.TenantTags)
}

func TestTenants(t *testing.T) {
	t.Run("AddAndDelete", TestTenantAddAndDelete)
	t.Run("AddGetAndDelete", TestTenantAddGetAndDelete)
	t.Run("GetAll", TestTenantGetAll)
	t.Run("GetByPartialName", TestTenantGetByPartialName)
	t.Run("Update", TestTenantUpdate)
}

func TestTenantAddAndDelete(t *testing.T) {
	tenantName := getRandomName()
	expected := createTenant(tenantName)
	actual := createTestTenant(t, nil, tenantName)

	defer cleanTenant(t, nil, actual.ID)

	assert.Equal(t, expected.Name, actual.Name, "tenant name doesn't match expected")
	assert.NotEmpty(t, actual.ID, "tenant doesn't contain an ID from the octopus server")
}

func TestTenantAddGetAndDelete(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	expected := createTestTenant(t, octopusClient, getRandomName())
	defer cleanTenant(t, octopusClient, expected.ID)

	actual, err := octopusClient.Tenants.GetByID(expected.ID)
	assert.NoError(t, err)
	isEqualTenants(t, expected, *actual)
}

func TestTenantGetAll(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	const count int = 32
	expected := map[string]model.Tenant{}
	for i := 0; i < count; i++ {
		resource := createTestTenant(t, octopusClient, getRandomName())
		defer cleanTenant(t, octopusClient, resource.ID)
		expected[resource.ID] = resource
	}

	resources, err := octopusClient.Tenants.GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, resources)
	assert.GreaterOrEqual(t, len(resources), count)

	for _, actual := range resources {
		_, ok := expected[actual.ID]
		if ok {
			isEqualTenants(t, expected[actual.ID], actual)
		}
	}
}

func TestTenantGetByPartialName(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	expected := createTestTenant(t, octopusClient, getRandomName())
	defer cleanTenant(t, octopusClient, expected.ID)

	resources, err := octopusClient.Tenants.GetByPartialName(expected.Name)
	assert.NoError(t, err)
	assert.NotNil(t, resources)

	for _, actual := range resources {
		isEqualTenants(t, expected, actual)
	}
}

func TestTenantUpdate(t *testing.T) {
	octopusClient := getOctopusClient()
	require.NotNil(t, octopusClient)

	expected := createTestTenant(t, octopusClient, getRandomName())
	defer cleanTenant(t, octopusClient, expected.ID)

	expected.Name = getRandomName()
	expected.Description = getRandomName()

	actual, err := octopusClient.Tenants.Update(expected)
	assert.NoError(t, err)
	assert.NotNil(t, actual)

	isEqualTenants(t, expected, *actual)
}
