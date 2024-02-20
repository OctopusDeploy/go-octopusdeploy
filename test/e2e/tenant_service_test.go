package e2e

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/tenants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/variables"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualTenants(t *testing.T, expected *tenants.Tenant, actual *tenants.Tenant) {
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

func CreateTestTenant(t *testing.T, octopusClient *client.Client, project *projects.Project, environment *environments.Environment) *tenants.Tenant {
	if octopusClient == nil {
		octopusClient = getOctopusClient()
	}
	require.NotNil(t, octopusClient)

	name := internal.GetRandomName()

	tenant := tenants.NewTenant(name)
	tenant.Description = internal.GetRandomName()

	if project != nil {
		tenant.ProjectEnvironments[project.ID] = []string{environment.ID}
	}

	createdTenant, err := octopusClient.Tenants.Add(tenant)
	require.NoError(t, err)
	require.NotNil(t, createdTenant)
	require.NotEmpty(t, createdTenant.GetID())

	return createdTenant
}

func DeleteTestTenant(t *testing.T, client *client.Client, tenant *tenants.Tenant) {
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

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup(t, client, projectGroup)

	project := CreateTestProject(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject(t, client, project)

	variable := CreateTestVariable(t, project.GetID(), internal.GetRandomName())
	require.NotNil(t, variable)

	actionTemplateParameter := CreateActionTemplateParameter()
	require.NotNil(t, actionTemplateParameter)
	project.Templates = append(project.Templates, actionTemplateParameter)
	project, err := client.Projects.Update(project)
	require.NotNil(t, project)
	require.NoError(t, err)

	libraryVariableSet := CreateLibraryVariableSet(t, client)
	require.NotNil(t, libraryVariableSet)
	defer DeleteLibraryVariableSet(t, client, libraryVariableSet)

	environment := CreateTestEnvironment(t, client)
	defer DeleteTestEnvironment(t, client, environment)

	tenant := CreateTestTenant(t, client, project, environment)
	defer DeleteTestTenant(t, client, tenant)

	missingVariablesQuery := variables.MissingVariablesQuery{}

	missingVariables, err := client.Tenants.GetMissingVariables(missingVariablesQuery)
	require.NoError(t, err)
	require.NotNil(t, missingVariables)

	tenantVariables := variables.NewTenantVariables(tenant.GetID())
	require.NotNil(t, tenantVariables)

	tenantVariables, err = client.Tenants.UpdateVariables(tenant, tenantVariables)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	tenantVariables, err = client.Tenants.GetVariables(tenant)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	propertyValue := core.NewPropertyValue("#{binding}", false)

	tenantVariables.ProjectVariables[project.GetID()].Variables[environment.GetID()][project.Templates[0].GetID()] = propertyValue
	tenantVariables, err = client.Tenants.UpdateVariables(tenant, tenantVariables)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	tenantVariables, err = client.Tenants.GetVariables(tenant)
	require.Equal(t, propertyValue.Value, tenantVariables.ProjectVariables[project.GetID()].Variables[environment.GetID()][project.Templates[0].GetID()].Value)

	tenantVariables, err = client.Tenants.GetVariables(tenant)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	actual, err := client.Tenants.GetByID(tenant.GetID())
	assert.NoError(t, err)
	AssertEqualTenants(t, tenant, actual)
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

	expected.Name = internal.GetRandomName()
	expected.Description = internal.GetRandomName()

	actual, err := octopusClient.Tenants.Update(expected)
	assert.NoError(t, err)
	assert.NotNil(t, actual)

	AssertEqualTenants(t, expected, actual)
}

// === NEW ===

func CreateTestTenant_NewClient(t *testing.T, client *client.Client, project *projects.Project, environment *environments.Environment) *tenants.Tenant {
	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	name := internal.GetRandomName()

	tenant := tenants.NewTenant(name)
	tenant.Description = internal.GetRandomName()

	if project != nil {
		tenant.ProjectEnvironments[project.ID] = []string{environment.ID}
	}

	createdTenant, err := tenants.Add(client, tenant)
	require.NoError(t, err)
	require.NotNil(t, createdTenant)
	require.NotEmpty(t, createdTenant.GetID())

	return createdTenant
}

func DeleteTestTenant_NewClient(t *testing.T, client *client.Client, tenant *tenants.Tenant) {
	require.NotNil(t, tenant)

	if client == nil {
		client = getOctopusClient()
	}
	require.NotNil(t, client)

	err := tenants.DeleteByID(client, tenant.SpaceID, tenant.GetID())
	assert.NoError(t, err)

	// verify the delete operation was successful
	deletedTenant, err := tenants.GetByID(client, tenant.SpaceID, tenant.GetID())
	assert.Error(t, err)
	assert.Nil(t, deletedTenant)
}

// TODO: still some calls to old client
func TestTenantAddGetAndDelete_NewClient(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	lifecycle := CreateTestLifecycle_NewClient(t, client)
	require.NotNil(t, lifecycle)
	defer DeleteTestLifecycle_NewClient(t, client, lifecycle)

	projectGroup := CreateTestProjectGroup_NewClient(t, client)
	require.NotNil(t, projectGroup)
	defer DeleteTestProjectGroup_NewClient(t, client, projectGroup)

	project := CreateTestProject_NewClient(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)
	defer DeleteTestProject_NewClient(t, client, project)

	variable := CreateTestVariable_NewClient(t, project.GetID(), internal.GetRandomName())
	require.NotNil(t, variable)

	actionTemplateParameter := CreateActionTemplateParameter()
	require.NotNil(t, actionTemplateParameter)
	project.Templates = append(project.Templates, actionTemplateParameter)
	project, err := projects.Update(client, project)
	require.NotNil(t, project)
	require.NoError(t, err)

	libraryVariableSet := CreateLibraryVariableSet_NewClient(t, client)
	require.NotNil(t, libraryVariableSet)
	defer DeleteLibraryVariableSet_NewClient(t, client, libraryVariableSet)

	environment := CreateTestEnvironment_NewClient(t, client)
	defer DeleteTestEnvironment_NewClient(t, client, environment)

	tenant := CreateTestTenant_NewClient(t, client, project, environment)
	defer DeleteTestTenant_NewClient(t, client, tenant)

	missingVariablesQuery := variables.MissingVariablesQuery{}

	missingVariables, err := client.Tenants.GetMissingVariables(missingVariablesQuery)
	require.NoError(t, err)
	require.NotNil(t, missingVariables)

	tenantVariables := variables.NewTenantVariables(tenant.GetID())
	require.NotNil(t, tenantVariables)

	tenantVariables, err = client.Tenants.UpdateVariables(tenant, tenantVariables)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	tenantVariables, err = client.Tenants.GetVariables(tenant)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	propertyValue := core.NewPropertyValue("#{binding}", false)

	tenantVariables.ProjectVariables[project.GetID()].Variables[environment.GetID()][project.Templates[0].GetID()] = propertyValue
	tenantVariables, err = client.Tenants.UpdateVariables(tenant, tenantVariables)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	tenantVariables, err = client.Tenants.GetVariables(tenant)
	require.Equal(t, propertyValue.Value, tenantVariables.ProjectVariables[project.GetID()].Variables[environment.GetID()][project.Templates[0].GetID()].Value)

	tenantVariables, err = client.Tenants.GetVariables(tenant)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	actual, err := tenants.GetByID(client, tenant.SpaceID, tenant.GetID())
	assert.NoError(t, err)
	AssertEqualTenants(t, tenant, actual)
}
