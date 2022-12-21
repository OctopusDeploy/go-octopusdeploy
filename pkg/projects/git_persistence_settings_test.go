package projects_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/projects"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestGitPersistenceSettingsNew(t *testing.T) {
	var basePath string
	var gitCredentials credentials.GitCredential
	var defaultBranch string
	var protectedBranchNamePatterns []string
	var url *url.URL

	gitPersistenceSettings := projects.NewGitPersistenceSettings(basePath, gitCredentials, defaultBranch, protectedBranchNamePatterns, url)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, projects.PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.Type())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath())
	require.Equal(t, gitCredentials, gitPersistenceSettings.Credential())
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch())
	require.Equal(t, protectedBranchNamePatterns, gitPersistenceSettings.ProtectedBranchNamePatterns())
	require.Equal(t, url, gitPersistenceSettings.URL())

	basePath = internal.GetRandomName()
	gitCredentials = credentials.NewAnonymous()
	defaultBranch = internal.GetRandomName()
	protectedBranchNamePatterns = []string{internal.GetRandomName(), internal.GetRandomName()}
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)

	gitPersistenceSettings = projects.NewGitPersistenceSettings(basePath, gitCredentials, defaultBranch, protectedBranchNamePatterns, url)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, projects.PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.Type())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath())
	require.Equal(t, gitCredentials, gitPersistenceSettings.Credential())
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch())
	require.Equal(t, protectedBranchNamePatterns, gitPersistenceSettings.ProtectedBranchNamePatterns())
	require.Equal(t, url, gitPersistenceSettings.URL())

	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	basePath = internal.GetRandomName()
	gitCredentials = credentials.NewUsernamePassword(username, password)
	defaultBranch = internal.GetRandomName()
	protectedBranchNamePatterns = []string{internal.GetRandomName(), internal.GetRandomName()}
	url, err = url.Parse("https://example.com/")
	require.NoError(t, err)

	gitPersistenceSettings = projects.NewGitPersistenceSettings(basePath, gitCredentials, defaultBranch, protectedBranchNamePatterns, url)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, projects.PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.Type())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath())
	require.Equal(t, gitCredentials, gitPersistenceSettings.Credential())
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch())
	require.Equal(t, protectedBranchNamePatterns, gitPersistenceSettings.ProtectedBranchNamePatterns())
	require.Equal(t, url, gitPersistenceSettings.URL())
}

func TestGitPersistenceSettingsMarshalJSONWithProtectedDefaultBranch(t *testing.T) {
	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	basePath := internal.GetRandomName()
	gitCredentials := credentials.NewUsernamePassword(username, password)
	defaultBranch := internal.GetRandomName()
	protectedBranchNamePatterns := []string{defaultBranch}
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)

	gitCredentialsAsJSON, err := json.Marshal(gitCredentials)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	expectedJson := fmt.Sprintf(`{
		"BasePath": "%s",
		"Credentials": %s,
		"DefaultBranch": "%s",
		"ProtectedBranchNamePatterns": [],
		"ProtectedDefaultBranch": true,
		"Type": "%s",
		"Url": "%s"
	}`, basePath, gitCredentialsAsJSON, defaultBranch, projects.PersistenceSettingsTypeVersionControlled, url.String())

	gitPersistenceSettings := projects.NewGitPersistenceSettings(basePath, gitCredentials, defaultBranch, protectedBranchNamePatterns, url)
	gitPersistenceSettingsAsJSON, err := json.Marshal(gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettingsAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(gitPersistenceSettingsAsJSON))
}

