package e2e

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/dashboard"
	"github.com/stretchr/testify/require"
)

// These tests read whatever the target instance happens to hold rather than
// creating a deployment of their own, so anything needing an actual deployment
// skips on an instance that has none.
func skipWithoutItems(t *testing.T, board *dashboard.Dashboard) {
	t.Helper()

	if len(board.Items) == 0 {
		t.Skip("the dashboard is empty; this instance has no deployments to assert on")
	}
}

// projectWithItems returns a project ID that has at least one dashboard item,
// along with its name, so these tests discover their fixtures from the server
// rather than hardcoding IDs that only exist on one instance.
func projectWithItems(t *testing.T, board *dashboard.Dashboard) (string, string) {
	t.Helper()

	for _, item := range board.Items {
		for _, project := range board.Projects {
			if project.GetID() == item.ProjectID {
				return project.GetID(), project.Name
			}
		}
	}

	t.Skip("no project on the dashboard has a deployment; cannot exercise filtering")
	return "", ""
}

func TestDashboardGetDynamicDashboardUnfiltered(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	board, err := client.Dashboards.GetDynamicDashboard(dashboard.DashboardDynamicQuery{})
	require.NoError(t, err)
	require.NotNil(t, board)
	require.False(t, board.IsFiltered)
	skipWithoutItems(t, board)

	// Every item must be resolvable through the reference data, otherwise a
	// caller cannot render it.
	projectIDs := map[string]bool{}
	for _, project := range board.Projects {
		projectIDs[project.GetID()] = true
	}
	environmentIDs := map[string]bool{}
	for _, environment := range board.Environments {
		environmentIDs[environment.GetID()] = true
	}

	for _, item := range board.Items {
		require.True(t, projectIDs[item.ProjectID], "item project %s missing from reference data", item.ProjectID)
		require.True(t, environmentIDs[item.EnvironmentID], "item environment %s missing from reference data", item.EnvironmentID)

		// Id and Links are returned per item and are what callers navigate by.
		require.NotEmpty(t, item.GetID())
		require.NotEmpty(t, item.Links)
		require.NotEmpty(t, item.Links["Self"])
	}
}

func TestDashboardGetDynamicDashboardFiltersByProjectID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	board, err := client.Dashboards.GetDynamicDashboard(dashboard.DashboardDynamicQuery{})
	require.NoError(t, err)
	projectID, _ := projectWithItems(t, board)

	filtered, err := client.Dashboards.GetDynamicDashboard(dashboard.DashboardDynamicQuery{
		Projects: []string{projectID},
	})
	require.NoError(t, err)
	require.True(t, filtered.IsFiltered)
	require.NotEmpty(t, filtered.Items)

	for _, item := range filtered.Items {
		require.Equal(t, projectID, item.ProjectID)
	}
	require.LessOrEqual(t, len(filtered.Items), len(board.Items))
}

// TestDashboardGetDynamicDashboardIgnoresProjectName pins the behaviour the
// query documents: the server matches IDs only, and a name yields an empty
// dashboard rather than an error. If a future server starts accepting names,
// this fails and the documentation needs revisiting.
func TestDashboardGetDynamicDashboardIgnoresProjectName(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	board, err := client.Dashboards.GetDynamicDashboard(dashboard.DashboardDynamicQuery{})
	require.NoError(t, err)
	_, projectName := projectWithItems(t, board)

	byName, err := client.Dashboards.GetDynamicDashboard(dashboard.DashboardDynamicQuery{
		Projects: []string{projectName},
	})
	require.NoError(t, err)
	require.True(t, byName.IsFiltered)
	require.Empty(t, byName.Items, "a project name matched items; the server contract may have changed")
}

func TestDashboardGetDynamicDashboardUnknownIDReturnsEmpty(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	board, err := client.Dashboards.GetDynamicDashboard(dashboard.DashboardDynamicQuery{
		Projects: []string{"Projects-999999999"},
	})
	require.NoError(t, err)
	require.True(t, board.IsFiltered)
	require.Empty(t, board.Items)
}

func TestDashboardGetDynamicDashboardIncludePrevious(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	current, err := client.Dashboards.GetDynamicDashboard(dashboard.DashboardDynamicQuery{})
	require.NoError(t, err)

	withPrevious, err := client.Dashboards.GetDynamicDashboard(dashboard.DashboardDynamicQuery{
		IncludePrevious: true,
	})
	require.NoError(t, err)

	// Previous deployments arrive inside Items flagged with IsPrevious, not in a
	// separate collection.
	require.GreaterOrEqual(t, len(withPrevious.Items), len(current.Items))
	for _, item := range current.Items {
		require.False(t, item.IsPrevious)
	}
	if len(withPrevious.Items) > len(current.Items) {
		previous := 0
		for _, item := range withPrevious.Items {
			if item.IsPrevious {
				previous++
			}
		}
		require.NotZero(t, previous)
	}
}

// TestDashboardDynamicDashboardModelsEveryServerField decodes the live response
// with unknown fields disallowed, so a field the server adds or renames fails
// here instead of silently arriving as a zero value.
func TestDashboardDynamicDashboardModelsEveryServerField(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	path, err := url.Parse("/api/" + client.GetSpaceID() + "/dashboard/dynamic?includePrevious=true")
	require.NoError(t, err)

	response, err := client.HttpSession().DoRawRequest(&http.Request{
		Method: http.MethodGet,
		URL:    path,
		Header: make(http.Header),
	})
	require.NoError(t, err)
	defer response.Body.Close()
	require.Equal(t, http.StatusOK, response.StatusCode)

	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields()

	board := &dashboard.Dashboard{}
	require.NoError(t, decoder.Decode(board), "the server returned a field the SDK does not model")

	// The decode above is the assertion. On an instance with no deployments it
	// only covers the envelope and whatever reference data exists, so log the
	// coverage rather than failing.
	if len(board.Items) == 0 {
		t.Log("dashboard is empty; item fields were not covered by this decode")
	}
}

