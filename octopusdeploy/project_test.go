package octopusdeploy

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestProjectNew(t *testing.T) {
	var lifecycleID string
	var name string
	var projectGroupID string

	project := NewProject(name, lifecycleID, projectGroupID)
	require.NotNil(t, project)
	require.Equal(t, name, project.Name)
	require.Equal(t, lifecycleID, project.LifecycleID)
	require.Equal(t, projectGroupID, project.ProjectGroupID)
}

func TestProjectMarshalJSON(t *testing.T) {
	lifecycleID := getRandomName()
	name := getRandomName()
	projectGroupID := getRandomName()

	expectedJson := fmt.Sprintf(`{
		"LifecycleId": "%s",
		"Name": "%s",
		"ProjectGroupId": "%s"
	}`, lifecycleID, name, projectGroupID)

	project := NewProject(name, lifecycleID, projectGroupID)
	projectAsJSON, err := json.Marshal(project)
	require.NoError(t, err)
	require.NotNil(t, projectAsJSON)
	jsonassert.New(t).Assertf(string(projectAsJSON), expectedJson)

	connectivityPolicy := NewConnectivityPolicy()
	connectivityPolicyAsJSON, err := json.Marshal(connectivityPolicy)
	require.NoError(t, err)
	require.NotNil(t, connectivityPolicy)

	project.ConnectivityPolicy = NewConnectivityPolicy()
	projectAsJSON, err = json.Marshal(project)
	require.NoError(t, err)
	require.NotNil(t, projectAsJSON)

	expectedJson = fmt.Sprintf(`{
		"LifecycleId": "%s",
		"Name": "%s",
		"ProjectConnectivityPolicy": %s,
		"ProjectGroupId": "%s"
	}`, lifecycleID, name, connectivityPolicyAsJSON, projectGroupID)

	jsonassert.New(t).Assertf(string(projectAsJSON), expectedJson)
}

func TestProjectUnmarshalJSON(t *testing.T) {
	lifecycleID := getRandomName()
	name := getRandomName()
	projectGroupID := getRandomName()

	inputJSON := fmt.Sprintf(`{
		"LifecycleId": "%s",
		"Name": "%s",
		"ProjectGroupId": "%s"
	}`, lifecycleID, name, projectGroupID)

	var project Project
	err := json.Unmarshal([]byte(inputJSON), &project)
	require.NoError(t, err)
	require.NotNil(t, project)
	require.Equal(t, lifecycleID, project.LifecycleID)
	require.Equal(t, name, project.Name)
	require.Equal(t, projectGroupID, project.ProjectGroupID)

	persistenceSettings := NewDatabasePersistenceSettings()
	persistenceSettingsAsJSON, err := json.Marshal(persistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, persistenceSettingsAsJSON)

	inputJSON = fmt.Sprintf(`{
		"LifecycleId": "%s",
		"Name": "%s",
		"PersistenceSettings": %s,
		"ProjectGroupId": "%s"
	}`, lifecycleID, name, persistenceSettingsAsJSON, projectGroupID)

	err = json.Unmarshal([]byte(inputJSON), &project)
	require.NoError(t, err)
	require.NotNil(t, project)
	require.Equal(t, lifecycleID, project.LifecycleID)
	require.Equal(t, name, project.Name)
	require.Equal(t, projectGroupID, project.ProjectGroupID)
	require.Equal(t, persistenceSettings, project.PersistenceSettings)

	password := NewSensitiveValue(getRandomName())
	username := getRandomName()

	basePath := getRandomName()
	credentials := NewUsernamePasswordGitCredential(username, password)
	defaultBranch := getRandomName()
	url, err := url.Parse("https://example.com/")
	require.NoError(t, err)

	gitPersistenceSettings := NewGitPersistenceSettings(basePath, credentials, defaultBranch, url)
	gitPersistenceSettingsAsJSON, err := json.Marshal(gitPersistenceSettings)
	require.NoError(t, err)
	require.NotNil(t, gitPersistenceSettingsAsJSON)

	inputJSON = fmt.Sprintf(`{
		"LifecycleId": "%s",
		"Name": "%s",
		"PersistenceSettings": %s,
		"ProjectGroupId": "%s"
	}`, lifecycleID, name, gitPersistenceSettingsAsJSON, projectGroupID)

	err = json.Unmarshal([]byte(inputJSON), &project)
	require.NoError(t, err)
	require.NotNil(t, project)
	require.Equal(t, lifecycleID, project.LifecycleID)
	require.Equal(t, name, project.Name)
	require.Equal(t, projectGroupID, project.ProjectGroupID)
	require.Equal(t, gitPersistenceSettings, project.PersistenceSettings)
}
