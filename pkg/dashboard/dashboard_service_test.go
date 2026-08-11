package dashboard

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createDashboardService(t *testing.T) *DashboardService {
	service := NewDashboardService(nil, constants.TestURIDashboard, constants.TestURIDashboardDynamic)
	services.NewServiceTests(t, service, constants.TestURIDashboard, constants.ServiceDashboardService)
	return service
}

func TestNewDashboardService(t *testing.T) {
	service := NewDashboardService(nil, constants.TestURIDashboard, constants.TestURIDashboardDynamic)
	require.NotNil(t, service)
	require.Equal(t, constants.TestURIDashboardDynamic, service.dashboardDynamicPath)
}

func TestDashboardServiceGetDynamicDashboardPath(t *testing.T) {
	service := createDashboardService(t)

	tests := []struct {
		name  string
		query DashboardDynamicQuery
		// expected query parameters, checked individually because the template
		// does not guarantee ordering
		expected map[string][]string
	}{
		{
			name:     "empty query returns the unfiltered dashboard",
			query:    DashboardDynamicQuery{},
			expected: map[string][]string{},
		},
		{
			name:     "single environment",
			query:    DashboardDynamicQuery{Environments: []string{"Environments-1"}},
			expected: map[string][]string{"environments": {"Environments-1"}},
		},
		{
			name:     "multiple projects",
			query:    DashboardDynamicQuery{Projects: []string{"Projects-1", "Projects-2"}},
			expected: map[string][]string{"projects": {"Projects-1,Projects-2"}},
		},
		{
			name:     "include previous",
			query:    DashboardDynamicQuery{IncludePrevious: true},
			expected: map[string][]string{"includePrevious": {"true"}},
		},
		{
			// IncludePrevious is omitempty, so false must not be sent; the
			// server treats its presence as opt-in.
			name:     "include previous false is omitted",
			query:    DashboardDynamicQuery{IncludePrevious: false},
			expected: map[string][]string{},
		},
		{
			name: "projects and environments together",
			query: DashboardDynamicQuery{
				Projects:        []string{"Projects-1"},
				Environments:    []string{"Environments-1", "Environments-2"},
				IncludePrevious: true,
			},
			expected: map[string][]string{
				"projects":        {"Projects-1"},
				"environments":    {"Environments-1,Environments-2"},
				"includePrevious": {"true"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path, err := service.getDynamicDashboardPath(test.query)
			require.NoError(t, err)

			parsed, err := url.Parse(path)
			require.NoError(t, err)
			assert.Equal(t, "/api/Spaces-1/dashboard/dynamic", parsed.Path)
			assert.Equal(t, url.Values(test.expected), parsed.Query())
		})
	}
}

func TestDashboardServiceGetDynamicDashboardPathRequiresLink(t *testing.T) {
	service := NewDashboardService(nil, constants.TestURIDashboard, "")

	path, err := service.getDynamicDashboardPath(DashboardDynamicQuery{})
	require.Equal(t, internal.CreateInvalidParameterError(constants.OperationGet, "dashboardDynamicPath"), err)
	require.Empty(t, path)
}

func TestDashboardServiceGetDynamicDashboardWithoutLinkReturnsError(t *testing.T) {
	service := NewDashboardService(nil, constants.TestURIDashboard, "")

	dashboard, err := service.GetDynamicDashboard(DashboardDynamicQuery{})
	require.Error(t, err)
	require.Nil(t, dashboard)
}

// TestDashboardDeserialization pins the response shape against a capture from a
// real server, so a field being renamed or re-nested shows up as a test failure
// rather than a silently empty column.
func TestDashboardDeserialization(t *testing.T) {
	const payload = `{
	  "Projects": [
	    {
	      "Id": "Projects-81",
	      "Name": "NJ Todo List K8s",
	      "IsDisabled": false,
	      "Slug": "nj-todo-list-k8s",
	      "ProjectGroupId": "ProjectGroups-1",
	      "EnvironmentIds": ["Environments-3", "Environments-1", "Environments-2"],
	      "TenantedDeploymentMode": "Untenanted",
	      "CanPerformUntenantedDeployment": true,
	      "Links": { "Self": "/api/Spaces-1/projects/Projects-81" }
	    }
	  ],
	  "ProjectGroups": [
	    {
	      "Id": "ProjectGroups-1",
	      "Name": "Default Project Group",
	      "EnvironmentIds": ["Environments-3", "Environments-1", "Environments-2"]
	    }
	  ],
	  "Environments": [
	    { "Id": "Environments-3", "Name": "Development" },
	    { "Id": "Environments-1", "Name": "Staging" },
	    { "Id": "Environments-2", "Name": "Production" }
	  ],
	  "Tenants": [
	    {
	      "Id": "Tenants-1",
	      "Name": "Aus-East",
	      "TenantTags": ["Region/Aus-East"],
	      "ProjectEnvironments": { "Projects-182": ["Environments-3"] },
	      "IsDisabled": false
	    }
	  ],
	  "Items": [
	    {
	      "Id": "Deployments-387",
	      "ProjectId": "Projects-81",
	      "EnvironmentId": "Environments-2",
	      "ReleaseId": "Releases-302",
	      "DeploymentId": "Deployments-387",
	      "TaskId": "ServerTasks-12388",
	      "TenantId": null,
	      "ChannelId": "Channels-103",
	      "ReleaseVersion": "2.16.0",
	      "CompletedTime": "2026-08-04T07:39:42.254+00:00",
	      "State": "Success",
	      "HasWarningsOrErrors": false,
	      "ErrorMessage": "",
	      "Duration": "59 seconds",
	      "IsCurrent": true,
	      "IsPrevious": false,
	      "IsCompleted": true,
	      "HasPendingInterruptions": false,
	      "HasPendingPreconditions": false,
	      "PendingInterruptionTypes": [],
	      "PendingPreconditionTypes": [],
	      "Links": {
	        "Self": "/api/Spaces-1/deployments/Deployments-387",
	        "Release": "/api/Spaces-1/releases/Releases-302",
	        "Task": "/api/tasks/ServerTasks-12388"
	      }
	    }
	  ],
	  "ProjectLimit": null,
	  "IsFiltered": false
	}`

	dashboard := &Dashboard{}
	require.NoError(t, json.Unmarshal([]byte(payload), dashboard))

	require.Len(t, dashboard.Items, 1)
	item := dashboard.Items[0]
	assert.Equal(t, "Projects-81", item.ProjectID)
	assert.Equal(t, "Environments-2", item.EnvironmentID)
	assert.Equal(t, "Releases-302", item.ReleaseID)
	assert.Equal(t, "2.16.0", item.ReleaseVersion)
	assert.Equal(t, "ServerTasks-12388", item.TaskID)
	assert.Equal(t, "Success", item.State)
	assert.Empty(t, item.TenantID)
	assert.True(t, item.IsCurrent)
	assert.False(t, item.IsPrevious)
	require.NotNil(t, item.CompletedTime)
	assert.Equal(t, 2026, item.CompletedTime.Year())

	// Each item carries its own Id and Links. Without them a caller cannot reach
	// the deployment, release or task without rebuilding paths from the IDs.
	assert.Equal(t, "Deployments-387", item.GetID())
	assert.Equal(t, "/api/Spaces-1/deployments/Deployments-387", item.Links["Self"])
	assert.Equal(t, "/api/tasks/ServerTasks-12388", item.Links["Task"])
	assert.False(t, item.HasPendingInterruptions)
	assert.False(t, item.HasPendingPreconditions)
	assert.Empty(t, item.PendingInterruptionTypes)
	assert.Empty(t, item.PendingPreconditionTypes)

	// The reference data is what makes the item IDs printable, so check each
	// lookup resolves rather than only that it parsed.
	require.Len(t, dashboard.Projects, 1)
	assert.Equal(t, "Projects-81", dashboard.Projects[0].GetID())
	assert.Equal(t, "NJ Todo List K8s", dashboard.Projects[0].Name)
	assert.Equal(t, "ProjectGroups-1", dashboard.Projects[0].ProjectGroupID)
	assert.Len(t, dashboard.Projects[0].EnvironmentIDs, 3)
	assert.Equal(t, "Untenanted", dashboard.Projects[0].TenantedDeploymentMode)

	require.Len(t, dashboard.ProjectGroups, 1)
	assert.Equal(t, "Default Project Group", dashboard.ProjectGroups[0].Name)

	// Environment order is the dashboard's own ordering, not alphabetical, and
	// callers rely on it to print columns Development -> Staging -> Production.
	require.Len(t, dashboard.Environments, 3)
	assert.Equal(t, []string{"Development", "Staging", "Production"},
		[]string{dashboard.Environments[0].Name, dashboard.Environments[1].Name, dashboard.Environments[2].Name})

	require.Len(t, dashboard.Tenants, 1)
	assert.Equal(t, "Tenants-1", dashboard.Tenants[0].GetID())
	assert.Equal(t, "Aus-East", dashboard.Tenants[0].Name)
	assert.Equal(t, []string{"Region/Aus-East"}, dashboard.Tenants[0].TenantTags)
	assert.Equal(t, []string{"Environments-3"}, dashboard.Tenants[0].ProjectEnvironments["Projects-182"])

	// A null ProjectLimit means uncapped, which callers must distinguish from a
	// limit of zero.
	assert.Nil(t, dashboard.ProjectLimit)
	assert.False(t, dashboard.IsFiltered)
}

func TestDashboardDeserializationWithProjectLimit(t *testing.T) {
	dashboard := &Dashboard{}
	require.NoError(t, json.Unmarshal([]byte(`{"ProjectLimit": 50, "IsFiltered": true}`), dashboard))

	require.NotNil(t, dashboard.ProjectLimit)
	assert.Equal(t, 50, *dashboard.ProjectLimit)
	assert.True(t, dashboard.IsFiltered)
}

func TestDashboardDeserializationTenantedItem(t *testing.T) {
	dashboard := &Dashboard{}
	require.NoError(t, json.Unmarshal([]byte(`{
	  "Items": [
	    {
	      "ProjectId": "Projects-182",
	      "EnvironmentId": "Environments-3",
	      "TenantId": "Tenants-1",
	      "ReleaseVersion": "1.0.0",
	      "State": "Success",
	      "IsCurrent": true
	    }
	  ]
	}`), dashboard))

	require.Len(t, dashboard.Items, 1)
	assert.Equal(t, "Tenants-1", dashboard.Items[0].TenantID)
	assert.Equal(t, "1.0.0", dashboard.Items[0].ReleaseVersion)
}
