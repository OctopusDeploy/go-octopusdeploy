package octopusdeploy

type KubernetesStandardAuthentication struct {
	AccountID          string `json:"AccountId,omitempty"`
	AuthenticationType string `json:"AuthenticationType"`
}

// NewKubernetesStandardAuthentication creates and initializes a Kubernetes AWS
// authentication.
func NewKubernetesStandardAuthentication(authenticationType string) *KubernetesStandardAuthentication {
	if len(authenticationType) == 0 {
		authenticationType = "KubernetesStandard"
	}

	return &KubernetesStandardAuthentication{
		AuthenticationType: authenticationType,
	}
}

// GetAuthenticationType returns the authentication type of this
// Kubernetes-based authentication.
func (k *KubernetesStandardAuthentication) GetAuthenticationType() string {
	return k.AuthenticationType
}

var _ IKubernetesAuthentication = &KubernetesStandardAuthentication{}
