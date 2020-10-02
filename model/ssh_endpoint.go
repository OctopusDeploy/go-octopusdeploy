package model

import (
	"strconv"
)

// SSHEndpoint contains the information necessary to communicate with an SSH
// endpoint. If a private key file is provided it will be used; otherwise we
// fall back to username/password.
type SSHEndpoint struct {
	AccountID          string `json:"AccountId,omitempty"`
	DotNetCorePlatform string `json:"DotNetCorePlatform"`
	Fingerprint        string `json:"Fingerprint"`
	Host               string `json:"Host"`
	Port               int    `json:"Port" validate:"hostname_port"`
	ProxyID            string `json:"ProxyId"`
	URI                string `json:"Uri" validate="uri"`

	endpoint
}

func NewSSHEndpoint(host string, port int, fingerprint string) *SSHEndpoint {
	resource := &SSHEndpoint{
		Fingerprint: fingerprint,
		Host:        host,
		Port:        port,
	}

	resource.CommunicationStyle = "Ssh"
	resource.URI = "ssh://" + host + ":" + strconv.Itoa(port) + "/"

	return resource
}
