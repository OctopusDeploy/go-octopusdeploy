package octopusdeploy

import (
	"encoding/json"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestSensitiveValueBehaviour(t *testing.T) {
	sv := SensitiveValue{}

	require.NotNil(t, sv)
	require.False(t, sv.HasValue)
	require.Nil(t, sv.Hint)
	require.Nil(t, sv.NewValue)

	svp := &SensitiveValue{}

	require.NotNil(t, svp)
	require.False(t, svp.HasValue)
	require.Nil(t, sv.Hint)
	require.Nil(t, sv.NewValue)

	sensitiveValueAsJSON, err := json.Marshal(svp)
	require.NoError(t, err)
	require.NotNil(t, sensitiveValueAsJSON)
	jsonassert.New(t).Assertf(string(sensitiveValueAsJSON), emptySensitiveValueAsJSON)

	svp = &SensitiveValue{
		HasValue: true,
	}

	require.NotNil(t, svp)
	require.True(t, svp.HasValue)
	require.Nil(t, sv.Hint)
	require.Nil(t, sv.NewValue)

	sensitiveValueAsJSON, err = json.Marshal(svp)
	require.NoError(t, err)
	require.NotNil(t, sensitiveValueAsJSON)
	jsonassert.New(t).Assertf(string(sensitiveValueAsJSON), emptyHasValueSensitiveValueAsJSON)

	svp = &SensitiveValue{
		HasValue: true,
		Hint:     String("this is the hint"),
		NewValue: String("test"),
	}

	require.NotNil(t, svp)
	require.True(t, svp.HasValue)
	require.Equal(t, "this is the hint", *svp.Hint)
	require.Equal(t, "test", *svp.NewValue)

	sensitiveValueAsJSON, err = json.Marshal(svp)
	require.NoError(t, err)
	require.NotNil(t, sensitiveValueAsJSON)
	jsonassert.New(t).Assertf(string(sensitiveValueAsJSON), testWithHintSensitiveValueAsJSON)
}

func TestNewSensitiveValueBehaviour(t *testing.T) {
	svp := NewSensitiveValue("")

	require.NotNil(t, svp)
	require.False(t, svp.HasValue)
	require.Nil(t, svp.Hint)
	require.Nil(t, svp.NewValue)

	svp = NewSensitiveValue("test")

	require.NotNil(t, svp)
	require.True(t, svp.HasValue)
	require.Nil(t, svp.Hint)
	require.Equal(t, "test", *svp.NewValue)
}

func TestNewSensitiveValueMarshalJSON(t *testing.T) {
	sensitiveValue := NewSensitiveValue("")

	sensitiveValueAsJSON, err := json.Marshal(sensitiveValue)
	require.NoError(t, err)
	require.NotNil(t, sensitiveValueAsJSON)
	jsonassert.New(t).Assertf(string(sensitiveValueAsJSON), emptySensitiveValueAsJSON)

	sensitiveValue = NewSensitiveValue("test")

	sensitiveValueAsJSON, err = json.Marshal(sensitiveValue)
	require.NoError(t, err)
	require.NotNil(t, sensitiveValueAsJSON)
	jsonassert.New(t).Assertf(string(sensitiveValueAsJSON), testNoHintSensitiveValueAsJSON)
}

func TestNewSensitiveValueUnmarshalJSON(t *testing.T) {
	var emptySensitiveValue SensitiveValue
	err := json.Unmarshal([]byte(emptySensitiveValueAsJSON), &emptySensitiveValue)

	require.NoError(t, err)
	require.NotNil(t, emptySensitiveValue)
	require.False(t, emptySensitiveValue.HasValue)
	require.Nil(t, emptySensitiveValue.NewValue)

	var testSensitiveValue SensitiveValue
	err = json.Unmarshal([]byte(testNoHintSensitiveValueAsJSON), &testSensitiveValue)

	require.NoError(t, err)
	require.NotNil(t, testSensitiveValue)
	require.True(t, testSensitiveValue.HasValue)
	require.Equal(t, "test", *testSensitiveValue.NewValue)
}

const emptySensitiveValueAsJSON string = `{
	"HasValue": false,
	"Hint": null,
	"NewValue": null
  }`

const emptyHasValueSensitiveValueAsJSON string = `{
	"HasValue": true,
	"Hint": null,
	"NewValue": null
  }`

const testNoHintSensitiveValueAsJSON string = `{
	"HasValue": true,
	"Hint": null,
	"NewValue": "test"
  }`

const testWithHintSensitiveValueAsJSON string = `{
	"HasValue": true,
	"Hint": "this is the hint",
	"NewValue": "test"
  }`
