package octopusdeploy

type KubernetesGcpAuthentication struct {
	ClusterName               string `json:"ClusterName,omitempty"`
	ImpersonateServiceAccount bool   `json:"ImpersonateServiceAccount,omitempty"`
	Project                   string `json:"Project,omitempty"`
	Region                    string `json:"Region,omitempty"`
	ServiceAccountEmails      string `json:"ServiceAccountEmails,omitempty"`
	UseVmServiceAccount       bool   `json:"UseVmServiceAccount,omitempty"`
	Zone                      string `json:"Zone,omitempty"`

	KubernetesStandardAuthentication
}

// NewKubernetesGcpAuthentication creates and initializes a Kubernetes GCP
// authentication.
func NewKubernetesGcpAuthentication() *KubernetesGcpAuthentication {
	return &KubernetesGcpAuthentication{
		KubernetesStandardAuthentication: *NewKubernetesStandardAuthentication("KubernetesGoogleCloud"),
	}
}
