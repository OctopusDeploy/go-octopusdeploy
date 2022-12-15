package machines

type KubernetesPodAuthentication struct {
	TokenPath          string `json:"TokenPath,omitempty"`
	AuthenticationType string `json:"AuthenticationType"`
}

// NewKubernetesPodAuthentication creates and initializes a Kubernetes
// authentication.
func NewKubernetesPodAuthentication() *KubernetesPodAuthentication {
	authenticationType := "KubernetesPodService"
	return &KubernetesPodAuthentication{
		AuthenticationType: authenticationType,
	}
}

// GetAuthenticationType returns the authentication type of this
// Kubernetes-based authentication.
func (k *KubernetesPodAuthentication) GetAuthenticationType() string {
	return k.AuthenticationType
}

var _ IKubernetesAuthentication = &KubernetesPodAuthentication{}
