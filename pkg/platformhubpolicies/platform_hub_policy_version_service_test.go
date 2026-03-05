package platformhubpolicies

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/stretchr/testify/require"
)

func TestPlatformHubPolicyVersionService_BuildPublishCommand(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	gitRef := "refs/heads/main"
	slug := "my_policy"
	version := "1.0.1"
	urlEncodedGitRef := "refs%2Fheads%2Fmain"

	// Act
	command, path, commandError := buildPublishCommand(client, gitRef, slug, version)

	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/%s/policies/%s/publish", urlEncodedGitRef, slug), path)
	require.Equal(t, version, command.Version)

	// Verify JSON serialization
	jsonBytes, jsonErr := json.Marshal(command)
	require.NoError(t, jsonErr)
	require.JSONEq(t, fmt.Sprintf(`{"Version":"%s"}`, version), string(jsonBytes))
}

func TestPlatformHubPolicyVersionService_BuildActivateCommand(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	version := PlatformHubPolicyVersion{
		Slug:    "my_policy",
		Version: "1.0.0",
	}

	// Act
	command, path, commandError := buildModifyVersionStatusCommand(client, version, true)

	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/policies/%s/versions/%s/modify-status", version.Slug, version.Version), path)
	require.True(t, command.IsActive)

	// Verify JSON serialization
	jsonBytes, jsonErr := json.Marshal(command)
	require.NoError(t, jsonErr)
	require.JSONEq(t, fmt.Sprintf(`{"IsActive":%t}`, true), string(jsonBytes))
}

func TestPlatformHubPolicyVersionService_BuildDeactivateCommand(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	version := PlatformHubPolicyVersion{
		Slug:    "my_policy",
		Version: "2.0.0",
	}

	// Act
	command, path, commandError := buildModifyVersionStatusCommand(client, version, false)

	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/policies/%s/versions/%s/modify-status", version.Slug, version.Version), path)
	require.False(t, command.IsActive)

	// Verify JSON serialization
	jsonBytes, jsonErr := json.Marshal(command)
	require.NoError(t, jsonErr)
	require.JSONEq(t, fmt.Sprintf(`{"IsActive":%t}`, false), string(jsonBytes))
}

func TestPlatformHubPolicyVersionService_BuildGetVersionsPath_All(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	// All parameters
	query := VersionsQuery{
		Slug: "my_policy",
		Skip: 10,
		Take: 5,
	}

	path, err := buildGetVersionsPath(client, query)

	require.NoError(t, err)
	require.Equal(t, "/api/platformhub/policies/my_policy/versions?skip=10&take=5", path)
}

func TestPlatformHubPolicyVersionService_BuildGetVersionsPath_OnlySlug(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	query := VersionsQuery{
		Slug: "my_policy",
	}

	// Act
	path, err := buildGetVersionsPath(client, query)

	require.NoError(t, err)
	require.Equal(t, "/api/platformhub/policies/my_policy/versions", path)
}
