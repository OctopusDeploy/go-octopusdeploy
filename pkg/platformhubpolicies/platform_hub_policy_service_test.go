package platformhubpolicies

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/newclient"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestPlatformHubPolicyService_BuildAddCommand_Valid(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	expectedPolicy := newPolicyBuilder().
		WithName("ValidPolicyName").
		WithGitRef("refs/heads/test/branch").
		WithDescription("Testing policy").
		WithViolationReason("Missing manual intervention")
	urlEncodedGitRef := "refs%2Fheads%2Ftest%2Fbranch"
	commitMessage := "Create new policy"

	// Act
	newPolicy := expectedPolicy.Build()
	command, path, commandError := buildAddCommand(client, *newPolicy, commitMessage)

	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/%s/policies", urlEncodedGitRef), path)

	testAssertUpsertCommand(t, command, expectedPolicy, commitMessage)
	testAssertUpsertCommandMarshalJSON(t, command, expectedPolicy, commitMessage)
}

func TestPlatformHubPolicyService_BuildAddCommand_Invalid(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	invalidPolicy := newPolicyBuilder().WithName("InvalidPolicyName").WithConditionsRego("").Build()

	_, _, invalidCommandError := buildAddCommand(client, *invalidPolicy, "commit invalid command")

	require.ErrorContains(t, invalidCommandError, "ConditionsRego")
}

func TestPlatformHubPolicyService_BuildUpdateCommand_Valid(t *testing.T) {
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
	command, path, commandError := buildUpdateCommand(client, *newPolicy, commitMessage)

	require.NoError(t, commandError)
	require.Equal(t, fmt.Sprintf("/api/platformhub/%s/policies/%s", urlEncodedGitRef, expectedPolicy.slug), path)

	testAssertUpsertCommand(t, command, expectedPolicy, commitMessage)
	testAssertUpsertCommandMarshalJSON(t, command, expectedPolicy, commitMessage)
}

func TestPlatformHubPolicyService_BuildUpdateCommand_Invalid(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})

	invalidPolicy := newPolicyBuilder().WithName("Invalid Action").WithViolationAction("").Build()
	_, _, invalidCommandError := buildUpdateCommand(client, *invalidPolicy, "commit invalid command")

	require.ErrorContains(t, invalidCommandError, "ViolationAction")
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
