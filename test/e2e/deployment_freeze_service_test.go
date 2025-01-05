package e2e

import (
	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deploymentfreezes"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/tenants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type testResources struct {
	environment *environments.Environment
	project     *projects.Project
	tenant      *tenants.Tenant
	cleanup     func()
}

func createTestResources(t *testing.T, client *client.Client) *testResources {
	env := CreateTestEnvironment_NewClient(t, client)
	require.NotNil(t, env)

	lifecycle := CreateTestLifecycle_NewClient(t, client)
	require.NotNil(t, lifecycle)

	projectGroup := CreateTestProjectGroup(t, client)
	require.NotNil(t, projectGroup)

	space := GetDefaultSpace(t, client)
	require.NotNil(t, space)

	project := CreateTestProject_NewClient(t, client, space, lifecycle, projectGroup)
	require.NotNil(t, project)

	tenant := CreateTestTenant_NewClient(t, client, project, env)
	require.NotNil(t, tenant)

	cleanup := func() {
		DeleteTestEnvironment_NewClient(t, client, env)
		DeleteTestProject_NewClient(t, client, project)
		DeleteTestLifecycle_NewClient(t, client, lifecycle)
		DeleteTestProjectGroup(t, client, projectGroup)
		DeleteTestTenant_NewClient(t, client, tenant)
	}

	return &testResources{
		environment: env,
		project:     project,
		tenant:      tenant,
		cleanup:     cleanup,
	}
}

func getTestDeploymentFreeze() *deploymentfreezes.DeploymentFreeze {
	name := internal.GetRandomName()
	startTime := time.Now().Add(time.Hour).UTC().Truncate(time.Second)
	endTime := startTime.Add(time.Hour * 24).UTC().Truncate(time.Second)

	return &deploymentfreezes.DeploymentFreeze{
		Name:                          name,
		Start:                         &startTime,
		End:                           &endTime,
		ProjectEnvironmentScope:       make(map[string][]string),
		TenantProjectEnvironmentScope: []deploymentfreezes.TenantProjectEnvironment{},
	}
}

func createTestDeploymentFreeze(t *testing.T, client *client.Client) *deploymentfreezes.DeploymentFreeze {
	freeze := getTestDeploymentFreeze()
	require.NotNil(t, freeze)

	createdFreeze, err := deploymentfreezes.Add(client, freeze)
	require.NoError(t, err)
	require.NotNil(t, createdFreeze)
	require.NotEmpty(t, createdFreeze.GetID())

	return createdFreeze
}

func cleanDeploymentFreeze(t *testing.T, client *client.Client, freezeID string) {
	freeze, err := deploymentfreezes.GetById(client, freezeID)
	if err != nil || freeze.GetID() == "" {
		return
	}
	err = deploymentfreezes.Delete(client, freeze)
	assert.NoError(t, err)
}

func TestDeploymentFreezeCRUD(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	resources := createTestResources(t, client)
	defer resources.cleanup()

	t.Run("Add and Delete", func(t *testing.T) {
		freeze := createTestDeploymentFreeze(t, client)
		require.NotNil(t, freeze)
		defer cleanDeploymentFreeze(t, client, freeze.GetID())

		err := deploymentfreezes.Delete(client, freeze)
		assert.NoError(t, err)

		_, err = deploymentfreezes.GetById(client, freeze.GetID())
		assert.Error(t, err)
	})

	t.Run("Update", func(t *testing.T) {
		freeze := createTestDeploymentFreeze(t, client)
		require.NotNil(t, freeze)
		defer cleanDeploymentFreeze(t, client, freeze.GetID())

		newName := internal.GetRandomName()
		freeze.Name = newName
		freeze.ProjectEnvironmentScope[resources.project.GetID()] = []string{resources.environment.GetID()}

		updatedFreeze, err := deploymentfreezes.Update(client, freeze)
		require.NoError(t, err)
		require.Equal(t, newName, updatedFreeze.Name)
		require.Contains(t, updatedFreeze.ProjectEnvironmentScope, resources.project.GetID())
	})

	t.Run("GetAll", func(t *testing.T) {
		freeze1 := createTestDeploymentFreeze(t, client)
		require.NotNil(t, freeze1)
		defer cleanDeploymentFreeze(t, client, freeze1.GetID())

		freeze2 := createTestDeploymentFreeze(t, client)
		require.NotNil(t, freeze2)
		defer cleanDeploymentFreeze(t, client, freeze2.GetID())

		allFreezes, err := deploymentfreezes.GetAll(client)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, allFreezes.Count, 2)

		freezeIds := map[string]bool{
			freeze1.GetID(): false,
			freeze2.GetID(): false,
		}
		for _, freeze := range allFreezes.DeploymentFreezes {
			if _, exists := freezeIds[freeze.GetID()]; exists {
				freezeIds[freeze.GetID()] = true
			}
		}
		for id, found := range freezeIds {
			assert.True(t, found, "Freeze %s not found in GetAll results", id)
		}
	})

	t.Run("Get Non-existent", func(t *testing.T) {
		_, err := deploymentfreezes.GetById(client, internal.GetRandomName())
		assert.Error(t, err)
	})
}

