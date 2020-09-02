package model

type SshMachineEndpoint struct {
	Fingerprint        string `json:Fingerprint`
	Host               string `json:Host`
	Port               int    `json:Port`
	ProxyID            string `json:ProxyId`
	DotNetCorePlatform string `json:DotNetCorePlatform`
}
