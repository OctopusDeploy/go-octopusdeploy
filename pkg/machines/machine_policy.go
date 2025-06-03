package machines

import (
	"encoding/json"
	"time"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type MachinePolicy struct {
	ConnectionConnectTimeout                      time.Duration                       `json:"ConnectionConnectTimeout" validate:"required,min=10s"`
	ConnectionRetryCountLimit                     int32                               `json:"ConnectionRetryCountLimit" validate:"required,gte=2"`
	ConnectionRetrySleepInterval                  time.Duration                       `json:"ConnectionRetrySleepInterval" validate:"required"`
	ConnectionRetryTimeLimit                      time.Duration                       `json:"ConnectionRetryTimeLimit" validate:"required,min=10s"`
	Description                                   string                              `json:"Description,omitempty"`
	IsDefault                                     bool                                `json:"IsDefault"`
	MachineCleanupPolicy                          *MachineCleanupPolicy               `json:"MachineCleanupPolicy"`
	MachineConnectivityPolicy                     *MachineConnectivityPolicy          `json:"MachineConnectivityPolicy"`
	MachineHealthCheckPolicy                      *MachineHealthCheckPolicy           `json:"MachineHealthCheckPolicy"`
	MachineUpdatePolicy                           *MachineUpdatePolicy                `json:"MachineUpdatePolicy"`
	MachinePackageCacheRetentionPolicy            *MachinePackageCacheRetentionPolicy `json:"MachinePackageCacheRetentionPolicy"`
	Name                                          string                              `json:"Name" validate:"required,notblank"`
	PollingRequestMaximumMessageProcessingTimeout time.Duration                       `json:"PollingRequestMaximumMessageProcessingTimeout,omitempty"`
	PollingRequestQueueTimeout                    time.Duration                       `json:"PollingRequestQueueTimeout" validate:"required"`
	SpaceID                                       string                              `json:"SpaceId,omitempty"`

	resources.Resource
}

func NewMachinePolicy(name string) *MachinePolicy {
	return &MachinePolicy{
		ConnectionConnectTimeout:                      time.Minute,
		ConnectionRetryCountLimit:                     5,
		ConnectionRetrySleepInterval:                  time.Second,
		ConnectionRetryTimeLimit:                      5 * time.Minute,
		MachineCleanupPolicy:                          NewMachineCleanupPolicy(),
		MachineConnectivityPolicy:                     NewMachineConnectivityPolicy(),
		MachineHealthCheckPolicy:                      NewMachineHealthCheckPolicy(),
		MachineUpdatePolicy:                           NewMachineUpdatePolicy(),
		MachinePackageCacheRetentionPolicy:            NewDefaultMachinePackageCacheRetentionPolicy(),
		Name:                                          name,
		PollingRequestMaximumMessageProcessingTimeout: 10 * time.Minute,
		PollingRequestQueueTimeout:                    2 * time.Minute,
		Resource:                                      *resources.NewResource(),
	}
}

// MarshalJSON returns a machine policy as its JSON encoding.
func (m *MachinePolicy) MarshalJSON() ([]byte, error) {
	machinePolicy := struct {
		ConnectionConnectTimeout                      string                              `json:"ConnectionConnectTimeout" validate:"required"`
		ConnectionRetryCountLimit                     int32                               `json:"ConnectionRetryCountLimit" validate:"required"`
		ConnectionRetrySleepInterval                  string                              `json:"ConnectionRetrySleepInterval" validate:"required"`
		ConnectionRetryTimeLimit                      string                              `json:"ConnectionRetryTimeLimit" validate:"required"`
		Description                                   string                              `json:"Description,omitempty"`
		IsDefault                                     bool                                `json:"IsDefault"`
		MachineCleanupPolicy                          *MachineCleanupPolicy               `json:"MachineCleanupPolicy"`
		MachineConnectivityPolicy                     *MachineConnectivityPolicy          `json:"MachineConnectivityPolicy"`
		MachineHealthCheckPolicy                      *MachineHealthCheckPolicy           `json:"MachineHealthCheckPolicy"`
		MachineUpdatePolicy                           *MachineUpdatePolicy                `json:"MachineUpdatePolicy"`
		MachinePackageCacheRetentionPolicy            *MachinePackageCacheRetentionPolicy `json:"MachinePackageCacheRetentionPolicy"`
		Name                                          string                              `json:"Name" validate:"required,notblank"`
		PollingRequestMaximumMessageProcessingTimeout string                              `json:"PollingRequestMaximumMessageProcessingTimeout,omitempty"`
		PollingRequestQueueTimeout                    string                              `json:"PollingRequestQueueTimeout" validate:"required"`
		SpaceID                                       string                              `json:"SpaceId,omitempty"`
		resources.Resource
	}{
		ConnectionConnectTimeout:                      ToTimeSpan(m.ConnectionConnectTimeout),
		ConnectionRetryCountLimit:                     m.ConnectionRetryCountLimit,
		ConnectionRetrySleepInterval:                  ToTimeSpan(m.ConnectionRetrySleepInterval),
		ConnectionRetryTimeLimit:                      ToTimeSpan(m.ConnectionRetryTimeLimit),
		Description:                                   m.Description,
		IsDefault:                                     m.IsDefault,
		MachineCleanupPolicy:                          m.MachineCleanupPolicy,
		MachineConnectivityPolicy:                     m.MachineConnectivityPolicy,
		MachineHealthCheckPolicy:                      m.MachineHealthCheckPolicy,
		MachineUpdatePolicy:                           m.MachineUpdatePolicy,
		MachinePackageCacheRetentionPolicy:            m.MachinePackageCacheRetentionPolicy,
		Name:                                          m.Name,
		PollingRequestMaximumMessageProcessingTimeout: ToTimeSpan(m.PollingRequestMaximumMessageProcessingTimeout),
		PollingRequestQueueTimeout:                    ToTimeSpan(m.PollingRequestQueueTimeout),
		SpaceID:                                       m.SpaceID,
		Resource:                                      m.Resource,
	}

	return json.Marshal(machinePolicy)
}

// UnmarshalJSON sets this Kubernetes endpoint to its representation in JSON.
func (m *MachinePolicy) UnmarshalJSON(data []byte) error {
	var fields struct {
		ConnectionConnectTimeout                      string                              `json:"ConnectionConnectTimeout" validate:"required"`
		ConnectionRetryCountLimit                     int32                               `json:"ConnectionRetryCountLimit" validate:"required"`
		ConnectionRetrySleepInterval                  string                              `json:"ConnectionRetrySleepInterval" validate:"required"`
		ConnectionRetryTimeLimit                      string                              `json:"ConnectionRetryTimeLimit" validate:"required"`
		Description                                   string                              `json:"Description,omitempty"`
		IsDefault                                     bool                                `json:"IsDefault"`
		MachineCleanupPolicy                          *MachineCleanupPolicy               `json:"MachineCleanupPolicy"`
		MachineConnectivityPolicy                     *MachineConnectivityPolicy          `json:"MachineConnectivityPolicy"`
		MachineHealthCheckPolicy                      *MachineHealthCheckPolicy           `json:"MachineHealthCheckPolicy"`
		MachineUpdatePolicy                           *MachineUpdatePolicy                `json:"MachineUpdatePolicy"`
		MachinePackageCacheRetentionPolicy            *MachinePackageCacheRetentionPolicy `json:"MachinePackageCacheRetentionPolicy"`
		Name                                          string                              `json:"Name"`
		PollingRequestMaximumMessageProcessingTimeout string                              `json:"PollingRequestMaximumMessageProcessingTimeout,omitempty"`
		PollingRequestQueueTimeout                    string                              `json:"PollingRequestQueueTimeout" validate:"required"`
		SpaceID                                       string                              `json:"SpaceId,omitempty"`
		resources.Resource
	}
	err := json.Unmarshal(data, &fields)
	if err != nil {
		return err
	}

	// validate JSON representation
	v := validator.New()
	err = v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.Struct(fields)
	if err != nil {
		return err
	}

	if len(fields.ConnectionConnectTimeout) > 0 {
		m.ConnectionConnectTimeout = FromTimeSpan(fields.ConnectionConnectTimeout)
	}

	if len(fields.ConnectionRetrySleepInterval) > 0 {
		m.ConnectionRetrySleepInterval = FromTimeSpan(fields.ConnectionRetrySleepInterval)
	}

	if len(fields.ConnectionRetryTimeLimit) > 0 {
		m.ConnectionRetryTimeLimit = FromTimeSpan(fields.ConnectionRetryTimeLimit)
	}

	if len(fields.PollingRequestMaximumMessageProcessingTimeout) > 0 {
		m.PollingRequestMaximumMessageProcessingTimeout = FromTimeSpan(fields.PollingRequestMaximumMessageProcessingTimeout)
	}

	if len(fields.PollingRequestQueueTimeout) > 0 {
		m.PollingRequestQueueTimeout = FromTimeSpan(fields.PollingRequestQueueTimeout)
	}

	m.ConnectionRetryCountLimit = fields.ConnectionRetryCountLimit
	m.Description = fields.Description
	m.IsDefault = fields.IsDefault
	m.MachineCleanupPolicy = fields.MachineCleanupPolicy
	m.MachineConnectivityPolicy = fields.MachineConnectivityPolicy
	m.MachineHealthCheckPolicy = fields.MachineHealthCheckPolicy
	m.MachineUpdatePolicy = fields.MachineUpdatePolicy
	m.MachinePackageCacheRetentionPolicy = fields.MachinePackageCacheRetentionPolicy
	m.Name = fields.Name
	m.SpaceID = fields.SpaceID
	m.Resource = fields.Resource

	return nil
}

// Validate checks the state of the machine policy and returns an error if
// invalid.
func (m *MachinePolicy) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	return v.Struct(m)
}
