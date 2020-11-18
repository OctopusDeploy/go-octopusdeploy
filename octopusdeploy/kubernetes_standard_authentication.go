package octopusdeploy

type KubernetesStandardAuthentication struct {
	AccountID string `json:"AccountId,omitempty"`

	kubernetesAuthentication
}

// NewKubernetesStandardAuthentication creates and initializes a Kubernetes AWS
// authentication.
func NewKubernetesStandardAuthentication(authenticationType string) *KubernetesStandardAuthentication {
	if len(authenticationType) == 0 {
		authenticationType = "KubernetesStandard"
	}

	return &KubernetesStandardAuthentication{
		kubernetesAuthentication: *newKubernetesAuthentication(authenticationType),
	}
}
