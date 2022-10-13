package resources

import (
	"encoding/json"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestAnonymousGitCredentialNew(t *testing.T) {
	anonymousGitCredential := credentials.NewAnonymous()
	require.NotNil(t, anonymousGitCredential)
	require.Equal(t, credentials.Type("Anonymous"), anonymousGitCredential.GetType())
}

func TestAnonymousGitCredentialMarshalJSON(t *testing.T) {
	expectedJson := `{
		"Type": "Anonymous"
	}`

	anonymousGitCredential := credentials.NewAnonymous()
	require.Equal(t, credentials.Type("Anonymous"), anonymousGitCredential.GetType())

	anonymousGitCredentialAsJSON, err := json.Marshal(anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, anonymousGitCredentialAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(anonymousGitCredentialAsJSON))
}

func TestAnonymousGitCredentialUnmarshalJSON(t *testing.T) {
	inputJSON := `{
		"Type": "Anonymous"
	}`

	var anonymousGitCredential credentials.Anonymous
	err := json.Unmarshal([]byte(inputJSON), &anonymousGitCredential)
	require.NoError(t, err)
	require.NotNil(t, anonymousGitCredential)
	require.Equal(t, credentials.Type("Anonymous"), anonymousGitCredential.GetType())
}
