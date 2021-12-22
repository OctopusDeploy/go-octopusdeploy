package octopusdeploy

import (
	"encoding/json"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestAnonymousGitCredentialNew(t *testing.T) {
	anonymousGitCredential := NewAnonymousGitCredential()
	require.NotNil(t, anonymousGitCredential)
	require.Equal(t, GitCredentialType("Anonymous"), anonymousGitCredential.GetType())
}

func TestAnonymousGitCredentialMarshalJSON(t *testing.T) {
	expectedJson := `{
		"Type": "Anonymous"
	}`

	anonymousGitCredential := NewAnonymousGitCredential()
	require.Equal(t, GitCredentialType("Anonymous"), anonymousGitCredential.GetType())

	anonymousGitCredentialAsJSON, err := json.Marshal(anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, anonymousGitCredentialAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(anonymousGitCredentialAsJSON))
}

func TestAnonymousGitCredentialUnmarshalJSON(t *testing.T) {
	inputJSON := `{
		"Type": "Anonymous"
	}`

	var anonymousGitCredential AnonymousGitCredential
	err := json.Unmarshal([]byte(inputJSON), &anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, anonymousGitCredential)
	require.Equal(t, GitCredentialType("Anonymous"), anonymousGitCredential.GetType())
}
