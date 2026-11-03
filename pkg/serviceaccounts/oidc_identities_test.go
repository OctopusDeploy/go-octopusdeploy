package serviceaccounts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOIDCIdentity_MarshalJSON_OmitsEmptyAudience(t *testing.T) {
	identity := NewOIDCIdentity("ServiceAccounts-1", "name", "issuer", "subject")

	data, err := json.Marshal(identity)
	assert.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)

	_, hasAudience := result["Audience"]
	assert.False(t, hasAudience, "Audience must be omitted from JSON when unset")
}

func TestOIDCIdentity_MarshalJSON_IncludesAudienceWhenSet(t *testing.T) {
	identity := NewOIDCIdentity("ServiceAccounts-1", "name", "issuer", "subject")
	identity.Audience = "api://custom"

	data, err := json.Marshal(identity)
	assert.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)

	assert.Equal(t, "api://custom", result["Audience"])
}

func TestOIDCIdentity_MarshalJSON_SendsNullClaimConditionsWhenUnset(t *testing.T) {
	identity := NewOIDCIdentity("ServiceAccounts-1", "name", "issuer", "subject")

	data, err := json.Marshal(identity)
	assert.NoError(t, err)

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)

	value, hasClaimConditions := result["ClaimConditions"]
	assert.True(t, hasClaimConditions)
	assert.Nil(t, value, "A nil slice must be sent as null so Octopus keeps the existing conditions")
}

func TestOIDCIdentity_MarshalJSON_SendsEmptyClaimConditionsToClearThem(t *testing.T) {
	identity := NewOIDCIdentity("ServiceAccounts-1", "name", "issuer", "subject")
	identity.ClaimConditions = []OIDCIdentityClaimCondition{}

	data, err := json.Marshal(identity)
	assert.NoError(t, err)

	assert.Contains(t, string(data), `"ClaimConditions":[]`)
}

func TestOIDCIdentity_JSONRoundTrip_ClaimConditions(t *testing.T) {
	identity := NewOIDCIdentity("ServiceAccounts-1", "name", "issuer", "subject")
	identity.ClaimConditions = []OIDCIdentityClaimCondition{
		{Claim: "repository_owner_id", Values: []string{"123456"}},
		{Claim: "environment", Values: []string{"prod", "staging-*"}},
	}

	data, err := json.Marshal(identity)
	assert.NoError(t, err)

	var result OIDCIdentity
	err = json.Unmarshal(data, &result)
	assert.NoError(t, err)

	assert.Equal(t, identity.ClaimConditions, result.ClaimConditions)
}
