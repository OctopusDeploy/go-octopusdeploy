package octopusdeploy

type KubernetesAzureAuthentication struct {
	ClusterName          string `json:"ClusterName,omitempty"`
	ClusterResourceGroup string `json:"ClusterResourceGroup,omitempty"`

	KubernetesStandardAuthentication
}

// NewKubernetesAzureAuthentication creates and initializes a Kubernetes Azure
// authentication.
func NewKubernetesAzureAuthentication() *KubernetesAzureAuthentication {
	return &KubernetesAzureAuthentication{
		KubernetesStandardAuthentication: *NewKubernetesStandardAuthentication("KubernetesAzure"),
	}
}
