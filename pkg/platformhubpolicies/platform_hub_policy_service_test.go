package platformhubpolicies

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestPlatformHubPolicyService_BuildAddCommand(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	// Policy is nil - should fail
	nilCommand, nilPath, commandNilError := buildAddCommand(client, nil, "commit nil command")
	require.ErrorContains(t, commandNilError, "policy")
	require.Nil(t, nilCommand)
	require.Empty(t, nilPath)

	// Invalid policy
	invalidPolicy := &PlatformHubPolicy{Name: "Invalid", GitRef: ""}
	invalidCommand, invalidPath, invalidCommandError := buildAddCommand(client, invalidPolicy, "commit invalid command")
	require.Error(t, invalidCommandError)
	require.Nil(t, invalidCommand)
	require.Empty(t, invalidPath)

	// Valid policy
	policy := newPolicyBuilder().
		GitRef("refs/heads/test/branch").
		Description("Testing policy").
		ViolationReason("Missing manual intervention")
	urlEncodedGitRef := "refs%2Fheads%2Ftest%2Fbranch"
	commitMessage := "Create new policy"

	newPolicy, policyError := policy.Create()
	require.NoError(t, policyError)

	// Act
	command, path, commandError := buildAddCommand(client, newPolicy, commitMessage)
	require.NoError(t, commandError)

	require.Equal(t, fmt.Sprintf("/api/platformhub/%s/policies", urlEncodedGitRef), path)

	testAssertUpsertCommand(t, command, policy, commitMessage)
	testAssertUpsertCommandMarshalJSON(t, command, policy, commitMessage)
}

func TestPlatformHubPolicyService_BuildUpdateCommand(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	// Policy is nil - should fail
	nilCommand, nilPath, commandNilError := buildUpdateCommand(client, nil, "commit nil command")
	require.ErrorContains(t, commandNilError, "policy")
	require.Nil(t, nilCommand)
	require.Empty(t, nilPath)

	// Invalid policy
	invalidPolicy := &PlatformHubPolicy{Name: "Invalid", GitRef: "main", ViolationAction: ""}
	invalidCommand, invalidPath, invalidCommandError := buildUpdateCommand(client, invalidPolicy, "commit invalid command")
	require.Error(t, invalidCommandError)
	require.Nil(t, invalidCommand)
	require.Empty(t, invalidPath)

	// Valid policy
	policy := newPolicyBuilder().Name("Valid Policy").
		GitRef("refs/heads/main").
		Slug("valid_policy").
		Description("Ok").
		ViolationReason("None")
	urlEncodedGitRef := "refs%2Fheads%2Fmain"
	commitMessage := "Update valid policy"

	newPolicy, policyError := policy.Create()
	require.NoError(t, policyError)

	// Act
	command, path, commandError := buildUpdateCommand(client, newPolicy, commitMessage)
	require.NoError(t, commandError)

	require.Equal(t, fmt.Sprintf("/api/platformhub/%s/policies/%s", urlEncodedGitRef, policy.slug), path)

	testAssertUpsertCommand(t, command, policy, commitMessage)
	testAssertUpsertCommandMarshalJSON(t, command, policy, commitMessage)
}

func testAssertUpsertCommand(t *testing.T, command *platformHubPolicyUpsertCommand, expected *policyBuilder, expectedCommit string) {
	require.NotNil(t, command)
	require.Equal(t, expected.name, command.GetName())
	require.Equal(t, expected.gitRef, command.GetGitRef())
	require.Equal(t, expected.slug, command.GetSlug())
	require.Equal(t, expected.scopeRego, command.GetScopeRego())
	require.Equal(t, expected.conditionsRego, command.GetConditionsRego())
	require.Equal(t, expected.violationAction, command.GetViolationAction())
	require.Equal(t, expectedCommit, command.ChangeDescription)
}

func testAssertUpsertCommandMarshalJSON(t *testing.T, command *platformHubPolicyUpsertCommand, expected *policyBuilder, expectedCommit string) {
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
