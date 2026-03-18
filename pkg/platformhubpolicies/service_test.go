package platformhubpolicies

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestPolicy_BuildAddCommand_Valid(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	expectedPolicy := newPolicyBuilder().
		WithName("ValidPolicyName").
		WithGitRef("refs/heads/test/branch").
		WithDescription("Testing policy").
		WithViolationReason("Missing manual intervention")
	urlEncodedGitRef := "refs%2Fheads%2Ftest%2Fbranch"
	commitMessage := "Create new policy"

	// Act
	policy := expectedPolicy.BuildCandidate()
	command, path, commandError := buildAddCommand(client, *policy, commitMessage)

	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/%s/policies", urlEncodedGitRef), path)

	testAssertUpsertCommand(t, command, expectedPolicy, commitMessage)
	testAssertUpsertCommandMarshalJSON(t, command, expectedPolicy, commitMessage)
}

func TestPolicy_BuildAddCommand_Invalid(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	invalidPolicy := newPolicyBuilder().WithName("InvalidPolicyName").WithConditionsRego("").BuildCandidate()

	_, _, invalidCommandError := buildAddCommand(client, *invalidPolicy, "commit invalid command")

	require.ErrorContains(t, invalidCommandError, "ConditionsRego")
}

func TestPolicy_BuildUpdateCommand_Valid(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	expectedPolicy := newPolicyBuilder().WithName("Valid Policy").
		WithGitRef("refs/heads/main").
		WithSlug("valid_policy").
		WithDescription("Ok").
		WithViolationReason("None")
	urlEncodedGitRef := "refs%2Fheads%2Fmain"
	commitMessage := "Update valid policy"

	// Act
	newPolicy := expectedPolicy.Build()
	command, path, commandError := buildUpdateCommand(client, newPolicy, commitMessage)

	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/%s/policies/%s", urlEncodedGitRef, expectedPolicy.slug), path)

	testAssertUpsertCommand(t, command, expectedPolicy, commitMessage)
	testAssertUpsertCommandMarshalJSON(t, command, expectedPolicy, commitMessage)
}

func TestPolicy_BuildUpdateCommand_Invalid(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	invalidPolicy := newPolicyBuilder().WithName("Invalid Action").WithViolationAction("").Build()
	_, _, invalidCommandError := buildUpdateCommand(client, invalidPolicy, "commit invalid command")

	require.ErrorContains(t, invalidCommandError, "ViolationAction")
}

func TestPolicy_BuildPublishCommand(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	gitRef := "refs/heads/main"
	slug := "my_policy"
	version := "1.0.1"
	urlEncodedGitRef := "refs%2Fheads%2Fmain"

	policy := newPolicyBuilder().WithGitRef(gitRef).WithSlug(slug).Build()

	// Act
	command, path, commandError := buildPublishCommand(client, policy, version)

	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/%s/policies/%s/publish", urlEncodedGitRef, slug), path)
	require.Equal(t, version, command.Version)

	// Verify JSON serialization
	jsonBytes, jsonErr := json.Marshal(command)
	require.NoError(t, jsonErr)
	require.JSONEq(t, fmt.Sprintf(`{"Version":"%s"}`, version), string(jsonBytes))
}

func TestPolicy_BuildActivateCommand(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	version := publishedPolicyVersion{
		Slug:    "my_policy",
		Version: "1.0.0",
	}

	// Act
	command, path, commandError := buildModifyVersionStatusCommand(client, &version, true)

	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/policies/%s/versions/%s/modify-status", version.Slug, version.Version), path)
	require.True(t, command.IsActive)

	// Verify JSON serialization
	jsonBytes, jsonErr := json.Marshal(command)
	require.NoError(t, jsonErr)
	require.JSONEq(t, fmt.Sprintf(`{"IsActive":%t}`, true), string(jsonBytes))
}

func TestPolicy_BuildDeactivateCommand(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	version := publishedPolicyVersion{
		Slug:    "my_policy",
		Version: "2.0.0",
	}

	// Act
	command, path, commandError := buildModifyVersionStatusCommand(client, &version, false)

	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/policies/%s/versions/%s/modify-status", version.Slug, version.Version), path)
	require.False(t, command.IsActive)

	// Verify JSON serialization
	jsonBytes, jsonErr := json.Marshal(command)
	require.NoError(t, jsonErr)
	require.JSONEq(t, fmt.Sprintf(`{"IsActive":%t}`, false), string(jsonBytes))
}

func TestPolicy_BuildGetVersionsPath_All(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	// All parameters
	query := PublishedPoliciesQuery{
		Slug: "my_policy",
		Skip: 10,
		Take: 5,
	}

	path, err := buildGetVersionsPath(client, query)

	require.NoError(t, err)
	require.Equal(t, "/api/platformhub/policies/my_policy/versions?skip=10&take=5", path)
}

func TestPolicy_BuildGetVersionsPath_OnlySlug(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	query := PublishedPoliciesQuery{
		Slug: "my_policy",
	}

	// Act
	path, err := buildGetVersionsPath(client, query)

	require.NoError(t, err)
	require.Equal(t, "/api/platformhub/policies/my_policy/versions", path)
}

func testAssertUpsertCommand(t *testing.T, command platformHubPolicyUpsertCommand, expected *policyBuilder, expectedCommit string) {
	require.NotNil(t, command)
	require.Equal(t, expected.name, command.GetName())
	require.Equal(t, expected.gitRef, command.GetGitRef())
	require.Equal(t, expected.slug, command.GetSlug())
	require.Equal(t, expected.scopeRego, command.GetScopeRego())
	require.Equal(t, expected.conditionsRego, command.GetConditionsRego())
	require.Equal(t, expected.violationAction, command.GetViolationAction())
	require.Equal(t, expectedCommit, command.ChangeDescription)
}

func testAssertUpsertCommandMarshalJSON(t *testing.T, command platformHubPolicyUpsertCommand, expected *policyBuilder, expectedCommit string) {
	jsonCommand, jsonError := json.Marshal(command)
	require.NoError(t, jsonError)
	require.NotNil(t, jsonCommand)

	expectedJson := fmt.Sprintf(`{
		"ChangeDescription": "%s",
		"Name": "%s",
		"GitRef": "%s",
		"Slug": "%s",
		"Description": "%s",
		"ScopeRego": "%s",
		"ConditionsRego": "%s",
		"ViolationReason": "%s",
		"ViolationAction": "%s"
	}`, expectedCommit, expected.name, expected.gitRef, expected.slug, expected.description, expected.scopeRego, expected.conditionsRego, expected.violationReason, expected.violationAction)

	jsonassert.New(t).Assertf(expectedJson, string(jsonCommand))
}