func TestGitPersistenceSettingsMarshalJSONWithProtectedDefaultBranchAsLastItem(t *testing.T) {
	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	basePath := internal.GetRandomName()
	gitCredentials := credentials.NewUsernamePassword(username, password)
	defaultBranch := internal.GetRandomName()
	protectedBranchNamePatterns := []string{"foo", defaultBranch}
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)

	gitCredentialsAsJSON, err := json.Marshal(gitCredentials)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	expectedJson := fmt.Sprintf(`{
		"BasePath": "%s",
		"Credentials": %s,
		"DefaultBranch": "%s",
		"ProtectedBranchNamePatterns": ["foo"],
		"ProtectedDefaultBranch": true,
		"Type": "%s",
		"Url": "%s"
	}`, basePath, gitCredentialsAsJSON, defaultBranch, projects.PersistenceSettingsTypeVersionControlled, url.String())

	gitPersistenceSettings := projects.NewGitPersistenceSettings(basePath, gitCredentials, defaultBranch, protectedBranchNamePatterns, url)
	gitPersistenceSettingsAsJSON, err := json.Marshal(gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettingsAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(gitPersistenceSettingsAsJSON))
}

func TestGitPersistenceSettingsMarshalJSONWithoutProtectedDefaultBranch(t *testing.T) {
	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	basePath := internal.GetRandomName()
	gitCredentials := credentials.NewUsernamePassword(username, password)
	defaultBranch := internal.GetRandomName()
	protectedBranchName := internal.GetRandomName()
	protectedBranchNamePatterns := []string{protectedBranchName}
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)

	gitCredentialsAsJSON, err := json.Marshal(gitCredentials)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	expectedJson := fmt.Sprintf(`{
		"BasePath": "%s",
		"Credentials": %s,
		"DefaultBranch": "%s",
		"ProtectedBranchNamePatterns": ["%s"],
		"ProtectedDefaultBranch": false,
		"Type": "%s",
		"Url": "%s"
	}`, basePath, gitCredentialsAsJSON, defaultBranch, protectedBranchName, projects.PersistenceSettingsTypeVersionControlled, url.String())

	gitPersistenceSettings := projects.NewGitPersistenceSettings(basePath, gitCredentials, defaultBranch, protectedBranchNamePatterns, url)
	gitPersistenceSettingsAsJSON, err := json.Marshal(gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettingsAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(gitPersistenceSettingsAsJSON))
}

func TestGitPersistenceSettingsUnmarshalJSONWithoutProtectedDefaultBranch(t *testing.T) {
	basePath := ""
	anonymousGitCredential := credentials.NewAnonymous()
	defaultBranch := ""
	var url *url.URL

	gitCredentialsAsJSON, err := json.Marshal(anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	inputJSON := fmt.Sprintf(`{
		"Type": "%s"
	}`, projects.PersistenceSettingsTypeVersionControlled)

	gitPersistenceSettings := projects.NewGitPersistenceSettings("", nil, "", []string{}, nil)
	err = json.Unmarshal([]byte(inputJSON), gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, projects.PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.Type())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath())
	require.Nil(t, gitPersistenceSettings.Credential())
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch())
	require.Equal(t, url, gitPersistenceSettings.URL())

	basePath = internal.GetRandomName()
	defaultBranch = internal.GetRandomName()

	gitCredentialsAsJSON, err = json.Marshal(anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	inputJSON = fmt.Sprintf(`{
		"BasePath": "%s",
		"DefaultBranch": "%s",
		"Type": "%s"
	}`, basePath, defaultBranch, projects.PersistenceSettingsTypeVersionControlled)

	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, projects.PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.Type())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath())
	require.Nil(t, gitPersistenceSettings.Credential())
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch())
	require.Equal(t, url, gitPersistenceSettings.URL())

	basePath = internal.GetRandomName()
	anonymousGitCredential = credentials.NewAnonymous()
	defaultBranch = internal.GetRandomName()
	url, err = url.Parse("https://example.com/")
	require.NoError(t, err)

	gitCredentialsAsJSON, err = json.Marshal(anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	inputJSON = fmt.Sprintf(`{
		"BasePath": "%s",
		"Credentials": %s,
		"DefaultBranch": "%s",
		"Type": "%s",
		"Url": "%s"
	}`, basePath, gitCredentialsAsJSON, defaultBranch, projects.PersistenceSettingsTypeVersionControlled, url.String())

	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, projects.PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.Type())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath())
	require.Equal(t, anonymousGitCredential, gitPersistenceSettings.Credential())
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch())
	require.Equal(t, url, gitPersistenceSettings.URL())

	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	basePath = internal.GetRandomName()
	defaultBranch = internal.GetRandomName()
	url, err = url.Parse("https://example.com/")
	usernamePasswordGitCredential := credentials.NewUsernamePassword(username, password)
	require.NoError(t, err)

	gitCredentialsAsJSON, err = json.Marshal(usernamePasswordGitCredential)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	inputJSON = fmt.Sprintf(`{
		"BasePath": "%s",
		"Credentials": %s,
		"DefaultBranch": "%s",
		"Type": "%s",
		"Url": "%s"
	}`, basePath, gitCredentialsAsJSON, defaultBranch, projects.PersistenceSettingsTypeVersionControlled, url.String())

	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, projects.PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.Type())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath())
	require.Equal(t, usernamePasswordGitCredential, gitPersistenceSettings.Credential())
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch())
	require.Equal(t, url, gitPersistenceSettings.URL())
}

