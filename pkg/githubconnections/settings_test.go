package githubconnections

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSettings(t *testing.T) {
	var requested string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requested = r.URL.RequestURI()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"CanUseGitHubApp":true,"CanUseTrustedFlow":false}`))
	}))
	defer server.Close()

	baseURL, err := url.Parse(server.URL + "/")
	require.NoError(t, err)
	client := newclient.NewClient(&newclient.HttpSession{HttpClient: server.Client(), BaseURL: baseURL})

	settings, err := GetSettings(client)
	require.NoError(t, err)

	assert.Equal(t, "/api/githubconnections/app/settings", requested)
	assert.True(t, settings.CanUseGitHubApp)
	assert.False(t, settings.CanUseTrustedFlow)
}
