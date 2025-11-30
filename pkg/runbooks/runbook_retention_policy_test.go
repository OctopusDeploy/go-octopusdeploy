package runbooks

import (
	"encoding/json"
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/require"
)

func TestCountBasedRunbookRetentionPolicyMarshalJSON(t *testing.T) {
	expectedJson := `{
		"Strategy": "Count",
		"QuantityToKeep": 10,
		"ShouldKeepForever": false,
		"Unit": "Days"
	}`

	runbookRetentionPolicy, err := NewCountBasedRunbookRetentionPolicy(10, RunbookRetentionUnitDays)
	require.NoError(t, err)
	require.NotNil(t, runbookRetentionPolicy)

	runbookRetentionPolicyAsJSON, err := json.Marshal(runbookRetentionPolicy)
	require.NoError(t, err)
	require.NotNil(t, runbookRetentionPolicyAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(runbookRetentionPolicyAsJSON))
}

func TestKeepForeverRunbookRetentionPolicyMarshalJSON(t *testing.T) {
	expectedJson := `{
		"Strategy": "Forever",
		"QuantityToKeep": 0,
		"ShouldKeepForever": true,
		"Unit": "Items"
	}`

	runbookRetentionPolicy := NewKeepForeverRunbookRetentionPolicy()
	require.NotNil(t, runbookRetentionPolicy)
	runbookRetentionPolicyAsJSON, err := json.Marshal(runbookRetentionPolicy)
	require.NoError(t, err)
	require.NotNil(t, runbookRetentionPolicyAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(runbookRetentionPolicyAsJSON))
}

func TestDefaultRunbookRetentionPolicyMarshalJSON(t *testing.T) {
	expectedJson := `{
		"Strategy": "Default",
		"QuantityToKeep": 100,
		"ShouldKeepForever": false,
		"Unit": "Items"
	}`

	runbookRetentionPolicy := NewDefaultRunbookRetentionPolicy()
	require.NotNil(t, runbookRetentionPolicy)
	runbookRetentionPolicyAsJSON, err := json.Marshal(runbookRetentionPolicy)
	require.NoError(t, err)
	require.NotNil(t, runbookRetentionPolicyAsJSON)

	jsonassert.New(t).Assertf(expectedJson, string(runbookRetentionPolicyAsJSON))
}

func TestCountBasedRunbookRetentionPolicyWithStrategyUnmarshalJSON(t *testing.T) {
	inputJSON := `{
		"Strategy": "Count",
		"QuantityToKeep": 10,
		"Unit": "Days"
	}`

	var runbookRetentionPolicy RunbookRetentionPolicy
	err := json.Unmarshal([]byte(inputJSON), &runbookRetentionPolicy)
	require.NoError(t, err)
	require.NotNil(t, runbookRetentionPolicy)
	require.Equal(t, RunbookRetentionStrategyCount, runbookRetentionPolicy.Strategy)
	require.Equal(t, int32(10), runbookRetentionPolicy.QuantityToKeep)
	require.Equal(t, RunbookRetentionUnitDays, runbookRetentionPolicy.Unit)
}

func TestKeepForeverRunbookRetentionPolicyWithStrategyUnmarshalJSON(t *testing.T) {
	inputJSON := `{
		"Strategy": "Forever",
		"QuantityToKeep": 0,
		"Unit": "Items",
		"ShouldKeepForever": true
	}`

	var runbookRetentionPolicy RunbookRetentionPolicy
	err := json.Unmarshal([]byte(inputJSON), &runbookRetentionPolicy)
	require.NoError(t, err)
	require.NotNil(t, runbookRetentionPolicy)
	require.Equal(t, RunbookRetentionStrategyForever, runbookRetentionPolicy.Strategy)
	require.Equal(t, int32(0), runbookRetentionPolicy.QuantityToKeep)
	require.Equal(t, RunbookRetentionUnitItems, runbookRetentionPolicy.Unit)
}

func TestDefaultRunbookRetentionPolicyWithStrategyUnmarshalJSON(t *testing.T) {
	inputJSON := `{
		"Strategy": "Default",
		"QuantityToKeep": 100,
		"Unit": "Items",
		"ShouldKeepForever": false
	}`

	var runbookRetentionPolicy RunbookRetentionPolicy
	err := json.Unmarshal([]byte(inputJSON), &runbookRetentionPolicy)
	require.NoError(t, err)
	require.NotNil(t, runbookRetentionPolicy)
	require.Equal(t, RunbookRetentionStrategyDefault, runbookRetentionPolicy.Strategy)
	require.Equal(t, int32(100), runbookRetentionPolicy.QuantityToKeep)
	require.Equal(t, RunbookRetentionUnitItems, runbookRetentionPolicy.Unit)
}

func TestCountBasedRunbookRetentionPolicyWithoutStrategyUnmarshalJSON(t *testing.T) {
	inputJSON := `{
		"QuantityToKeep": 10,
		"Unit": "Days",
		"ShouldKeepForever": false
	}`

	var runbookRetentionPolicy RunbookRetentionPolicy
	err := json.Unmarshal([]byte(inputJSON), &runbookRetentionPolicy)
	require.NoError(t, err)
	require.NotNil(t, runbookRetentionPolicy)
	require.Equal(t, RunbookRetentionStrategyCount, runbookRetentionPolicy.Strategy)
	require.Equal(t, int32(10), runbookRetentionPolicy.QuantityToKeep)
	require.Equal(t, RunbookRetentionUnitDays, runbookRetentionPolicy.Unit)
}

func TestKeepForeverRunbookRetentionPolicyWithoutStrategyUnmarshalJSON(t *testing.T) {
	inputJSON := `{
		"QuantityToKeep": 0,
		"Unit": "Items",
		"ShouldKeepForever": true
	}`

	var runbookRetentionPolicy RunbookRetentionPolicy
	err := json.Unmarshal([]byte(inputJSON), &runbookRetentionPolicy)
	require.NoError(t, err)
	require.NotNil(t, runbookRetentionPolicy)
	require.Equal(t, RunbookRetentionStrategyForever, runbookRetentionPolicy.Strategy)
	require.Equal(t, int32(0), runbookRetentionPolicy.QuantityToKeep)
	require.Equal(t, RunbookRetentionUnitItems, runbookRetentionPolicy.Unit)
}

func TestDefaultRunbookRetentionPolicyWithoutStrategyUnmarshalJSON(t *testing.T) {
	inputJSON := `{
		"QuantityToKeep": 100,
		"Unit": "Items",
		"ShouldKeepForever": false
	}`

	var runbookRetentionPolicy RunbookRetentionPolicy
	err := json.Unmarshal([]byte(inputJSON), &runbookRetentionPolicy)
	require.NoError(t, err)
	require.NotNil(t, runbookRetentionPolicy)
	require.Equal(t, RunbookRetentionStrategyDefault, runbookRetentionPolicy.Strategy)
	require.Equal(t, int32(100), runbookRetentionPolicy.QuantityToKeep)
	require.Equal(t, RunbookRetentionUnitItems, runbookRetentionPolicy.Unit)
}
