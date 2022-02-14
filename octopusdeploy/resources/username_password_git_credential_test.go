package resources

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestUsernamePasswordGitCredentialNew(t *testing.T) {
	var password *SensitiveValue
	var username string

	usernamePasswordGitCredential := NewUsernamePasswordGitCredential(username, password)
	require.NotNil(t, usernamePasswordGitCredential)
	require.Equal(t, GitCredentialType("UsernamePassword"), usernamePasswordGitCredential.GetType())
	require.Equal(t, password, usernamePasswordGitCredential.Password)
	require.Equal(t, username, usernamePasswordGitCredential.Username)

	password = NewSensitiveValue(getRandomName())
	username = getRandomName()
	usernamePasswordGitCredential = NewUsernamePasswordGitCredential(username, password)
	require.Equal(t, password, usernamePasswordGitCredential.Password)
	require.Equal(t, username, usernamePasswordGitCredential.Username)
}

func TestUsernamePasswordGitCredentialMarshalJSON(t *testing.T) {
	password := NewSensitiveValue(getRandomName())
	username := getRandomName()

	passwordAsJSON, err := json.Marshal(password)
	require.NoError(t, err)
	require.NotNil(t, passwordAsJSON)

	expectedJson := fmt.Sprintf(`{
		"Password": %s,
		"Type": "UsernamePassword",
		"Username": "%s"
	}`, passwordAsJSON, username)

	usernamePasswordGitCredential := NewUsernamePasswordGitCredential(username, password)
	require.NotNil(t, usernamePasswordGitCredential)
	require.Equal(t, GitCredentialType("UsernamePassword"), usernamePasswordGitCredential.GetType())
	require.Equal(t, password, usernamePasswordGitCredential.Password)
	require.Equal(t, username, usernamePasswordGitCredential.Username)

	usernamePasswordGitCredentialAsJSON, err := json.Marshal(usernamePasswordGitCredential)
	require.NoError(t, err)
	require.NotNil(t, usernamePasswordGitCredentialAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(usernamePasswordGitCredentialAsJSON))
}

func TestUsernamePasswordGitCredentialUnmarshalJSON(t *testing.T) {
	password := NewSensitiveValue(getRandomName())
	username := getRandomName()

	passwordAsJSON, err := json.Marshal(password)
	require.NoError(t, err)
	require.NotNil(t, passwordAsJSON)

	inputJSON := fmt.Sprintf(`{
		"Password": %s,
		"Type": "UsernamePassword",
		"Username": "%s"
	}`, passwordAsJSON, username)

	var usernamePasswordGitCredential UsernamePasswordGitCredential
	err = json.Unmarshal([]byte(inputJSON), &usernamePasswordGitCredential)
	require.NoError(t, err)
	require.NotNil(t, usernamePasswordGitCredential)
	require.Equal(t, GitCredentialType("UsernamePassword"), usernamePasswordGitCredential.GetType())
	require.Equal(t, password, usernamePasswordGitCredential.Password)
	require.Equal(t, username, usernamePasswordGitCredential.Username)
}
