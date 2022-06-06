package octopusdeploy

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// machine represents a machine on which Octopus can orchestrate actions. May
// either be a deployment target or a worker. It may either correspond to a
// running instance of the tentacle agent ("TentacleActive" and
// "TentaclePassive" communication styles), or an agentless target accessed
// using a protocol like SSH.
type machine struct {
	Endpoint          IEndpoint `json:"Endpoint" validate:"required"`
	HasLatestCalamari bool      `json:"HasLatestCalamari"`
	HealthStatus      string    `json:"HealthStatus,omitempty" validate:"omitempty,oneof=HasWarnings Healthy Unavailable Unhealthy Unknown"`
	IsDisabled        bool      `json:"IsDisabled"`
	IsInProcess       bool      `json:"IsInProcess"`
	MachinePolicyID   string    `json:"MachinePolicyId,omitempty"`
	Name              string    `json:"Name"`
	OperatingSystem   string    `json:"OperatingSystem,omitempty"`
	ShellName         string    `json:"ShellName,omitempty"`
	ShellVersion      string    `json:"ShellVersion,omitempty"`
	Status            string    `json:"Status,omitempty" validate:"omitempty,oneof=CalamariNeedsUpgrade Disabled NeedsUpgrade Offline Online Unknown"`
	StatusSummary     string    `json:"StatusSummary,omitempty"`
	Thumbprint        string    `json:"Thumbprint,omitempty"`
	URI               string    `json:"Uri,omitempty" validate:"omitempty,uri"`

	resource
}

func newMachine(name string, endpoint IEndpoint) *machine {
	return &machine{
		Endpoint:        endpoint,
		Name:            name,
		OperatingSystem: "Unknown",
		ShellName:       "Unknown",
		ShellVersion:    "Unknown",
		resource:        *newResource(),
	}
}

// MarshalJSON returns a machine as its JSON encoding.
func (m *machine) MarshalJSON() ([]byte, error) {
	machine := struct {
		Endpoint          IEndpoint `json:"Endpoint" validate:"required"`
		HasLatestCalamari bool      `json:"HasLatestCalamari"`
		HealthStatus      string    `json:"HealthStatus,omitempty" validate:"omitempty,oneof=HasWarnings Healthy Unavailable Unhealthy Unknown"`
		IsDisabled        bool      `json:"IsDisabled"`
		IsInProcess       bool      `json:"IsInProcess"`
		MachinePolicyID   string    `json:"MachinePolicyId,omitempty"`
		Name              string    `json:"Name,omitempty"`
		OperatingSystem   string    `json:"OperatingSystem,omitempty"`
		ShellName         string    `json:"ShellName,omitempty"`
		ShellVersion      string    `json:"ShellVersion,omitempty"`
		Status            string    `json:"Status,omitempty" validate:"omitempty,oneof=CalamariNeedsUpgrade Disabled NeedsUpgrade Offline Online Unknown"`
		StatusSummary     string    `json:"StatusSummary,omitempty"`
		Thumbprint        string    `json:"Thumbprint,omitempty"`
		URI               string    `json:"Uri,omitempty" validate:"omitempty,uri"`
		resource
	}{
		Endpoint:          m.Endpoint,
		HasLatestCalamari: m.HasLatestCalamari,
		HealthStatus:      m.HealthStatus,
		IsDisabled:        m.IsDisabled,
		IsInProcess:       m.IsInProcess,
		MachinePolicyID:   m.MachinePolicyID,
		Name:              m.Name,
		OperatingSystem:   m.OperatingSystem,
		ShellName:         m.ShellName,
		ShellVersion:      m.ShellVersion,
		Status:            m.Status,
		StatusSummary:     m.StatusSummary,
		Thumbprint:        m.Thumbprint,
		URI:               m.URI,
		resource:          m.resource,
	}

	return json.Marshal(machine)
}

