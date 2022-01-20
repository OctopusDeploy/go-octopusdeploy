package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/stretchr/testify/require"
)

func TestTenantVariablesServiceGetAll(t *testing.T) {
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

	environment := CreateTestEnvironment(t, client)
	defer DeleteTestEnvironment(t, client, environment)

	tenant := CreateTestTenant(t, client, project, environment)
	defer DeleteTestTenant(t, client, tenant)

	tenantVariables := octopusdeploy.NewTenantVariables(tenant.GetID())
	require.NotNil(t, tenantVariables)

	tenantVariables, err := client.Tenants.UpdateVariables(tenant, tenantVariables)
	require.NoError(t, err)
	require.NotNil(t, tenantVariables)

	allTenantVariables, err := client.TenantVariables.GetAll()
	require.NoError(t, err)
	require.NotNil(t, allTenantVariables)
	require.True(t, len(allTenantVariables) > 0)

	for k, v := range allTenantVariables {
		require.NotNil(t, k)
		require.NotNil(t, v)
	}
}
