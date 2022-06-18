package resources

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/projects"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestUsernamePasswordGitCredentialNew(t *testing.T) {
	var password *core.SensitiveValue
	var username string

	usernamePasswordGitCredential := projects.NewUsernamePasswordGitCredential(username, password)
	require.NotNil(t, usernamePasswordGitCredential)
	require.Equal(t, projects.GitCredentialType("UsernamePassword"), usernamePasswordGitCredential.GetType())
	require.Equal(t, password, usernamePasswordGitCredential.Password)
	require.Equal(t, username, usernamePasswordGitCredential.Username)

	password = core.NewSensitiveValue(internal.GetRandomName())
	username = internal.GetRandomName()
	usernamePasswordGitCredential = projects.NewUsernamePasswordGitCredential(username, password)
	require.Equal(t, password, usernamePasswordGitCredential.Password)
	require.Equal(t, username, usernamePasswordGitCredential.Username)
}

func TestUsernamePasswordGitCredentialMarshalJSON(t *testing.T) {
	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	passwordAsJSON, err := json.Marshal(password)
	require.NoError(t, err)
	require.NotNil(t, passwordAsJSON)

	expectedJson := fmt.Sprintf(`{
		"Password": %s,
		"Type": "UsernamePassword",
		"Username": "%s"
	}`, passwordAsJSON, username)

	usernamePasswordGitCredential := projects.NewUsernamePasswordGitCredential(username, password)
	require.NotNil(t, usernamePasswordGitCredential)
	require.Equal(t, projects.GitCredentialType("UsernamePassword"), usernamePasswordGitCredential.GetType())
	require.Equal(t, password, usernamePasswordGitCredential.Password)
	require.Equal(t, username, usernamePasswordGitCredential.Username)

	usernamePasswordGitCredentialAsJSON, err := json.Marshal(usernamePasswordGitCredential)
	require.NoError(t, err)
	require.NotNil(t, usernamePasswordGitCredentialAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(usernamePasswordGitCredentialAsJSON))
}

func TestUsernamePasswordGitCredentialUnmarshalJSON(t *testing.T) {
	password := core.NewSensitiveValue(internal.GetRandomName())
	username := internal.GetRandomName()

	passwordAsJSON, err := json.Marshal(password)
	require.NoError(t, err)
	require.NotNil(t, passwordAsJSON)

	inputJSON := fmt.Sprintf(`{
		"Password": %s,
		"Type": "UsernamePassword",
		"Username": "%s"
	}`, passwordAsJSON, username)

	var usernamePasswordGitCredential projects.UsernamePasswordGitCredential
	err = json.Unmarshal([]byte(inputJSON), &usernamePasswordGitCredential)
	require.NoError(t, err)
	require.NotNil(t, usernamePasswordGitCredential)
	require.Equal(t, projects.GitCredentialType("UsernamePassword"), usernamePasswordGitCredential.GetType())
	require.Equal(t, password, usernamePasswordGitCredential.Password)
	require.Equal(t, username, usernamePasswordGitCredential.Username)
}
