package octopusdeploy

// kubernetesAuthentication is the base definition of Kubernetes-based
// authentication.
type kubernetesAuthentication struct {
	AuthenticationType string `json:"AuthenticationType" validate:"required,oneof=KubernetesAws KubernetesAzure KubernetesCertificate KubernetesStandard None"`
}

// newKubernetesAuthentication creates and initializes a new Kubernetes-based
// authentication.
func newKubernetesAuthentication(authenticationType string) *kubernetesAuthentication {
	kubernetesAuthentication := &kubernetesAuthentication{
		AuthenticationType: authenticationType,
	}
	return kubernetesAuthentication
}

// GetAuthenticationType returns the authentication type of this
// Kubernetes-based authenication.
func (e kubernetesAuthentication) GetAuthenticationType() string {
	return e.AuthenticationType
}

var _ IKubernetesAuthentication = &kubernetesAuthentication{}
