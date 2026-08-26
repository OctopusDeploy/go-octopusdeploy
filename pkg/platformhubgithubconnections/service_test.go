package platformhubgithubconnections

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/githubconnections"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// newTestClient returns a client pointed at a server that records the requested URI and
// replies with the given payloads, one per request in order.
func newTestClient(t *testing.T, requested *[]string, payloads ...string) newclient.Client {
	t.Helper()

	call := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*requested = append(*requested, r.URL.RequestURI())
		payload := payloads[len(payloads)-1]
		if call < len(payloads) {
			payload = payloads[call]
		}
		call++
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(payload))
	}))
	t.Cleanup(server.Close)

	baseURL, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	return newclient.NewClient(&newclient.HttpSession{HttpClient: server.Client(), BaseURL: baseURL})
}

func TestList(t *testing.T) {
	const payload = `{
	  "Connections": [
	    {
	      "Id": "GitHubAppConnections-1",
	      "Status": "Connected",
	      "Installation": {
	        "InstallationId": "12345",
	        "AccountId": "678",
	        "AccountLogin": "OctopusDeploy",
	        "AccountAvatarUrl": "https://avatars.githubusercontent.com/u/678",
	        "AccountType": "Organization",
	        "AllRepositories": false
	      }
	    },
	    { "Id": "GitHubAppConnections-2", "Status": "InstallationSuspended" }
	  ],
	  "ItemsPerPage": 30,
	  "NumberOfPages": 1,
	  "TotalResults": 2
	}`

	var requested []string
	client := newTestClient(t, &requested, payload)

	result, err := List(client, 0, 30)
	require.NoError(t, err)

	// skip and take are [Required] on the server contract, so a zero skip must still be sent.
	require.Len(t, requested, 1)
	assert.Equal(t, "/api/platformhub/githubconnections/connections?skip=0&take=30", requested[0])

	require.Len(t, result.Connections, 2)
	assert.Equal(t, 2, result.TotalResults)
	assert.Equal(t, "GitHubAppConnections-1", result.Connections[0].ID)
	assert.Equal(t, githubconnections.ConnectionStatusConnected, result.Connections[0].Status)
	require.NotNil(t, result.Connections[0].Installation)
	assert.Equal(t, "OctopusDeploy", result.Connections[0].Installation.AccountLogin)
	assert.Equal(t, "Organization", result.Connections[0].Installation.AccountType)
	assert.False(t, result.Connections[0].Installation.AllRepositories)
	assert.Equal(t, githubconnections.ConnectionStatusInstallationSuspended, result.Connections[1].Status)
	assert.Nil(t, result.Connections[1].Installation)
}

func TestGetByID(t *testing.T) {
	const payload = `{
	  "Id": "GitHubAppConnections-1",
	  "Status": "Connected",
	  "StatusUserMessage": "All good",
	  "Installation": { "InstallationId": "12345", "AccountLogin": "OctopusDeploy", "AccountType": "Organization" },
	  "Repositories": [
	    {
	      "RepositoryId": "R_1",
	      "RepositoryName": "hub",
	      "IsAdmin": true,
	      "IsPrivate": true,
	      "Visibility": "private",
	      "Language": "Go",
	      "GitUrl": "https://githubconnections.com/OctopusDeploy/hub.git",
	      "DefaultBranch": "main"
	    }
	  ],
	  "UnknownRepositories": [{ "RepositoryId": "R_2" }]
	}`

	var requested []string
	client := newTestClient(t, &requested, payload)

	connection, err := GetByID(client, "GitHubAppConnections-1")
	require.NoError(t, err)

	assert.Equal(t, "/api/platformhub/githubconnections/connections/GitHubAppConnections-1", requested[0])
	assert.Equal(t, githubconnections.ConnectionStatusConnected, connection.Status)
	assert.Equal(t, "All good", connection.StatusUserMessage)
	require.Len(t, connection.Repositories, 1)
	assert.Equal(t, "https://github.com/OctopusDeploy/hub.git", connection.Repositories[0].GitURL)
	assert.Equal(t, "main", connection.Repositories[0].DefaultBranch)
	require.Len(t, connection.UnknownRepositories, 1)
	assert.Equal(t, "R_2", connection.UnknownRepositories[0].RepositoryID)
}

func TestGetByIDWithEmptyID(t *testing.T) {
	var requested []string
	client := newTestClient(t, &requested, `{}`)

	connection, err := GetByID(client, "")
	require.Error(t, err)
	require.Nil(t, connection)
	assert.Empty(t, requested)
}

func TestGetRepositories(t *testing.T) {
	const payload = `{
	  "Repositories": [
	    { "RepositoryId": "R_1", "RepositoryName": "hub", "GitUrl": "https://githubconnections.com/OctopusDeploy/hub.git", "DefaultBranch": "main" },
	    { "RepositoryId": "R_2", "RepositoryName": "other", "GitUrl": "https://githubconnections.com/OctopusDeploy/other.git", "DefaultBranch": "trunk" }
	  ]
	}`

	var requested []string
	client := newTestClient(t, &requested, payload)

	repositories, err := GetRepositories(client, "GitHubAppConnections-1")
	require.NoError(t, err)

	assert.Equal(t, "/api/platformhub/githubconnections/connections/GitHubAppConnections-1/repositories", requested[0])
	require.Len(t, repositories, 2)
	assert.Equal(t, "trunk", repositories[1].DefaultBranch)
}

func TestGetRepositoriesWithEmptyConnectionID(t *testing.T) {
	var requested []string
	client := newTestClient(t, &requested, `{}`)

	repositories, err := GetRepositories(client, "")
	require.Error(t, err)
	require.Nil(t, repositories)
	assert.Empty(t, requested)
}
