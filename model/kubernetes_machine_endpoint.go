package model

type KubernetesMachineEndpoint struct {
	ClusterCertificate  string                         `json:ClusterCertificate`
	ClusterUrl          string                         `json:ClusterUrl`
	Namespace           int                            `json:Namespace`
	SkipTlsVerification string                         `json:SkipTlsVerification`
	ProxyID             string                         `json:ProxyId`
	DefaultWorkerPoolID string                         `json:DefaultWorkerPoolId`
	Authentication      *MachineEndpointAuthentication `json:"Authentication"`
	RunningInContainer  bool                           `json:RunningInContainer`
	Container           DeploymentActionContainer      `json:Container`
}
