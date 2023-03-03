package machines

type KubernetesGcpAuthentication struct {
	ClusterName               string `json:"ClusterName,omitempty"`
	ImpersonateServiceAccount bool   `json:"ImpersonateServiceAccount"`
	Project                   string `json:"Project,omitempty"`
	Region                    string `json:"Region,omitempty"`
	ServiceAccountEmails      string `json:"ServiceAccountEmails,omitempty"`
	UseVmServiceAccount       bool   `json:"UseVmServiceAccount"`
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
