package model

type MachineEndpoint struct {
	CommunicationStyle     string                         `json:"CommunicationStyle"`
	ProxyID                *string                        `json:"ProxyId"`
	Thumbprint             string                         `json:"Thumbprint"`
	TentacleVersionDetails MachineTentacleVersionDetails  `json:"TentacleVersionDetails"`
	URI                    string                         `json:"Uri"`                 // This is not in the spec doc, but it shows up and needs to be kept in sync
	ClusterCertificate     string                         `json:"ClusterCertificate"`  // Kubernetes (not in spec doc)
	ClusterURL             string                         `json:"ClusterUrl"`          // Kubernetes (not in spec doc)
	Namespace              string                         `json:"Namespace"`           // Kubernetes (not in spec doc)
	SkipTLSVerification    string                         `json:"SkipTlsVerification"` // Kubernetes (not in spec doc)
	DefaultWorkerPoolID    string                         `json:"DefaultWorkerPoolId"`
	Authentication         *MachineEndpointAuthentication `json:"Authentication"`
	Resource
}
