package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type MachinePolicies struct {
	Items []MachinePolicy `json:"Items"`
	PagedResults
}

type MachinePolicy struct {
	ConnectionConnectTimeout                      string                     `json:"ConnectionConnectTimeout,omitempty"`
	ConnectionRetryCountLimit                     int32                      `json:"ConnectionRetryCountLimit,omitempty"`
	ConnectionRetrySleepInterval                  string                     `json:"ConnectionRetrySleepInterval,omitempty"`
	ConnectionRetryTimeLimit                      string                     `json:"ConnectionRetryTimeLimit,omitempty"`
	Description                                   string                     `json:"Description,omitempty"`
	IsDefault                                     bool                       `json:"IsDefault"`
	MachineCleanupPolicy                          *MachineCleanupPolicy      `json:"MachineCleanupPolicy,omitempty"`
	MachineConnectivityPolicy                     *MachineConnectivityPolicy `json:"MachineConnectivityPolicy,omitempty"`
	MachineHealthCheckPolicy                      *MachineHealthCheckPolicy  `json:"MachineHealthCheckPolicy,omitempty"`
	MachineUpdatePolicy                           *MachineUpdatePolicy       `json:"MachineUpdatePolicy,omitempty"`
	Name                                          string                     `json:"Name,omitempty"`
	PollingRequestMaximumMessageProcessingTimeout string                     `json:"PollingRequestMaximumMessageProcessingTimeout,omitempty"`
	PollingRequestQueueTimeout                    string                     `json:"PollingRequestQueueTimeout,omitempty"`
	SpaceID                                       string                     `json:"SpaceId,omitempty"`

	Resource
}

// GetID returns the ID value of the MachinePolicy.
func (resource MachinePolicy) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of this MachinePolicy.
func (resource MachinePolicy) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this MachinePolicy was changed.
func (resource MachinePolicy) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this MachinePolicy.
func (resource MachinePolicy) GetLinks() map[string]string {
	return resource.Links
}

// Validate checks the state of the MachinePolicy and returns an error if invalid.
func (resource MachinePolicy) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		return err
	}

	return nil
}

var _ ResourceInterface = &MachinePolicy{}
