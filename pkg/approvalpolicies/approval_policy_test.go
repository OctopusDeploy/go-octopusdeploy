package approvalpolicies

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApprovalPolicyTagScopeMarshalRoundTrip(t *testing.T) {
	policy := NewApprovalPolicy("Production approvals")
	policy.ID = "ApprovalPolicies-1"
	policy.SpaceID = "Spaces-1"
	policy.Description = "Requires two approvers"
	policy.ScopingStrategy = ApprovalPolicyScopingStrategyTag
	policy.MinimumApproversRequired = 2
	policy.AllowSelfApproval = false
	policy.IsDisabled = false
	policy.TagScopes = []ApprovalPolicyTagScope{
		{ProjectTags: []string{"tenant/prod"}, EnvironmentTags: []string{"env/prod"}},
	}
	policy.IdScopes = []ApprovalPolicyIdScope{}
	policy.ApprovingUserIds = []string{"Users-1"}
	policy.ApprovingTeamIds = []string{"Teams-1"}

	data, err := json.Marshal(policy)
	require.NoError(t, err)

	expected := `{
		"Id": "ApprovalPolicies-1",
		"SpaceId": "Spaces-1",
		"Name": "Production approvals",
		"Description": "Requires two approvers",
		"ScopingStrategy": "Tag",
		"TagScopes": [{"ProjectTags":["tenant/prod"],"EnvironmentTags":["env/prod"]}],
		"IdScopes": [],
		"MinimumApproversRequired": 2,
		"AllowSelfApproval": false,
		"IsDisabled": false,
		"ApprovingUserIds": ["Users-1"],
		"ApprovingTeamIds": ["Teams-1"]
	}`
	require.JSONEq(t, expected, string(data))
}

func TestApprovalPolicyIdScopeUnmarshal(t *testing.T) {
	payload := `{
		"Id": "ApprovalPolicies-2",
		"SpaceId": "Spaces-1",
		"Name": "Id scoped",
		"ScopingStrategy": "Id",
		"TagScopes": [],
		"IdScopes": [{"Id":"scope-1","ProjectId":"Projects-1","EnvironmentIds":["Environments-1","Environments-2"]}],
		"MinimumApproversRequired": 1,
		"AllowSelfApproval": true,
		"IsDisabled": true,
		"ApprovingUserIds": [],
		"ApprovingTeamIds": ["Teams-9"]
	}`
	var policy ApprovalPolicy
	require.NoError(t, json.Unmarshal([]byte(payload), &policy))

	require.Equal(t, "ApprovalPolicies-2", policy.GetID())
	require.Equal(t, ApprovalPolicyScopingStrategyId, policy.ScopingStrategy)
	require.Len(t, policy.IdScopes, 1)
	require.Equal(t, "Projects-1", policy.IdScopes[0].ProjectId)
	require.Equal(t, []string{"Environments-1", "Environments-2"}, policy.IdScopes[0].EnvironmentIds)
	require.True(t, policy.AllowSelfApproval)
	require.True(t, policy.IsDisabled)
	require.Equal(t, "Id scoped", policy.GetName())
}