func TestDeploymentFreezeRecurringSchedules(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	resources := createTestResources(t, client)
	defer resources.cleanup()

	testCases := []struct {
		name     string
		schedule *deploymentfreezes.RecurringSchedule
		validate func(*testing.T, *deploymentfreezes.DeploymentFreeze)
	}{
		{
			name: "Daily Schedule",
			schedule: &deploymentfreezes.RecurringSchedule{
				Type:                deploymentfreezes.Daily,
				Unit:                2,
				EndType:             deploymentfreezes.AfterOccurrences,
				EndAfterOccurrences: 5,
			},
			validate: func(t *testing.T, freeze *deploymentfreezes.DeploymentFreeze) {
				require.Equal(t, deploymentfreezes.Daily, freeze.RecurringSchedule.Type)
				require.Equal(t, deploymentfreezes.AfterOccurrences, freeze.RecurringSchedule.EndType)
				require.Equal(t, 5, freeze.RecurringSchedule.EndAfterOccurrences)
				require.Equal(t, 2, freeze.RecurringSchedule.Unit)
			},
		},
		{
			name: "Weekly Schedule",
			schedule: &deploymentfreezes.RecurringSchedule{
				Type:                deploymentfreezes.Weekly,
				Unit:                24,
				EndType:             deploymentfreezes.AfterOccurrences,
				EndAfterOccurrences: 5,
				DaysOfWeek:          []string{"Monday", "Wednesday", "Friday"},
			},
			validate: func(t *testing.T, freeze *deploymentfreezes.DeploymentFreeze) {
				require.Equal(t, deploymentfreezes.Weekly, freeze.RecurringSchedule.Type)
				require.Equal(t, []string{"Monday", "Wednesday", "Friday"}, freeze.RecurringSchedule.DaysOfWeek)
			},
		},
		{
			name: "Monthly Schedule",
			schedule: &deploymentfreezes.RecurringSchedule{
				Type:                deploymentfreezes.Monthly,
				Unit:                1,
				EndType:             deploymentfreezes.Never,
				MonthlyScheduleType: "DayOfMonth",
				DayOfWeek:           "Thursday",
				DayNumberOfMonth:    "1",
			},
			validate: func(t *testing.T, freeze *deploymentfreezes.DeploymentFreeze) {
				require.Equal(t, deploymentfreezes.Monthly, freeze.RecurringSchedule.Type)
				require.Equal(t, "DayOfMonth", freeze.RecurringSchedule.MonthlyScheduleType)
				require.Equal(t, "Thursday", freeze.RecurringSchedule.DayOfWeek)
			},
		},
		{
			name: "Yearly Schedule",
			schedule: &deploymentfreezes.RecurringSchedule{
				Type:      deploymentfreezes.Annually,
				Unit:      1,
				EndType:   deploymentfreezes.OnDate,
				EndOnDate: ptr(time.Now().AddDate(1, 0, 0).UTC().Truncate(time.Second)),
			},
			validate: func(t *testing.T, freeze *deploymentfreezes.DeploymentFreeze) {
				require.Equal(t, deploymentfreezes.Annually, freeze.RecurringSchedule.Type)
				require.Equal(t, 1, freeze.RecurringSchedule.Unit)
				require.Equal(t, deploymentfreezes.OnDate, freeze.RecurringSchedule.EndType)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			freeze := getTestDeploymentFreeze()
			freeze.ProjectEnvironmentScope[resources.project.GetID()] = []string{resources.environment.GetID()}
			freeze.RecurringSchedule = tc.schedule

			createdFreeze, err := deploymentfreezes.Add(client, freeze)
			require.NoError(t, err)
			defer cleanDeploymentFreeze(t, client, createdFreeze.GetID())

			tc.validate(t, createdFreeze)
		})
	}
}

func TestDeploymentFreezeScopeQueries(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	resources := createTestResources(t, client)
	defer resources.cleanup()

	t.Run("Tenant and Project Query", func(t *testing.T) {
		freeze := getTestDeploymentFreeze()
		freeze.ProjectEnvironmentScope[resources.project.GetID()] = []string{resources.environment.GetID()}
		freeze.TenantProjectEnvironmentScope = []deploymentfreezes.TenantProjectEnvironment{
			{
				TenantId:      resources.tenant.GetID(),
				ProjectId:     resources.project.GetID(),
				EnvironmentId: resources.environment.GetID(),
			},
		}

		createdFreeze, err := deploymentfreezes.Add(client, freeze)
		require.NoError(t, err)
		defer cleanDeploymentFreeze(t, client, createdFreeze.GetID())

		query := deploymentfreezes.DeploymentFreezeQuery{
			TenantIds:  []string{resources.tenant.GetID()},
			ProjectIds: []string{resources.project.GetID()},
			Skip:       0,
			Take:       30,
		}

		result, err := deploymentfreezes.Get(client, query)
		require.NoError(t, err)
		require.Greater(t, result.Count, 0)

		var found bool
		for _, f := range result.DeploymentFreezes {
			if f.GetID() == createdFreeze.GetID() {
				found = true
				require.Contains(t, f.ProjectEnvironmentScope, resources.project.GetID())
				require.Contains(t, f.TenantProjectEnvironmentScope, deploymentfreezes.TenantProjectEnvironment{
					TenantId:      resources.tenant.GetID(),
					ProjectId:     resources.project.GetID(),
					EnvironmentId: resources.environment.GetID(),
				})
				break
			}
		}
		assert.True(t, found)
	})
}

func ptr[T any](v T) *T {
	return &v
}
