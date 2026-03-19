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
	policy := expectedPolicy.BuildDraft()
	command, path, commandError := buildAddCommand(client, *policy, commitMessage)

	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/%s/policies", urlEncodedGitRef), path)

	testAssertUpsertCommand(t, command, expectedPolicy, commitMessage)
	testAssertUpsertCommandMarshalJSON(t, command, expectedPolicy, commitMessage)
}

func TestPolicy_BuildAddCommand_Invalid(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	invalidPolicy := newPolicyBuilder().WithName("InvalidPolicyName").WithConditionsRego("").BuildDraft()

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
	version := publishedPolicyVersion{
		Slug:    "my_policy",
		Version: "1.0.0",
	}

	assertModifyVersionStatusCommand(t, version, true)
}

func TestPolicy_BuildDeactivateCommand(t *testing.T) {
	version := publishedPolicyVersion{
		Slug:    "my_policy",
		Version: "2.0.0",
	}

	assertModifyVersionStatusCommand(t, version, false)
}

func assertModifyVersionStatusCommand(t *testing.T, version publishedPolicyVersion, activated bool) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	// Act
	command, path, commandError := buildModifyVersionStatusCommand(client, &version, activated)

	// Assert
	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/policies/%s/versions/%s/modify-status", version.Slug, version.Version), path)
	require.Equal(t, command.IsActive, activated)

	// Verify JSON serialization
	jsonBytes, jsonErr := json.Marshal(command)
	require.NoError(t, jsonErr)
	require.JSONEq(t, fmt.Sprintf(`{"IsActive":%t}`, activated), string(jsonBytes))
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
	require.Equal(t, "/api/platformhub/policies/my_policy/versions/v2?skip=10&take=5", path)
}

func TestPolicy_BuildGetVersionsPath_OnlySlug(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	query := PublishedPoliciesQuery{
		Slug: "my_policy",
	}

	// Act
	path, err := buildGetVersionsPath(client, query)

	require.NoError(t, err)
	require.Equal(t, "/api/platformhub/policies/my_policy/versions/v2", path)
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

func TestPoliciesQueryResult_UnmarshalJSON(t *testing.T) {
	inputJSON := `{
		"Policies": [
			{
				"GitRef": "refs/heads/main",
				"Slug": "policy_one",
				"Name": "Policy One",
				"Description": "First policy",
				"ScopeRego": "package scope1",
				"ConditionsRego": "package cond1",
				"ViolationAction": "block",
				"ViolationReason": "reason1"
			},
			{
				"GitRef": "refs/heads/main",
				"Slug": "policy_two",
				"Name": "Policy Two",
				"ScopeRego": "package scope2",
				"ConditionsRego": "package cond2",
				"ViolationAction": "warn"
			}
		],
		"ItemsPerPage": 30,
		"FilteredItemsCount": 2,
		"TotalItemsCount": 5
	}`

	var result PoliciesQueryResult
	err := json.Unmarshal([]byte(inputJSON), &result)

	require.NoError(t, err)
	require.Equal(t, 30, result.ItemsPerPage)
	require.Equal(t, 2, result.FilteredItemsCount)
	require.Equal(t, 5, result.TotalItemsCount)
	require.Len(t, result.Policies, 2)

	require.Equal(t, "Policy One", result.Policies[0].GetName())
	require.Equal(t, "policy_one", result.Policies[0].GetSlug())
	require.Equal(t, "refs/heads/main", result.Policies[0].GetGitRef())
	require.Equal(t, "First policy", result.Policies[0].GetDescription())
	require.Equal(t, "block", result.Policies[0].GetViolationAction())
	require.Equal(t, "reason1", result.Policies[0].GetViolationReason())

	require.Equal(t, "Policy Two", result.Policies[1].GetName())
	require.Equal(t, "policy_two", result.Policies[1].GetSlug())
	require.Equal(t, "warn", result.Policies[1].GetViolationAction())
	require.Equal(t, "", result.Policies[1].GetDescription())
}

func TestPublishedPoliciesQueryResult_UnmarshalJSON(t *testing.T) {
	inputJSON := `{
		"Items": [
			{
				"Id": "pv-001",
				"Slug": "policy_one",
				"Version": "1.0.0",
				"PublishedDate": "2026-03-15T10:30:00Z",
				"GitRef": "refs/heads/main",
				"GitCommit": "abc123",
				"Name": "Policy One",
				"Description": "First version",
				"ViolationReason": "Missing step",
				"ViolationAction": "block",
				"RegoScope": "package scope1",
				"RegoConditions": "package cond1",
				"IsActive": true
			},
			{
				"Id": "pv-002",
				"Slug": "policy_one",
				"Version": "2.0.0",
				"PublishedDate": "2026-03-18T14:00:00Z",
				"GitRef": "refs/heads/main",
				"GitCommit": "def456",
				"Name": "Policy One v2",
				"ViolationAction": "warn",
				"RegoScope": "package scope2",
				"RegoConditions": "package cond2",
				"IsActive": false
			}
		],
		"ItemsPerPage": 20,
		"TotalResults": 2
	}`

	var result PublishedPoliciesQueryResult
	err := json.Unmarshal([]byte(inputJSON), &result)

	require.NoError(t, err)
	require.Equal(t, 20, result.ItemsPerPage)
	require.Equal(t, 2, result.TotalResults)
	require.Len(t, result.Items, 2)

	v1 := result.Items[0]
	require.Equal(t, "pv-001", v1.GetID())
	require.Equal(t, "policy_one", v1.GetSlug())
	require.Equal(t, "1.0.0", v1.GetVersion())
	require.Equal(t, "abc123", v1.GetGitCommit())
	require.Equal(t, "refs/heads/main", v1.GetGitRef())
	require.Equal(t, "Policy One", v1.GetName())
	require.Equal(t, "First version", v1.GetDescription())
	require.Equal(t, "Missing step", v1.GetViolationReason())
	require.Equal(t, "block", v1.GetViolationAction())
	require.Equal(t, "package scope1", v1.GetScopeRego())
	require.Equal(t, "package cond1", v1.GetConditionsRego())
	require.True(t, v1.IsActivated())

	v2 := result.Items[1]
	require.Equal(t, "pv-002", v2.GetID())
	require.Equal(t, "2.0.0", v2.GetVersion())
	require.Equal(t, "def456", v2.GetGitCommit())
	require.Equal(t, "Policy One v2", v2.GetName())
	require.Equal(t, "warn", v2.GetViolationAction())
	require.Equal(t, "", v2.GetDescription())
	require.False(t, v2.IsActivated())
}