func TestGitPersistenceSettingsUnmarshalJSONWithProtectedBranchNotDefault(t *testing.T) {
	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	basePath := internal.GetRandomName()
	defaultBranch := internal.GetRandomName()
	url, err := url.Parse("https://example.com/")
	usernamePasswordGitCredential := credentials.NewUsernamePassword(username, password)
	require.NoError(t, err)

	gitCredentialsAsJSON, err := json.Marshal(usernamePasswordGitCredential)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	inputJSON := fmt.Sprintf(`{
		"BasePath": "%s",
		"Credentials": %s,
		"DefaultBranch": "%s",
		"ProtectedBranchNamePatterns": ["foo"],
		"ProtectedDefaultBranch": false,
		"Type": "%s",
		"Url": "%s"
	}`, basePath, gitCredentialsAsJSON, defaultBranch, projects.PersistenceSettingsTypeVersionControlled, url.String())

	gitPersistenceSettings := projects.NewGitPersistenceSettings("", nil, "", []string{}, nil)
	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, projects.PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.Type())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath())
	require.Equal(t, usernamePasswordGitCredential, gitPersistenceSettings.Credential())
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch())
	require.Equal(t, url, gitPersistenceSettings.URL())
	require.Equal(t, []string{"foo"}, gitPersistenceSettings.ProtectedBranchNamePatterns())
}

func TestGitPersistenceSettingsUnmarshalJSONWithProtectedBranchIsDefault(t *testing.T) {
	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	basePath := internal.GetRandomName()
	defaultBranch := internal.GetRandomName()
	url, err := url.Parse("https://example.com/")
	usernamePasswordGitCredential := credentials.NewUsernamePassword(username, password)
	require.NoError(t, err)

	gitCredentialsAsJSON, err := json.Marshal(usernamePasswordGitCredential)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	inputJSON := fmt.Sprintf(`{
		"BasePath": "%s",
		"Credentials": %s,
		"DefaultBranch": "%s",
		"ProtectedBranchNamePatterns": [],
		"ProtectedDefaultBranch": true,
		"Type": "%s",
		"Url": "%s"
	}`, basePath, gitCredentialsAsJSON, defaultBranch, projects.PersistenceSettingsTypeVersionControlled, url.String())

	gitPersistenceSettings := projects.NewGitPersistenceSettings("", nil, "", []string{}, nil)
	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, projects.PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.Type())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath())
	require.Equal(t, usernamePasswordGitCredential, gitPersistenceSettings.Credential())
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch())
	require.Equal(t, url, gitPersistenceSettings.URL())
	require.Equal(t, []string{defaultBranch}, gitPersistenceSettings.ProtectedBranchNamePatterns())
}
