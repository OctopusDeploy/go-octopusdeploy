package octopusdeploy

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/go-playground/validator/v10"
)

// SSHEndpoint contains the information necessary to communicate with an SSH
// endpoint. If a private key file is provided it will be used; otherwise we
// fall back to username/password.
type SSHEndpoint struct {
	AccountID          string `json:"AccountId,omitempty"`
	CommunicationStyle string `json:"CommunicationStyle" validate:"required,eq=Ssh"`
	DotNetCorePlatform string
	Fingerprint        string
	Host               string
	ProxyID            string `json:"ProxyId,omitempty"`
	Port               int
	URI                *url.URL `json:"Uri"`

	resource
}

// NewSSHEndpoint creates and initializes a new SSH endpoint.
func NewSSHEndpoint(host string, port int, fingerprint string) *SSHEndpoint {
	sshEndpoint := &SSHEndpoint{
		CommunicationStyle: "Ssh",
		Fingerprint:        fingerprint,
		Host:               host,
		Port:               port,
		resource:           *newResource(),
	}

	url, _ := url.Parse("ssh://" + host + ":" + strconv.Itoa(port) + "/")
	sshEndpoint.URI = url

	return sshEndpoint
}

// GetAccountID returns the account ID associated with this SSH endpoint.
func (s *SSHEndpoint) GetAccountID() string {
	return s.AccountID
}

// GetCommunicationStyle returns the communication style of this endpoint.
func (s *SSHEndpoint) GetCommunicationStyle() string {
	return s.CommunicationStyle
}

// GetFingerprint returns the fingerprint associated with this SSH endpoint.
func (s *SSHEndpoint) GetFingerprint() string {
	return s.Fingerprint
}

// GetHost returns the host associated with this SSH endpoint.
func (s *SSHEndpoint) GetHost() string {
	return s.Host
}

// GetProxyID returns the proxy ID associated with this SSH endpoint.
func (s *SSHEndpoint) GetProxyID() string {
	return s.ProxyID
}

func (s *SSHEndpoint) MarshalJSON() ([]byte, error) {
	sshEndpoint := struct {
		AccountID          string `json:"AccountId,omitempty"`
		ComunicationStyle  string `json:"CommunicationStyle" validate:"required,eq=Ssh"`
		DotNetCorePlatform string
		Fingerprint        string
		Host               string
		Port               int
		ProxyID            string `json:"ProxyId,omitempty"`
		URI                string `json:"Uri"`
		resource
	}{
		AccountID:          s.AccountID,
		ComunicationStyle:  s.CommunicationStyle,
		DotNetCorePlatform: s.DotNetCorePlatform,
		Fingerprint:        s.Fingerprint,
		Host:               s.Host,
		Port:               s.Port,
		ProxyID:            s.ProxyID,
		URI:                s.URI.String(),
		resource:           s.resource,
	}

	return json.Marshal(sshEndpoint)
}

// SetProxyID sets the proxy ID associated with this SSH endpoint.
func (s *SSHEndpoint) SetProxyID(proxyID string) {
	s.ProxyID = proxyID
}

// UnmarshalJSON sets this SSH endpoint to its representation in JSON.
func (s *SSHEndpoint) UnmarshalJSON(b []byte) error {
	var fields struct {
		AccountID          string `json:"AccountId,omitempty"`
		CommunicationStyle string `json:"CommunicationStyle" validate:"required,eq=Ssh"`
		DotNetCorePlatform string
		Fingerprint        string
		Host               string
		Port               int
		ProxyID            string `json:"ProxyId,omitempty"`
		URI                string `json:"Uri"`
		resource
	}
	err := json.Unmarshal(b, &fields)
	if err != nil {
		return err
	}

	// validate JSON representation
	validate := validator.New()
	err = validate.Struct(fields)
	if err != nil {
		return err
	}

	// return error if unable to parse URI string
	u, err := url.Parse(fields.URI)
	if err != nil {
		return err
	}

	// set endpoint fields
	s.AccountID = fields.AccountID
	s.CommunicationStyle = fields.CommunicationStyle
	s.DotNetCorePlatform = fields.DotNetCorePlatform
	s.Fingerprint = fields.Fingerprint
	s.Host = fields.Host
	s.Port = fields.Port
	s.ProxyID = fields.ProxyID
	s.resource = fields.resource
	s.URI = u

	return nil
}

// Validate checks the state of the SSH endpoint and returns an error if
// invalid.
func (s *SSHEndpoint) Validate() error {
	return validator.New().Struct(s)
}

var _ IEndpoint = &SSHEndpoint{}
var _ IEndpointWithAccount = &SSHEndpoint{}
var _ IEndpointWithFingerprint = &SSHEndpoint{}
var _ IEndpointWithHostname = &SSHEndpoint{}
var _ IEndpointWithProxy = &SSHEndpoint{}
