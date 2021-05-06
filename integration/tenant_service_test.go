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

func CreateTestTenant(t *testing.T, octopusClient *octopusdeploy.Client, project *octopusdeploy.Project, environment *octopusdeploy.Environment) *octopusdeploy.Tenant {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	name := getRandomName()

	tenant := octopusdeploy.NewTenant(name)
	tenant.Description = getRandomName()

	if project != nil {
		tenant.ProjectEnvironments[project.ID] = []string{environment.ID}
	}

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

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	libraryVariableSet := CreateLibraryVariableSet(t, client)
	require.NotNil(t, libraryVariableSet)
	defer DeleteLibraryVariableSet(t, client, libraryVariableSet)

	environment := CreateTestEnvironment(t, client)
	defer DeleteTestEnvironment(t, client, environment)

	expected := CreateTestTenant(t, client, project, environment)
	defer DeleteTestTenant(t, client, expected)

	missingVariablesQuery := octopusdeploy.MissingVariablesQuery{}

	tenantMissingVariables, err := client.Tenants.GetMissingVariables(missingVariablesQuery)
	require.NoError(t, err)
	require.NotNil(t, tenantMissingVariables)

	tenantVariables := octopusdeploy.NewTenantVariable(expected.GetID())
	require.NotNil(t, tenantVariables)

	tenantVariables, err = client.Tenants.UpdateVariables(expected, tenantVariables)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	tenantVariables, err = client.Tenants.GetVariables(expected)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	tenantVariables, err = client.Tenants.UpdateVariables(expected, tenantVariables)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	actual, err := client.Tenants.GetByID(expected.GetID())
	assert.NoError(t, err)
	AssertEqualTenants(t, expected, actual)
}

func TestTenantServiceGetAll(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	// create 10 test tenants (to be deleted)
	for i := 0; i < 10; i++ {
		tenant := CreateTestTenant(t, client, nil, nil)
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

	expected := CreateTestTenant(t, client, nil, nil)
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

	expected := CreateTestTenant(t, octopusClient, nil, nil)
	defer DeleteTestTenant(t, octopusClient, expected)

	expected.Name = getRandomName()
	expected.Description = getRandomName()

	actual, err := octopusClient.Tenants.Update(expected)
	assert.NoError(t, err)
	assert.NotNil(t, actual)

	AssertEqualTenants(t, expected, actual)
}
