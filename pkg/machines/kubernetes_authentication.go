package machines

type IKubernetesAuthentication interface {
	GetAuthenticationType() string
}

type KubernetesAuthentication struct {
	AccountID                 string `json:"AccountId,omitempty"`
	AdminLogin                string `json:"AdminLogin,omitempty"`
	AssumedRoleARN            string `json:"AssumedRoleArn,omitempty"`
	AssumedRoleSession        string `json:"AssumedRoleSession,omitempty"`
	AssumeRole                bool   `json:"AssumeRole"`
	AssumeRoleSessionDuration int    `json:"AssumeRoleSessionDurationSeconds,omitempty"`
	AssumeRoleExternalID      string `json:"AssumeRoleExternalId,omitempty"`
	AuthenticationType        string `json:"AuthenticationType,omitempty"`
	ClientCertificate         string `json:"ClientCertificate,omitempty"`
	ClusterName               string `json:"ClusterName,omitempty"`
	ClusterResourceGroup      string `json:"ClusterResourceGroup,omitempty"`
	ImpersonateServiceAccount bool   `json:"ImpersonateServiceAccount"`
	Project                   string `json:"Project,omitempty"`
	Region                    string `json:"Region,omitempty"`
	ServiceAccountEmails      string `json:"ServiceAccountEmails,omitempty"`
	UseInstanceRole           bool   `json:"UseInstanceRole"`
	UseVmServiceAccount       bool   `json:"UseVmServiceAccount"`
	Zone                      string `json:"Zone,omitempty"`
	TokenPath                 string `json:"TokenPath,omitempty"`
}

// GetAuthenticationType returns the authentication type of this
// Kubernetes-based authentication.
func (k *KubernetesAuthentication) GetAuthenticationType() string {
	return k.AuthenticationType
}

var _ IKubernetesAuthentication = &KubernetesAuthentication{}
