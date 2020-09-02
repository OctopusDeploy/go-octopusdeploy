package model

type KubernetesMachineEndpoint struct {
	ClusterCertificate  string                         `json:"ClusterCertificate,omitempty"`
	ClusterURL          string                         `json:"ClusterUrl,omitempty" validate:"omitempty,url"`
	Namespace           string                         `json:"Namespace,omitempty"`
	SkipTLSVerification string                         `json:"SkipTlsVerification,omitempty"`
	Authentication      *MachineEndpointAuthentication `json:"Authentication,omitempty"`
	RunningInContainer  bool                           `json:"RunningInContainer"`
	Container           DeploymentActionContainer      `json:"Container"`
}
