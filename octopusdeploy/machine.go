package octopusdeploy

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

// machines defines a collection of machines with built-in support for paged
// results from the API.
type machines struct {
	Items []machine `json:"Items"`
	PagedResults
}

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
	err := json.Unmarshal(b, &fields)
	if err != nil {
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
	err = json.Unmarshal(b, &machine)
	if err != nil {
		return err
	}

	var endpoint *json.RawMessage
	var endpointProperties map[string]*json.RawMessage
	var communicationStyle string

	if machine["Endpoint"] != nil {
		endpointValue := machine["Endpoint"]

		err = json.Unmarshal(*endpointValue, &endpoint)
		if err != nil {
			return err
		}

		err = json.Unmarshal(*endpoint, &endpointProperties)
		if err != nil {
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
		err := json.Unmarshal(*endpoint, &azureCloudServiceEndpoint)
		if err != nil {
			return err
		}
		m.Endpoint = azureCloudServiceEndpoint
	case "AzureServiceFabricCluster":
		var azureServiceFabricEndpoint *AzureServiceFabricEndpoint
		err := json.Unmarshal(*endpoint, &azureServiceFabricEndpoint)
		if err != nil {
			return err
		}
		m.Endpoint = azureServiceFabricEndpoint
	case "AzureWebApp":
		var azureWebAppEndpoint *AzureWebAppEndpoint
		err := json.Unmarshal(*endpoint, &azureWebAppEndpoint)
		if err != nil {
			return err
		}
		m.Endpoint = azureWebAppEndpoint
	case "Kubernetes":
		var kubernetesEndpoint *KubernetesEndpoint
		err := json.Unmarshal(*endpoint, &kubernetesEndpoint)
		if err != nil {
			return err
		}
		m.Endpoint = kubernetesEndpoint
	case "None":
		var cloudRegionEndpoint *CloudRegionEndpoint
		err := json.Unmarshal(*endpoint, &cloudRegionEndpoint)
		if err != nil {
			return err
		}
		m.Endpoint = cloudRegionEndpoint
	case "OfflineDrop":
		var offlinePackageDropEndpoint *OfflinePackageDropEndpoint
		err := json.Unmarshal(*endpoint, &offlinePackageDropEndpoint)
		if err != nil {
			return err
		}
		m.Endpoint = offlinePackageDropEndpoint
	case "Ssh":
		var sshEndpoint *SSHEndpoint
		err := json.Unmarshal(*endpoint, &sshEndpoint)
		if err != nil {
			return err
		}
		m.Endpoint = sshEndpoint
	case "TentacleActive":
		var pollingTentacleEndpoint *PollingTentacleEndpoint
		err := json.Unmarshal(*endpoint, &pollingTentacleEndpoint)
		if err != nil {
			return err
		}
		m.Endpoint = pollingTentacleEndpoint
	case "TentaclePassive":
		var listeningTentacleEndpoint *ListeningTentacleEndpoint
		err := json.Unmarshal(*endpoint, &listeningTentacleEndpoint)
		if err != nil {
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
