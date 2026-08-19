package machinepolicies

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator/v10"
)

type MachineHealthCheckPolicy struct {
	BashHealthCheckPolicy       *MachineScriptPolicy `json:"BashHealthCheckPolicy" validate:"required"`
	HealthCheckCron             string               `json:"HealthCheckCron,omitempty"`
	HealthCheckCronTimezone     string               `json:"HealthCheckCronTimezone" validate:"required"`
	HealthCheckInterval         time.Duration        `json:"HealthCheckInterval,omitempty"`
	HealthCheckType             string               `json:"HealthCheckType" validate:"required,oneof=OnlyConnectivity RunScript"`
	PowerShellHealthCheckPolicy *MachineScriptPolicy `json:"PowerShellHealthCheckPolicy" validate:"required"`
}

func NewMachineHealthCheckPolicy() *MachineHealthCheckPolicy {
	return &MachineHealthCheckPolicy{
		BashHealthCheckPolicy:       NewMachineScriptPolicy(),
		HealthCheckCronTimezone:     "UTC",
		HealthCheckInterval:         24 * time.Hour,
		HealthCheckType:             "RunScript",
		PowerShellHealthCheckPolicy: NewMachineScriptPolicy(),
	}
}

// MarshalJSON returns a machine health check policy as its JSON encoding.
func (m *MachineHealthCheckPolicy) MarshalJSON() ([]byte, error) {
	// The server reads an absent interval as "no health checks" but "00:00:00" as "check constantly".
	healthCheckInterval := ""
	if m.HealthCheckInterval > 0 {
		healthCheckInterval = ToTimeSpan(m.HealthCheckInterval)
	}

	machineHealthCheckPolicy := struct {
		BashHealthCheckPolicy       *MachineScriptPolicy `json:"BashHealthCheckPolicy,omitempty"`
		HealthCheckCron             string               `json:"HealthCheckCron,omitempty"`
		HealthCheckCronTimezone     string               `json:"HealthCheckCronTimezone,omitempty"`
		HealthCheckInterval         string               `json:"HealthCheckInterval,omitempty"`
		HealthCheckType             string               `json:"HealthCheckType" validate:"required,oneof=OnlyConnectivity RunScript"`
		PowerShellHealthCheckPolicy *MachineScriptPolicy `json:"PowerShellHealthCheckPolicy,omitempty"`
	}{
		BashHealthCheckPolicy:       m.BashHealthCheckPolicy,
		HealthCheckCron:             m.HealthCheckCron,
		HealthCheckCronTimezone:     m.HealthCheckCronTimezone,
		HealthCheckInterval:         healthCheckInterval,
		HealthCheckType:             m.HealthCheckType,
		PowerShellHealthCheckPolicy: m.PowerShellHealthCheckPolicy,
	}

	return json.Marshal(machineHealthCheckPolicy)
}

// UnmarshalJSON sets this machine health check policy to its representation in
// JSON.
func (m *MachineHealthCheckPolicy) UnmarshalJSON(data []byte) error {
	var fields struct {
		BashHealthCheckPolicy       *MachineScriptPolicy `json:"BashHealthCheckPolicy,omitempty"`
		HealthCheckCron             string               `json:"HealthCheckCron,omitempty"`
		HealthCheckCronTimezone     string               `json:"HealthCheckCronTimezone,omitempty"`
		HealthCheckInterval         string               `json:"HealthCheckInterval,omitempty"`
		HealthCheckType             string               `json:"HealthCheckType" validate:"required,oneof=OnlyConnectivity RunScript"`
		PowerShellHealthCheckPolicy *MachineScriptPolicy `json:"PowerShellHealthCheckPolicy,omitempty"`
	}
	err := json.Unmarshal(data, &fields)
	if err != nil {
		return err
	}

	// validate JSON representation
	v := validator.New()
	err = v.Struct(fields)
	if err != nil {
		return err
	}

	// An absent interval means no health checks, so it must clear any existing value.
	m.HealthCheckInterval = FromTimeSpan(fields.HealthCheckInterval)

	m.BashHealthCheckPolicy = fields.BashHealthCheckPolicy
	m.HealthCheckCron = fields.HealthCheckCron
	m.HealthCheckCronTimezone = fields.HealthCheckCronTimezone
	m.HealthCheckType = fields.HealthCheckType
	m.PowerShellHealthCheckPolicy = fields.PowerShellHealthCheckPolicy

	return nil
}

// Validate checks the state of the machine health check policy and returns an
// error if invalid.
func (m *MachineHealthCheckPolicy) Validate() error {
	return validator.New().Struct(m)
}
