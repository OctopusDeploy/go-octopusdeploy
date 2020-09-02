package model

type SshMachineEndpoint struct {
	Fingerprint        string  `json:"Fingerprint,omitempty"`
	Host               string  `json:"Host,omitempty"`
	Port               *uint16 `json:"Port,omitempty"`
	DotNetCorePlatform string  `json:"DotNetCorePlatform,omitempty"`
}
