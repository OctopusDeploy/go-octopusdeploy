package ratelimitingpolicies

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRateLimitingPolicyScopeTypeJsonMarshal(t *testing.T) {
	cases := map[RateLimitingPolicyScopeType]string{
		Unauthenticated:    `"Unauthenticated"`,
		AuthenticatedHuman: `"AuthenticatedHuman"`,
		AuthenticatedAgent: `"AuthenticatedAgent"`,
	}
	for scope, expected := range cases {
		jsonValue, err := json.Marshal(scope)
		require.NoError(t, err)
		require.JSONEq(t, expected, string(jsonValue))

		var enumValue RateLimitingPolicyScopeType
		require.NoError(t, json.Unmarshal(jsonValue, &enumValue))
		require.Equal(t, scope, enumValue)
	}
}

func TestRateLimitingPolicyScopeTypeJsonUnmarshalInvalid(t *testing.T) {
	var scope RateLimitingPolicyScopeType
	require.Error(t, json.Unmarshal([]byte(`"NotAScope"`), &scope))
}

func TestRateLimitingPolicyMarshalRoundTrip(t *testing.T) {
	policy := RateLimitingPolicy{
		ID:                "RateLimitingPolicies-2",
		IsBuiltIn:         true,
		Name:              "Authenticated requests",
		IsEnabled:         true,
		ScopeType:         AuthenticatedHuman,
		RequestsPerMinute: 10_000,
		BurstLimit:        5_000,
		AuditMode:         true,
	}

	data, err := json.Marshal(policy)
	require.NoError(t, err)

	expected := `{
		"Id": "RateLimitingPolicies-2",
		"IsBuiltIn": true,
		"Name": "Authenticated requests",
		"IsEnabled": true,
		"ScopeType": "AuthenticatedHuman",
		"RequestsPerMinute": 10000,
		"BurstLimit": 5000,
		"AuditMode": true
	}`
	require.JSONEq(t, expected, string(data))

	var received RateLimitingPolicy
	require.NoError(t, json.Unmarshal(data, &received))
	require.Equal(t, policy, received)
}

func TestModifyRateLimitingPolicyCommandMarshal(t *testing.T) {
	command := ModifyRateLimitingPolicyCommand{
		ID:                "RateLimitingPolicies-1",
		Name:              "Changed",
		IsEnabled:         true,
		ScopeType:         Unauthenticated,
		RequestsPerMinute: 123,
		BurstLimit:        456,
		AuditMode:         true,
	}

	data, err := json.Marshal(command)
	require.NoError(t, err)
	require.JSONEq(t, `{
		"Name": "Changed",
		"IsEnabled": true,
		"ScopeType": "Unauthenticated",
		"RequestsPerMinute": 123,
		"BurstLimit": 456,
		"AuditMode": true
	}`, string(data))
}

func TestListRateLimitingPoliciesResponseUnmarshal(t *testing.T) {
	payload := `{
		"ItemType": "RateLimitingPolicy",
		"TotalResults": 2,
		"ItemsPerPage": 30,
		"NumberOfPages": 1,
		"LastPageNumber": 0,
		"Items": [
			{
				"Id": "RateLimitingPolicies-1",
				"Name": "Human",
				"IsBuiltIn": true,
				"ScopeType": "AuthenticatedHuman",
				"IsEnabled": true,
				"RequestsPerMinute": 1000,
				"BurstLimit": 50,
				"AuditMode": true
			},
			{
				"Id": "RateLimitingPolicies-2",
				"Name": "Agent",
				"IsBuiltIn": true,
				"ScopeType": "AuthenticatedAgent",
				"IsEnabled": false,
				"RequestsPerMinute": 500,
				"BurstLimit": 25,
				"AuditMode": false
			}
		]
	}`

	var response ListRateLimitingPoliciesResponse
	require.NoError(t, json.Unmarshal([]byte(payload), &response))

	require.Equal(t, 2, response.TotalResults)
	require.Len(t, response.Items, 2)
	require.Equal(t, AuthenticatedHuman, response.Items[0].ScopeType)
	require.Equal(t, AuthenticatedAgent, response.Items[1].ScopeType)
	require.False(t, response.Items[1].IsEnabled)
	require.True(t, response.Items[0].AuditMode)
	require.False(t, response.Items[1].AuditMode)
}
