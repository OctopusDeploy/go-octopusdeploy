package platformhubgitcredential

import (
	"encoding/json"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/credentials"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/require"
)

func createClient() newclient.Client {
	return newclient.NewClient(&newclient.HttpSession{})
}

func TestTemplateV2Expansion(t *testing.T) {
	cache := createClient().URITemplateCache()

	listOrCreate, err := cache.Expand(templateV2, map[string]any{})
	require.NoError(t, err)
	require.Equal(t, "/api/platformhub/git-credentials/v2", listOrCreate)

	byIDOrModify, err := cache.Expand(templateV2, map[string]any{"id": "GitCredentials-1"})
	require.NoError(t, err)
	require.Equal(t, "/api/platformhub/git-credentials/GitCredentials-1/v2", byIDOrModify)
}

func TestPlatformHubGitCredentialUnmarshalUsernamePassword(t *testing.T) {
	inputJSON := `{
		"Id": "GitCredentials-1",
		"Name": "user-pass",
		"Details": { "Type": "UsernamePassword", "Username": "admin", "Password": { "HasValue": true } },
		"RepositoryRestrictions": { "Enabled": false, "AllowedRepositories": [] },
		"Links": {}
	}`

	var resource PlatformHubGitCredential
	require.NoError(t, json.Unmarshal([]byte(inputJSON), &resource))

	require.Equal(t, credentials.GitCredentialTypeUsernamePassword, resource.Details.Type())
	_, ok := resource.Details.(*credentials.UsernamePassword)
	require.True(t, ok)
}

func TestPlatformHubGitCredentialUnmarshalSshKey(t *testing.T) {
	inputJSON := `{
		"Id": "GitCredentials-2",
		"Name": "ssh",
		"Details": { "Type": "SshKey", "Username": "git", "PrivateKey": { "HasValue": true }, "PrivateKeyFingerprint": "SHA256:abc" },
		"RepositoryRestrictions": { "Enabled": false, "AllowedRepositories": [] },
		"Links": {}
	}`

	var resource PlatformHubGitCredential
	require.NoError(t, json.Unmarshal([]byte(inputJSON), &resource))

	require.Equal(t, credentials.GitCredentialTypeSshKey, resource.Details.Type())
	sshKey, ok := resource.Details.(*credentials.SshKey)
	require.True(t, ok)
	require.Equal(t, "SHA256:abc", sshKey.PrivateKeyFingerprint)
}

func TestAddV2NilCredential(t *testing.T) {
	response, err := AddV2(createClient(), nil)
	require.Equal(t, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential"), err)
	require.Nil(t, response)
}

func TestGetByIDV2EmptyID(t *testing.T) {
	resource, err := GetByIDV2(createClient(), "")
	require.Equal(t, internal.CreateRequiredParameterIsEmptyOrNilError("id"), err)
	require.Nil(t, resource)
}

func TestUpdateV2NilCredential(t *testing.T) {
	err := UpdateV2(createClient(), nil)
	require.Equal(t, internal.CreateRequiredParameterIsEmptyOrNilError("platformHubGitCredential"), err)
}
