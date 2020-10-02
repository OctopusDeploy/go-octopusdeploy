package model

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type KubernetesEndpoint struct {
	Authentication      EndpointAuthentication
	ClusterCertificate  string
	ClusterURL          url.URL `json:"ClusterUrl" validate:"required,url"`
	CommunicationStyle  string  `validate:"required,eq=Kubernetes"`
	Container           DeploymentActionContainer
	DefaultWorkerPoolID string `json:"DefaultWorkerPoolId" validate:"required"`
	Namespace           string
	ProxyID             string `json:"ProxyId"`
	RunningInContainer  bool
	SkipTLSVerification string `json:"SkipTlsVerification"`

	endpoint
}

func NewKubernetesEndpoint(clusterURL url.URL, defaultWorkerPoolID string) *KubernetesEndpoint {
	// TODO: validate clusterURL
	// TODO: validate defaultWorkerPoolID
	kubernetesEndpoint := &KubernetesEndpoint{
		ClusterURL:          clusterURL,
		CommunicationStyle:  "Kubernetes",
		DefaultWorkerPoolID: defaultWorkerPoolID,
	}
	kubernetesEndpoint.endpoint.CommunicationStyle = kubernetesEndpoint.CommunicationStyle
	return kubernetesEndpoint
}

// GetID returns the ID value of the Kubernetes endpoint.
func (resource KubernetesEndpoint) GetID() string {
	return resource.ID
}

// GetLastModifiedBy returns the name of the account that modified the value of
// this Kubernetes endpoint.
func (resource KubernetesEndpoint) GetLastModifiedBy() string {
	return resource.LastModifiedBy
}

// GetLastModifiedOn returns the time when the value of this Kubernetes
// endpoint was changed.
func (resource KubernetesEndpoint) GetLastModifiedOn() *time.Time {
	return resource.LastModifiedOn
}

// GetLinks returns the associated links with the value of this Kubernetes
// endpoint.
func (resource KubernetesEndpoint) GetLinks() map[string]string {
	return resource.Links
}

func (e KubernetesEndpoint) MarshalJSON() ([]byte, error) {
	resource := struct {
		Authentication      EndpointAuthentication
		ClusterCertificate  string
		ClusterURL          string `json:"ClusterUrl"`
		CommunicationStyle  string
		Container           DeploymentActionContainer
		DefaultWorkerPoolID string `json:"DefaultWorkerPoolId"`
		Namespace           string
		ProxyID             string `json:"ProxyId"`
		RunningInContainer  bool
		SkipTLSVerification string `json:"SkipTlsVerification"`
		endpoint
	}{
		Authentication:      e.Authentication,
		ClusterCertificate:  e.ClusterCertificate,
		ClusterURL:          e.ClusterURL.String(),
		CommunicationStyle:  e.CommunicationStyle,
		Container:           e.Container,
		DefaultWorkerPoolID: e.DefaultWorkerPoolID,
		Namespace:           e.Namespace,
		ProxyID:             e.ProxyID,
		RunningInContainer:  e.RunningInContainer,
		SkipTLSVerification: e.SkipTLSVerification,
		endpoint:            e.endpoint,
	}

	return json.Marshal(resource)
}

func (e *KubernetesEndpoint) UnmarshalJSON(j []byte) error {
	s := strings.TrimSpace(string(j))

	var rawStrings map[string]interface{}
	err := json.Unmarshal([]byte(s), &rawStrings)
	if err != nil {
		return err
	}

	for k, v := range rawStrings {
		if k == "Authentication" {
			authentication := v.(map[string]interface{})
			for authKey, authValue := range authentication {
				if authKey == "AccountId" {
					e.Authentication.AccountID = authValue.(string)
				}
				if authKey == "ClientCertificate" {
					e.Authentication.ClientCertificate = authValue.(string)
				}
				if authKey == "AuthenticationType" {
					e.Authentication.AuthenticationType = authValue.(string)
				}
			}
		}
		if k == "ClusterCertificate" {
			e.ClusterCertificate = v.(string)
		}
		if k == "ClusterUrl" {
			rawURL := v.(string)
			u, err := url.Parse(rawURL)
			if err != nil {
				return err
			}
			e.ClusterURL = *u
		}
		if k == "CommunicationStyle" {
			e.CommunicationStyle = v.(string)
		}
		if k == "Container" {
			container := v.(map[string]interface{})
			for containerKey, containerValue := range container {
				if containerKey == "Image" {
					if containerValue != nil {
						image := containerValue.(string)
						e.Container.Image = &image
					}
				}
				if containerKey == "FeedId" {
					if containerValue != nil {
						feedID := containerValue.(string)
						e.Container.FeedID = &feedID
					}
				}
			}
		}
		if k == "DefaultWorkerPoolId" {
			e.DefaultWorkerPoolID = v.(string)
		}
		if k == "Id" {
			if v != nil {
				e.ID = v.(string)
			}
		}
		if k == "LastModifiedOn" {
			if v != nil {
				e.LastModifiedOn = v.(*time.Time)
			}
		}
		if k == "LastModifiedBy" {
			if v != nil {
				e.LastModifiedBy = v.(string)
			}
		}
		if k == "Links" {
			links := v.(map[string]interface{})
			if len(links) > 0 {
				fmt.Println(links)
			}
		}
		if k == "Namespace" {
			e.Namespace = v.(string)
		}
		if k == "ProxyId" {
			e.ProxyID = v.(string)
		}
		if k == "RunningInContainer" {
			e.RunningInContainer = v.(bool)
		}
		if k == "SkipTlsVerification" {
			e.SkipTLSVerification = v.(string)
		}
	}

	return nil
}

// Validate checks the state of the Kubernetes endpoint and returns an error if
// invalid.
func (resource KubernetesEndpoint) Validate() error {
	validate := validator.New()
	err := validate.Struct(resource)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}

		return err
	}

	return nil
}

var _ ResourceInterface = &KubernetesEndpoint{}
var _ EndpointInterface = &KubernetesEndpoint{}
