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
	policy := newPolicyBuilder().GitRef("refs/heads/test/branch")
	urlEncodedGitRef := "refs%2Fheads%2Ftest%2Fbranch"
	commitMessage := "Create new policy"

	newPolicy, policyError := policy.Create()
	require.NoError(t, policyError)

	// Act
	command, path, commandError := buildAddCommand(client, newPolicy, commitMessage)
	require.NoError(t, commandError)

	require.Equal(t, fmt.Sprintf("/api/platformhub/%s/policies", urlEncodedGitRef), *path)

	require.NotNil(t, command)
	require.Equal(t, policy.name, command.GetName())
	require.Equal(t, policy.gitRef, command.GetGitRef())
	require.Equal(t, policy.slug, command.GetSlug())
	require.Equal(t, policy.scopeRego, command.GetScopeRego())
	require.Equal(t, policy.conditionsRego, command.GetConditionsRego())
	require.Equal(t, policy.violationAction, command.GetViolationAction())
	require.Equal(t, commitMessage, command.ChangeDescription)
}

func TestPlatformHubPolicyService_AddCommandMarshalJSON(t *testing.T) {
	var client = newclient.NewClient(&newclient.HttpSession{})
	policy := newPolicyBuilder().GitRef("main").Description("Testing policy").ViolationReason("Missing manual intervention")
	commitMessage := "Add new policy as JSON"

	newPolicy, policyError := policy.Create()
	require.NoError(t, policyError)

	command, _, commandError := buildAddCommand(client, newPolicy, commitMessage)
	require.NoError(t, commandError)

	// Act
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
	}`, commitMessage, policy.name, policy.gitRef, policy.slug, policy.description, policy.scopeRego, policy.conditionsRego, policy.violationReason, policy.violationAction)

	jsonassert.New(t).Assertf(expectedJson, string(jsonCommand))
}
