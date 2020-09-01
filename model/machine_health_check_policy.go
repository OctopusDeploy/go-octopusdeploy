package model

type MachineHealthCheckPolicy struct {
	TentacleEndpointHealthCheckPolicy map[string]string `json:"TentacleEndpointHealthCheckPolicy"`
	SSHEndpointHealthCheckPolicy      map[string]string `json:"SshEndpointHealthCheckPolicy"`
	HealthCheckInterval               string            `json:"HealthCheckInterval"`
}
