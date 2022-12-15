package machines

type KubernetesAzureAuthentication struct {
	ClusterName          string `json:"ClusterName,omitempty"`
	ClusterResourceGroup string `json:"ClusterResourceGroup,omitempty"`
	AdminLogin           bool   `json:"AdminLogin,omitempty"`

	KubernetesStandardAuthentication
}

// NewKubernetesAzureAuthentication creates and initializes a Kubernetes Azure
// authentication.
func NewKubernetesAzureAuthentication() *KubernetesAzureAuthentication {
	return &KubernetesAzureAuthentication{
		KubernetesStandardAuthentication: *NewKubernetesStandardAuthentication("KubernetesAzure"),
	}
}
