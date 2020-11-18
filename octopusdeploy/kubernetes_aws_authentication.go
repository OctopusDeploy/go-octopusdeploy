package octopusdeploy

type KubernetesAwsAuthentication struct {
	AssumedRoleARN            string `json:"AssumedRoleArn,omitempty"`
	AssumedRoleSession        string `json:"AssumedRoleSession,omitempty"`
	AssumeRole                bool   `json:"AssumeRole,omitempty"`
	AssumeRoleExternalID      string `json:"AssumeRoleExternalId,omitempty"`
	AssumeRoleSessionDuration int    `json:"AssumeRoleSessionDurationSeconds,omitempty"`
	ClusterName               string `json:"ClusterName,omitempty"`
	UseInstanceRole           bool   `json:"UseInstanceRole,omitempty"`

	KubernetesStandardAuthentication
}

// NewKubernetesAwsAuthentication creates and initializes a Kubernetes AWS
// authentication.
func NewKubernetesAwsAuthentication() *KubernetesAwsAuthentication {
	return &KubernetesAwsAuthentication{
		KubernetesStandardAuthentication: *NewKubernetesStandardAuthentication("KubernetesAws"),
	}
}
