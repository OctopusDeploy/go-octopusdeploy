package credentials

import (
	"encoding/json"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/require"
)

func TestTemplateV2Expansion(t *testing.T) {
	cache := newclient.NewClient(&newclient.HttpSession{}).URITemplateCache()

	listOrCreate, err := cache.Expand(templateV2, map[string]any{"spaceId": "Spaces-1"})
	require.NoError(t, err)
	require.Equal(t, "/api/Spaces-1/git-credentials/v2", listOrCreate)

	byIDOrModify, err := cache.Expand(templateV2, map[string]any{"spaceId": "Spaces-1", "id": "GitCredentials-1"})
	require.NoError(t, err)
	require.Equal(t, "/api/Spaces-1/git-credentials/GitCredentials-1/v2", byIDOrModify)

	withQuery, err := cache.Expand(templateV2, map[string]any{"spaceId": "Spaces-1", "name": "foo", "skip": 5, "take": 10})
	require.NoError(t, err)
	require.Equal(t, "/api/Spaces-1/git-credentials/v2?skip=5&take=10&name=foo", withQuery)
}

func TestCreateGitCredentialResponseV2Unmarshal(t *testing.T) {
	var response CreateGitCredentialResponseV2
	err := json.Unmarshal([]byte(`{ "Id": "GitCredentials-1" }`), &response)
	require.NoError(t, err)
	require.Equal(t, "GitCredentials-1", response.ID)
}

func TestGetGitCredentialByIdResponseV2Unmarshal(t *testing.T) {
	inputJSON := `{
		"GitCredential": {
			"Id": "GitCredentials-1",
			"SpaceId": "Spaces-1",
			"Name": "my-credential",
			"Details": { "Type": "SshKey", "Username": "git", "PrivateKey": { "HasValue": true } },
			"RepositoryRestrictions": { "Enabled": false, "AllowedRepositories": [] },
			"Links": {}
		}
	}`

	var response getGitCredentialByIdResponseV2
	err := json.Unmarshal([]byte(inputJSON), &response)
	require.NoError(t, err)
	require.NotNil(t, response.GitCredential)
	require.Equal(t, "GitCredentials-1", response.GitCredential.GetID())
	require.Equal(t, GitCredentialTypeSshKey, response.GitCredential.Details.Type())
}
