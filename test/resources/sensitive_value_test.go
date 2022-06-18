package resources

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/pkg/core"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestSensitiveValueBehaviour(t *testing.T) {
	sv := core.SensitiveValue{}

	require.NotNil(t, sv)
	require.False(t, sv.HasValue)
	require.Nil(t, sv.Hint)
	require.Nil(t, sv.NewValue)

	svp := &core.SensitiveValue{}

	require.NotNil(t, svp)
	require.False(t, svp.HasValue)
	require.Nil(t, sv.Hint)
	require.Nil(t, sv.NewValue)

	sensitiveValueAsJSON, err := json.Marshal(svp)
	require.NoError(t, err)
	require.NotNil(t, sensitiveValueAsJSON)
	jsonassert.New(t).Assertf(string(sensitiveValueAsJSON), emptySensitiveValueAsJSON)

	svp = &core.SensitiveValue{
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

	hasValue := true
	hint := internal.GetRandomName()
	newValue := internal.GetRandomName()

	svp = &core.SensitiveValue{
		HasValue: hasValue,
		Hint:     &hint,
		NewValue: &newValue,
	}

	require.NotNil(t, svp)
	require.Equal(t, hasValue, svp.HasValue)
	require.Equal(t, hint, *svp.Hint)
	require.Equal(t, newValue, *svp.NewValue)

	sensitiveValueAsJSON, err = json.Marshal(svp)
	require.NoError(t, err)
	require.NotNil(t, sensitiveValueAsJSON)

	testWithHintSensitiveValueAsJSON := fmt.Sprintf(`{
		"HasValue": %v,
		"Hint": "%s",
		"NewValue": "%s"
  	}`, hasValue, hint, newValue)

	jsonassert.New(t).Assertf(string(sensitiveValueAsJSON), testWithHintSensitiveValueAsJSON)
}

func TestNewSensitiveValueBehaviour(t *testing.T) {
	svp := core.NewSensitiveValue("")

	require.NotNil(t, svp)
	require.False(t, svp.HasValue)
	require.Nil(t, svp.Hint)
	require.Nil(t, svp.NewValue)

	newValue := internal.GetRandomName()

	svp = core.NewSensitiveValue(newValue)

	require.NotNil(t, svp)
	require.True(t, svp.HasValue)
	require.Nil(t, svp.Hint)
	require.Equal(t, newValue, *svp.NewValue)
}

func TestNewSensitiveValueMarshalJSON(t *testing.T) {
	sensitiveValue := core.NewSensitiveValue("")

	sensitiveValueAsJSON, err := json.Marshal(sensitiveValue)
	require.NoError(t, err)
	require.NotNil(t, sensitiveValueAsJSON)
	jsonassert.New(t).Assertf(string(sensitiveValueAsJSON), emptySensitiveValueAsJSON)

	sensitiveValue = core.NewSensitiveValue("test")

	sensitiveValueAsJSON, err = json.Marshal(sensitiveValue)
	require.NoError(t, err)
	require.NotNil(t, sensitiveValueAsJSON)
	jsonassert.New(t).Assertf(string(sensitiveValueAsJSON), testNoHintSensitiveValueAsJSON)
}

func TestNewSensitiveValueUnmarshalJSON(t *testing.T) {
	var emptySensitiveValue core.SensitiveValue
	err := json.Unmarshal([]byte(emptySensitiveValueAsJSON), &emptySensitiveValue)

	require.NoError(t, err)
	require.NotNil(t, emptySensitiveValue)
	require.False(t, emptySensitiveValue.HasValue)
	require.Nil(t, emptySensitiveValue.NewValue)

	var testSensitiveValue core.SensitiveValue
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
