package octopusdeploy

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type KubernetesEndpoint struct {
	Authentication      EndpointAuthentication    `json:"Authentication,omitempty"`
	ClusterCertificate  string                    `json:"ClusterCertificate,omitempty"`
	ClusterURL          *url.URL                  `json:"ClusterUrl" validate:"required,url"`
	CommunicationStyle  string                    `json:"CommunicationStyle" validate:"required,eq=Kubernetes"`
	Container           DeploymentActionContainer `json:"Container,omitempty"`
	DefaultWorkerPoolID string                    `json:"DefaultWorkerPoolId,omitempty"`
	Namespace           string                    `json:"Namespace,omitempty"`
	ProxyID             string                    `json:"ProxyId,omitempty"`
	RunningInContainer  bool                      `json:"RunningInContainer"`
	SkipTLSVerification bool                      `json:"SkipTlsVerification"`

	resource
}

// NewKubernetesEndpoint creates and initializes a new Kubernetes endpoint.
func NewKubernetesEndpoint(clusterURL *url.URL) *KubernetesEndpoint {
	return &KubernetesEndpoint{
		ClusterURL:         clusterURL,
		CommunicationStyle: "Kubernetes",
		resource:           *newResource(),
	}
}

// GetCommunicationStyle returns the communication style of this endpoint.
func (s *KubernetesEndpoint) GetCommunicationStyle() string {
	return s.CommunicationStyle
}

// GetDefaultWorkerPoolID returns the default worker pool ID of this Kubernetes
// endpoint.
func (k *KubernetesEndpoint) GetDefaultWorkerPoolID() string {
	return k.DefaultWorkerPoolID
}

// GetProxyID returns the proxy ID associated with this Kubernetes endpoint.
func (k *KubernetesEndpoint) GetProxyID() string {
	return k.ProxyID
}

// SetDefaultWorkerPoolID sets the default worker pool ID of this Kubernetes
// endpoint.
func (k *KubernetesEndpoint) SetDefaultWorkerPoolID(defaultWorkerPoolID string) {
	k.DefaultWorkerPoolID = defaultWorkerPoolID
}

// SetProxyID sets the proxy ID associated with this Kubernetes endpoint.
func (k *KubernetesEndpoint) SetProxyID(proxyID string) {
	k.ProxyID = proxyID
}

// MarshalJSON returns a Kubernetes endpoint as its JSON encoding.
func (k *KubernetesEndpoint) MarshalJSON() ([]byte, error) {
	kubernetesEndpoint := struct {
		Authentication      EndpointAuthentication    `json:"Authentication,omitempty"`
		ClusterCertificate  string                    `json:"ClusterCertificate,omitempty"`
		ClusterURL          string                    `json:"ClusterUrl"`
		CommunicationStyle  string                    `json:"CommunicationStyle" validate:"required,eq=Kubernetes"`
		Container           DeploymentActionContainer `json:"Container,omitempty"`
		DefaultWorkerPoolID string                    `json:"DefaultWorkerPoolId"`
		Namespace           string                    `json:"Namespace,omitempty"`
		ProxyID             string                    `json:"ProxyId,omitempty"`
		RunningInContainer  bool                      `json:"RunningInContainer"`
		SkipTLSVerification string                    `json:"SkipTlsVerification"`
		resource
	}{
		Authentication:      k.Authentication,
		ClusterCertificate:  k.ClusterCertificate,
		ClusterURL:          k.ClusterURL.String(),
		CommunicationStyle:  k.CommunicationStyle,
		Container:           k.Container,
		DefaultWorkerPoolID: k.DefaultWorkerPoolID,
		Namespace:           k.Namespace,
		ProxyID:             k.ProxyID,
		RunningInContainer:  k.RunningInContainer,
		SkipTLSVerification: strings.Title(strconv.FormatBool(k.SkipTLSVerification)),
		resource:            k.resource,
	}

	return json.Marshal(kubernetesEndpoint)
}

// UnmarshalJSON sets this Kubernetes endpoint to its representation in JSON.
func (k *KubernetesEndpoint) UnmarshalJSON(data []byte) error {
	var fields struct {
		Authentication      EndpointAuthentication    `json:"Authentication,omitempty"`
		ClusterCertificate  string                    `json:"ClusterCertificate,omitempty"`
		ClusterURL          string                    `json:"ClusterUrl" validate:"required,url"`
		CommunicationStyle  string                    `json:"CommunicationStyle" validate:"required,eq=Kubernetes"`
		Container           DeploymentActionContainer `json:"Container,omitempty"`
		DefaultWorkerPoolID string                    `json:"DefaultWorkerPoolId"`
		Namespace           string                    `json:"Namespace,omitempty"`
		ProxyID             string                    `json:"ProxyId,omitempty"`
		RunningInContainer  bool                      `json:"RunningInContainer"`
		SkipTLSVerification string                    `json:"SkipTlsVerification"`
		resource
	}
	err := json.Unmarshal(data, &fields)
	if err != nil {
		return err
	}

	// validate JSON representation
	validate := validator.New()
	err = validate.Struct(fields)
	if err != nil {
		return err
	}

	// return error if unable to parse cluster URL string
	u, err := url.Parse(fields.ClusterURL)
	if err != nil {
		return err
	}

	if !isEmpty(fields.SkipTLSVerification) {
		skipTLSVerification, err := strconv.ParseBool(fields.SkipTLSVerification)
		if err != nil {
			return err
		}
		k.SkipTLSVerification = skipTLSVerification
	}

	k.Authentication = fields.Authentication
	k.ClusterCertificate = fields.ClusterCertificate
	k.ClusterURL = u
	k.CommunicationStyle = fields.CommunicationStyle
	k.Container = fields.Container
	k.DefaultWorkerPoolID = fields.DefaultWorkerPoolID
	k.Namespace = fields.Namespace
	k.ProxyID = fields.ProxyID
	k.RunningInContainer = fields.RunningInContainer
	k.resource = fields.resource

	return nil
}

// Validate checks the state of the Kubernetes endpoint and returns an error if
// invalid.
func (k *KubernetesEndpoint) Validate() error {
	return validator.New().Struct(k)
}

var _ IEndpoint = &KubernetesEndpoint{}
var _ IEndpointWithProxy = &KubernetesEndpoint{}
var _ IRunsOnAWorker = &KubernetesEndpoint{}
