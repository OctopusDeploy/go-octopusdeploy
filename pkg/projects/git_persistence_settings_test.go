package projects

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestGitPersistenceSettingsNew(t *testing.T) {
	var basePath string
	var conversionState *ConversionState
	var gitCredentials credentials.IGitCredential
	var defaultBranch string
	var protectedBranchNamePatterns []string
	var url *url.URL

	gitPersistenceSettings := NewGitPersistenceSettings(basePath, conversionState, gitCredentials, defaultBranch, protectedBranchNamePatterns, url)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, conversionState, gitPersistenceSettings.ConversionState)
	require.Equal(t, gitCredentials, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, protectedBranchNamePatterns, gitPersistenceSettings.ProtectedBranchNamePatterns)
	require.Equal(t, url, gitPersistenceSettings.URL)

	basePath = internal.GetRandomName()
	conversionState = NewConversionState(false)
	gitCredentials = credentials.NewAnonymous()
	defaultBranch = internal.GetRandomName()
	protectedBranchNamePatterns = []string{internal.GetRandomName(), internal.GetRandomName()}
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)

	gitPersistenceSettings = NewGitPersistenceSettings(basePath, conversionState, gitCredentials, defaultBranch, protectedBranchNamePatterns, url)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, conversionState, gitPersistenceSettings.ConversionState)
	require.Equal(t, gitCredentials, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, protectedBranchNamePatterns, gitPersistenceSettings.ProtectedBranchNamePatterns)
	require.Equal(t, url, gitPersistenceSettings.URL)

	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	basePath = internal.GetRandomName()
	conversionState = NewConversionState(true)
	gitCredentials = credentials.NewUsernamePassword(username, password)
	defaultBranch = internal.GetRandomName()
	protectedBranchNamePatterns = []string{internal.GetRandomName(), internal.GetRandomName()}
	url, err = url.Parse("https://example.com/")
	require.NoError(t, err)

	gitPersistenceSettings = NewGitPersistenceSettings(basePath, conversionState, gitCredentials, defaultBranch, protectedBranchNamePatterns, url)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, conversionState, gitPersistenceSettings.ConversionState)
	require.Equal(t, gitCredentials, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, protectedBranchNamePatterns, gitPersistenceSettings.ProtectedBranchNamePatterns)
	require.Equal(t, url, gitPersistenceSettings.URL)
}

func TestGitPersistenceSettingsMarshalJSON(t *testing.T) {
	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	isConversionState := false

	basePath := internal.GetRandomName()
	conversionState := NewConversionState(isConversionState)
	gitCredentials := credentials.NewUsernamePassword(username, password)
	defaultBranch := internal.GetRandomName()
	protectedBranchNamePatterns := []string{}
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)

	conversionStateAsJSON, err := json.Marshal(conversionState)
	require.NoError(t, err)
	require.NotNil(t, conversionStateAsJSON)

	gitCredentialsAsJSON, err := json.Marshal(gitCredentials)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	expectedJson := fmt.Sprintf(`{
		"BasePath": "%s",
		"ConversionState": %s,
		"Credentials": %s,
		"DefaultBranch": "%s",
		"ProtectedBranchNamePatterns": [],
		"Type": "%s",
		"Url": "%s"
	}`, basePath, conversionStateAsJSON, gitCredentialsAsJSON, defaultBranch, PersistenceSettingsTypeVersionControlled, url.String())

	gitPersistenceSettings := NewGitPersistenceSettings(basePath, conversionState, gitCredentials, defaultBranch, protectedBranchNamePatterns, url)
	gitPersistenceSettingsAsJSON, err := json.Marshal(gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettingsAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(gitPersistenceSettingsAsJSON))
}

func TestGitPersistenceSettingsUnmarshalJSON(t *testing.T) {
	basePath := ""
	var anonymousGitCredential credentials.IGitCredential
	defaultBranch := ""
	var url *url.URL

	gitCredentialsAsJSON, err := json.Marshal(anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	inputJSON := fmt.Sprintf(`{
		"Type": "%s"
	}`, PersistenceSettingsTypeVersionControlled)

	var gitPersistenceSettings GitPersistenceSettings
	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, anonymousGitCredential, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, url, gitPersistenceSettings.URL)

	basePath = internal.GetRandomName()
	defaultBranch = internal.GetRandomName()

	gitCredentialsAsJSON, err = json.Marshal(anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, gitCredentialsAsJSON)

	inputJSON = fmt.Sprintf(`{
		"BasePath": "%s",
		"DefaultBranch": "%s",
		"Type": "%s"
	}`, basePath, defaultBranch, PersistenceSettingsTypeVersionControlled)

	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, anonymousGitCredential, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, url, gitPersistenceSettings.URL)

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
	}`, basePath, gitCredentialsAsJSON, defaultBranch, PersistenceSettingsTypeVersionControlled, url.String())

	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, anonymousGitCredential, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, url, gitPersistenceSettings.URL)

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
	}`, basePath, gitCredentialsAsJSON, defaultBranch, PersistenceSettingsTypeVersionControlled, url.String())

	err = json.Unmarshal([]byte(inputJSON), &gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettings)
	require.Equal(t, PersistenceSettingsTypeVersionControlled, gitPersistenceSettings.GetType())
	require.Equal(t, basePath, gitPersistenceSettings.BasePath)
	require.Equal(t, usernamePasswordGitCredential, gitPersistenceSettings.Credentials)
	require.Equal(t, defaultBranch, gitPersistenceSettings.DefaultBranch)
	require.Equal(t, url, gitPersistenceSettings.URL)
}
