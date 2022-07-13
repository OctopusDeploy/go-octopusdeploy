package projects

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestGitPersistenceSettingsNew(t *testing.T) {
	var basePath string
	var credentials IGitCredential
	var defaultBranch string
	var url *url.URL

	gitPersistenceSettings := NewGitPersistenceSettings(basePath, credentials, defaultBranch, url)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, "VersionControlled", gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, credentials, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, url, gitPersistenceSettings.URL)

	basePath = internal.GetRandomName()
	credentials = NewAnonymousGitCredential()
	defaultBranch = internal.GetRandomName()
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)

	gitPersistenceSettings = NewGitPersistenceSettings(basePath, credentials, defaultBranch, url)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, "VersionControlled", gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, credentials, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, url, gitPersistenceSettings.URL)

	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	basePath = internal.GetRandomName()
	credentials = NewUsernamePasswordGitCredential(username, password)
	defaultBranch = internal.GetRandomName()
	url, err = url.Parse("https://example.com/")
	require.NoError(t, err)

	gitPersistenceSettings = NewGitPersistenceSettings(basePath, credentials, defaultBranch, url)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, "VersionControlled", gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, credentials, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, url, gitPersistenceSettings.URL)
}

func TestGitPersistenceSettingsMarshalJSON(t *testing.T) {
	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	basePath := internal.GetRandomName()
	credentials := NewUsernamePasswordGitCredential(username, password)
	defaultBranch := internal.GetRandomName()
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)

	credentialsAsJSON, err := json.Marshal(credentials)
	require.NoError(t, err)
	require.NotNil(t, credentialsAsJSON)

	expectedJson := fmt.Sprintf(`{
		"BasePath": "%s",
		"Credentials": %s,
		"DefaultBranch": "%s",
		"Type": "VersionControlled",
		"Url": "%s"
	}`, basePath, credentialsAsJSON, defaultBranch, url.String())

	gitPersistenceSettings := NewGitPersistenceSettings(basePath, credentials, defaultBranch, url)
	gitPersistenceSettingsAsJSON, err := json.Marshal(gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettingsAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(gitPersistenceSettingsAsJSON))
}

func TestGitPersistenceSettingsUnmarshalJSON(t *testing.T) {
	basePath := ""
	var anonymousGitCredential IGitCredential
	defaultBranch := ""
	var url *url.URL

	credentialsAsJSON, err := json.Marshal(anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, credentialsAsJSON)

	inputJSON := `{
		"Type": "VersionControlled"
	}`

	var gitPersistenceSettings GitPersistenceSettings
	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, "VersionControlled", gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, anonymousGitCredential, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, url, gitPersistenceSettings.URL)

	basePath = internal.GetRandomName()
	defaultBranch = internal.GetRandomName()

	credentialsAsJSON, err = json.Marshal(anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, credentialsAsJSON)

	inputJSON = fmt.Sprintf(`{
		"BasePath": "%s",
		"DefaultBranch": "%s",
		"Type": "VersionControlled"
	}`, basePath, defaultBranch)

	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, "VersionControlled", gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, anonymousGitCredential, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, url, gitPersistenceSettings.URL)

	basePath = internal.GetRandomName()
	anonymousGitCredential = NewAnonymousGitCredential()
	defaultBranch = internal.GetRandomName()
	url, err = url.Parse("https://example.com/")
	require.NoError(t, err)

	credentialsAsJSON, err = json.Marshal(anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, credentialsAsJSON)

	inputJSON = fmt.Sprintf(`{
		"BasePath": "%s",
		"Credentials": %s,
		"DefaultBranch": "%s",
		"Type": "VersionControlled",
		"Url": "%s"
	}`, basePath, credentialsAsJSON, defaultBranch, url.String())

	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, "VersionControlled", gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, anonymousGitCredential, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, url, gitPersistenceSettings.URL)

	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	basePath = internal.GetRandomName()
	defaultBranch = internal.GetRandomName()
	url, err = url.Parse("https://example.com/")
	usernamePasswordGitCredential := NewUsernamePasswordGitCredential(username, password)
	require.NoError(t, err)

	credentialsAsJSON, err = json.Marshal(usernamePasswordGitCredential)
	require.NoError(t, err)
	require.NotNil(t, credentialsAsJSON)

	inputJSON = fmt.Sprintf(`{
		"BasePath": "%s",
		"Credentials": %s,
		"DefaultBranch": "%s",
		"Type": "VersionControlled",
		"Url": "%s"
	}`, basePath, credentialsAsJSON, defaultBranch, url.String())

	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, "VersionControlled", gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, usernamePasswordGitCredential, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, url, gitPersistenceSettings.URL)
}
