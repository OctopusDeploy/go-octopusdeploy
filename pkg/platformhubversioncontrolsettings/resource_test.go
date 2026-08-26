package platformhubversioncontrolsettings

import (
	"encoding/json"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/stretchr/testify/require"
)

func TestResource_UnmarshalJSON_Anonymous(t *testing.T) {
	const payload = `{
		"Url": "https://githubconnections.com/OctopusDeploy/hub.git",
		"DefaultBranch": "main",
		"BasePath": ".octopus/",
		"Credentials": { "Type": "Anonymous" }
	}`

	var resource Resource
	require.NoError(t, json.Unmarshal([]byte(payload), &resource))

	require.Equal(t, "https://github.com/OctopusDeploy/hub.git", resource.URL)
	require.Equal(t, "main", resource.DefaultBranch)
	require.Equal(t, ".octopus/", resource.BasePath)

	anonymous, ok := resource.Credentials.(*credentials.Anonymous)
	require.True(t, ok)
	require.Equal(t, credentials.GitCredentialTypeAnonymous, anonymous.Type())
}

func TestResource_UnmarshalJSON_UsernamePassword(t *testing.T) {
	const payload = `{
		"Url": "https://githubconnections.com/OctopusDeploy/hub.git",
		"DefaultBranch": "main",
		"BasePath": ".octopus/",
		"Credentials": { "Type": "UsernamePassword", "Username": "octobob", "Password": { "HasValue": true } }
	}`

	var resource Resource
	require.NoError(t, json.Unmarshal([]byte(payload), &resource))

	usernamePassword, ok := resource.Credentials.(*credentials.UsernamePassword)
	require.True(t, ok)
	require.Equal(t, credentials.GitCredentialTypeUsernamePassword, usernamePassword.Type())
	require.Equal(t, "octobob", usernamePassword.Username)
	require.NotNil(t, usernamePassword.Password)
	require.True(t, usernamePassword.Password.HasValue)
}

func TestResource_UnmarshalJSON_GitHubApp(t *testing.T) {
	const payload = `{
		"Url": "https://githubconnections.com/OctopusDeploy/hub.git",
		"DefaultBranch": "main",
		"BasePath": ".octopus/",
		"Credentials": { "Type": "GitHub", "Id": "GitHubAppConnections-1" }
	}`

	var resource Resource
	require.NoError(t, json.Unmarshal([]byte(payload), &resource))

	gitHubApp, ok := resource.Credentials.(*credentials.GitHubApp)
	require.True(t, ok)
	require.Equal(t, credentials.GitCredentialTypeGitHubApp, gitHubApp.Type())
	require.Equal(t, "GitHubAppConnections-1", gitHubApp.ID)
}

func TestResource_RoundTrip(t *testing.T) {
	testCases := []struct {
		name        string
		credentials credentials.GitCredential
	}{
		{"anonymous", credentials.NewAnonymous()},
		{"username password", credentials.NewUsernamePassword("octobob", core.NewSensitiveValue("secret"))},
		{"githubconnections app", credentials.NewGitHubApp("GitHubAppConnections-1")},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			original := NewResource("https://github.com/OctopusDeploy/hub.git", testCase.credentials, "main", ".octopus/")

			data, err := json.Marshal(original)
			require.NoError(t, err)

			var actual Resource
			require.NoError(t, json.Unmarshal(data, &actual))

			require.Equal(t, original.URL, actual.URL)
			require.Equal(t, original.DefaultBranch, actual.DefaultBranch)
			require.Equal(t, original.BasePath, actual.BasePath)
			require.NotNil(t, actual.Credentials)
			require.Equal(t, original.Credentials.Type(), actual.Credentials.Type())
			require.Equal(t, original.Credentials, actual.Credentials)
		})
	}
}
