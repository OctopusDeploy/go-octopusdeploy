package dashboard

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/constants"
	"github.com/dghubble/sling"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The path tests expand the template without calling the method; this covers
// the URL GetDashboard actually requests, and the response deserialising.
func TestDashboardServiceGetDashboardRoundTrip(t *testing.T) {
	const payload = `{
	  "Projects": [{ "Id": "Projects-81", "Name": "NJ Todo List K8s" }],
	  "Environments": [{ "Id": "Environments-2", "Name": "Production" }],
	  "Items": [
	    {
	      "Id": "Deployments-387",
	      "ProjectId": "Projects-81",
	      "EnvironmentId": "Environments-2",
	      "ReleaseVersion": "2.16.0",
	      "State": "Success",
	      "IsCurrent": true,
	      "Links": { "Self": "/api/Spaces-1/deployments/Deployments-387" }
	    }
	  ],
	  "ProjectLimit": null,
	  "IsFiltered": true
	}`

	var requested string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requested = r.URL.RequestURI()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(payload))
	}))
	defer server.Close()

	base := sling.New().Client(server.Client()).Base(server.URL+"/").Set("Accept", "application/json")
	service := NewDashboardService(base, constants.TestURIDashboard, constants.TestURIDashboardDynamic)

	board, err := service.GetDashboard(DashboardQuery{ProjectID: "Projects-81", IncludeLatest: true})
	require.NoError(t, err)

	parsed, err := url.Parse(requested)
	require.NoError(t, err)
	assert.Equal(t, "/api/Spaces-1/dashboard", parsed.Path)
	assert.Equal(t, url.Values{
		"projectId": {"Projects-81"},
		"highestLatestVersionPerProjectAndEnvironment": {"true"},
	}, parsed.Query())

	require.Len(t, board.Items, 1)
	assert.Equal(t, "Deployments-387", board.Items[0].GetID())
	assert.Equal(t, "Projects-81", board.Items[0].ProjectID)
	assert.Equal(t, "/api/Spaces-1/deployments/Deployments-387", board.Items[0].Links["Self"])
	assert.True(t, board.IsFiltered)
}
