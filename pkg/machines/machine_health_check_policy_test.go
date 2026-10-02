package machines

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalOmitsZeroHealthCheckInterval(t *testing.T) {
	policy := NewMachineHealthCheckPolicy()
	policy.HealthCheckInterval = 0

	data, err := json.Marshal(policy)
	require.NoError(t, err)

	var fields map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &fields))

	_, present := fields["HealthCheckInterval"]
	assert.False(t, present, "a zero interval must be omitted, not sent as \"00:00:00\"")
}

func TestMarshalKeepsNonZeroHealthCheckInterval(t *testing.T) {
	policy := NewMachineHealthCheckPolicy()
	policy.HealthCheckInterval = 24 * time.Hour

	data, err := json.Marshal(policy)
	require.NoError(t, err)

	var fields map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &fields))

	assert.Equal(t, "01.00:00:00", fields["HealthCheckInterval"])
}

func TestMarshalCronOnlyOmitsInterval(t *testing.T) {
	policy := NewMachineHealthCheckPolicy()
	policy.HealthCheckInterval = 0
	policy.HealthCheckCron = "0 0 0 1 1 * 2099"

	data, err := json.Marshal(policy)
	require.NoError(t, err)

	var fields map[string]interface{}
	require.NoError(t, json.Unmarshal(data, &fields))

	assert.Equal(t, "0 0 0 1 1 * 2099", fields["HealthCheckCron"])
	_, present := fields["HealthCheckInterval"]
	assert.False(t, present, "cron-only policy must not also carry an interval")
}

func TestHealthCheckIntervalSurvivesRoundTrip(t *testing.T) {
	policy := NewMachineHealthCheckPolicy()
	policy.HealthCheckInterval = 24 * time.Hour

	data, err := json.Marshal(policy)
	require.NoError(t, err)

	var decoded MachineHealthCheckPolicy
	require.NoError(t, json.Unmarshal(data, &decoded))

	assert.Equal(t, 24*time.Hour, decoded.HealthCheckInterval)
}

func TestAbsentHealthCheckIntervalDecodesAsZero(t *testing.T) {
	body := `{
		"BashHealthCheckPolicy": {"RunType": "InheritFromDefault"},
		"PowerShellHealthCheckPolicy": {"RunType": "InheritFromDefault"},
		"HealthCheckCronTimezone": "UTC",
		"HealthCheckType": "RunScript"
	}`

	// Seeded with the 24h default so the assertion fails if an absent interval is ignored.
	decoded := NewMachineHealthCheckPolicy()
	require.NoError(t, json.Unmarshal([]byte(body), decoded))

	assert.Equal(t, time.Duration(0), decoded.HealthCheckInterval, "a server response with no interval means no health checks")
	assert.Empty(t, decoded.HealthCheckCron)
}
