package machines

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestMachineHealthCheckPolicyMarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		interval time.Duration
		expected string
	}{
		{"interval is written as a time span", 24 * time.Hour, `"HealthCheckInterval":"01.00:00:00"`},
		{"sub-day interval is written as a time span", time.Hour, `"HealthCheckInterval":"01:00:00"`},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			policy := NewMachineHealthCheckPolicy()
			policy.HealthCheckInterval = testCase.interval

			data, err := json.Marshal(policy)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !strings.Contains(string(data), testCase.expected) {
				t.Errorf("marshalled policy %s does not contain %s", data, testCase.expected)
			}
		})
	}
}

// A zero interval means the health check schedule is "Never", which the server represents by the
// absence of both HealthCheckInterval and HealthCheckCron. Writing "00:00:00" instead is read back
// as an interval of zero minutes, not as Never.
func TestMachineHealthCheckPolicyMarshalJSONOmitsZeroInterval(t *testing.T) {
	policy := NewMachineHealthCheckPolicy()
	policy.HealthCheckInterval = 0

	data, err := json.Marshal(policy)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.Contains(string(data), "HealthCheckInterval") {
		t.Errorf("marshalled policy %s should not contain HealthCheckInterval", data)
	}
}

func TestMachineHealthCheckPolicyUnmarshalJSONWithoutInterval(t *testing.T) {
	data := `{
		"PowerShellHealthCheckPolicy": {"RunType": "Inline", "ScriptBody": ""},
		"BashHealthCheckPolicy": {"RunType": "Inline", "ScriptBody": ""},
		"HealthCheckCronTimezone": "UTC",
		"HealthCheckType": "RunScript"
	}`

	policy := &MachineHealthCheckPolicy{}
	if err := json.Unmarshal([]byte(data), policy); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if policy.HealthCheckInterval != 0 {
		t.Errorf("HealthCheckInterval = %s, want 0", policy.HealthCheckInterval)
	}
}
