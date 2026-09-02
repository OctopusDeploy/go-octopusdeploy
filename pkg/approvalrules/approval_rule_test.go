package approvalrules

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestApprovalRuleTagScopeMarshalRoundTrip(t *testing.T) {
	rule := NewApprovalRule("Production approvals")
	rule.ID = "ApprovalRules-1"
	rule.SpaceID = "Spaces-1"
	rule.Description = "Requires two approvers"
	rule.ScopingStrategy = ApprovalRuleScopingStrategyTag
	rule.TenantApprovalStrategy = TenantApprovalStrategyPerTenant
	rule.MinimumApproversRequired = 2
	rule.AllowSelfApproval = false
	rule.IsDisabled = false
	rule.TagScopes = []ApprovalRuleTagScope{
		{ProjectTags: []string{"tenant/prod"}, EnvironmentTags: []string{"env/prod"}},
	}
	rule.IdScopes = []ApprovalRuleIdScope{}
	rule.ApprovingUserIds = []string{"Users-1"}
	rule.ApprovingTeamIds = []string{"Teams-1"}

	data, err := json.Marshal(rule)
	require.NoError(t, err)

	expected := `{
		"Id": "ApprovalRules-1",
		"SpaceId": "Spaces-1",
		"Name": "Production approvals",
		"Description": "Requires two approvers",
		"ScopingStrategy": "Tag",
		"TenantApprovalStrategy": "PerTenant",
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

// TestApprovalRuleIdScopeUnmarshal decodes a payload in the shape the server
// actually returns, including the scope id prefix and tenant approval strategy.
func TestApprovalRuleIdScopeUnmarshal(t *testing.T) {
	payload := `{
		"SpaceId": "Spaces-1",
		"Id": "ApprovalRules-1",
		"Name": "Test Approval",
		"Description": "",
		"ScopingStrategy": "Id",
		"TenantApprovalStrategy": "PerRelease",
		"TagScopes": [],
		"IdScopes": [
			{
				"Id": "ApprovalRuleIdScopes-1",
				"ProjectId": "Projects-622",
				"EnvironmentIds": ["Environments-102", "Environments-101"]
			}
		],
		"MinimumApproversRequired": 1,
		"AllowSelfApproval": true,
		"IsDisabled": false,
		"ApprovingUserIds": ["Users-61"],
		"ApprovingTeamIds": []
	}`
	var rule ApprovalRule
	require.NoError(t, json.Unmarshal([]byte(payload), &rule))

	require.Equal(t, "ApprovalRules-1", rule.GetID())
	require.Equal(t, "Test Approval", rule.GetName())
	require.Equal(t, ApprovalRuleScopingStrategyId, rule.ScopingStrategy)
	require.Equal(t, TenantApprovalStrategyPerRelease, rule.TenantApprovalStrategy)
	require.Len(t, rule.IdScopes, 1)
	require.Equal(t, "ApprovalRuleIdScopes-1", rule.IdScopes[0].Id)
	require.Equal(t, "Projects-622", rule.IdScopes[0].ProjectId)
	require.Equal(t, []string{"Environments-102", "Environments-101"}, rule.IdScopes[0].EnvironmentIds)
	require.Equal(t, 1, rule.MinimumApproversRequired)
	require.True(t, rule.AllowSelfApproval)
	require.False(t, rule.IsDisabled)
	require.Equal(t, []string{"Users-61"}, rule.ApprovingUserIds)
}

func TestNewApprovalRuleDefaults(t *testing.T) {
	rule := NewApprovalRule("example")

	require.Equal(t, "example", rule.Name)
	require.Equal(t, ApprovalRuleScopingStrategyTag, rule.ScopingStrategy)
	require.Equal(t, TenantApprovalStrategyPerRelease, rule.TenantApprovalStrategy)
	require.Equal(t, 2, rule.MinimumApproversRequired)
	require.NotNil(t, rule.TagScopes)
	require.NotNil(t, rule.IdScopes)
	require.NotNil(t, rule.ApprovingUserIds)
	require.NotNil(t, rule.ApprovingTeamIds)
}

// TestApprovalRuleOmitsUnsetMinimumApprovers guards that an unset value is left
// out of the request. Zero is never valid, and the server rejects it outright,
// whereas omitting the field applies the server default.
func TestApprovalRuleOmitsUnsetMinimumApprovers(t *testing.T) {
	rule := NewApprovalRule("example")
	rule.MinimumApproversRequired = 0

	data, err := json.Marshal(rule)
	require.NoError(t, err)
	require.NotContains(t, string(data), "MinimumApproversRequired")

	rule.MinimumApproversRequired = 3
	data, err = json.Marshal(rule)
	require.NoError(t, err)
	require.Contains(t, string(data), `"MinimumApproversRequired":3`)
}

func TestTenantApprovalStrategyValues(t *testing.T) {
	cases := map[TenantApprovalStrategy]string{
		TenantApprovalStrategyPerRelease: "PerRelease",
		TenantApprovalStrategyPerTenant:  "PerTenant",
	}
	for strategy, want := range cases {
		rule := NewApprovalRule("example")
		rule.TenantApprovalStrategy = strategy

		data, err := json.Marshal(rule)
		require.NoError(t, err)
		require.Contains(t, string(data), `"TenantApprovalStrategy":"`+want+`"`)
	}
}
