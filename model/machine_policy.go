package model

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
