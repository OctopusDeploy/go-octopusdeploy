package releases

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReleaseTemplateGitResourceUnmarshalJSON(t *testing.T) {
	payload := `{
		"ActionName": "Update Argo Manifests",
		"RepositoryUri": "https://github.com/OctopusDeploy/manifests.git",
		"DefaultBranch": "main",
		"IsResolvable": true,
		"Name": "",
		"FilePathFilters": ["deployment.yaml"],
		"GitCredentialId": "GitCredentials-1",
		"GitHubConnectionId": "GitHubConnections-1",
		"GitResourceSelectedLastRelease": {
			"GitRef": "refs/heads/main",
			"GitCommit": "d65a219f86d95cc193050efa2e7a0cd2d314e9a2"
		}
	}`

	var gitResource ReleaseTemplateGitResource
	require.NoError(t, json.Unmarshal([]byte(payload), &gitResource))

	require.Equal(t, "Update Argo Manifests", gitResource.ActionName)
	require.Equal(t, "https://github.com/OctopusDeploy/manifests.git", gitResource.RepositoryUri)
	require.Equal(t, "main", gitResource.DefaultBranch)
	require.True(t, gitResource.IsResolvable)
	require.Equal(t, []string{"deployment.yaml"}, gitResource.FilePathFilters)
	require.Equal(t, "GitCredentials-1", gitResource.GitCredentialId)
	require.Equal(t, "GitHubConnections-1", gitResource.GitHubConnectionId)
	require.Equal(t, "refs/heads/main", gitResource.GitResourceSelectedLastRelease.GitRef)
	require.Equal(t, "d65a219f86d95cc193050efa2e7a0cd2d314e9a2", gitResource.GitResourceSelectedLastRelease.GitCommit)
}

func TestReleaseTemplateGitResourceMarshalJSON(t *testing.T) {
	gitResource := ReleaseTemplateGitResource{
		ActionName:         "Update Argo Manifests",
		GitCredentialId:    "GitCredentials-1",
		GitHubConnectionId: "GitHubConnections-1",
	}

	data, err := json.Marshal(gitResource)
	require.NoError(t, err)

	var roundTripped map[string]any
	require.NoError(t, json.Unmarshal(data, &roundTripped))

	require.Equal(t, "GitCredentials-1", roundTripped["GitCredentialId"])
	require.Equal(t, "GitHubConnections-1", roundTripped["GitHubConnectionId"])
	require.NotContains(t, roundTripped, "NuGetPackageId")
}
