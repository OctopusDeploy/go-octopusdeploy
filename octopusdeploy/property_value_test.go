package octopusdeploy

import (
	"encoding/json"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPropertyValueBehaviour(t *testing.T) {
	pv := SensitiveValue{}

	require.NotNil(t, pv)
	require.False(t, pv.HasValue)
	require.Len(t, pv.NewValue, 0)

	pvp := &SensitiveValue{}

	require.NotNil(t, pvp)
	require.False(t, pvp.HasValue)
	require.Len(t, pvp.NewValue, 0)
}

func TestNewPropertyValueBehaviour(t *testing.T) {
	pvp := NewPropertyValue("", false)
	require.NotNil(t, pvp)
	require.False(t, pvp.IsSensitive)
	require.Nil(t, pvp.SensitiveValue)
	require.Len(t, pvp.Value, 0)

	pvp = NewPropertyValue("", true)
	require.NotNil(t, pvp)
	require.True(t, pvp.IsSensitive)
	require.NotNil(t, pvp.SensitiveValue)
	require.True(t, pvp.SensitiveValue.HasValue)
	require.Len(t, pvp.SensitiveValue.NewValue, 0)
	require.Len(t, pvp.Value, 0)

	pvp = NewPropertyValue("test", false)
	require.NotNil(t, pvp)
	require.False(t, pvp.IsSensitive)
	require.Nil(t, pvp.SensitiveValue)
	require.Equal(t, "test", pvp.Value)

	pvp = NewPropertyValue("sensitive value", true)
	require.NotNil(t, pvp)
	require.True(t, pvp.IsSensitive)
	require.NotNil(t, pvp.SensitiveValue)
	require.True(t, pvp.SensitiveValue.HasValue)
	require.Equal(t, "sensitive value", pvp.SensitiveValue.NewValue)
	require.Len(t, pvp.Value, 0)
}

func TestNewPropertyValueMarshalJSON(t *testing.T) {
	propertyValue := NewPropertyValue("", true)

	propertyValueAsJSON, err := json.Marshal(propertyValue)
	require.NoError(t, err)
	require.NotNil(t, propertyValueAsJSON)

	jsonassert.New(t).Assertf(string(propertyValueAsJSON), emptySensitiveValueAsJSON)

	propertyValue = NewPropertyValue("test", true)

	propertyValueAsJSON, err = json.Marshal(propertyValue)
	require.NoError(t, err)
	require.NotNil(t, propertyValueAsJSON)

	jsonassert.New(t).Assertf(string(propertyValueAsJSON), testSensitivePropertyValueAsJSON)
}

func TestNewPropertyValueUnmarshalJSON(t *testing.T) {
	var emptyPropertyValue SensitiveValue
	err := json.Unmarshal([]byte(emptySensitiveValueAsJSON), &emptyPropertyValue)
	require.NoError(t, err)
	require.NotNil(t, emptyPropertyValue)

	assert.False(t, emptyPropertyValue.HasValue)
	assert.Len(t, emptyPropertyValue.NewValue, 0)

	var testPropertyValue SensitiveValue
	err = json.Unmarshal([]byte(testSensitivePropertyValueAsJSON), &testPropertyValue)
	require.NoError(t, err)
	require.NotNil(t, testPropertyValue)

	assert.True(t, testPropertyValue.HasValue)
	assert.Equal(t, "test", testPropertyValue.NewValue)
}

const emptyPropertyValueAsJSON string = `{
	"IsSensitive": false
	"SensitiveValue": null,
	"Value": ""
  }`

const emptySensitivePropertyValueAsJSON string = `{
	"IsSensitive": true
	"SensitiveValue": {
		"HasValue": false,
		"NewValue": null
	},
	"Value": ""
  }`

const testSensitivePropertyValueAsJSON string = `{
	"IsSensitive": true
	"SensitiveValue": {
		"HasValue": true,
		"NewValue": "test"
	},
	"Value": ""
  }`

const testNonSensitivePropertyValueAsJSON string = `{
	"IsSensitive": false
	"SensitiveValue": null,
	"Value": "non-sensitive value"
  }`
