package machines

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/internal"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/deployments"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type KubernetesEndpoint struct {
	Authentication         IKubernetesAuthentication              `json:"Authentication,omitempty"`
	ClusterCertificate     string                                 `json:"ClusterCertificate,omitempty"`
	ClusterCertificatePath string                                 `json:"ClusterCertificatePath,omitempty"`
	ClusterURL             *url.URL                               `json:"ClusterUrl" validate:"required,url"`
	Container              *deployments.DeploymentActionContainer `json:"Container,omitempty"`
	DefaultWorkerPoolID    string                                 `json:"DefaultWorkerPoolId,omitempty"`
	Namespace              string                                 `json:"Namespace,omitempty"`
	ProxyID                string                                 `json:"ProxyId,omitempty"`
	RunningInContainer     bool                                   `json:"RunningInContainer"`
	SkipTLSVerification    bool                                   `json:"SkipTlsVerification"`

	endpoint
}

// NewKubernetesEndpoint creates and initializes a new Kubernetes endpoint.
func NewKubernetesEndpoint(clusterURL *url.URL) *KubernetesEndpoint {
	return &KubernetesEndpoint{
		ClusterURL: clusterURL,
		endpoint:   *newEndpoint("Kubernetes"),
	}
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
		Authentication         IKubernetesAuthentication              `json:"Authentication,omitempty"`
		ClusterCertificate     string                                 `json:"ClusterCertificate,omitempty"`
		ClusterCertificatePath string                                 `json:"ClusterCertificatePath,omitempty"`
		ClusterURL             string                                 `json:"ClusterUrl"`
		Container              *deployments.DeploymentActionContainer `json:"Container,omitempty"`
		DefaultWorkerPoolID    string                                 `json:"DefaultWorkerPoolId"`
		Namespace              string                                 `json:"Namespace,omitempty"`
		ProxyID                string                                 `json:"ProxyId,omitempty"`
		RunningInContainer     bool                                   `json:"RunningInContainer"`
		SkipTLSVerification    string                                 `json:"SkipTlsVerification"`
		endpoint
	}{
		Authentication:         k.Authentication,
		ClusterCertificate:     k.ClusterCertificate,
		ClusterCertificatePath: k.ClusterCertificatePath,
		ClusterURL:             k.ClusterURL.String(),
		Container:              k.Container,
		DefaultWorkerPoolID:    k.DefaultWorkerPoolID,
		Namespace:              k.Namespace,
		ProxyID:                k.ProxyID,
		RunningInContainer:     k.RunningInContainer,
		SkipTLSVerification:    cases.Title(language.Und, cases.NoLower).String(strconv.FormatBool(k.SkipTLSVerification)),
		endpoint:               k.endpoint,
	}

	return json.Marshal(kubernetesEndpoint)
}

// UnmarshalJSON sets this Kubernetes endpoint to its representation in JSON.
func (k *KubernetesEndpoint) UnmarshalJSON(data []byte) error {
	var fields struct {
		ClusterCertificate     string                                 `json:"ClusterCertificate,omitempty"`
		ClusterCertificatePath string                                 `json:"ClusterCertificatePath,omitempty"`
		ClusterURL             string                                 `json:"ClusterUrl"`
		Container              *deployments.DeploymentActionContainer `json:"Container,omitempty"`
		DefaultWorkerPoolID    string                                 `json:"DefaultWorkerPoolId"`
		Namespace              string                                 `json:"Namespace,omitempty"`
		ProxyID                string                                 `json:"ProxyId,omitempty"`
		RunningInContainer     bool                                   `json:"RunningInContainer"`
		SkipTLSVerification    string                                 `json:"SkipTlsVerification"`
		endpoint
	}

	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	// validate JSON representation
	validate := validator.New()
	if err := validate.Struct(fields); err != nil {
		return err
	}

	// return error if unable to parse cluster URL string
	u, err := url.Parse(fields.ClusterURL)
	if err != nil {
		return err
	}

	if !internal.IsEmpty(fields.SkipTLSVerification) {
		skipTLSVerification, err := strconv.ParseBool(fields.SkipTLSVerification)
		if err != nil {
			return err
		}
		k.SkipTLSVerification = skipTLSVerification
	}

	k.ClusterCertificate = fields.ClusterCertificate
	k.ClusterCertificatePath = fields.ClusterCertificatePath
	k.ClusterURL = u
	k.Container = fields.Container
	k.DefaultWorkerPoolID = fields.DefaultWorkerPoolID
	k.Namespace = fields.Namespace
	k.ProxyID = fields.ProxyID
	k.RunningInContainer = fields.RunningInContainer
	k.endpoint = fields.endpoint

	var kubernetesEndpoint map[string]*json.RawMessage
	if err := json.Unmarshal(data, &kubernetesEndpoint); err != nil {
		return err
	}

	var authentication *json.RawMessage
	var authenticationProperties map[string]*json.RawMessage
	var authenticationType string

	if kubernetesEndpoint["Authentication"] != nil {
		authenticationValue := kubernetesEndpoint["Authentication"]

		if err = json.Unmarshal(*authenticationValue, &authentication); err != nil {
			return err
		}

		if err = json.Unmarshal(*authentication, &authenticationProperties); err != nil {
			return err
		}

		if authenticationProperties["AuthenticationType"] != nil {
			at := authenticationProperties["AuthenticationType"]
			json.Unmarshal(*at, &authenticationType)
		}
	}

	switch authenticationType {
	case "KubernetesAws":
		var kubernetesAwsAuthentication *KubernetesAwsAuthentication
		if err := json.Unmarshal(*authentication, &kubernetesAwsAuthentication); err != nil {
			return err
		}
		k.Authentication = kubernetesAwsAuthentication
	case "KubernetesAzure":
		var kubernetesAzureAuthentication *KubernetesAzureAuthentication
		if err := json.Unmarshal(*authentication, &kubernetesAzureAuthentication); err != nil {
			return err
		}
		k.Authentication = kubernetesAzureAuthentication
	case "KubernetesCertificate":
		var kubernetesCertificateAuthentication *KubernetesCertificateAuthentication
		if err := json.Unmarshal(*authentication, &kubernetesCertificateAuthentication); err != nil {
			return err
		}
		k.Authentication = kubernetesCertificateAuthentication
	case "KubernetesGoogleCloud":
		var kubernetesGcpAuthentication *KubernetesGcpAuthentication
		if err := json.Unmarshal(*authentication, &kubernetesGcpAuthentication); err != nil {
			return err
		}
		k.Authentication = kubernetesGcpAuthentication
	case "KubernetesStandard":
		var kubernetesStandardAuthentication *KubernetesStandardAuthentication
		if err := json.Unmarshal(*authentication, &kubernetesStandardAuthentication); err != nil {
			return err
		}
		k.Authentication = kubernetesStandardAuthentication
	case "KubernetesPodService":
		var kubernetesPodAuthentication *KubernetesPodAuthentication
		if err := json.Unmarshal(*authentication, &kubernetesPodAuthentication); err != nil {
			return err
		}
		k.Authentication = kubernetesPodAuthentication
	}

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