// UnmarshalJSON sets this machine to its representation in JSON.
func (m *machine) UnmarshalJSON(b []byte) error {
	var fields struct {
		HasLatestCalamari bool   `json:"HasLatestCalamari"`
		HealthStatus      string `json:"HealthStatus,omitempty" validate:"omitempty,oneof=HasWarnings Healthy Unavailable Unhealthy Unknown"`
		IsDisabled        bool   `json:"IsDisabled"`
		IsInProcess       bool   `json:"IsInProcess"`
		MachinePolicyID   string `json:"MachinePolicyId,omitempty"`
		Name              string `json:"Name,omitempty"`
		OperatingSystem   string `json:"OperatingSystem,omitempty"`
		ShellName         string `json:"ShellName,omitempty"`
		ShellVersion      string `json:"ShellVersion,omitempty"`
		Status            string `json:"Status,omitempty" validate:"omitempty,oneof=CalamariNeedsUpgrade Disabled NeedsUpgrade Offline Online Unknown"`
		StatusSummary     string `json:"StatusSummary,omitempty"`
		Thumbprint        string `json:"Thumbprint,omitempty"`
		URI               string `json:"Uri,omitempty" validate:"omitempty,uri"`
		resource
	}
	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}

	m.HasLatestCalamari = fields.HasLatestCalamari
	m.HealthStatus = fields.HealthStatus
	m.IsDisabled = fields.IsDisabled
	m.IsInProcess = fields.IsInProcess
	m.MachinePolicyID = fields.MachinePolicyID
	m.Name = fields.Name
	m.OperatingSystem = fields.OperatingSystem
	m.ShellName = fields.ShellName
	m.ShellVersion = fields.ShellVersion
	m.Status = fields.Status
	m.StatusSummary = fields.StatusSummary
	m.Thumbprint = fields.Thumbprint
	m.URI = fields.URI
	m.resource = fields.resource

	var machine map[string]*json.RawMessage
	if err := json.Unmarshal(b, &machine); err != nil {
		return err
	}

	var endpoint *json.RawMessage
	var endpointProperties map[string]*json.RawMessage
	var communicationStyle string

	if machine["Endpoint"] != nil {
		endpointValue := machine["Endpoint"]

		if err := json.Unmarshal(*endpointValue, &endpoint); err != nil {
			return err
		}

		if err := json.Unmarshal(*endpoint, &endpointProperties); err != nil {
			return err
		}

		if endpointProperties["CommunicationStyle"] != nil {
			cs := endpointProperties["CommunicationStyle"]
			json.Unmarshal(*cs, &communicationStyle)
		}
	}

	switch communicationStyle {
	case "AzureCloudService":
		var azureCloudServiceEndpoint *AzureCloudServiceEndpoint
		if err := json.Unmarshal(*endpoint, &azureCloudServiceEndpoint); err != nil {
			return err
		}
		m.Endpoint = azureCloudServiceEndpoint
	case "AzureServiceFabricCluster":
		var azureServiceFabricEndpoint *AzureServiceFabricEndpoint
		if err := json.Unmarshal(*endpoint, &azureServiceFabricEndpoint); err != nil {
			return err
		}
		m.Endpoint = azureServiceFabricEndpoint
	case "AzureWebApp":
		var azureWebAppEndpoint *AzureWebAppEndpoint
		if err := json.Unmarshal(*endpoint, &azureWebAppEndpoint); err != nil {
			return err
		}
		m.Endpoint = azureWebAppEndpoint
	case "Kubernetes":
		var kubernetesEndpoint *KubernetesEndpoint
		if err := json.Unmarshal(*endpoint, &kubernetesEndpoint); err != nil {
			return err
		}
		m.Endpoint = kubernetesEndpoint
	case "None":
		var cloudRegionEndpoint *CloudRegionEndpoint
		if err := json.Unmarshal(*endpoint, &cloudRegionEndpoint); err != nil {
			return err
		}
		m.Endpoint = cloudRegionEndpoint
	case "OfflineDrop":
		var offlinePackageDropEndpoint *OfflinePackageDropEndpoint
		if err := json.Unmarshal(*endpoint, &offlinePackageDropEndpoint); err != nil {
			return err
		}
		m.Endpoint = offlinePackageDropEndpoint
	case "Ssh":
		var sshEndpoint *SSHEndpoint
		if err := json.Unmarshal(*endpoint, &sshEndpoint); err != nil {
			return err
		}
		m.Endpoint = sshEndpoint
	case "TentacleActive":
		var pollingTentacleEndpoint *PollingTentacleEndpoint
		if err := json.Unmarshal(*endpoint, &pollingTentacleEndpoint); err != nil {
			return err
		}
		m.Endpoint = pollingTentacleEndpoint
	case "TentaclePassive":
		var listeningTentacleEndpoint *ListeningTentacleEndpoint
		if err := json.Unmarshal(*endpoint, &listeningTentacleEndpoint); err != nil {
			return err
		}
		m.Endpoint = listeningTentacleEndpoint
	}

	return nil
}

// Validate checks the state of the machine and returns an error if invalid.
func (m machine) Validate() error {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}

		return err
	}

	return nil
}
