package resources

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/machines"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMachineHealthCheckPolicyMarshalJSON(t *testing.T) {
	healthCheckInterval, err := time.ParseDuration("12h34m56s")
	require.NotNil(t, healthCheckInterval)
	require.NoError(t, err)

	machineHealthCheckPolicy := machines.NewMachineHealthCheckPolicy()
	machineHealthCheckPolicy.BashHealthCheckPolicy = machines.NewMachineScriptPolicy()
	machineHealthCheckPolicy.HealthCheckCron = "health-check-cron"
	machineHealthCheckPolicy.HealthCheckCronTimezone = "UTC"
	machineHealthCheckPolicy.HealthCheckInterval = healthCheckInterval
	machineHealthCheckPolicy.HealthCheckType = "OnlyConnectivity"
	machineHealthCheckPolicy.PowerShellHealthCheckPolicy = machines.NewMachineScriptPolicy()

	jsonEncoding, err := json.Marshal(machineHealthCheckPolicy)
	require.NoError(t, err)
	require.NotNil(t, jsonEncoding)

	actual := string(jsonEncoding)

	jsonassert.New(t).Assertf(actual, machineHealthCheckPolicyAsJSON)

}

func TestMachineHealthCheckPolicyUnmarshalJSON(t *testing.T) {
	var machineHealthCheckPolicy machines.MachineHealthCheckPolicy
	err := json.Unmarshal([]byte(machineHealthCheckPolicyAsJSON), &machineHealthCheckPolicy)
	require.NoError(t, err)
	require.NotNil(t, machineHealthCheckPolicy)

	healthCheckInterval, err := time.ParseDuration("12h34m56s")
	require.NotNil(t, healthCheckInterval)
	require.NoError(t, err)

	bashHealthCheckPolicy := machines.NewMachineScriptPolicy()
	powerShellHealthCheckPolicy := machines.NewMachineScriptPolicy()

	assert.Equal(t, bashHealthCheckPolicy, machineHealthCheckPolicy.BashHealthCheckPolicy)
	assert.Equal(t, "health-check-cron", machineHealthCheckPolicy.HealthCheckCron)
	assert.Equal(t, "UTC", machineHealthCheckPolicy.HealthCheckCronTimezone)
	assert.Equal(t, healthCheckInterval, machineHealthCheckPolicy.HealthCheckInterval)
	assert.Equal(t, "OnlyConnectivity", machineHealthCheckPolicy.HealthCheckType)
	assert.Equal(t, powerShellHealthCheckPolicy, machineHealthCheckPolicy.PowerShellHealthCheckPolicy)
}

const machineHealthCheckPolicyAsJSON string = `{
  "PowerShellHealthCheckPolicy": {
    "RunType": "InheritFromDefault",
    "ScriptBody": null
  },
  "BashHealthCheckPolicy": {
    "RunType": "InheritFromDefault",
    "ScriptBody": null
  },
  "HealthCheckInterval": "12:34:56",
  "HealthCheckCron": "health-check-cron",
  "HealthCheckCronTimezone": "UTC",
  "HealthCheckType": "OnlyConnectivity"
}`
