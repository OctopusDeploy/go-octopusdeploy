package model

type KubernetesMachineEndpoint struct {
	ClusterCertificate  string                         `json:ClusterCertificate`
	ClusterUrl          string                         `json:ClusterUrl validate:"omitempty,url`
	Namespace           string                         `json:Namespace`
	SkipTlsVerification string                         `json:SkipTlsVerification`
	ProxyID             string                         `json:ProxyId`
	Authentication      *MachineEndpointAuthentication `json:"Authentication"`
	RunningInContainer  bool                           `json:RunningInContainer`
	Container           DeploymentActionContainer      `json:Container`
}
