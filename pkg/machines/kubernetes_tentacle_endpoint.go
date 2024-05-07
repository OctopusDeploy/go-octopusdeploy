package machines

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/url"
)

type KubernetesTentacleEndpoint struct {
	TentacleEndpointConfiguration *TentacleEndpointConfiguration `json:"TentacleEndpointConfiguration" validate:"required"`
	KubernetesAgentDetails        *KubernetesAgentDetails        `json:"KubernetesAgentDetails,omitempty"`
	UpgradeLocked                 bool                           `json:"UpgradeLocked"`
	DefaultNamespace              string                         `json:"DefaultNamespace,omitempty"`

	endpoint `validate:"required"`
}

func NewKubernetesTentacleEndpoint(uri *url.URL, thumbprint string, upgradeLocked bool, communicationsMode string, defaultNamespace string) *KubernetesTentacleEndpoint {
	return &KubernetesTentacleEndpoint{
		TentacleEndpointConfiguration: &TentacleEndpointConfiguration{
			Thumbprint:        thumbprint,
			URI:               uri,
			CommunicationMode: communicationsMode,
		},
		UpgradeLocked:    upgradeLocked,
		DefaultNamespace: defaultNamespace,
		endpoint:         *newEndpoint("KubernetesTentacle"),
	}
}

func (l KubernetesTentacleEndpoint) MarshalJSON() ([]byte, error) {
	te := struct {
		TentacleEndpointConfiguration *TentacleEndpointConfiguration `json:"TentacleEndpointConfiguration" validate:"required"`
		KubernetesAgentDetails        *KubernetesAgentDetails        `json:"KubernetesAgentDetails,omitempty"`
		UpgradeLocked                 bool                           `json:"UpgradeLocked"`
		DefaultNamespace              string                         `json:"DefaultNamespace,omitempty"`

		endpoint
	}{
		TentacleEndpointConfiguration: l.TentacleEndpointConfiguration,
		KubernetesAgentDetails:        l.KubernetesAgentDetails,
		endpoint:                      l.endpoint,
	}

	return json.Marshal(te)
}

// UnmarshalJSON sets this tentacle endpoint to its representation in JSON.
func (l *KubernetesTentacleEndpoint) UnmarshalJSON(b []byte) error {
	var fields struct {
		TentacleEndpointConfiguration *TentacleEndpointConfiguration `json:"TentacleEndpointConfiguration" validate:"required"`
		KubernetesAgentDetails        *KubernetesAgentDetails        `json:"KubernetesAgentDetails,omitempty"`
		UpgradeLocked                 bool                           `json:"UpgradeLocked"`
		DefaultNamespace              string                         `json:"DefaultNamespace,omitempty"`

		endpoint
	}

	if err := json.Unmarshal(b, &fields); err != nil {
		return err
	}

	// validate JSON representation
	validate := validator.New()
	if err := validate.Struct(fields); err != nil {
		return err
	}

	l.TentacleEndpointConfiguration = fields.TentacleEndpointConfiguration
	l.KubernetesAgentDetails = fields.KubernetesAgentDetails
	l.UpgradeLocked = fields.UpgradeLocked
	l.DefaultNamespace = fields.DefaultNamespace
	l.endpoint = fields.endpoint

	return nil
}