func TestDashboardGetDynamicDashboardFiltersByEnvironmentID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	board, err := client.Dashboards.GetDynamicDashboard(dashboard.DashboardDynamicQuery{})
	require.NoError(t, err)
	skipWithoutItems(t, board)

	environmentID := board.Items[0].EnvironmentID
	filtered, err := client.Dashboards.GetDynamicDashboard(dashboard.DashboardDynamicQuery{
		Environments: []string{environmentID},
	})
	require.NoError(t, err)
	require.True(t, filtered.IsFiltered)

	for _, item := range filtered.Items {
		require.Equal(t, environmentID, item.EnvironmentID)
	}
}

func TestDashboardGetDashboardUnfiltered(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	board, err := client.Dashboards.GetDashboard(dashboard.DashboardQuery{})
	require.NoError(t, err)
	require.NotNil(t, board)
	require.False(t, board.IsFiltered)
	skipWithoutItems(t, board)

	projectIDs := map[string]bool{}
	for _, project := range board.Projects {
		projectIDs[project.GetID()] = true
	}

	for _, item := range board.Items {
		require.True(t, projectIDs[item.ProjectID], "item project %s missing from reference data", item.ProjectID)
		require.NotEmpty(t, item.GetID())
		require.NotEmpty(t, item.Links["Self"])
	}
}

func TestDashboardGetDashboardFiltersByProjectID(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	board, err := client.Dashboards.GetDashboard(dashboard.DashboardQuery{})
	require.NoError(t, err)
	projectID, projectName := projectWithItems(t, board)

	filtered, err := client.Dashboards.GetDashboard(dashboard.DashboardQuery{ProjectID: projectID})
	require.NoError(t, err)
	require.True(t, filtered.IsFiltered)
	require.NotEmpty(t, filtered.Items)
	for _, item := range filtered.Items {
		require.Equal(t, projectID, item.ProjectID)
	}

	// As with the dynamic dashboard, the server matches IDs only.
	byName, err := client.Dashboards.GetDashboard(dashboard.DashboardQuery{ProjectID: projectName})
	require.NoError(t, err)
	require.Empty(t, byName.Items, "a project name matched items; the server contract may have changed")
}

// TestDashboardGetDashboardIncludeLatest exercises the parameter that carried
// only a url tag before #442 and so never reached the server. Neither parameter
// narrows the result, so what is assertable is the shape the name promises: the
// highest version per project and environment, meaning one item per cell.
func TestDashboardGetDashboardIncludeLatest(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	board, err := client.Dashboards.GetDashboard(dashboard.DashboardQuery{IncludeLatest: true, ShowAll: true})
	require.NoError(t, err)
	require.NotNil(t, board)
	require.False(t, board.IsFiltered)
	skipWithoutItems(t, board)

	projectIDs := map[string]bool{}
	for _, project := range board.Projects {
		projectIDs[project.GetID()] = true
	}
	environmentIDs := map[string]bool{}
	for _, environment := range board.Environments {
		environmentIDs[environment.GetID()] = true
	}
	tenantIDs := map[string]bool{}
	for _, tenant := range board.Tenants {
		tenantIDs[tenant.GetID()] = true
	}

	cells := map[string]bool{}
	for _, item := range board.Items {
		require.True(t, projectIDs[item.ProjectID], "item project %s missing from reference data", item.ProjectID)
		require.True(t, environmentIDs[item.EnvironmentID], "item environment %s missing from reference data", item.EnvironmentID)
		if item.TenantID != "" {
			require.True(t, tenantIDs[item.TenantID], "item tenant %s missing from reference data", item.TenantID)
		}

		// Enough of a cell to render: which release is deployed, how it went,
		// and a link to the deployment behind it.
		require.NotEmpty(t, item.ReleaseVersion)
		require.NotEmpty(t, item.State)
		require.NotEmpty(t, item.TaskID)
		require.NotEmpty(t, item.Links["Self"])
		if item.IsCompleted {
			require.NotNil(t, item.CompletedTime)
		}

		// This endpoint reports what is deployed now; previous deployments are
		// only reachable through the dynamic dashboard's IncludePrevious.
		require.True(t, item.IsCurrent)
		require.False(t, item.IsPrevious)

		// highestLatestVersionPerProjectAndEnvironment resolves each cell to one
		// release, so a repeated cell means it did not do what its name says.
		cell := fmt.Sprintf("%s/%s/%s", item.ProjectID, item.EnvironmentID, item.TenantID)
		require.False(t, cells[cell], "two items for cell %s", cell)
		cells[cell] = true
	}
}

func TestDashboardGetDashboardFiltersByTenant(t *testing.T) {
	client := getOctopusClient()
	require.NotNil(t, client)

	board, err := client.Dashboards.GetDashboard(dashboard.DashboardQuery{})
	require.NoError(t, err)

	tenantID := ""
	for _, item := range board.Items {
		if item.TenantID != "" {
			tenantID = item.TenantID
			break
		}
	}
	if tenantID == "" {
		t.Skip("no tenanted deployment on the dashboard")
	}

	filtered, err := client.Dashboards.GetDashboard(dashboard.DashboardQuery{SelectedTenants: []string{tenantID}})
	require.NoError(t, err)
	require.True(t, filtered.IsFiltered)
	require.NotEmpty(t, filtered.Items)
	for _, item := range filtered.Items {
		require.Equal(t, tenantID, item.TenantID)
	}
}
