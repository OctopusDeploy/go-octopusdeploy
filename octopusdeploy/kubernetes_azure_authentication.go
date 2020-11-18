package octopusdeploy

type KubernetesAzureAuthentication struct {
	ClusterName          string `json:"ClusterName,omitempty"`
	ClusterResourceGroup string `json:"ClusterResourceGroup,omitempty"`
	AdminLogin           string `json:"AdminLogin,omitempty"`

	KubernetesStandardAuthentication
}

// NewKubernetesAzureAuthentication creates and initializes a Kubernetes Azure
// authentication.
func NewKubernetesAzureAuthentication() *KubernetesAzureAuthentication {
	return &KubernetesAzureAuthentication{
		KubernetesStandardAuthentication: *NewKubernetesStandardAuthentication("KubernetesAzure"),
	}
}
