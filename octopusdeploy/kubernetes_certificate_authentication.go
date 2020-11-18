package octopusdeploy

type KubernetesCertificateAuthentication struct {
	AuthenticationType string `json:"AuthenticationType"`
	ClientCertificate  string `json:"ClientCertificate,omitempty"`
}

// NewKubernetesCertificateAuthentication creates and initializes a Kubernetes
// certificate authentication.
func NewKubernetesCertificateAuthentication() *KubernetesCertificateAuthentication {
	return &KubernetesCertificateAuthentication{
		AuthenticationType: "KubernetesCertificate",
	}
}

// GetAuthenticationType returns the authentication type of this
// Kubernetes-based authentication.
func (k *KubernetesCertificateAuthentication) GetAuthenticationType() string {
	return k.AuthenticationType
}

var _ IKubernetesAuthentication = &KubernetesCertificateAuthentication{}
