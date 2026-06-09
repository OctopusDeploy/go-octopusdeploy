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
