package octopusdeploy

type KubernetesCertificateAuthentication struct {
	ClientCertificate string `json:"ClientCertificate,omitempty"`

	kubernetesAuthentication
}

// NewKubernetesCertificateAuthentication creates and initializes a Kubernetes
// certificate authentication.
func NewKubernetesCertificateAuthentication() *KubernetesCertificateAuthentication {
	return &KubernetesCertificateAuthentication{
		kubernetesAuthentication: *newKubernetesAuthentication("KubernetesCertificate"),
	}
}
